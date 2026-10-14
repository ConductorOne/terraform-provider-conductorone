package provider

import (
	"context"
	"encoding/json"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
)

const multiStepJSON = `{"provisionSteps":[{"delegated":{"appId":"app2","entitlementId":"ent2"}}]}`

func delegatedPolicy() *tfTypes.AppEntitlementProvisionPolicy {
	return &tfTypes.AppEntitlementProvisionPolicy{
		DelegatedProvision: &tfTypes.DelegatedProvision{
			AppID:         types.StringValue("app2"),
			EntitlementID: types.StringValue("ent2"),
		},
		MultiStep: jsontypes.NewNormalizedNull(),
	}
}

func multiStepPolicy() *tfTypes.AppEntitlementProvisionPolicy {
	return &tfTypes.AppEntitlementProvisionPolicy{MultiStep: jsontypes.NewNormalizedValue(multiStepJSON)}
}

// testConfig renders model as a tfsdk.Config against the real resource schema.
func testConfig(t *testing.T, model *AppEntitlementResourceModel) tfsdk.Config {
	t.Helper()
	ctx := context.Background()
	schemaResp := &resource.SchemaResponse{}
	NewAppEntitlementResource().Schema(ctx, resource.SchemaRequest{}, schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatalf("schema: %v", schemaResp.Diagnostics)
	}
	state := tfsdk.State{Schema: schemaResp.Schema}
	state.Raw = state.Raw.Copy()
	if diags := state.Set(ctx, model); diags.HasError() {
		t.Fatalf("encoding config: %v", diags)
	}
	return tfsdk.Config{Schema: schemaResp.Schema, Raw: state.Raw}
}

// newModel returns a model with the scalar fields ToUpdateSDKType reads left null.
func newModel(provision, deprovisioner *tfTypes.AppEntitlementProvisionPolicy) *AppEntitlementResourceModel {
	return &AppEntitlementResourceModel{
		AppID:               types.StringValue("app1"),
		ID:                  types.StringValue("ent1"),
		DisplayName:         types.StringValue("E"),
		ProvisionPolicy:     provision,
		DeprovisionerPolicy: deprovisioner,
	}
}

// requestArms returns the sorted oneof arms present in the serialized update
// entitlement for the policy under key.
func requestArms(t *testing.T, m *AppEntitlementResourceModel, key string) []string {
	t.Helper()
	body, err := json.Marshal(m.ToUpdateSDKType())
	if err != nil {
		t.Fatal(err)
	}
	var entitlement map[string]json.RawMessage
	if err := json.Unmarshal(body, &entitlement); err != nil {
		t.Fatal(err)
	}
	raw, ok := entitlement[key]
	if !ok {
		return nil
	}
	var policy map[string]json.RawMessage
	if err := json.Unmarshal(raw, &policy); err != nil {
		t.Fatal(err)
	}
	arms := make([]string, 0, len(policy))
	for arm := range policy {
		arms = append(arms, arm)
	}
	sort.Strings(arms)
	return arms
}

// The request serializes the provision policy as provisionerPolicy.
const (
	provisionKey     = "provisionerPolicy"
	deprovisionerKey = "deprovisionerPolicy"
)

func assertArms(t *testing.T, m *AppEntitlementResourceModel, key string, want ...string) {
	t.Helper()
	got := requestArms(t, m, key)
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("%s arms = %v, want %v", key, got, want)
	}
}

func prune(t *testing.T, cfg, merged *AppEntitlementResourceModel) {
	t.Helper()
	diags := pruneAppEntitlementPolicyOneofs(context.Background(), testConfig(t, cfg), merged)
	if diags.HasError() {
		t.Fatalf("prune: %v", diags)
	}
}

// merge() seeds the model from prior state, so the replaced arm sits next to the
// new one. Each case starts from that two-arm model.
func TestPruneAppEntitlementPolicyOneofsSwitchesArm(t *testing.T) {
	tests := []struct {
		name     string
		cfg      func() *tfTypes.AppEntitlementProvisionPolicy
		wantArms []string
	}{
		{"delegated to multi_step", multiStepPolicy, []string{"multiStep"}},
		{"multi_step to delegated", delegatedPolicy, []string{"delegated"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twoArms := func() *tfTypes.AppEntitlementProvisionPolicy {
				p := delegatedPolicy()
				p.MultiStep = jsontypes.NewNormalizedValue(multiStepJSON)
				return p
			}
			merged := newModel(twoArms(), twoArms())

			// Before pruning the request carries both arms: the bug.
			assertArms(t, merged, provisionKey, "delegated", "multiStep")
			assertArms(t, merged, deprovisionerKey, "delegated", "multiStep")

			prune(t, newModel(tt.cfg(), tt.cfg()), merged)

			assertArms(t, merged, provisionKey, tt.wantArms...)
			assertArms(t, merged, deprovisionerKey, tt.wantArms...)
		})
	}
}

// Prune is per policy: switching one must not touch the other.
func TestPruneAppEntitlementPolicyOneofsPoliciesIndependent(t *testing.T) {
	twoArms := func() *tfTypes.AppEntitlementProvisionPolicy {
		p := delegatedPolicy()
		p.MultiStep = jsontypes.NewNormalizedValue(multiStepJSON)
		return p
	}
	merged := newModel(twoArms(), delegatedPolicy())
	prune(t, newModel(multiStepPolicy(), nil), merged)

	assertArms(t, merged, provisionKey, "multiStep")
	assertArms(t, merged, deprovisionerKey, "delegated")
}

// An omitted policy block in config means nothing was chosen, so the stored arm
// is what the API already has and must still be sent.
func TestPruneAppEntitlementPolicyOneofsOmittedBlockKeepsStoredArm(t *testing.T) {
	merged := newModel(delegatedPolicy(), multiStepPolicy())
	prune(t, newModel(nil, nil), merged)

	assertArms(t, merged, provisionKey, "delegated")
	assertArms(t, merged, deprovisionerKey, "multiStep")
}

func TestPruneAppEntitlementPolicyOneofsRejectsTwoArmsInConfig(t *testing.T) {
	both := delegatedPolicy()
	both.MultiStep = jsontypes.NewNormalizedValue(multiStepJSON)

	merged := newModel(delegatedPolicy(), nil)
	diags := pruneAppEntitlementPolicyOneofs(context.Background(), testConfig(t, newModel(both, nil)), merged)
	if !diags.HasError() {
		t.Fatal("expected an error for two arms in config")
	}
	detail := diags.Errors()[0].Detail()
	for _, want := range []string{"delegated_provision", "multi_step"} {
		if !strings.Contains(detail, want) {
			t.Errorf("error detail %q does not name %s", detail, want)
		}
	}
}
