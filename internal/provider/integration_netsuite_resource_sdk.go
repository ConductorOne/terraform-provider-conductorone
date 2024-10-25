// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const netsuiteCatalogID = "2OF3kwc6iUfypKssyNr9S7IP9oe"

func (r *IntegrationNetsuiteResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(netsuiteCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("NetSuite"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationNetsuiteResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(netsuiteCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	configOut, configSet := r.getConfig()
	if !configSet {
		return nil, fmt.Errorf("config must be set for create request")
	}

	out := shared.ConnectorServiceCreateRequest{
		CatalogID: catalogID,
		UserIds:   userIds,
		Config: &shared.ConnectorServiceCreateRequestConfig{
			AtType: sdk.String(envConfigType),
			AdditionalProperties: map[string]interface{}{
				"configuration": configOut,
			},
		},
	}
	return &out, nil
}

func (r *IntegrationNetsuiteResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	configValues := r.populateConfig()

	configOut := make(map[string]interface{})
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}

	out := shared.Connector{
		DisplayName: sdk.String("NetSuite"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(netsuiteCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationNetsuiteResourceModel) populateConfig() map[string]interface{} {
	netsuiteAccountId := new(string)
	if !r.NetsuiteAccountId.IsUnknown() && !r.NetsuiteAccountId.IsNull() {
		*netsuiteAccountId = r.NetsuiteAccountId.ValueString()
	} else {
		netsuiteAccountId = nil
	}

	netsuiteConsumerKey := new(string)
	if !r.NetsuiteConsumerKey.IsUnknown() && !r.NetsuiteConsumerKey.IsNull() {
		*netsuiteConsumerKey = r.NetsuiteConsumerKey.ValueString()
	} else {
		netsuiteConsumerKey = nil
	}

	netsuiteConsumerSecret := new(string)
	if !r.NetsuiteConsumerSecret.IsUnknown() && !r.NetsuiteConsumerSecret.IsNull() {
		*netsuiteConsumerSecret = r.NetsuiteConsumerSecret.ValueString()
	} else {
		netsuiteConsumerSecret = nil
	}

	netsuiteTokenKey := new(string)
	if !r.NetsuiteTokenKey.IsUnknown() && !r.NetsuiteTokenKey.IsNull() {
		*netsuiteTokenKey = r.NetsuiteTokenKey.ValueString()
	} else {
		netsuiteTokenKey = nil
	}

	netsuiteTokenSecret := new(string)
	if !r.NetsuiteTokenSecret.IsUnknown() && !r.NetsuiteTokenSecret.IsNull() {
		*netsuiteTokenSecret = r.NetsuiteTokenSecret.ValueString()
	} else {
		netsuiteTokenSecret = nil
	}

	configValues := map[string]interface{}{
		"netsuite_account_id":      netsuiteAccountId,
		"netsuite_consumer_key":    netsuiteConsumerKey,
		"netsuite_consumer_secret": netsuiteConsumerSecret,
		"netsuite_token_key":       netsuiteTokenKey,
		"netsuite_token_secret":    netsuiteTokenSecret,
	}

	return configValues
}

func (r *IntegrationNetsuiteResourceModel) getConfig() (map[string]interface{}, bool) {
	configValues := r.populateConfig()
	configOut := make(map[string]interface{})
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}
	return configOut, configSet
}

func (r *IntegrationNetsuiteResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationNetsuiteResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationNetsuiteResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
	if resp == nil {
		return
	}
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

	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["netsuite_account_id"]; ok {
					if val, ok := v.(string); ok {
						r.NetsuiteAccountId = types.StringValue(val)
					}
				}

				if v, ok := values["netsuite_consumer_key"]; ok {
					if val, ok := v.(string); ok {
						r.NetsuiteConsumerKey = types.StringValue(val)
					}
				}

				if v, ok := values["netsuite_token_key"]; ok {
					if val, ok := v.(string); ok {
						r.NetsuiteTokenKey = types.StringValue(val)
					}
				}

			}
		}
	}
}

func (r *IntegrationNetsuiteResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationNetsuiteResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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

	if resp.Config != nil && *resp.Config.AtType == envConfigType {
		if config, ok := resp.Config.AdditionalProperties.(map[string]interface{}); ok {
			if values, ok := config["configuration"].(map[string]interface{}); ok {
				if v, ok := values["netsuite_account_id"]; ok {
					if val, ok := v.(string); ok {
						r.NetsuiteAccountId = types.StringValue(val)
					}
				}

				if v, ok := values["netsuite_consumer_key"]; ok {
					if val, ok := v.(string); ok {
						r.NetsuiteConsumerKey = types.StringValue(val)
					}
				}

				if v, ok := values["netsuite_token_key"]; ok {
					if val, ok := v.(string); ok {
						r.NetsuiteTokenKey = types.StringValue(val)
					}
				}

			}
		}
	}
}
