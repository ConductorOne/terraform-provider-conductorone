package provider

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
)

// TestPolicyOneofTypesStayInSync guards the deliberate duplication in
// provision_policy_oneof.go. The prune helpers are written twice because
// ProvisionPolicy and DeprovisionerPolicy are distinct structs with identical
// arms; if the spec ever adds or renames an arm on one, this fails and the four
// functions need the same edit.
func TestPolicyOneofTypesStayInSync(t *testing.T) {
	provision := reflect.TypeOf(tfTypes.ProvisionPolicy{})
	deprovisioner := reflect.TypeOf(tfTypes.DeprovisionerPolicy{})

	if got, want := deprovisioner.NumField(), provision.NumField(); got != want {
		t.Fatalf("DeprovisionerPolicy has %d fields, ProvisionPolicy has %d", got, want)
	}
	for i := 0; i < provision.NumField(); i++ {
		p, d := provision.Field(i), deprovisioner.Field(i)
		if p.Name != d.Name {
			t.Errorf("field %d: ProvisionPolicy has %q, DeprovisionerPolicy has %q", i, p.Name, d.Name)
		}
		if p.Tag.Get("tfsdk") != d.Tag.Get("tfsdk") {
			t.Errorf("field %s: tfsdk tag %q vs %q", p.Name, p.Tag.Get("tfsdk"), d.Tag.Get("tfsdk"))
		}
	}
}

func TestKeepOnlyProvisionPolicyBranchDropsTheStaleArm(t *testing.T) {
	// The IGA-3483 shape: state contributed unconfigured, config selected manual,
	// and merge() left both on the model.
	merged := &tfTypes.ProvisionPolicy{
		ManualProvision:       &tfTypes.ManualProvision{},
		UnconfiguredProvision: &tfTypes.UnconfiguredProvision{},
	}
	configured := &tfTypes.ProvisionPolicy{ManualProvision: &tfTypes.ManualProvision{}}

	branch, diags := selectedProvisionPolicyBranch("provision_policy", configured)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	keepOnlyProvisionPolicyBranch(merged, branch)

	if merged.ManualProvision == nil {
		t.Error("configured arm was dropped")
	}
	if merged.UnconfiguredProvision != nil {
		t.Error("stale arm survived; the request would carry two oneof arms and the API returns 400")
	}
}

func TestKeepOnlyDeprovisionerPolicyBranchDropsTheStaleArm(t *testing.T) {
	merged := &tfTypes.DeprovisionerPolicy{
		ConnectorProvision:    &tfTypes.ConnectorProvision{},
		UnconfiguredProvision: &tfTypes.UnconfiguredProvision{},
	}
	configured := &tfTypes.DeprovisionerPolicy{ConnectorProvision: &tfTypes.ConnectorProvision{}}

	branch, diags := selectedDeprovisionerPolicyBranch("deprovisioner_policy", configured)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	keepOnlyDeprovisionerPolicyBranch(merged, branch)

	if merged.ConnectorProvision == nil {
		t.Error("configured arm was dropped")
	}
	if merged.UnconfiguredProvision != nil {
		t.Error("stale arm survived")
	}
}

func TestEmptyConfigSelectionLeavesTheStoredArmAlone(t *testing.T) {
	// Config omits the policy block, so there is nothing to choose: whatever
	// single arm the remote already had is what should be sent back.
	merged := &tfTypes.ProvisionPolicy{UnconfiguredProvision: &tfTypes.UnconfiguredProvision{}}

	branch, diags := selectedProvisionPolicyBranch("provision_policy", nil)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	keepOnlyProvisionPolicyBranch(merged, branch)

	if merged.UnconfiguredProvision == nil {
		t.Error("stored arm was cleared even though the practitioner configured nothing")
	}
}

