// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const snowflakeCatalogID = "24cjzH6sY7x4huNU1JbaxkvzVm9"

func (r *IntegrationSnowflakeResourceModel) ToCreateSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
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

func (r *IntegrationSnowflakeResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}

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

	configValues := map[string]*string{
		"snowflake_account":   snowflakeAccount,
		"snowflake_username":  snowflakeUsername,
		"snowflake_password":  snowflakePassword,
		"snowflake_user_role": snowflakeUserRole,
	}

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
		DisplayName: sdk.String("Snowflake"),
		AppID:       sdk.String(r.AppID.ValueString()),
		CatalogID:   sdk.String(snowflakeCatalogID),
		ID:          sdk.String(r.ID.ValueString()),
		UserIds:     userIds,
		Config:      makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationSnowflakeResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *IntegrationSnowflakeResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateSDKType()
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
					r.SnowflakeAccount = types.StringValue(v.(string))
				}

				if v, ok := values["snowflake_username"]; ok {
					r.SnowflakeUsername = types.StringValue(v.(string))
				}

				if v, ok := values["snowflake_password"]; ok {
					r.SnowflakePassword = types.StringValue(v.(string))
				}

				if v, ok := values["snowflake_user_role"]; ok {
					r.SnowflakeUserRole = types.StringValue(v.(string))
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
					r.SnowflakeAccount = types.StringValue(v.(string))
				}

				if v, ok := values["snowflake_username"]; ok {
					r.SnowflakeUsername = types.StringValue(v.(string))
				}

				if v, ok := values["snowflake_password"]; ok {
					r.SnowflakePassword = types.StringValue(v.(string))
				}

				if v, ok := values["snowflake_user_role"]; ok {
					r.SnowflakeUserRole = types.StringValue(v.(string))
				}

			}
		}
	}
}