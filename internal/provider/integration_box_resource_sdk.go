// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
    "fmt"
	
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const boxCatalogID = "2TBFfg3QmqMKniWXDCzNapZ8Ope"

func (r *IntegrationBoxResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(boxCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Box"),
		CatalogID: catalogID,
		UserIds:   userIds,
	}
	return &out
}

func (r *IntegrationBoxResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(boxCatalogID)
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

func (r *IntegrationBoxResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
	    DisplayName: sdk.String("Box"),
		AppID:     sdk.String(r.AppID.ValueString()),
		CatalogID: sdk.String(boxCatalogID),
		ID:        sdk.String(r.ID.ValueString()),
		UserIds:   userIds,
		Config: makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationBoxResourceModel) populateConfig() map[string]*string {
     boxClientId := new(string)
if !r.BoxClientId.IsUnknown() && !r.BoxClientId.IsNull() {
*boxClientId = r.BoxClientId.ValueString()
} else {
boxClientId = nil
}

        boxClientSecret := new(string)
if !r.BoxClientSecret.IsUnknown() && !r.BoxClientSecret.IsNull() {
*boxClientSecret = r.BoxClientSecret.ValueString()
} else {
boxClientSecret = nil
}

        boxEnterpriseId := new(string)
if !r.BoxEnterpriseId.IsUnknown() && !r.BoxEnterpriseId.IsNull() {
*boxEnterpriseId = r.BoxEnterpriseId.ValueString()
} else {
boxEnterpriseId = nil
}

        

    	configValues := map[string]*string{
    	"box_client_id": boxClientId,
"box_client_secret": boxClientSecret,
"box_enterprise_id": boxEnterpriseId,

    	}

    	return configValues
}

func (r *IntegrationBoxResourceModel) getConfig() (map[string]string, bool) {
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
	return configOut, configSet
}

func (r *IntegrationBoxResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationBoxResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationBoxResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
               if v, ok := values["box_client_id"]; ok {
r.BoxClientId = types.StringValue(v.(string))
}

               
               if v, ok := values["box_enterprise_id"]; ok {
r.BoxEnterpriseId = types.StringValue(v.(string))
}

               
           }
       }
    }
}

func (r *IntegrationBoxResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationBoxResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
                  if v, ok := values["box_client_id"]; ok {
r.BoxClientId = types.StringValue(v.(string))
}

                  
                  if v, ok := values["box_enterprise_id"]; ok {
r.BoxEnterpriseId = types.StringValue(v.(string))
}

                  
              }
          }
       }
}
