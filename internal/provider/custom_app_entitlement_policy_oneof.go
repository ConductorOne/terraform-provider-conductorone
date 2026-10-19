package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
)

// policyArm names one arm of the policy oneof by its Terraform attribute name.
type policyArm string

const (
	policyArmNone            policyArm = ""
	policyArmAction          policyArm = "action"
	policyArmConnector       policyArm = "connector"
	policyArmDelegated       policyArm = "delegated"
	policyArmDevicePlacement policyArm = "device_placement"
	policyArmExternalTicket  policyArm = "external_ticket"
	policyArmManual          policyArm = "manual"
	policyArmMultiStep       policyArm = "multi_step"
	policyArmUnconfigured    policyArm = "unconfigured"
	policyArmWebhook         policyArm = "webhook"
)

// pruneCustomAppEntitlementPolicyOneofs narrows provision_policy and
// deprovisioner_policy on data to the single arm set in configuration.
//
// The generated custom resource models each policy as a oneof whose arms are
// independent optional fields, and merge() loads prior state before overlaying
// the plan, so the arm being replaced survives next to its replacement and both
// are serialized. The API rejects a body carrying two arms.
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

	provision, provisionDiags := selectedPolicyArm("provision_policy", configured.ProvisionPolicy)
	diags.Append(provisionDiags...)
	deprovisioner, deprovisionerDiags := selectedPolicyArm("deprovisioner_policy", (*tfTypes.ProvisionPolicy)(configured.DeprovisionerPolicy))
	diags.Append(deprovisionerDiags...)
	if diags.HasError() {
		return diags
	}

	keepOnlyPolicyArm(data.ProvisionPolicy, provision)
	keepOnlyPolicyArm((*tfTypes.ProvisionPolicy)(data.DeprovisionerPolicy), deprovisioner)
	return diags
}

// selectedPolicyArm returns the one arm set on p. More than one is an error:
// choosing among them would apply a policy the practitioner did not write.
func selectedPolicyArm(attribute string, p *tfTypes.ProvisionPolicy) (policyArm, diag.Diagnostics) {
	var diags diag.Diagnostics
	if p == nil {
		return policyArmNone, diags
	}

	var set []policyArm
	if p.Action != nil {
		set = append(set, policyArmAction)
	}
	if p.Connector != nil {
		set = append(set, policyArmConnector)
	}
	if p.Delegated != nil {
		set = append(set, policyArmDelegated)
	}
	if p.DevicePlacement != nil {
		set = append(set, policyArmDevicePlacement)
	}
	if p.ExternalTicket != nil {
		set = append(set, policyArmExternalTicket)
	}
	if p.Manual != nil {
		set = append(set, policyArmManual)
	}
	if !p.MultiStep.IsNull() && !p.MultiStep.IsUnknown() {
		set = append(set, policyArmMultiStep)
	}
	if p.Unconfigured != nil {
		set = append(set, policyArmUnconfigured)
	}
	if p.Webhook != nil {
		set = append(set, policyArmWebhook)
	}

	switch len(set) {
	case 0:
		return policyArmNone, diags
	case 1:
		return set[0], diags
	}

	names := make([]string, 0, len(set))
	for _, arm := range set {
		names = append(names, string(arm))
	}
	diags.AddError(
		fmt.Sprintf("Conflicting %s configuration", attribute),
		fmt.Sprintf(
			"Only one arm may be configured, but the configuration sets %s. The API models this as a oneof and rejects a body carrying more than one arm.",
			strings.Join(names, ", "),
		),
	)
	return policyArmNone, diags
}

// keepOnlyPolicyArm nils every arm of p except keep. A nil p or an empty keep
// leaves p untouched (a config naming no arm keeps the stored arm).
func keepOnlyPolicyArm(p *tfTypes.ProvisionPolicy, keep policyArm) {
	if p == nil || keep == policyArmNone {
		return
	}
	if keep != policyArmAction {
		p.Action = nil
	}
	if keep != policyArmConnector {
		p.Connector = nil
	}
	if keep != policyArmDelegated {
		p.Delegated = nil
	}
	if keep != policyArmDevicePlacement {
		p.DevicePlacement = nil
	}
	if keep != policyArmExternalTicket {
		p.ExternalTicket = nil
	}
	if keep != policyArmManual {
		p.Manual = nil
	}
	if keep != policyArmMultiStep {
		p.MultiStep = jsontypes.NewNormalizedNull()
	}
	if keep != policyArmUnconfigured {
		p.Unconfigured = nil
	}
	if keep != policyArmWebhook {
		p.Webhook = nil
	}
}
