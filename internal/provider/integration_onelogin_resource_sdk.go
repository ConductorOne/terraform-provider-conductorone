// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const oneloginCatalogID = "2DAwxvPhmPBAYWkXMMK35X9lcDD"

func (r *IntegrationOneloginResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(oneloginCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("OneLogin"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationOneloginResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	configValues := r.populateConfig()

	configOut := make(map[string]string)
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = *configValue
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}

	out := shared.Connector{
		DisplayName: sdk.String("OneLogin"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(oneloginCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationOneloginResourceModel) populateConfig() map[string]*string {
	oneloginDomain := new(string)
	if !r.OneloginDomain.IsUnknown() && !r.OneloginDomain.IsNull() {
		*oneloginDomain = r.OneloginDomain.ValueString()
	} else {
		oneloginDomain = nil
	}

	oauthClientCredGrantClientId := new(string)
	if !r.OauthClientCredGrantClientId.IsUnknown() && !r.OauthClientCredGrantClientId.IsNull() {
		*oauthClientCredGrantClientId = r.OauthClientCredGrantClientId.ValueString()
	} else {
		oauthClientCredGrantClientId = nil
	}

	oauthClientCredGrantClientSecret := new(string)
	if !r.OauthClientCredGrantClientSecret.IsUnknown() && !r.OauthClientCredGrantClientSecret.IsNull() {
		*oauthClientCredGrantClientSecret = r.OauthClientCredGrantClientSecret.ValueString()
	} else {
		oauthClientCredGrantClientSecret = nil
	}

	configValues := map[string]*string{
		"onelogin_domain":                       oneloginDomain,
		"oauth_client_cred_grant_client_id":     oauthClientCredGrantClientId,
		"oauth_client_cred_grant_client_secret": oauthClientCredGrantClientSecret,
	}

	return configValues
}

func (r *IntegrationOneloginResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationOneloginResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationOneloginResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["onelogin_domain"]; ok {
					r.OneloginDomain = types.StringValue(v.(string))
				}

				if v, ok := values["oauth_client_cred_grant_client_id"]; ok {
					r.OauthClientCredGrantClientId = types.StringValue(v.(string))
				}

			}
		}
	}
}

func (r *IntegrationOneloginResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationOneloginResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["onelogin_domain"]; ok {
					r.OneloginDomain = types.StringValue(v.(string))
				}

				if v, ok := values["oauth_client_cred_grant_client_id"]; ok {
					r.OauthClientCredGrantClientId = types.StringValue(v.(string))
				}

			}
		}
	}
}
