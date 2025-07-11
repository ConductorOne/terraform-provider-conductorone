// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/conductorone/terraform-provider-conductorone/internal/provider/typeconvert"
	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppResourcesDataSourceModel) ToSharedSearchAppResourcesRequest(ctx context.Context) (*shared.SearchAppResourcesRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	appID := new(string)
	if !r.AppID.IsUnknown() && !r.AppID.IsNull() {
		*appID = r.AppID.ValueString()
	} else {
		appID = nil
	}
	var appUserIds []string
	if r.AppUserIds != nil {
		appUserIds = make([]string, 0, len(r.AppUserIds))
		for _, appUserIdsItem := range r.AppUserIds {
			appUserIds = append(appUserIds, appUserIdsItem.ValueString())
		}
	}
	excludeDeletedResourceBindings := new(bool)
	if !r.ExcludeDeletedResourceBindings.IsUnknown() && !r.ExcludeDeletedResourceBindings.IsNull() {
		*excludeDeletedResourceBindings = r.ExcludeDeletedResourceBindings.ValueBool()
	} else {
		excludeDeletedResourceBindings = nil
	}
	var excludeResourceIds []string
	if r.ExcludeResourceIds != nil {
		excludeResourceIds = make([]string, 0, len(r.ExcludeResourceIds))
		for _, excludeResourceIdsItem := range r.ExcludeResourceIds {
			excludeResourceIds = append(excludeResourceIds, excludeResourceIdsItem.ValueString())
		}
	}
	var excludeResourceTypeTraitIds []string
	if r.ExcludeResourceTypeTraitIds != nil {
		excludeResourceTypeTraitIds = make([]string, 0, len(r.ExcludeResourceTypeTraitIds))
		for _, excludeResourceTypeTraitIdsItem := range r.ExcludeResourceTypeTraitIds {
			excludeResourceTypeTraitIds = append(excludeResourceTypeTraitIds, excludeResourceTypeTraitIdsItem.ValueString())
		}
	}
	pageSize := new(int)
	if !r.PageSize.IsUnknown() && !r.PageSize.IsNull() {
		*pageSize = int(r.PageSize.ValueInt32())
	} else {
		pageSize = nil
	}
	pageToken := new(string)
	if !r.PageToken.IsUnknown() && !r.PageToken.IsNull() {
		*pageToken = r.PageToken.ValueString()
	} else {
		pageToken = nil
	}
	query := new(string)
	if !r.Query.IsUnknown() && !r.Query.IsNull() {
		*query = r.Query.ValueString()
	} else {
		query = nil
	}
	var refs []shared.AppResourceRef
	if r.Refs != nil {
		refs = make([]shared.AppResourceRef, 0, len(r.Refs))
		for _, refsItem := range r.Refs {
			appId1 := new(string)
			if !refsItem.AppID.IsUnknown() && !refsItem.AppID.IsNull() {
				*appId1 = refsItem.AppID.ValueString()
			} else {
				appId1 = nil
			}
			appResourceTypeID := new(string)
			if !refsItem.AppResourceTypeID.IsUnknown() && !refsItem.AppResourceTypeID.IsNull() {
				*appResourceTypeID = refsItem.AppResourceTypeID.ValueString()
			} else {
				appResourceTypeID = nil
			}
			id := new(string)
			if !refsItem.ID.IsUnknown() && !refsItem.ID.IsNull() {
				*id = refsItem.ID.ValueString()
			} else {
				id = nil
			}
			refs = append(refs, shared.AppResourceRef{
				AppID:             appId1,
				AppResourceTypeID: appResourceTypeID,
				ID:                id,
			})
		}
	}
	var resourceIds []string
	if r.ResourceIds != nil {
		resourceIds = make([]string, 0, len(r.ResourceIds))
		for _, resourceIdsItem := range r.ResourceIds {
			resourceIds = append(resourceIds, resourceIdsItem.ValueString())
		}
	}
	var resourceTypeIds []string
	if r.ResourceTypeIds != nil {
		resourceTypeIds = make([]string, 0, len(r.ResourceTypeIds))
		for _, resourceTypeIdsItem := range r.ResourceTypeIds {
			resourceTypeIds = append(resourceTypeIds, resourceTypeIdsItem.ValueString())
		}
	}
	var resourceTypeTraitIds []string
	if r.ResourceTypeTraitIds != nil {
		resourceTypeTraitIds = make([]string, 0, len(r.ResourceTypeTraitIds))
		for _, resourceTypeTraitIdsItem := range r.ResourceTypeTraitIds {
			resourceTypeTraitIds = append(resourceTypeTraitIds, resourceTypeTraitIdsItem.ValueString())
		}
	}
	out := shared.SearchAppResourcesRequest{
		AppID:                          appID,
		AppUserIds:                     appUserIds,
		ExcludeDeletedResourceBindings: excludeDeletedResourceBindings,
		ExcludeResourceIds:             excludeResourceIds,
		ExcludeResourceTypeTraitIds:    excludeResourceTypeTraitIds,
		PageSize:                       pageSize,
		PageToken:                      pageToken,
		Query:                          query,
		Refs:                           refs,
		ResourceIds:                    resourceIds,
		ResourceTypeIds:                resourceTypeIds,
		ResourceTypeTraitIds:           resourceTypeTraitIds,
	}

	return &out, diags
}

