package provider

import (
	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppEntitlementProxyBindingResourceModel) ToSharedCreateAppEntitlementProxyRequest() *shared.CreateAppEntitlementProxyRequest {
	dstAppEntitlementID := new(string)
	if !r.DstAppEntitlementID.IsUnknown() && !r.DstAppEntitlementID.IsNull() {
		*dstAppEntitlementID = r.DstAppEntitlementID.ValueString()
	} else {
		dstAppEntitlementID = nil
	}
	dstAppID := new(string)
	if !r.DstAppID.IsUnknown() && !r.DstAppID.IsNull() {
		*dstAppID = r.DstAppID.ValueString()
	} else {
		dstAppID = nil
	}
	srcAppID := new(string)
	if !r.SrcAppID.IsUnknown() && !r.SrcAppID.IsNull() {
		*srcAppID = r.SrcAppID.ValueString()
	} else {
		srcAppID = nil
	}
	srcAppEntitlementID := new(string)
	if !r.SrcAppEntitlementID.IsUnknown() && !r.SrcAppEntitlementID.IsNull() {
		*srcAppEntitlementID = r.SrcAppEntitlementID.ValueString()
	} else {
		srcAppEntitlementID = nil
	}

	out := shared.CreateAppEntitlementProxyRequest{
		DstAppEntitlementID: dstAppEntitlementID,
		DstAppID:            dstAppID,
		SrcAppID:            srcAppID,
		SrcAppEntitlementID: srcAppEntitlementID,
	}
	return &out
}

func (r *AppEntitlementProxyBindingResourceModel) RefreshFromSharedAppEntitlementProxy(resp *shared.AppEntitlementProxy) {
	if resp != nil {
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
		r.DstAppEntitlementID = types.StringPointerValue(resp.DstAppEntitlementID)
		r.DstAppID = types.StringPointerValue(resp.DstAppID)
		r.Implicit = types.BoolPointerValue(resp.Implicit)
		r.SrcAppEntitlementID = types.StringPointerValue(resp.SrcAppEntitlementID)
		r.SrcAppID = types.StringPointerValue(resp.SrcAppID)
		r.SystemBuiltin = types.BoolPointerValue(resp.SystemBuiltin)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}

func (r *AppEntitlementProxyBindingResourceModel) ToSharedDeleteAppEntitlementProxyRequest() *shared.DeleteAppEntitlementProxyRequest {
	dstAppEntitlementID := new(string)
	if !r.DstAppEntitlementID.IsUnknown() && !r.DstAppEntitlementID.IsNull() {
		*dstAppEntitlementID = r.DstAppEntitlementID.ValueString()
	} else {
		dstAppEntitlementID = nil
	}
	dstAppID := new(string)
	if !r.DstAppID.IsUnknown() && !r.DstAppID.IsNull() {
		*dstAppID = r.DstAppID.ValueString()
	} else {
		dstAppID = nil
	}
	out := shared.DeleteAppEntitlementProxyRequest{
		DstAppEntitlementID: dstAppEntitlementID,
		DstAppID:            dstAppID,
	}
	return &out
}
