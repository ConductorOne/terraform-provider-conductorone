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
	policyArmNone           policyArm = ""
	policyArmAction         policyArm = "action_provision"
	policyArmConnector      policyArm = "connector_provision"
	policyArmDelegated      policyArm = "delegated_provision"
	policyArmExternalTicket policyArm = "external_ticket_provision"
	policyArmManual         policyArm = "manual_provision"
	policyArmMultiStep      policyArm = "multi_step"
	policyArmUnconfigured   policyArm = "unconfigured_provision"
	policyArmWebhook        policyArm = "webhook_provision"
)

// pruneAppEntitlementPolicyOneofs narrows provision_policy and
// deprovisioner_policy on data to the single arm set in configuration.
//
// The API models each policy as a protobuf oneof and rejects a body carrying two
// arms. The generated SDK models every arm as an independent optional field, and
// merge() loads prior state before overlaying the plan, so the arm being replaced
// survives next to its replacement and both are serialized.
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

	provision, provisionDiags := selectedPolicyArm("provision_policy", configured.ProvisionPolicy)
	diags.Append(provisionDiags...)
	deprovisioner, deprovisionerDiags := selectedPolicyArm("deprovisioner_policy", configured.DeprovisionerPolicy)
	diags.Append(deprovisionerDiags...)
	if diags.HasError() {
		return diags
	}

	keepOnlyPolicyArm(data.ProvisionPolicy, provision)
	keepOnlyPolicyArm(data.DeprovisionerPolicy, deprovisioner)
	return diags
}

// selectedPolicyArm returns the one arm set on p. More than one is an error:
// choosing among them would apply a policy the practitioner did not write.
func selectedPolicyArm(attribute string, p *tfTypes.AppEntitlementProvisionPolicy) (policyArm, diag.Diagnostics) {
	var diags diag.Diagnostics
	if p == nil {
		return policyArmNone, diags
	}

	var set []policyArm
	if p.ActionProvision != nil {
		set = append(set, policyArmAction)
	}
	if p.ConnectorProvision != nil {
		set = append(set, policyArmConnector)
	}
	if p.DelegatedProvision != nil {
		set = append(set, policyArmDelegated)
	}
	if p.ExternalTicketProvision != nil {
		set = append(set, policyArmExternalTicket)
	}
	if p.ManualProvision != nil {
		set = append(set, policyArmManual)
	}
	if !p.MultiStep.IsNull() && !p.MultiStep.IsUnknown() {
		set = append(set, policyArmMultiStep)
	}
	if p.UnconfiguredProvision != nil {
		set = append(set, policyArmUnconfigured)
	}
	if p.WebhookProvision != nil {
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

func keepOnlyPolicyArm(p *tfTypes.AppEntitlementProvisionPolicy, keep policyArm) {
	if p == nil || keep == policyArmNone {
		return
	}
	if keep != policyArmAction {
		p.ActionProvision = nil
	}
	if keep != policyArmConnector {
		p.ConnectorProvision = nil
	}
	if keep != policyArmDelegated {
		p.DelegatedProvision = nil
	}
	if keep != policyArmExternalTicket {
		p.ExternalTicketProvision = nil
	}
	if keep != policyArmManual {
		p.ManualProvision = nil
	}
	if keep != policyArmMultiStep {
		p.MultiStep = jsontypes.NewNormalizedNull()
	}
	if keep != policyArmUnconfigured {
		p.UnconfiguredProvision = nil
	}
	if keep != policyArmWebhook {
		p.WebhookProvision = nil
	}
}
