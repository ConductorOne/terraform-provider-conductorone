// File: integration_baton_resource_sdk.go. Written by jirwin
package provider

import (
	"strconv"
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const oktaCatalogID = "23w9L3qudsiSZQJ8jUP1KYyQqVW"

func (r *IntegrationOktaResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(oktaCatalogID)
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

func (r *IntegrationOktaResourceModel) ToUpdateSDKType() *shared.Connector {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	oktaDomain := new(string)
	if !r.OktaDomain.IsUnknown() && !r.OktaDomain.IsNull() {
		*oktaDomain = r.OktaDomain.ValueString()
	} else {
		oktaDomain = nil
	}
	oktaApiKey := new(string)
	if !r.OktaApiKey.IsUnknown() && !r.OktaApiKey.IsNull() {
		*oktaApiKey = r.OktaApiKey.ValueString()
	} else {
		oktaApiKey = nil
	}
	oktaDontSyncInactiveApps := new(string)
	if !r.OktaDontSyncInactiveApps.IsUnknown() && !r.OktaDontSyncInactiveApps.IsNull() {
		*oktaDontSyncInactiveApps = strconv.FormatBool(r.OktaDontSyncInactiveApps.ValueBool())
	} else {
		oktaDontSyncInactiveApps = nil
	}
	oktaExtractAwsSamlRoles := new(string)
	if !r.OktaExtractAwsSamlRoles.IsUnknown() && !r.OktaExtractAwsSamlRoles.IsNull() {
		*oktaExtractAwsSamlRoles = strconv.FormatBool(r.OktaExtractAwsSamlRoles.ValueBool())
	} else {
		oktaExtractAwsSamlRoles = nil
	}

	config := makeConnectorConfig(map[string]interface{}{
		"okta_domain":                  oktaDomain,
		"okta_api_key":                 oktaApiKey,
		"okta_dont_sync_inactive_apps": oktaDontSyncInactiveApps,
		"okta_extract_aws_saml_roles":  oktaExtractAwsSamlRoles,
	})

	out := shared.Connector{
		AppID:     sdk.String(r.AppID.ValueString()),
		CatalogID: sdk.String(oktaCatalogID),
		ID:        sdk.String(r.ID.ValueString()),
		UserIds:   userIds,
		Config:    config,
	}
	return &out
}

func (r *IntegrationOktaResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationOktaResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationOktaResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
			if v, ok := config["okta_domain"]; ok {
				r.OktaDomain = types.StringValue(v.(string))
			}
			if v, ok := config["okta_api_key"]; ok {
				r.OktaApiKey = types.StringValue(v.(string))
			}
			if v, ok := config["okta_dont_sync_inactive_apps"]; ok {
				bv, err := strconv.ParseBool(v.(string))
				if err != nil {
					r.OktaDontSyncInactiveApps = types.BoolValue(false)
				}
				r.OktaDontSyncInactiveApps = types.BoolValue(bv)
			}
			if v, ok := config["okta_okta_extract_aws_saml_roles"]; ok {
				bv, err := strconv.ParseBool(v.(string))
				if err != nil {
					r.OktaExtractAwsSamlRoles = types.BoolValue(false)
				}
				r.OktaExtractAwsSamlRoles = types.BoolValue(bv)
			}
		}
	}
}

func (r *IntegrationOktaResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationOktaResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
