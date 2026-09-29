package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *MCPServerResourceModel) configuredToolApproval() *types.String {
	if r.ExternalConfig != nil {
		return &r.ExternalConfig.RequireToolApproval
	}
	if r.HostedConfig != nil {
		return &r.HostedConfig.RequireToolApproval
	}
	return nil
}

// Refresh only the configured approval leaf; the surrounding config contains
// credentials that the API does not return. Omission leaves the override unmanaged.
func (r *MCPServerResourceModel) refreshToolApproval(view *shared.MCPServerView) {
	approval := r.configuredToolApproval()
	if view == nil || approval == nil || approval.IsNull() || approval.IsUnknown() {
		return
	}
	actual := types.StringNull()
	if view.RequireToolApproval != nil {
		actual = types.StringValue(string(*view.RequireToolApproval))
	}
	*approval = actual
}
