// File: integration_baton_resource_sdk.go. Written by jirwin
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const batonCatalogID = "2HErg7BRmNdHgX8sTCWVq4E1fyh"

func (r *IntegrationBatonResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(batonCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		CatalogID: catalogID,
		UserIds:   userIds,
	}
	return &out
}

func (r *IntegrationBatonResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	out := shared.Connector{
		DisplayName: sdk.String("Baton"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(asanaCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
	}

	return &out, false
}

func (r *IntegrationBatonResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationBatonResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationBatonResourceModel) RefreshFromGetResponse(resp *shared.ConnectorServiceGetResponse) {
	if resp.ConnectorView == nil {
		return
	}
	if resp.ConnectorView.Connector == nil {
		return
	}
	if resp.ConnectorView.Connector.AppID != nil {
		r.AppID = types.StringValue(*resp.ConnectorView.Connector.AppID)
	} else {
		r.AppID = types.StringNull()
	}
	if resp.ConnectorView.Connector.CreatedAt != nil {
		r.CreatedAt = types.StringValue(resp.ConnectorView.Connector.CreatedAt.Format(time.RFC3339))
	} else {
		r.CreatedAt = types.StringNull()
	}
	if resp.ConnectorView.Connector.DeletedAt != nil {
		r.DeletedAt = types.StringValue(resp.ConnectorView.Connector.DeletedAt.Format(time.RFC3339))
	} else {
		r.DeletedAt = types.StringNull()
	}
	if resp.ConnectorView.Connector.ID != nil {
		r.ID = types.StringValue(*resp.ConnectorView.Connector.ID)
	} else {
		r.ID = types.StringNull()
	}
	if resp.ConnectorView.Connector.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.ConnectorView.Connector.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	r.UserIds = nil
	for _, v := range resp.ConnectorView.Connector.UserIds {
		r.UserIds = append(r.UserIds, types.StringValue(v))
	}
}

func (r *IntegrationBatonResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
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
	r.UserIds = nil
	for _, v := range resp.UserIds {
		r.UserIds = append(r.UserIds, types.StringValue(v))
	}
}

func (r *IntegrationBatonResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromCreateResponse(resp)
}
