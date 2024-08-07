// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

const riskLevelAttributeTypeID = "26G1Et6zd5AzsSaQLkKuPvLtxOb"

func (r *RiskLevelModel) ToCreateSDKType() *shared.CreateAttributeValueRequest {
	attributeTypeId := riskLevelAttributeTypeID
	value := new(string)
	if !r.Value.IsUnknown() && !r.Value.IsNull() {
		*value = r.Value.ValueString()
	} else {
		value = nil
	}
	out := shared.CreateAttributeValueRequest{
		AttributeTypeID: &attributeTypeId,
		Value:           value,
	}
	return &out
}

func (r *RiskLevelModel) ToGetSDKType() *shared.CreateAttributeValueRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *RiskLevelModel) ToDeleteSDKType() *shared.CreateAttributeValueRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *RiskLevelModel) RefreshFromGetResponse(resp *shared.AttributeValue) {
	if resp.AttributeTypeID != nil {
		r.AttributeTypeID = types.StringValue(*resp.AttributeTypeID)
	} else {
		r.AttributeTypeID = types.StringNull()
	}
	if resp.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	if resp.Value != nil {
		r.Value = types.StringValue(*resp.Value)
	} else {
		r.Value = types.StringNull()
	}
}

func (r *RiskLevelModel) RefreshFromCreateResponse(resp *shared.AttributeValue) {
	r.RefreshFromGetResponse(resp)
}