func (r *AppResourcesDataSourceModel) RefreshFromSharedSearchAppResourcesResponse(ctx context.Context, resp *shared.SearchAppResourcesResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.Expanded != nil {
		}
		if resp.List != nil {
			if r.List == nil {
				r.List = []tfTypes.AppResourceView{}
				if len(r.List) > len(resp.List) {
					r.List = r.List[:len(resp.List)]
				}
			}
			initListCount := len(r.List)
			for listCount, listItem := range resp.List {
				listCount = initListCount + listCount
				var list tfTypes.AppResourceView
				if listItem.AppResource == nil {
					list.AppResource = nil
				} else {
					list.AppResource = &tfTypes.AppResource{}
					list.AppResource.AppID = types.StringPointerValue(listItem.AppResource.AppID)
					list.AppResource.AppResourceTypeID = types.StringPointerValue(listItem.AppResource.AppResourceTypeID)
					list.AppResource.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(listItem.AppResource.CreatedAt))
					list.AppResource.DeletedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(listItem.AppResource.DeletedAt))
					list.AppResource.Description = types.StringPointerValue(listItem.AppResource.Description)
					list.AppResource.DisplayName = types.StringPointerValue(listItem.AppResource.DisplayName)
					list.AppResource.GrantCount = types.StringPointerValue(listItem.AppResource.GrantCount)
					list.AppResource.ID = types.StringPointerValue(listItem.AppResource.ID)
					list.AppResource.MatchBatonID = types.StringPointerValue(listItem.AppResource.MatchBatonID)
					list.AppResource.ParentAppResourceID = types.StringPointerValue(listItem.AppResource.ParentAppResourceID)
					list.AppResource.ParentAppResourceTypeID = types.StringPointerValue(listItem.AppResource.ParentAppResourceTypeID)
					if listItem.AppResource.SecretTrait == nil {
						list.AppResource.SecretTrait = nil
					} else {
						list.AppResource.SecretTrait = &tfTypes.SecretTrait{}
						list.AppResource.SecretTrait.IdentityAppUserID = types.StringPointerValue(listItem.AppResource.SecretTrait.IdentityAppUserID)
						list.AppResource.SecretTrait.LastUsedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(listItem.AppResource.SecretTrait.LastUsedAt))
						list.AppResource.SecretTrait.SecretCreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(listItem.AppResource.SecretTrait.SecretCreatedAt))
						list.AppResource.SecretTrait.SecretExpiresAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(listItem.AppResource.SecretTrait.SecretExpiresAt))
					}
					list.AppResource.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(listItem.AppResource.UpdatedAt))
				}
				if listCount+1 > len(r.List) {
					r.List = append(r.List, list)
				} else {
					r.List[listCount].AppResource = list.AppResource
				}
			}
		}
		r.NextPageToken = types.StringPointerValue(resp.NextPageToken)
	}

	return diags
}
