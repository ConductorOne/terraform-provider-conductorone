package provider

import (
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppEntitlementOwnerResourceModel) ToCreateSDKType() *shared.SetAppEntitlementOwnersRequest {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.SetAppEntitlementOwnersRequest{
		UserIds: userIds,
	}
	return &out
}

func (r *AppEntitlementOwnerResourceModel) ToDeleteSDKType(owners []shared.User) *shared.SetAppEntitlementOwnersRequest {
	userIdsMap := make(map[string]bool)
	var remainingOwners []string
	for _, userId := range r.UserIds {
		userIdsMap[userId.ValueString()] = true
	}
	for _, owner := range owners {
		ownerID := types.StringPointerValue(owner.ID).ValueString()
		if !userIdsMap[ownerID] {
			remainingOwners = append(remainingOwners, ownerID)
		}
	}
	out := shared.SetAppEntitlementOwnersRequest{
		UserIds: remainingOwners,
	}
	return &out
}

func (r *AppEntitlementOwnerResourceModel) RefreshFromCreateResponse(resp *shared.SetAppEntitlementOwnersResponse) {
}

func (r *AppEntitlementOwnerResourceModel) RefreshFromReadResponse(owners []shared.User) {
	r.UserIds = nil
	for _, owner := range owners {
		r.UserIds = append(r.UserIds, types.StringPointerValue(owner.ID))
	}
}
