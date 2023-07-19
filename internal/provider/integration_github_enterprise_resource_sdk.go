// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const githubEnterpriseCatalogID = "2GjkbaX83ElxXZ7PUvBTjaCuNs4"

func (r *IntegrationGithubEnterpriseResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(githubEnterpriseCatalogID)
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

func (r *IntegrationGithubEnterpriseResourceModel) ToUpdateSDKType() *shared.Connector {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

	githubInstanceUrl := new(string)
	if !r.GithubInstanceUrl.IsUnknown() && !r.GithubInstanceUrl.IsNull() {
		*githubInstanceUrl = r.GithubInstanceUrl.ValueString()
	} else {
		githubInstanceUrl = nil
	}

	githubAccessToken := new(string)
	if !r.GithubAccessToken.IsUnknown() && !r.GithubAccessToken.IsNull() {
		*githubAccessToken = r.GithubAccessToken.ValueString()
	} else {
		githubAccessToken = nil
	}

	config := makeConnectorConfig(map[string]interface{}{
		"github_instance_url": githubInstanceUrl,
		"github_access_token": githubAccessToken,
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

func (r *IntegrationGithubEnterpriseResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationGithubEnterpriseResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationGithubEnterpriseResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
			if v, ok := config["github_instance_url"]; ok {
				r.GithubInstanceUrl = types.StringValue(v.(string))
			}

			if v, ok := config["github_access_token"]; ok {
				r.GithubAccessToken = types.StringValue(v.(string))
			}

		}
	}
}

func (r *IntegrationGithubEnterpriseResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationGithubEnterpriseResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
