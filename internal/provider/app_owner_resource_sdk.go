package provider

import (
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppOwnerResourceModel) ToCreateSDKType() *shared.SetAppOwnersRequest {
	var userIds []string = nil
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.SetAppOwnersRequest{
		UserIds: userIds,
	}
	return &out
}

func (r *AppOwnerResourceModel) RefreshFromCreateResponse(resp *shared.SetAppOwnersResponse) {
}

func (r *AppOwnerResourceModel) RefreshFromReadResponse(owners []shared.User) {
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
