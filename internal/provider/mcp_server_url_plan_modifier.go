package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mcpServerURLRequiresReplace() planmodifier.String {
	return mcpServerURLPlanModifier{}
}

type mcpServerURLPlanModifier struct{}

func (m mcpServerURLPlanModifier) Description(context.Context) string {
	return "Replaces the server when its configured endpoint changes."
}

func (m mcpServerURLPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m mcpServerURLPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() || req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}
	if req.PlanValue.Equal(req.StateValue) {
		return
	}

	if req.StateValue.IsNull() {
		var endpointURL types.String
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("endpoint_url"), &endpointURL)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if endpointURL.Equal(req.PlanValue) {
			return
		}
	}

	resp.RequiresReplace = true
}
