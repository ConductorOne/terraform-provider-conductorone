// File: integration_baton_resource_sdk.go. Written by jirwin
package provider

import (
	"time"

	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const batonCatalogID = "2HErg7BRmNdHgX8sTCWVq4E1fyh"

func strptr(s string) *string {
	return &s
}

func (r *IntegrationBatonResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := strptr(batonCatalogID)
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	displayName := new(string)
	if !r.DisplayName.IsUnknown() && !r.DisplayName.IsNull() {
		*displayName = r.DisplayName.ValueString()
	} else {
		displayName = nil
	}
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		CatalogID:   catalogID,
		Description: description,
		DisplayName: displayName,
		UserIds:     userIds,
	}
	return &out
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
	// TODO(jirwin): Need to figure out how to read the config. Some config values are write only.
	// if r.ConnectorView.Connector.Config == nil {
	// 	r.ConnectorView.Connector.Config = &ConnectorConfig1{}
	// }
	// if resp.ConnectorView.Connector.Config == nil {
	// 	r.ConnectorView.Connector.Config = nil
	// } else {
	// 	r.ConnectorView.Connector.Config = &ConnectorConfig1{}
	// 	if resp.ConnectorView.Connector.Config.AtType != nil {
	// 		r.ConnectorView.Connector.Config.AtType = types.StringValue(*resp.ConnectorView.Connector.Config.AtType)
	// 	} else {
	// 		r.ConnectorView.Connector.Config.AtType = types.StringNull()
	// 	}
	// 	if r.ConnectorView.Connector.Config.AdditionalProperties.IsUnknown() {
	// 		if resp.ConnectorView.Connector.Config.AdditionalProperties == nil {
	// 			r.ConnectorView.Connector.Config.AdditionalProperties = types.StringNull()
	// 		} else {
	// 			additionalPropertiesResult, _ := json.Marshal(resp.ConnectorView.Connector.Config.AdditionalProperties)
	// 			r.ConnectorView.Connector.Config.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
	// 		}
	// 	}
	// }
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
	if resp.ConnectorView.Connector.Description != nil {
		r.Description = types.StringValue(*resp.ConnectorView.Connector.Description)
	} else {
		r.Description = types.StringNull()
	}
	if resp.ConnectorView.Connector.DisplayName != nil {
		r.DisplayName = types.StringValue(*resp.ConnectorView.Connector.DisplayName)
	} else {
		r.DisplayName = types.StringNull()
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
	// if r.Config == nil {
	// 	r.Config = &ConnectorConfig{}
	// }
	// if resp.Config == nil {
	// 	r.Config = nil
	// } else {
	// 	r.Config = &ConnectorConfig{}
	// 	if resp.Config.AtType != nil {
	// 		r.Config.AtType = types.StringValue(*resp.Config.AtType)
	// 	} else {
	// 		r.Config.AtType = types.StringNull()
	// 	}
	// 	if r.Config.AdditionalProperties.IsUnknown() {
	// 		if resp.Config.AdditionalProperties == nil {
	// 			r.Config.AdditionalProperties = types.StringNull()
	// 		} else {
	// 			additionalPropertiesResult, _ := json.Marshal(resp.Config.AdditionalProperties)
	// 			r.Config.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
	// 		}
	// 	}
	// }
	// if r.ConnectorStatus == nil {
	// 	r.ConnectorStatus = &ConnectorStatus{}
	// }
	// if resp.ConnectorStatus == nil {
	// 	r.ConnectorStatus = nil
	// } else {
	// 	r.ConnectorStatus = &ConnectorStatus{}
	// 	if resp.ConnectorStatus.CompletedAt != nil {
	// 		r.ConnectorStatus.CompletedAt = types.StringValue(resp.ConnectorStatus.CompletedAt.Format(time.RFC3339))
	// 	} else {
	// 		r.ConnectorStatus.CompletedAt = types.StringNull()
	// 	}
	// 	if resp.ConnectorStatus.LastError != nil {
	// 		r.ConnectorStatus.LastError = types.StringValue(*resp.ConnectorStatus.LastError)
	// 	} else {
	// 		r.ConnectorStatus.LastError = types.StringNull()
	// 	}
	// 	if resp.ConnectorStatus.StartedAt != nil {
	// 		r.ConnectorStatus.StartedAt = types.StringValue(resp.ConnectorStatus.StartedAt.Format(time.RFC3339))
	// 	} else {
	// 		r.ConnectorStatus.StartedAt = types.StringNull()
	// 	}
	// 	if resp.ConnectorStatus.Status != nil {
	// 		r.ConnectorStatus.Status = types.StringValue(string(*resp.ConnectorStatus.Status))
	// 	} else {
	// 		r.ConnectorStatus.Status = types.StringNull()
	// 	}
	// 	if resp.ConnectorStatus.UpdatedAt != nil {
	// 		r.ConnectorStatus.UpdatedAt = types.StringValue(resp.ConnectorStatus.UpdatedAt.Format(time.RFC3339))
	// 	} else {
	// 		r.ConnectorStatus.UpdatedAt = types.StringNull()
	// 	}
	// }
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
	if resp.Description != nil {
		r.Description = types.StringValue(*resp.Description)
	} else {
		r.Description = types.StringNull()
	}
	if resp.DisplayName != nil {
		r.DisplayName = types.StringValue(*resp.DisplayName)
	} else {
		r.DisplayName = types.StringNull()
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
