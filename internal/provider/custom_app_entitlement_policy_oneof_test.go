package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
)

func emptyCustomPolicy() *tfTypes.ProvisionPolicy {
	return &tfTypes.ProvisionPolicy{MultiStep: jsontypes.NewNormalizedNull()}
}

func delegatedCustomPolicy() *tfTypes.ProvisionPolicy {
	return &tfTypes.ProvisionPolicy{
		Delegated: &tfTypes.DelegatedProvision{
			AppID:         types.StringValue("app2"),
			EntitlementID: types.StringValue("ent2"),
		},
		MultiStep: jsontypes.NewNormalizedNull(),
	}
}

func multiStepCustomPolicy() *tfTypes.ProvisionPolicy {
	return &tfTypes.ProvisionPolicy{MultiStep: jsontypes.NewNormalizedValue(multiStepJSON)}
}

func twoArmCustomPolicy() *tfTypes.ProvisionPolicy {
	p := delegatedCustomPolicy()
	p.MultiStep = jsontypes.NewNormalizedValue(multiStepJSON)
	return p
}

func delegatedCustomDeprovisionerPolicy() *tfTypes.CreateAppEntitlementRequestDeprovisionerPolicy {
	return &tfTypes.CreateAppEntitlementRequestDeprovisionerPolicy{
		Delegated: &tfTypes.DelegatedProvision{
			AppID:         types.StringValue("app2"),
			EntitlementID: types.StringValue("ent2"),
		},
		MultiStep: jsontypes.NewNormalizedNull(),
	}
}

func twoArmCustomDeprovisionerPolicy() *tfTypes.CreateAppEntitlementRequestDeprovisionerPolicy {
	p := delegatedCustomDeprovisionerPolicy()
	p.MultiStep = jsontypes.NewNormalizedValue(multiStepJSON)
	return p
}

func TestSelectedCustomAppEntitlementPolicyArm(t *testing.T) {
	tests := []struct {
		name    string
		policy  *tfTypes.ProvisionPolicy
		want    string
		wantErr bool
	}{
		{"nil", nil, "", false},
		{"zero arms", emptyCustomPolicy(), "", false},
		{"delegated", delegatedCustomPolicy(), "delegated", false},
		{"multi_step", multiStepCustomPolicy(), "multi_step", false},
		{"two arms", twoArmCustomPolicy(), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, diags := selectedOneofArm("provision_policy", customAppEntitlementPolicyArms(tt.policy))
			if tt.wantErr {
				if !diags.HasError() {
					t.Fatalf("expected error, got arm %q", got)
				}
				return
			}
			if diags.HasError() {
				t.Fatalf("unexpected error: %v", diags)
			}
			if got != tt.want {
				t.Fatalf("arm = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestKeepOnlyCustomAppEntitlementPolicyArm(t *testing.T) {
	p := twoArmCustomPolicy()
	keepOnlyOneofArm(customAppEntitlementPolicyArms(p), "delegated")
	if p.Delegated == nil {
		t.Fatal("delegated arm should be kept")
	}
	if !p.MultiStep.IsNull() {
		t.Fatal("multi_step arm should be nulled")
	}

	p = twoArmCustomPolicy()
	keepOnlyOneofArm(customAppEntitlementPolicyArms(p), "multi_step")
	if p.Delegated != nil {
		t.Fatal("delegated arm should be nil")
	}
	if p.MultiStep.IsNull() {
		t.Fatal("multi_step arm should be kept")
	}

	// empty keep leaves the policy untouched
	p = twoArmCustomPolicy()
	keepOnlyOneofArm(customAppEntitlementPolicyArms(p), "")
	if p.Delegated == nil || p.MultiStep.IsNull() {
		t.Fatal("empty keep should leave the policy untouched")
	}

	// nil policy is a no-op
	keepOnlyOneofArm(customAppEntitlementPolicyArms(nil), "delegated")
}

func testCustomConfig(t *testing.T, model *CustomAppEntitlementResourceModel) tfsdk.Config {
	t.Helper()
	ctx := context.Background()
	schemaResp := &resource.SchemaResponse{}
	NewCustomAppEntitlementResource().Schema(ctx, resource.SchemaRequest{}, schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatalf("schema: %v", schemaResp.Diagnostics)
	}
	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(ctx, model); diags.HasError() {
		t.Fatalf("encoding config: %v", diags)
	}
	return tfsdk.Config{Schema: schemaResp.Schema, Raw: state.Raw}
}

func TestPruneCustomAppEntitlementPolicyOneofs(t *testing.T) {
	// merged model carries both arms (the bug), config selects delegated only.
	merged := &CustomAppEntitlementResourceModel{
		AppID:               types.StringValue("app1"),
		ID:                  types.StringValue("ent1"),
		DisplayName:         types.StringValue("E"),
		ProvisionPolicy:     twoArmCustomPolicy(),
		DeprovisionerPolicy: twoArmCustomDeprovisionerPolicy(),
	}
	cfg := testCustomConfig(t, &CustomAppEntitlementResourceModel{
		AppID:               types.StringValue("app1"),
		ID:                  types.StringValue("ent1"),
		DisplayName:         types.StringValue("E"),
		ProvisionPolicy:     delegatedCustomPolicy(),
		DeprovisionerPolicy: delegatedCustomDeprovisionerPolicy(),
	})

	diags := pruneCustomAppEntitlementPolicyOneofs(context.Background(), cfg, merged)
	if diags.HasError() {
		t.Fatalf("prune: %v", diags)
	}
	if merged.ProvisionPolicy.Delegated == nil || !merged.ProvisionPolicy.MultiStep.IsNull() {
		t.Fatal("provision_policy should keep only delegated")
	}
	if merged.DeprovisionerPolicy.Delegated == nil || !merged.DeprovisionerPolicy.MultiStep.IsNull() {
		t.Fatal("deprovisioner_policy should keep only delegated")
	}
}
