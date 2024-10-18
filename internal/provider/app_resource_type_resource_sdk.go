// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *AppResourceTypeResourceModel) ToSharedCreateManuallyManagedResourceTypeRequest() *shared.CreateManuallyManagedResourceTypeRequest {
	name := new(string)
	if !r.DisplayName.IsUnknown() && !r.DisplayName.IsNull() {
		*name = r.DisplayName.ValueString()
	} else {
		name = nil
	}
	resourceType := new(shared.ResourceType)
	if !r.ResourceType.IsUnknown() && !r.ResourceType.IsNull() {
		*resourceType = shared.ResourceType(r.ResourceType.ValueString())
	} else {
		resourceType = nil
	}
	out := shared.CreateManuallyManagedResourceTypeRequest{
		Name:         name,
		ResourceType: resourceType,
	}
	return &out
}

func (r *AppResourceTypeResourceModel) RefreshFromSharedAppResourceType(resp *shared.AppResourceType) {
	if resp != nil {
		r.AppID = types.StringPointerValue(resp.AppID)
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		if resp.DeletedAt != nil {
			r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339Nano))
		} else {
			r.DeletedAt = types.StringNull()
		}
		r.DisplayName = types.StringPointerValue(resp.DisplayName)
		r.ID = types.StringPointerValue(resp.ID)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}

func (r *AppResourceTypeResourceModel) ToSharedAppResourceTypeInput() *shared.AppResourceTypeInput {
	displayName := new(string)
	if !r.DisplayName.IsUnknown() && !r.DisplayName.IsNull() {
		*displayName = r.DisplayName.ValueString()
	} else {
		displayName = nil
	}
	out := shared.AppResourceTypeInput{
		DisplayName: displayName,
	}
	return &out
}

func (r *AppResourceTypeResourceModel) ToSharedDeleteManuallyManagedResourceTypeRequest() *shared.DeleteManuallyManagedResourceTypeRequest {
	out := shared.DeleteManuallyManagedResourceTypeRequest{}
	return &out
}