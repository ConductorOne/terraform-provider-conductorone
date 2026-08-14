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

// provisionBranch identifies one arm of the ProvisionPolicy oneof by its
// Terraform attribute name.
type provisionBranch string

const (
	provisionBranchNone           provisionBranch = ""
	provisionBranchAction         provisionBranch = "action_provision"
	provisionBranchConnector      provisionBranch = "connector_provision"
	provisionBranchDelegated      provisionBranch = "delegated_provision"
	provisionBranchExternalTicket provisionBranch = "external_ticket_provision"
	provisionBranchManual         provisionBranch = "manual_provision"
	provisionBranchMultiStep      provisionBranch = "multi_step"
	provisionBranchUnconfigured   provisionBranch = "unconfigured_provision"
	provisionBranchWebhook        provisionBranch = "webhook_provision"
)

// The two entry points below narrow both policy oneofs on data to the single arm
// the practitioner configured.
//
// The API models these as a protobuf oneof and rejects a body carrying two arms
// with `oneof c1.api.policy.v1.ProvisionPolicy.typ is already set`. The generated
// SDK instead models every arm as an independent optional field, and merge() loads
// prior state before overlaying the plan — so the arm being replaced survives in
// the merged model alongside its replacement and both get serialized.
//
// Configuration is the only input where an arm the practitioner did not write is
// reliably null: in the plan, an unset arm of a Computed+Optional object holds the
// prior state value. ConflictsWith cannot substitute for this, because it
// validates configuration and the surviving arm comes from state.
//
// The arms have to stay Computed. Dropping Computed also stops the stale arm, but
// then a config that omits the policy block entirely cannot carry the stored value
// forward and Terraform marks the whole object unknown on every plan — a permanent
// diff, verified against a live tenant.
//
// An empty config selection leaves data untouched: nothing was chosen, so whichever
// single arm the remote already had is the correct thing to send.

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

	provision, provisionDiags := selectedProvisionPolicyBranch("provision_policy", configured.ProvisionPolicy)
	diags.Append(provisionDiags...)
	deprovisioner, deprovisionerDiags := selectedDeprovisionerPolicyBranch("deprovisioner_policy", configured.DeprovisionerPolicy)
	diags.Append(deprovisionerDiags...)
	if diags.HasError() {
		return diags
	}

	keepOnlyProvisionPolicyBranch(data.ProvisionPolicy, provision)
	keepOnlyDeprovisionerPolicyBranch(data.DeprovisionerPolicy, deprovisioner)

	return diags
}

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

	provision, provisionDiags := selectedProvisionPolicyBranch("provision_policy", configured.ProvisionPolicy)
	diags.Append(provisionDiags...)
	deprovisioner, deprovisionerDiags := selectedDeprovisionerPolicyBranch("deprovisioner_policy", configured.DeprovisionerPolicy)
	diags.Append(deprovisionerDiags...)
	if diags.HasError() {
		return diags
	}

	keepOnlyProvisionPolicyBranch(data.ProvisionPolicy, provision)
	keepOnlyDeprovisionerPolicyBranch(data.DeprovisionerPolicy, deprovisioner)

	return diags
}

// The four functions below are duplicated per type rather than unified:
// tfTypes.ProvisionPolicy and tfTypes.DeprovisionerPolicy are distinct structs
// with identical arms, and Go generics cannot range over struct fields.
// TestPolicyOneofTypesStayInSync fails if the two shapes diverge, which is the
// signal to update all four.

func selectedProvisionPolicyBranch(attribute string, p *tfTypes.ProvisionPolicy) (provisionBranch, diag.Diagnostics) {
	if p == nil {
		return provisionBranchNone, nil
	}
	var configured []provisionBranch
	if p.ActionProvision != nil {
		configured = append(configured, provisionBranchAction)
	}
	if p.ConnectorProvision != nil {
		configured = append(configured, provisionBranchConnector)
	}
	if p.DelegatedProvision != nil {
		configured = append(configured, provisionBranchDelegated)
	}
	if p.ExternalTicketProvision != nil {
		configured = append(configured, provisionBranchExternalTicket)
	}
	if p.ManualProvision != nil {
		configured = append(configured, provisionBranchManual)
	}
	if !p.MultiStep.IsNull() && !p.MultiStep.IsUnknown() {
		configured = append(configured, provisionBranchMultiStep)
	}
	if p.UnconfiguredProvision != nil {
		configured = append(configured, provisionBranchUnconfigured)
	}
	if p.WebhookProvision != nil {
		configured = append(configured, provisionBranchWebhook)
	}
	return onlyConfiguredBranch(attribute, configured)
}

