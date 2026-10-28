package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// configureCustomEntitlementPolicyAliases gives both names the same configurable
// schema. ConflictsWith inspects configuration, not the computed values in state.
func configureCustomEntitlementPolicyAliases(s *schema.Schema) diag.Diagnostics {
	canonical, ok := s.Attributes["provision_policy"].(schema.SingleNestedAttribute)
	if !ok {
		return diag.Diagnostics{diag.NewErrorDiagnostic("Invalid policy schema", "The generated provision_policy attribute must be a single nested attribute.")}
	}
	canonical.MarkdownDescription += "\n\nConfigure either `provision_policy` or `provisioner_policy`, but not both. Both refer to the same provisioning policy."
	canonical.PlanModifiers = append(canonical.PlanModifiers, objectplanmodifier.UseStateForUnknown())
	canonical.Validators = append(canonical.Validators, objectvalidator.ConflictsWith(path.MatchRoot("provisioner_policy")))
	s.Attributes["provision_policy"] = canonical
	alias := canonical
	alias.Validators = []validator.Object{objectvalidator.ConflictsWith(path.MatchRoot("provision_policy"))}
	alias.DeprecationMessage = ""
	s.Attributes["provisioner_policy"] = alias
	return nil
}

func selectedCustomPolicyAlias(ctx context.Context, config tfsdk.Config) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var a, b types.Object
	diags.Append(config.GetAttribute(ctx, path.Root("provision_policy"), &a)...)
	diags.Append(config.GetAttribute(ctx, path.Root("provisioner_policy"), &b)...)
	if diags.HasError() {
		return "", diags
	}
	if a.IsUnknown() || b.IsUnknown() {
		diags.AddError("Unknown policy alias", "The configured policy must be known before writing the entitlement.")
		return "", diags
	}
	if !a.IsNull() && !b.IsNull() {
		diags.AddError("Conflicting policy aliases", "Configure only one of provision_policy and provisioner_policy.")
		return "", diags
	}
	if !a.IsNull() {
		return "provision_policy", diags
	}
	if !b.IsNull() {
		return "provisioner_policy", diags
	}
	return "", diags
}

func selectCustomEntitlementPolicyForWrite(ctx context.Context, config tfsdk.Config, data *CustomAppEntitlementResourceModel) diag.Diagnostics {
	selected, diags := selectedCustomPolicyAlias(ctx, config)
	if !diags.HasError() && selected == "provisioner_policy" {
		data.ProvisionPolicy = data.ProvisionerPolicy
	}
	return diags
}

// Refresh the original alias when it exists in prior state; leave a missing
// alias unset so creating with the other name does not manufacture state.
func syncCustomEntitlementPolicyAliases(data *CustomAppEntitlementResourceModel) {
	if data != nil && data.ProvisionPolicy != nil {
		data.ProvisionPolicy = data.ProvisionerPolicy
	}
}
