// File: integration_baton_resource_sdk.go. Written by jirwin
package provider

import (
	"fmt"
	"time"

	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	oktaCatalogID = "23w9L3qudsiSZQJ8jUP1KYyQqVW"
	envConfigType = "type.googleapis.com/c1.api.app.v1.EnvConfig"
)

func (r *IntegrationOktaResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := strptr(oktaCatalogID)
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

func (r *IntegrationOktaResourceModel) ToUpdateSDKType() *shared.Connector {
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

	config := &shared.ConnectorConfig{
		AtType: strptr(envConfigType),
		AdditionalProperties: map[string]interface{}{
			"configuration": map[string]interface{}{
				"okta_domain":                  oktaDomain,
				"okta_api_key":                 oktaApiKey,
				"okta_dont_sync_inactive_apps": "false",
				"okta_extract_aws_saml_roles":  "true",
			},
		},
	}

	out := shared.Connector{
		AppID:       strptr(r.AppID.ValueString()),
		CatalogID:   strptr(oktaCatalogID),
		ID:          strptr(r.ID.ValueString()),
		Description: description,
		DisplayName: displayName,
		UserIds:     userIds,
		Config:      config,
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

func (r *IntegrationOktaResourceModel) RefreshFromGetResponse(resp *shared.ConnectorServiceGetResponse) {
	fmt.Println("RefreshFromGetResponse")
	spew.Dump(resp)
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
	if resp.ConnectorView.Connector.Config != nil && *resp.ConnectorView.Connector.Config.AtType == envConfigType {
		if config, ok := resp.ConnectorView.Connector.Config.AdditionalProperties.(map[string]interface{}); ok {
			if v, ok := config["okta_domain"]; ok {
				r.OktaDomain = types.StringValue(v.(string))
			}
			if v, ok := config["okta_api_key"]; ok {
				r.OktaApiKey = types.StringValue(v.(string))
			}
		}

	}
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
