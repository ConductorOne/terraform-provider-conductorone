// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"fmt"

	"time"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const snowflakeCatalogID = "24cjzH6sY7x4huNU1JbaxkvzVm9"

func (r *IntegrationSnowflakeResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(snowflakeCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Snowflake"),
		CatalogID:   catalogID,
		UserIds:     userIds,
	}
	return &out
}

func (r *IntegrationSnowflakeResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(snowflakeCatalogID)
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

func (r *IntegrationSnowflakeResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
		DisplayName: sdk.String("Snowflake"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(snowflakeCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationSnowflakeResourceModel) populateConfig() map[string]interface{} {
	snowflakeAccount := new(string)
	if !r.SnowflakeAccount.IsUnknown() && !r.SnowflakeAccount.IsNull() {
		*snowflakeAccount = r.SnowflakeAccount.ValueString()
	} else {
		snowflakeAccount = nil
	}

	snowflakeUsername := new(string)
	if !r.SnowflakeUsername.IsUnknown() && !r.SnowflakeUsername.IsNull() {
		*snowflakeUsername = r.SnowflakeUsername.ValueString()
	} else {
		snowflakeUsername = nil
	}

	snowflakePassword := new(string)
	if !r.SnowflakePassword.IsUnknown() && !r.SnowflakePassword.IsNull() {
		*snowflakePassword = r.SnowflakePassword.ValueString()
	} else {
		snowflakePassword = nil
	}

	snowflakeUserRole := new(string)
	if !r.SnowflakeUserRole.IsUnknown() && !r.SnowflakeUserRole.IsNull() {
		*snowflakeUserRole = r.SnowflakeUserRole.ValueString()
	} else {
		snowflakeUserRole = nil
	}

	configValues := map[string]interface{}{
		"snowflake_account":   snowflakeAccount,
		"snowflake_username":  snowflakeUsername,
		"snowflake_password":  snowflakePassword,
		"snowflake_user_role": snowflakeUserRole,
	}

	return configValues
}

func (r *IntegrationSnowflakeResourceModel) getConfig() (map[string]interface{}, bool) {
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

func (r *IntegrationSnowflakeResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSnowflakeResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationSnowflakeResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
				if v, ok := values["snowflake_account"]; ok {
					if val, ok := v.(string); ok {
						r.SnowflakeAccount = types.StringValue(val)
					}
				}

				if v, ok := values["snowflake_username"]; ok {
					if val, ok := v.(string); ok {
						r.SnowflakeUsername = types.StringValue(val)
					}
				}

				if v, ok := values["snowflake_user_role"]; ok {
					if val, ok := v.(string); ok {
						r.SnowflakeUserRole = types.StringValue(val)
					}
				}

			}
		}
	}
}

func (r *IntegrationSnowflakeResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationSnowflakeResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
				if v, ok := values["snowflake_account"]; ok {
					if val, ok := v.(string); ok {
						r.SnowflakeAccount = types.StringValue(val)
					}
				}

				if v, ok := values["snowflake_username"]; ok {
					if val, ok := v.(string); ok {
						r.SnowflakeUsername = types.StringValue(val)
					}
				}

				if v, ok := values["snowflake_user_role"]; ok {
					if val, ok := v.(string); ok {
						r.SnowflakeUserRole = types.StringValue(val)
					}
				}

			}
		}
	}
}
