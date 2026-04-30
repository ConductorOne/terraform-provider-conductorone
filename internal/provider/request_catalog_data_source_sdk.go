package provider

import (
	"context"
	"github.com/conductorone/terraform-provider-conductorone/internal/provider/typeconvert"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *RequestCatalogDataSourceModel) RefreshFromSharedRequestCatalog(ctx context.Context, resp *shared.RequestCatalog) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreatedAt))
		r.CreatedByUserID = types.StringPointerValue(resp.CreatedByUserID)
		r.DeletedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.DeletedAt))
		r.Description = types.StringPointerValue(resp.Description)
		r.DisplayName = types.StringPointerValue(resp.DisplayName)
		if resp.EnrollmentBehavior != nil {
			r.EnrollmentBehavior = types.StringValue(string(*resp.EnrollmentBehavior))
		} else {
			r.EnrollmentBehavior = types.StringNull()
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.Published = types.BoolPointerValue(resp.Published)
		r.RequestBundle = types.BoolPointerValue(resp.RequestBundle)
		if resp.UnenrollmentBehavior != nil {
			r.UnenrollmentBehavior = types.StringValue(string(*resp.UnenrollmentBehavior))
		} else {
			r.UnenrollmentBehavior = types.StringNull()
		}
		if resp.UnenrollmentEntitlementBehavior != nil {
			r.UnenrollmentEntitlementBehavior = types.StringValue(string(*resp.UnenrollmentEntitlementBehavior))
		} else {
			r.UnenrollmentEntitlementBehavior = types.StringNull()
		}
		r.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.UpdatedAt))
		r.VisibleToEveryone = types.BoolPointerValue(resp.VisibleToEveryone)
	}

	return diags
}

func (r *RequestCatalogDataSourceModel) RefreshFromSharedRequestCatalogManagementServiceGetResponse(ctx context.Context, resp *shared.RequestCatalogManagementServiceGetResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		if resp.RequestCatalogView != nil {
			diags.Append(r.RefreshFromSharedRequestCatalog(ctx, resp.RequestCatalogView.RequestCatalog)...)

			if diags.HasError() {
				return diags
			}
		}
	}

	return diags
}

func (r *RequestCatalogDataSourceModel) ToOperationsC1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest(ctx context.Context) (*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var id string
	id = r.ID.ValueString()

	out := operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest{
		ID: id,
	}

	return &out, diags
}
