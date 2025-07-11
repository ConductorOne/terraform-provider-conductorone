// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppGroupApproval struct {
	AllowSelfApproval types.Bool     `tfsdk:"allow_self_approval"`
	AppGroupID        types.String   `tfsdk:"app_group_id"`
	AppID             types.String   `tfsdk:"app_id"`
	Fallback          types.Bool     `tfsdk:"fallback"`
	FallbackUserIds   []types.String `tfsdk:"fallback_user_ids"`
}
