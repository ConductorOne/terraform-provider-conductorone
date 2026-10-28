package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ resource.ResourceWithModifyPlan = (*CustomAppEntitlementResource)(nil)

// ModifyPlan removes the spurious update caused by an unconfigured, formerly
// computed-only provisioner_policy. The framework marks computed values
// unknown before its attribute plan modifiers restore that alias from state.
// Only restore unknowns when every configured value is known and the entire
// resulting plan matches prior state; real changes and deferred config retain
// their original plan.
func (r *CustomAppEntitlementResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	alias, diags := selectedCustomPolicyAlias(ctx, req.Config)
	if diags.HasError() || alias != "provision_policy" {
		return
	}
	restored, ok := unchangedPolicyAliasPlan(req.Config.Raw, resp.Plan.Raw, req.State.Raw)
	if ok {
		resp.Plan.Raw = restored
	}
}

func unchangedPolicyAliasPlan(config, plan, state tftypes.Value) (tftypes.Value, bool) {
	if config.IsNull() || plan.IsNull() || state.IsNull() || !config.IsKnown() || !plan.IsKnown() || !state.IsFullyKnown() {
		return plan, false
	}
	unknownConfig := false
	err := tftypes.Walk(config, func(_ *tftypes.AttributePath, v tftypes.Value) (bool, error) {
		if !v.IsKnown() {
			unknownConfig = true
			return false, nil
		}
		return true, nil
	})
	if err != nil || unknownConfig {
		return plan, false
	}
	candidate, err := tftypes.Transform(plan, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
		if v.IsKnown() {
			return v, nil
		}
		old, _, err := tftypes.WalkAttributePath(state, p)
		if err != nil {
			return v, err
		}
		prior, ok := old.(tftypes.Value)
		if !ok || !prior.IsKnown() || !prior.Type().Equal(v.Type()) {
			return v, nil
		}
		return prior, nil
	})
	if err != nil || !candidate.Equal(state) {
		return plan, false
	}
	return candidate, true
}
