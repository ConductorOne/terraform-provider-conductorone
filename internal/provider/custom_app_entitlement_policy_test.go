package provider

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
)

// These tests lock the IGA-4347 contract for conductorone_custom_app_entitlement:
// one configured write source (`provision_policy`) that reaches the API, and the
// pre-existing computed `provisioner_policy` retained as a deprecated read-only
// mirror so no published attribute path is removed or renamed.

func findAttribute(t *testing.T, attrs []*tfprotov6.SchemaAttribute, name string) *tfprotov6.SchemaAttribute {
	t.Helper()
	for _, a := range attrs {
		if a.Name == name {
			return a
		}
	}
	return nil
}

// TestCustomAppEntitlementProvisionPolicySchemaSurface locks the public schema:
// the writable subtree no longer forces replacement, and the legacy computed
// attribute keeps its path, its computed-only shape and gains a deprecation.
func TestCustomAppEntitlementProvisionPolicySchemaSurface(t *testing.T) {
	ctx := context.Background()
	ps := providerserver.NewProtocol6(New("test")())()
	resp, err := ps.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %s", err)
	}
	rs := resp.ResourceSchemas["conductorone_custom_app_entitlement"]
	if rs == nil {
		t.Fatal("conductorone_custom_app_entitlement schema missing")
	}

	writable := findAttribute(t, rs.Block.Attributes, "provision_policy")
	legacy := findAttribute(t, rs.Block.Attributes, "provisioner_policy")
	if writable == nil {
		t.Fatal("provision_policy attribute missing")
	}
	if legacy == nil {
		t.Fatal("legacy provisioner_policy attribute was removed; it must stay as a deprecated computed mirror")
	}
	if !writable.Optional || !writable.Computed {
		t.Errorf("provision_policy must stay Optional+Computed, got optional=%v computed=%v", writable.Optional, writable.Computed)
	}
	if legacy.Optional || legacy.Required || !legacy.Computed {
		t.Errorf("legacy provisioner_policy must stay computed-only, got optional=%v required=%v computed=%v",
			legacy.Optional, legacy.Required, legacy.Computed)
	}
	if !legacy.Deprecated {
		t.Error("legacy provisioner_policy must be marked deprecated")
	}

	// The maintained patch must keep the replacement modifiers out of the
	// writable subtree, and must not disturb the identity constraints.
	src, err := os.ReadFile("custom_app_entitlement_resource.go")
	if err != nil {
		t.Fatalf("read generated resource: %s", err)
	}
	text := string(src)
	start := strings.Index(text, `"provision_policy": schema.SingleNestedAttribute{`)
	end := strings.Index(text, `"provisioner_policy": schema.SingleNestedAttribute{`)
	if start < 0 || end < 0 || end < start {
		t.Fatalf("could not locate the policy subtrees (regeneration drift?): start=%d end=%d", start, end)
	}
	if n := strings.Count(text[start:end], "RequiresReplaceIfConfigured"); n != 0 {
		t.Errorf("provision_policy subtree still forces replacement in %d places", n)
	}
	if n := strings.Count(text[:start], "RequiresReplaceIfConfigured"); n == 0 {
		t.Error("identity replacement constraints (app_resource_id/app_resource_type_id) are missing")
	}
}

// TestCustomAppEntitlementConfiguredPolicyIsTheWriteSource proves the update
// request is built from the configured attribute, never from the server-read
// mirror, and that the wire names are unchanged.
func TestCustomAppEntitlementConfiguredPolicyIsTheWriteSource(t *testing.T) {
	ctx := context.Background()
	m := &CustomAppEntitlementResourceModel{
		AppID: types.StringValue("app-1"),
		ProvisionPolicy: &tfTypes.ProvisionPolicy{
			Manual: &tfTypes.ManualProvision{UserIds: []types.String{types.StringValue("configured-user")}},
		},
		// The read-side mirror holds server truth; it must never be serialized.
		ProvisionerPolicy: &tfTypes.ProvisionPolicy{
			Manual: &tfTypes.ManualProvision{UserIds: []types.String{types.StringValue("server-value")}},
		},
	}

	update, diags := m.ToSharedUpdateAppEntitlementRequest(ctx)
	if diags.HasError() {
		t.Fatalf("update diagnostics: %v", diags)
	}
	b, err := json.Marshal(update.Entitlement)
	if err != nil {
		t.Fatalf("marshal update: %s", err)
	}
	payload := string(b)
	if !strings.Contains(payload, `"provisionerPolicy"`) {
		t.Errorf("update must serialize the entity wire name provisionerPolicy, got %s", payload)
	}
	if !strings.Contains(payload, "configured-user") {
		t.Errorf("update did not serialize the configured policy, got %s", payload)
	}
	if strings.Contains(payload, "server-value") {
		t.Errorf("update serialized the deprecated read mirror instead of the configured policy, got %s", payload)
	}

	create, diags := m.ToSharedCreateAppEntitlementRequest(ctx)
	if diags.HasError() {
		t.Fatalf("create diagnostics: %v", diags)
	}
	cb, err := json.Marshal(create)
	if err != nil {
		t.Fatalf("marshal create: %s", err)
	}
	if !strings.Contains(string(cb), `"provisionPolicy"`) || !strings.Contains(string(cb), "configured-user") {
		t.Errorf("create must serialize the configured policy under provisionPolicy, got %s", cb)
	}
}

