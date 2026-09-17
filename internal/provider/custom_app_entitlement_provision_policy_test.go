package provider

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
)

// customAppEntitlementSchema builds the resource schema without a configured
// client. Schema() is pure and needs no provider configuration.
func customAppEntitlementSchema(t *testing.T) schema.Schema {
	t.Helper()
	r := &CustomAppEntitlementResource{}
	var resp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("schema diagnostics: %v", resp.Diagnostics)
	}
	return resp.Schema
}

// collectReplacePaths walks a nested schema subtree and returns the dotted
// paths of every node carrying a plan modifier whose description names
// replacement. IGA-4347 is exactly this: the writable provision policy subtree
// forced replacement on any configured change.
func collectReplacePaths(prefix string, attrs map[string]schema.Attribute, out *[]string) {
	for name, attr := range attrs {
		p := name
		if prefix != "" {
			p = prefix + "." + name
		}
		var descs []string
		switch a := attr.(type) {
		case schema.SingleNestedAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
			collectReplacePaths(p, a.Attributes, out)
		case schema.ListNestedAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
			collectReplacePaths(p, a.NestedObject.Attributes, out)
		case schema.SetNestedAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
			collectReplacePaths(p, a.NestedObject.Attributes, out)
		case schema.MapNestedAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
			collectReplacePaths(p, a.NestedObject.Attributes, out)
		case schema.StringAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
		case schema.BoolAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
		case schema.Int64Attribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
		case schema.Float64Attribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
		case schema.NumberAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
		case schema.ListAttribute:
			for _, pm := range a.PlanModifiers {
				descs = append(descs, pm.Description(context.Background()))
			}
		}
		for _, d := range descs {
			if strings.Contains(strings.ToLower(d), "replac") {
				*out = append(*out, p)
				break
			}
		}
	}
}

// TestCustomAppEntitlementProvisionPolicySchema locks the IGA-4347 schema
// contract: one writable provision_policy attribute that maps to the entity's
// provisionerPolicy for read/update, no replacement plan modifiers anywhere in
// its subtree, and no leftover computed provisioner_policy attribute.
func TestCustomAppEntitlementProvisionPolicySchema(t *testing.T) {
	s := customAppEntitlementSchema(t)

	raw, ok := s.Attributes["provision_policy"]
	if !ok {
		t.Fatal("provision_policy attribute missing from custom_app_entitlement schema")
	}
	if _, exists := s.Attributes["provisioner_policy"]; exists {
		t.Fatal("computed provisioner_policy attribute still present; the create/read split was not unified")
	}

	pp, ok := raw.(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("provision_policy is %T, want schema.SingleNestedAttribute", raw)
	}
	if !pp.Optional {
		t.Error("provision_policy must be Optional so existing HCL keeps working")
	}
	if !pp.Computed {
		t.Error("provision_policy must be Computed so API-returned defaults do not force a perpetual diff")
	}

	var replacePaths []string
	collectReplacePaths("provision_policy", pp.Attributes, &replacePaths)
	if len(replacePaths) != 0 {
		t.Errorf("provision_policy subtree must not force replacement; found %d replacement modifiers: %v", len(replacePaths), replacePaths)
	}
}

