package annotations

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	keyManagedBy    = "managed_by"
	keyIacWorkspace = "iac_workspace"
)

const (
	managedByTerraform = "terraform"
	managedByOpenTofu  = "opentofu"
)

// PlanModifier merges managed_by and iac_workspace into the annotations
// map. User-supplied values win.
//
// Stable across provider upgrades by design — tool versions are not
// auto-set, matching the prior art (k8s app.kubernetes.io/managed-by,
// GCP labels, AWS resource tags). Customers who want to track a version
// can set iac_tool_version in .tf themselves.
//
// iac_resource_address is not auto-set: the resource block name is not
// reachable from planmodifier.Map in the current Plugin Framework.
func PlanModifier() planmodifier.Map {
	return iacAnnotationsModifier{}
}

type iacAnnotationsModifier struct{}

func (m iacAnnotationsModifier) Description(_ context.Context) string {
	return "merges managed_by and iac_workspace into the annotations map; user-supplied values win"
}

func (m iacAnnotationsModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m iacAnnotationsModifier) PlanModifyMap(
	ctx context.Context,
	req planmodifier.MapRequest,
	resp *planmodifier.MapResponse,
) {
	elements := map[string]string{}
	if !req.PlanValue.IsNull() && !req.PlanValue.IsUnknown() {
		diags := req.PlanValue.ElementsAs(ctx, &elements, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	setIfMissing(elements, keyManagedBy, iacToolName())
	if ws := os.Getenv("TF_WORKSPACE"); ws != "" {
		setIfMissing(elements, keyIacWorkspace, ws)
	}

	value, diags := types.MapValueFrom(ctx, types.StringType, elements)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.PlanValue = value
}

func setIfMissing(m map[string]string, k, v string) {
	if v == "" {
		return
	}
	if _, present := m[k]; present {
		return
	}
	m[k] = v
}

func iacToolName() string {
	if isOpenTofu() {
		return managedByOpenTofu
	}
	return managedByTerraform
}

func isOpenTofu() bool {
	if v := os.Getenv("OPENTOFU"); v == "true" || v == "1" {
		return true
	}
	if v := os.Getenv("TF_FORK"); v == "opentofu" || v == "tofu" {
		return true
	}
	return false
}
