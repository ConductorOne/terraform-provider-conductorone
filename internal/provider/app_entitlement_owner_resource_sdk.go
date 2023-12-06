// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
)

func (r *AppEntitlementOwnerResourceModel) ToCreateSDKType() *shared.SetAppEntitlementOwnersRequest {
	var userIds []string = nil
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.SetAppEntitlementOwnersRequest{
		UserIds: userIds,
	}
	return &out
}

func (r *AppEntitlementOwnerResourceModel) RefreshFromCreateResponse(resp *shared.SetAppEntitlementOwnersResponse) {

}
