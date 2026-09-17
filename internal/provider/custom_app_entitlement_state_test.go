package provider

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// legacyProviderStateJSON builds a raw state payload shaped like state written
// by the released provider (v1.6.0), which carried BOTH provision_policy (the
// configured create-only attribute) and provisioner_policy (the computed
// read/update attribute). The IGA-4347 fix removes the computed
// provisioner_policy attribute by unifying it into provision_policy, so the
// upgrade path must drop the stale key rather than fail.
//
// The attribute names are read from the CURRENT schema so the fixture stays
// valid as the entity grows; the removed key is added explicitly.
func legacyProviderStateJSON(t *testing.T, schema *tfprotov6.Schema) []byte {
	t.Helper()

	values := map[string]any{
		"id":           "2X4Vq2wLhQ9c8mVvTq3bNk7dRsp",
		"app_id":       "2X4Vq2wLhQ9c8mVvTq3bNk7dRsp",
		"display_name": "legacy entitlement",
	}
	for _, attr := range schema.Block.Attributes {
		if _, ok := values[attr.Name]; !ok {
			values[attr.Name] = nil
		}
	}
	// The configured policy as written by v1.6.0.
	values["provision_policy"] = map[string]any{
		"manual_provision": map[string]any{"user_ids": []any{"2X4Vq2wLhQ9c8mVvTq3bNk7dRsp"}},
	}
	// The computed attribute that the fix removes.
	values["provisioner_policy"] = map[string]any{
		"manual": map[string]any{"user_ids": []any{"server-managed-value"}},
	}

	raw, err := json.Marshal(values)
	if err != nil {
		t.Fatalf("marshal legacy state: %s", err)
	}
	return raw
}

// TestCustomAppEntitlementLegacyStateUpgrades proves the IGA-4347 schema change
// does not brick previously-written state: a state payload that still contains
// the removed computed provisioner_policy attribute upgrades without error, and
// the upgraded state carries only the unified provision_policy attribute.
//
// This exercises the framework's real passthrough upgrade path
// (terraform-plugin-framework internal/fwserver/server_upgraderesourcestate.go),
// which unmarshals prior state with IgnoreUndefinedAttributes, so a removed
// attribute is dropped silently instead of producing a diagnostic.
func TestCustomAppEntitlementLegacyStateUpgrades(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	ps := providerserver.NewProtocol6(New("test")())()

	providerSchema, err := ps.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %s", err)
	}
	if len(providerSchema.Diagnostics) > 0 {
		t.Fatalf("GetProviderSchema diagnostics: %v", providerSchema.Diagnostics)
	}

	const typeName = "conductorone_custom_app_entitlement"
	rs, ok := providerSchema.ResourceSchemas[typeName]
	if !ok {
		t.Fatalf("resource schema %q not found", typeName)
	}

	resp, err := ps.UpgradeResourceState(ctx, &tfprotov6.UpgradeResourceStateRequest{
		TypeName: typeName,
		Version:  rs.Version,
		RawState: &tfprotov6.RawState{JSON: legacyProviderStateJSON(t, rs)},
	})
	if err != nil {
		t.Fatalf("UpgradeResourceState: %s", err)
	}
	for _, d := range resp.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("upgrading legacy state produced an error diagnostic: %s: %s", d.Summary, d.Detail)
		}
	}
	if resp.UpgradedState == nil {
		t.Fatalf("UpgradeResourceState returned no upgraded state (json=%d msgpack=%d)", 0, 0)
	}

	schemaType := rs.ValueType()

	upgradedValue, err := resp.UpgradedState.Unmarshal(schemaType)
	if err != nil {
		t.Fatalf("unmarshal upgraded state (json=%d bytes, msgpack=%d bytes): %s", len(resp.UpgradedState.JSON), len(resp.UpgradedState.MsgPack), err)
	}

	var upgraded map[string]tftypes.Value
	if err := upgradedValue.As(&upgraded); err != nil {
		t.Fatalf("decode upgraded state object: %s", err)
	}
	if _, present := upgraded["provisioner_policy"]; present {
		t.Error("upgraded state still contains the removed computed provisioner_policy attribute")
	}
	if _, present := upgraded["provision_policy"]; !present {
		t.Error("upgraded state lost the unified provision_policy attribute")
	}
}
