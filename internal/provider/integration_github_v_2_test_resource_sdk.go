// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"strings"
	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const githubV2TestCatalogID = "2x97DGRBEL69ood9pTWO3IpF7Ip"

func (r *IntegrationGithubV2TestResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(githubV2TestCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("GitHub v2 Test"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationGithubV2TestResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(githubV2TestCatalogID)
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

func (r *IntegrationGithubV2TestResourceModel) ToUpdateSDKType() (*shared.ConnectorInput, bool) {
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
			configOut[key] = makeStringValue(configValue)
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}

	out := shared.ConnectorInput{
		DisplayName: sdk.String("GitHub v2 Test"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(githubV2TestCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationGithubV2TestResourceModel) populateConfig() map[string]interface{} {
	configValues := make(map[string]interface{})

	githubAccessToken := new(string)
	if !r.GithubAccessToken.IsUnknown() && !r.GithubAccessToken.IsNull() {
		*githubAccessToken = r.GithubAccessToken.ValueString()
		configValues["github_access_token"] = githubAccessToken
	}

	githubOrgList := make([]string, 0)
	for _, item := range r.GithubOrgList {
		githubOrgList = append(githubOrgList, item.ValueString())
	}
	if len(githubOrgList) > 0 {
		configValues["github_org_list"] = strings.Join(githubOrgList, ",")
	}

	return configValues
}

func (r *IntegrationGithubV2TestResourceModel) getConfig() (map[string]interface{}, bool) {
	configValues := r.populateConfig()
	configOut := make(map[string]interface{})
	configSet := false
	for key, configValue := range configValues {
		configOut[key] = ""
		if configValue != nil {
			configOut[key] = makeStringValue(configValue)
			configSet = true
		}
	}
	if !configSet {
		configOut = nil
	}
	return configOut, configSet
}

func (r *IntegrationGithubV2TestResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationGithubV2TestResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationGithubV2TestResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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

				r.GithubOrgList = nil
				if v, ok := values["github_org_list"]; ok {
					if val, ok := v.(string); ok {
						tmpList := strings.Split(val, ",")
						for _, item := range tmpList {
							r.GithubOrgList = append(r.GithubOrgList, types.StringValue(item))
						}
					}
				}

			}
		}
	}
}

func (r *IntegrationGithubV2TestResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationGithubV2TestResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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

				r.GithubOrgList = nil
				if v, ok := values["github_org_list"]; ok {
					if val, ok := v.(string); ok {
						tmpList := strings.Split(val, ",")
						for _, item := range tmpList {
							r.GithubOrgList = append(r.GithubOrgList, types.StringValue(item))
						}
					}
				}

			}
		}
	}
}
