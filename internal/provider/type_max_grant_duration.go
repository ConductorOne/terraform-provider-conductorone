package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type MaxGrantDuration struct {
	DurationGrant types.String                 `tfsdk:"duration_grant"`
	DurationUnset *AppEntitlementDurationUnset `tfsdk:"duration_unset"`
}