// TestCustomAppEntitlementPolicyOneofPruning proves only the configured union arm
// is serialized and that no update mask is sent today.
func TestCustomAppEntitlementPolicyOneofPruning(t *testing.T) {
	ctx := context.Background()
	m := &CustomAppEntitlementResourceModel{
		AppID:           types.StringValue("app-1"),
		ProvisionPolicy: &tfTypes.ProvisionPolicy{Connector: &tfTypes.ConnectorProvision{}},
		ProvisionerPolicy: &tfTypes.ProvisionPolicy{
			Manual: &tfTypes.ManualProvision{UserIds: []types.String{types.StringValue("stale")}},
		},
	}
	update, diags := m.ToSharedUpdateAppEntitlementRequest(ctx)
	if diags.HasError() {
		t.Fatalf("diagnostics: %v", diags)
	}
	b, err := json.Marshal(update.Entitlement)
	if err != nil {
		t.Fatalf("marshal: %s", err)
	}
	payload := string(b)
	if !strings.Contains(payload, `"connector"`) {
		t.Errorf("configured connector arm missing from the update payload: %s", payload)
	}
	for _, stale := range []string{"stale", `"manual"`} {
		if strings.Contains(payload, stale) {
			t.Errorf("stale arm or value %q leaked into the update payload: %s", stale, payload)
		}
	}
	if update.UpdateMask != nil {
		t.Errorf("update unexpectedly set a mask %q", *update.UpdateMask)
	}
}

// TestCustomAppEntitlementLegacyStateUpgrades drives the real
// UpgradeResourceState RPC with a state payload written by the released provider,
// which carries both policy attributes, and asserts the upgrade is lossless.
func TestCustomAppEntitlementLegacyStateUpgrades(t *testing.T) {
	ctx := context.Background()
	ps := providerserver.NewProtocol6(New("test")())()
	resp, err := ps.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %s", err)
	}
	const typeName = "conductorone_custom_app_entitlement"
	rs := resp.ResourceSchemas[typeName]
	if rs == nil {
		t.Fatal("schema missing")
	}

	values := map[string]any{"id": "ent-1", "app_id": "app-1", "display_name": "legacy"}
	for _, a := range rs.Block.Attributes {
		if _, ok := values[a.Name]; !ok {
			values[a.Name] = nil
		}
	}
	user := "2X4Vq2wLhQ9c8mVvTq3bNk7dRsp"
	values["provision_policy"] = map[string]any{"manual": map[string]any{"user_ids": []any{user}}}
	values["provisioner_policy"] = map[string]any{"manual": map[string]any{"user_ids": []any{user}}}
	raw, err := json.Marshal(values)
	if err != nil {
		t.Fatalf("marshal raw state: %s", err)
	}

	upgraded, err := ps.UpgradeResourceState(ctx, &tfprotov6.UpgradeResourceStateRequest{
		TypeName: typeName,
		Version:  rs.Version,
		RawState: &tfprotov6.RawState{JSON: raw},
	})
	if err != nil {
		t.Fatalf("UpgradeResourceState: %s", err)
	}
	for _, d := range upgraded.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("legacy state produced an error diagnostic: %s: %s", d.Summary, d.Detail)
		}
	}
	if upgraded.UpgradedState == nil {
		t.Fatal("no upgraded state returned")
	}
	value, err := upgraded.UpgradedState.Unmarshal(rs.ValueType())
	if err != nil {
		t.Fatalf("unmarshal upgraded state: %s", err)
	}
	var decoded map[string]tftypes.Value
	if err := value.As(&decoded); err != nil {
		t.Fatalf("decode upgraded state: %s", err)
	}
	for _, name := range []string{"provision_policy", "provisioner_policy"} {
		if _, ok := decoded[name]; !ok {
			t.Errorf("upgraded state lost the %s attribute", name)
		}
	}
}
