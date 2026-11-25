package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
)

// pruneCustomAppEntitlementPolicyOneofs narrows provision_policy,
// deprovisioner_policy, and the duration_grant/duration_unset pair on data to
// the single arm set in configuration.
//
// The deprovisioner field uses the generated CreateAppEntitlementRequestDeprovisionerPolicy
// type, which is structurally identical to ProvisionPolicy; the conversion below
// fails to compile if a future regen makes the two structs diverge.
func pruneCustomAppEntitlementPolicyOneofs(ctx context.Context, config tfsdk.Config, data *CustomAppEntitlementResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if data == nil {
		return diags
	}

	var configured *CustomAppEntitlementResourceModel
	diags.Append(config.Get(ctx, &configured)...)
	if diags.HasError() || configured == nil {
		return diags
	}

	provision, provisionDiags := selectedOneofArm("provision_policy", customAppEntitlementPolicyArms(configured.ProvisionPolicy))
	diags.Append(provisionDiags...)
	deprovisioner, deprovisionerDiags := selectedOneofArm("deprovisioner_policy",
		customAppEntitlementPolicyArms((*tfTypes.ProvisionPolicy)(configured.DeprovisionerPolicy)))
	diags.Append(deprovisionerDiags...)
	if diags.HasError() {
		return diags
	}

	keepOnlyOneofArm(customAppEntitlementPolicyArms(data.ProvisionPolicy), provision)
	keepOnlyOneofArm(customAppEntitlementPolicyArms((*tfTypes.ProvisionPolicy)(data.DeprovisionerPolicy)), deprovisioner)
	pruneCustomAppEntitlementDurationOneof(configured, data)
	return diags
}

// pruneCustomAppEntitlementDurationOneof drops the duration arm that config
// does not set. duration_unset is Computed (the API always returns {} when no
// grant duration is set), so merge() copies a prior {} next to a newly
// configured duration_grant and the API rejects the body with
// "oneof max_grant_duration is already set". When config sets neither arm,
// both are left as merged so an unchanged duration is not cleared.
func pruneCustomAppEntitlementDurationOneof(configured, data *CustomAppEntitlementResourceModel) {
	if !configured.DurationGrant.IsNull() && !configured.DurationGrant.IsUnknown() {
		data.DurationUnset = nil
	}
	if configured.DurationUnset != nil {
		data.DurationGrant = types.StringNull()
	}
}

// customAppEntitlementPolicyArms lists the arms of p by Terraform attribute name.
// A nil p has no arms.
func customAppEntitlementPolicyArms(p *tfTypes.ProvisionPolicy) []oneofArm {
	if p == nil {
		return nil
	}
	return []oneofArm{
		{"action", func() bool { return p.Action != nil }, func() { p.Action = nil }},
		{"connector", func() bool { return p.Connector != nil }, func() { p.Connector = nil }},
		{"delegated", func() bool { return p.Delegated != nil }, func() { p.Delegated = nil }},
		{"device_placement", func() bool { return p.DevicePlacement != nil }, func() { p.DevicePlacement = nil }},
		{"external_ticket", func() bool { return p.ExternalTicket != nil }, func() { p.ExternalTicket = nil }},
		{"manual", func() bool { return p.Manual != nil }, func() { p.Manual = nil }},
		{"multi_step", func() bool { return !p.MultiStep.IsNull() && !p.MultiStep.IsUnknown() }, func() { p.MultiStep = jsontypes.NewNormalizedNull() }},
		{"unconfigured", func() bool { return p.Unconfigured != nil }, func() { p.Unconfigured = nil }},
		{"webhook", func() bool { return p.Webhook != nil }, func() { p.Webhook = nil }},
	}
}