func TestConfiguringTwoArmsIsAnErrorRatherThanASilentChoice(t *testing.T) {
	// The API models the policy as a oneof, so a body carrying two arms is
	// rejected. Keeping the first arm in source order would instead apply one
	// half of the configuration and drop the other without a plan showing it.
	configured := &tfTypes.DeprovisionerPolicy{
		ActionProvision: &tfTypes.ActionProvision{},
		ManualProvision: &tfTypes.ManualProvision{},
	}

	branch, diags := selectedDeprovisionerPolicyBranch("deprovisioner_policy", configured)

	if !diags.HasError() {
		t.Fatal("expected an error for a configuration setting two arms")
	}
	if branch != provisionBranchNone {
		t.Errorf("branch = %q, want none: an ambiguous configuration must not select an arm", branch)
	}
	detail := diags.Errors()[0].Detail()
	for _, arm := range []string{"action_provision", "manual_provision"} {
		if !strings.Contains(detail, arm) {
			t.Errorf("error detail %q does not name %q", detail, arm)
		}
	}
}

func TestProvisionPolicyReportsEveryConfiguredArm(t *testing.T) {
	configured := &tfTypes.ProvisionPolicy{
		ConnectorProvision:    &tfTypes.ConnectorProvision{},
		UnconfiguredProvision: &tfTypes.UnconfiguredProvision{},
		WebhookProvision:      &tfTypes.WebhookProvision{},
	}

	_, diags := selectedProvisionPolicyBranch("provision_policy", configured)

	if !diags.HasError() {
		t.Fatal("expected an error for a configuration setting three arms")
	}
	detail := diags.Errors()[0].Detail()
	for _, arm := range []string{"connector_provision", "unconfigured_provision", "webhook_provision"} {
		if !strings.Contains(detail, arm) {
			t.Errorf("error detail %q does not name %q", detail, arm)
		}
	}
}

// TestPolicyArmsConflictWithEveryOtherArm covers the schema half of the same
// defect: without a complete ConflictsWith set, Terraform accepts a
// configuration carrying two arms and Create serializes both.
func TestPolicyArmsConflictWithEveryOtherArm(t *testing.T) {
	ctx := context.Background()
	appEntitlement := &resource.SchemaResponse{}
	(&AppEntitlementResource{}).Schema(ctx, resource.SchemaRequest{}, appEntitlement)
	customAppEntitlement := &resource.SchemaResponse{}
	(&CustomAppEntitlementResource{}).Schema(ctx, resource.SchemaRequest{}, customAppEntitlement)

	resources := map[string]schema.Schema{
		"conductorone_app_entitlement":        appEntitlement.Schema,
		"conductorone_custom_app_entitlement": customAppEntitlement.Schema,
	}
	for resourceName, resourceSchema := range resources {
		for _, policyName := range []string{"provision_policy", "deprovisioner_policy"} {
			policy, ok := resourceSchema.Attributes[policyName].(schema.SingleNestedAttribute)
			if !ok {
				t.Fatalf("%s.%s is %T, want schema.SingleNestedAttribute", resourceName, policyName, resourceSchema.Attributes[policyName])
			}
			for name, attribute := range policy.Attributes {
				described := describeValidators(t, attribute)
				for sibling := range policy.Attributes {
					if sibling == name {
						continue
					}
					if !strings.Contains(described, sibling) {
						t.Errorf("%s.%s.%s does not conflict with %s: a configuration setting both is accepted",
							resourceName, policyName, name, sibling)
					}
				}
			}
		}
	}
}

func describeValidators(t *testing.T, attribute schema.Attribute) string {
	t.Helper()
	ctx := context.Background()
	var described []string
	switch typed := attribute.(type) {
	case schema.SingleNestedAttribute:
		for _, v := range typed.Validators {
			described = append(described, v.Description(ctx))
		}
	case schema.StringAttribute:
		for _, v := range typed.Validators {
			described = append(described, v.Description(ctx))
		}
	default:
		t.Fatalf("unhandled attribute type %T", attribute)
	}
	return strings.Join(described, "\n")
}
