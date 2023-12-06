// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"time"
)

func (r *ConnectorCredentialDataSourceModel) RefreshFromGetResponse(resp *shared.ConnectorCredential) {
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
	}
	if resp.ClientID != nil {
		r.ClientID = types.StringValue(*resp.ClientID)
	} else {
		r.ClientID = types.StringNull()
	}
	if resp.ConnectorID != nil {
		r.ConnectorID = types.StringValue(*resp.ConnectorID)
	} else {
		r.ConnectorID = types.StringNull()
	}
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
	if resp.DisplayName != nil {
		r.DisplayName = types.StringValue(*resp.DisplayName)
	} else {
		r.DisplayName = types.StringNull()
	}
	if resp.ExpiresTime != nil {
		r.ExpiresTime = types.StringValue(resp.ExpiresTime.Format(time.RFC3339Nano))
	} else {
		r.ExpiresTime = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.LastUsedAt != nil {
		r.LastUsedAt = types.StringValue(resp.LastUsedAt.Format(time.RFC3339Nano))
	} else {
		r.LastUsedAt = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	} else {
		r.UpdatedAt = types.StringNull()
	}
}