// onlyConfiguredBranch rejects a configuration carrying more than one arm.
// Choosing one of them applies a policy the practitioner did not write, and the
// arm that loses is dropped without a plan ever showing it.
func onlyConfiguredBranch(attribute string, configured []provisionBranch) (provisionBranch, diag.Diagnostics) {
	var diags diag.Diagnostics
	switch len(configured) {
	case 0:
		return provisionBranchNone, diags
	case 1:
		return configured[0], diags
	}

	names := make([]string, 0, len(configured))
	for _, branch := range configured {
		names = append(names, string(branch))
	}
	diags.AddError(
		fmt.Sprintf("Conflicting %s configuration", attribute),
		fmt.Sprintf(
			"Only one provisioning arm may be configured, but the configuration sets %s. The API models this as a oneof and rejects a body carrying more than one arm.",
			strings.Join(names, ", "),
		),
	)
	return provisionBranchNone, diags
}

func keepOnlyProvisionPolicyBranch(p *tfTypes.ProvisionPolicy, keep provisionBranch) {
	if p == nil || keep == provisionBranchNone {
		return
	}
	if keep != provisionBranchAction {
		p.ActionProvision = nil
	}
	if keep != provisionBranchConnector {
		p.ConnectorProvision = nil
	}
	if keep != provisionBranchDelegated {
		p.DelegatedProvision = nil
	}
	if keep != provisionBranchExternalTicket {
		p.ExternalTicketProvision = nil
	}
	if keep != provisionBranchManual {
		p.ManualProvision = nil
	}
	if keep != provisionBranchMultiStep {
		p.MultiStep = jsontypes.NewNormalizedNull()
	}
	if keep != provisionBranchUnconfigured {
		p.UnconfiguredProvision = nil
	}
	if keep != provisionBranchWebhook {
		p.WebhookProvision = nil
	}
}

func selectedDeprovisionerPolicyBranch(attribute string, p *tfTypes.DeprovisionerPolicy) (provisionBranch, diag.Diagnostics) {
	if p == nil {
		return provisionBranchNone, nil
	}
	var configured []provisionBranch
	if p.ActionProvision != nil {
		configured = append(configured, provisionBranchAction)
	}
	if p.ConnectorProvision != nil {
		configured = append(configured, provisionBranchConnector)
	}
	if p.DelegatedProvision != nil {
		configured = append(configured, provisionBranchDelegated)
	}
	if p.ExternalTicketProvision != nil {
		configured = append(configured, provisionBranchExternalTicket)
	}
	if p.ManualProvision != nil {
		configured = append(configured, provisionBranchManual)
	}
	if !p.MultiStep.IsNull() && !p.MultiStep.IsUnknown() {
		configured = append(configured, provisionBranchMultiStep)
	}
	if p.UnconfiguredProvision != nil {
		configured = append(configured, provisionBranchUnconfigured)
	}
	if p.WebhookProvision != nil {
		configured = append(configured, provisionBranchWebhook)
	}
	return onlyConfiguredBranch(attribute, configured)
}

func keepOnlyDeprovisionerPolicyBranch(p *tfTypes.DeprovisionerPolicy, keep provisionBranch) {
	if p == nil || keep == provisionBranchNone {
		return
	}
	if keep != provisionBranchAction {
		p.ActionProvision = nil
	}
	if keep != provisionBranchConnector {
		p.ConnectorProvision = nil
	}
	if keep != provisionBranchDelegated {
		p.DelegatedProvision = nil
	}
	if keep != provisionBranchExternalTicket {
		p.ExternalTicketProvision = nil
	}
	if keep != provisionBranchManual {
		p.ManualProvision = nil
	}
	if keep != provisionBranchMultiStep {
		p.MultiStep = jsontypes.NewNormalizedNull()
	}
	if keep != provisionBranchUnconfigured {
		p.UnconfiguredProvision = nil
	}
	if keep != provisionBranchWebhook {
		p.WebhookProvision = nil
	}
}
