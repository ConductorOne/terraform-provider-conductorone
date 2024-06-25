package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

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

func (r *AppEntitlementOwnerResourceModel) ToDeleteSDKType() *shared.SetAppEntitlementOwnersRequest {
	out := shared.SetAppEntitlementOwnersRequest{
		UserIds: []string{},
	}
	return &out
}

func (r *AppEntitlementOwnerResourceModel) RefreshFromCreateResponse(resp *shared.SetAppEntitlementOwnersResponse) {
}

func (r *AppEntitlementOwnerResourceModel) RefreshFromReadResponse(owners []shared.User) {
	userIdsMap := make(map[string]bool)
	ownerIds := make([]types.String, 0)
	needsUpdate := false

	// If the number of owners is different than the number of user ids currently set as owners, we need to update the resource.
	if len(r.UserIds) != len(owners) {
		needsUpdate = true
	}

	for _, userId := range r.UserIds {
		userIdsMap[userId.ValueString()] = true
	}
	for _, owner := range owners {
		ownerID := types.StringPointerValue(owner.ID)
		// If the owner is not in the current state of user ids, we need to update the resource.
		if !userIdsMap[ownerID.ValueString()] {
			needsUpdate = true
		}
		ownerIds = append(ownerIds, ownerID)
	}

	if needsUpdate {
		r.UserIds = ownerIds
	}
}
