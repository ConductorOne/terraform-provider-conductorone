package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeOktaDomainPlanModifier rewrites the planned value of an Okta domain
// attribute to its normalized form (stripping a trailing "-admin" from the
// hostname), so the planned value matches the value the C1 API returns. This
// resolves the "Provider produced inconsistent result after apply" error and
// prevents perpetual plan drift for admin-form domains.
func normalizeOktaDomainPlanModifier() planmodifier.String {
	return &normalizeOktaDomainPlanModifierType{}
}

type normalizeOktaDomainPlanModifierType struct{}

func (m *normalizeOktaDomainPlanModifierType) Description(_ context.Context) string {
	return "Normalizes the Okta domain to its canonical (non -admin) form"
}

func (m *normalizeOktaDomainPlanModifierType) MarkdownDescription(_ context.Context) string {
	return "Normalizes the Okta domain to its canonical (non -admin) form"
}

func (m *normalizeOktaDomainPlanModifierType) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Only normalize a known, non-null planned value. Leave unknown (computed
	// during apply) and null values untouched.
	if req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		resp.PlanValue = req.PlanValue
		return
	}
	resp.PlanValue = types.StringValue(normalizeOktaDomain(req.PlanValue.ValueString()))
}
