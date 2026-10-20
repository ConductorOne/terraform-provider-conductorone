package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
)

// pruneAppEntitlementPolicyOneofs narrows provision_policy and
// deprovisioner_policy on data to the single arm set in configuration.
//
// Configuration is the only input where an arm the practitioner did not write is
// reliably null: every arm is Computed, so in the plan an unset arm holds the
// prior state value or is unknown. Dropping Computed from the arms would also stop
// the stale arm, but a config that omits the policy block could then no longer
// carry the stored value forward. A config naming no arm leaves data untouched.
func pruneAppEntitlementPolicyOneofs(ctx context.Context, config tfsdk.Config, data *AppEntitlementResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if data == nil {
		return diags
	}

	var configured *AppEntitlementResourceModel
	diags.Append(config.Get(ctx, &configured)...)
	if diags.HasError() || configured == nil {
		return diags
	}

	provision, provisionDiags := selectedOneofArm("provision_policy", appEntitlementPolicyArms(configured.ProvisionPolicy))
	diags.Append(provisionDiags...)
	deprovisioner, deprovisionerDiags := selectedOneofArm("deprovisioner_policy", appEntitlementPolicyArms(configured.DeprovisionerPolicy))
	diags.Append(deprovisionerDiags...)
	if diags.HasError() {
		return diags
	}

	keepOnlyOneofArm(appEntitlementPolicyArms(data.ProvisionPolicy), provision)
	keepOnlyOneofArm(appEntitlementPolicyArms(data.DeprovisionerPolicy), deprovisioner)
	return diags
}

// appEntitlementPolicyArms lists the arms of p by Terraform attribute name.
// A nil p has no arms.
func appEntitlementPolicyArms(p *tfTypes.AppEntitlementProvisionPolicy) []oneofArm {
	if p == nil {
		return nil
	}
	return []oneofArm{
		{"action_provision", func() bool { return p.ActionProvision != nil }, func() { p.ActionProvision = nil }},
		{"connector_provision", func() bool { return p.ConnectorProvision != nil }, func() { p.ConnectorProvision = nil }},
		{"delegated_provision", func() bool { return p.DelegatedProvision != nil }, func() { p.DelegatedProvision = nil }},
		{"external_ticket_provision", func() bool { return p.ExternalTicketProvision != nil }, func() { p.ExternalTicketProvision = nil }},
		{"manual_provision", func() bool { return p.ManualProvision != nil }, func() { p.ManualProvision = nil }},
		{"multi_step", func() bool { return !p.MultiStep.IsNull() && !p.MultiStep.IsUnknown() }, func() { p.MultiStep = jsontypes.NewNormalizedNull() }},
		{"unconfigured_provision", func() bool { return p.UnconfiguredProvision != nil }, func() { p.UnconfiguredProvision = nil }},
		{"webhook_provision", func() bool { return p.WebhookProvision != nil }, func() { p.WebhookProvision = nil }},
	}
}