// TestCustomAppEntitlementProvisionPolicySerialization locks the create/read/
// update wire mapping for a configured policy:
//   - create serializes under the create-request wire name `provisionPolicy`
//   - update serializes the entity under the read/update wire name
//     `provisionerPolicy` (Go symbol ProvisionPolicy), never a Terraform alias
//   - read reconciles the API's ProvisionPolicy back into the same
//     provision_policy model field
func TestCustomAppEntitlementProvisionPolicySerialization(t *testing.T) {
	ctx := context.Background()

	m := &CustomAppEntitlementResourceModel{
		AppID: types.StringValue("app-1"),
		ProvisionPolicy: &tfTypes.ProvisionPolicy{
			Manual: &tfTypes.ManualProvision{
				UserIds: []types.String{types.StringValue("user-1")},
			},
		},
	}

	create, diags := m.ToSharedCreateAppEntitlementRequest(ctx)
	if diags.HasError() {
		t.Fatalf("create request diagnostics: %v", diags)
	}
	if create.ProvisionPolicy == nil {
		t.Fatal("create request did not carry the configured provision policy")
	}
	if create.ProvisionPolicy.Manual == nil || len(create.ProvisionPolicy.Manual.UserIds) != 1 || create.ProvisionPolicy.Manual.UserIds[0] != "user-1" {
		t.Fatalf("create policy not serialized from config: %+v", create.ProvisionPolicy.Manual)
	}
	createJSON, err := json.Marshal(create)
	if err != nil {
		t.Fatalf("marshal create request: %v", err)
	}
	if !strings.Contains(string(createJSON), `"provisionPolicy"`) {
		t.Errorf("create request must serialize the create wire name provisionPolicy; got %s", createJSON)
	}

	update, diags := m.ToSharedUpdateAppEntitlementRequest(ctx)
	if diags.HasError() {
		t.Fatalf("update request diagnostics: %v", diags)
	}
	if update.Entitlement == nil || update.Entitlement.ProvisionPolicy == nil {
		t.Fatal("update request did not carry the configured provision policy")
	}
	if update.Entitlement.ProvisionPolicy.Manual == nil || len(update.Entitlement.ProvisionPolicy.Manual.UserIds) != 1 || update.Entitlement.ProvisionPolicy.Manual.UserIds[0] != "user-1" {
		t.Fatalf("update policy not serialized from config: %+v", update.Entitlement.ProvisionPolicy.Manual)
	}
	updateJSON, err := json.Marshal(update.Entitlement)
	if err != nil {
		t.Fatalf("marshal update entitlement: %v", err)
	}
	if !strings.Contains(string(updateJSON), `"provisionerPolicy"`) {
		t.Errorf("update entitlement must serialize the entity wire name provisionerPolicy; got %s", updateJSON)
	}
	if strings.Contains(string(updateJSON), `"provisionPolicy"`) {
		t.Errorf("update entitlement leaked the create-only wire name provisionPolicy; got %s", updateJSON)
	}

	var read shared.AppEntitlement
	read.ProvisionPolicy = &shared.AppEntitlementProvisionPolicy{
		Manual: &shared.ManualProvision{UserIds: []string{"user-1"}},
	}
	var refreshed CustomAppEntitlementResourceModel
	if d := refreshed.RefreshFromSharedAppEntitlement(ctx, &read); d.HasError() {
		t.Fatalf("refresh diagnostics: %v", d)
	}
	if refreshed.ProvisionPolicy == nil || refreshed.ProvisionPolicy.Manual == nil ||
		len(refreshed.ProvisionPolicy.Manual.UserIds) != 1 || refreshed.ProvisionPolicy.Manual.UserIds[0].ValueString() != "user-1" {
		t.Fatalf("read did not reconcile API policy into provision_policy: %+v", refreshed.ProvisionPolicy)
	}
}

// TestCustomAppEntitlementProvisionPolicyOneofSwitch locks the oneof contract:
// a policy edit serializes exactly the configured arm and never a stale one,
// and the update carries no changed-field mask today (C1's maskCoversField
// treats an absent/empty mask as "write every supplied field", so the policy
// still persists; a future precise mask must name `provisioner_policy`).
func TestCustomAppEntitlementProvisionPolicyOneofSwitch(t *testing.T) {
	ctx := context.Background()

	m := &CustomAppEntitlementResourceModel{
		AppID: types.StringValue("app-1"),
		ProvisionPolicy: &tfTypes.ProvisionPolicy{
			Connector: &tfTypes.ConnectorProvision{},
		},
	}

	update, diags := m.ToSharedUpdateAppEntitlementRequest(ctx)
	if diags.HasError() {
		t.Fatalf("update diagnostics: %v", diags)
	}
	b, err := json.Marshal(update.Entitlement)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(b), `"connector"`) {
		t.Errorf("configured connector arm missing from update payload: %s", b)
	}
	for _, stale := range []string{`"manual"`, `"delegated"`, `"webhook"`, `"action"`} {
		if strings.Contains(string(b), stale) {
			t.Errorf("stale oneof arm %s serialized in update payload: %s", stale, b)
		}
	}
	if update.UpdateMask != nil {
		t.Errorf("update unexpectedly set a mask (%q); the IGA-4347 fix relies on the absent-mask write-all contract, and a precise mask would have to name provisioner_policy", *update.UpdateMask)
	}
}

// TestAppEntitlementProvisionerPolicyWireNamePreserved proves the overlay
// renamed only the Terraform/Go symbol, not the JSON wire contract. The entity
// read and update payloads must still use `provisionerPolicy`.
func TestAppEntitlementProvisionerPolicyWireNamePreserved(t *testing.T) {
	cases := map[string]any{
		"AppEntitlement":      shared.AppEntitlement{ProvisionPolicy: &shared.AppEntitlementProvisionPolicy{}},
		"AppEntitlementInput": shared.AppEntitlementInput{ProvisionPolicy: &shared.AppEntitlementProvisionPolicy{}},
	}
	for name, v := range cases {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("%s marshal: %v", name, err)
		}
		if !strings.Contains(string(b), `"provisionerPolicy"`) {
			t.Errorf("%s must keep the entity wire name provisionerPolicy; got %s", name, b)
		}
	}

	create := shared.CreateAppEntitlementRequest{ProvisionPolicy: &shared.ProvisionPolicy{}}
	b, err := json.Marshal(create)
	if err != nil {
		t.Fatalf("create marshal: %v", err)
	}
	if !strings.Contains(string(b), `"provisionPolicy"`) {
		t.Errorf("create request must keep the create wire name provisionPolicy; got %s", b)
	}
}
