// Generated by tf-integration-gen. DO NOT EDIT.
package provider

import (
    "fmt"
	
	"time"

	"conductorone/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

const bitbucketCatalogID = "2TBFfduj1grwkNxIV2OGUgFUgYV"

func (r *IntegrationBitbucketResourceModel) ToCreateDelegatedSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	catalogID := sdk.String(bitbucketCatalogID)
	userIds := make([]string, 0)
	for _, userIdsItem := range r.UserIds {
		userIds = append(userIds, userIdsItem.ValueString())
	}
	out := shared.ConnectorServiceCreateDelegatedRequest{
		DisplayName: sdk.String("Bitbucket"),
		CatalogID: catalogID,
		UserIds:   userIds,
	}
	return &out
}

func (r *IntegrationBitbucketResourceModel) ToCreateSDKType() (*shared.ConnectorServiceCreateRequest, error) {
	catalogID := sdk.String(bitbucketCatalogID)
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

func (r *IntegrationBitbucketResourceModel) ToUpdateSDKType() (*shared.Connector, bool) {
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
	    DisplayName: sdk.String("Bitbucket"),
		AppID:     sdk.String(r.AppID.ValueString()),
		CatalogID: sdk.String(bitbucketCatalogID),
		ID:        sdk.String(r.ID.ValueString()),
		UserIds:   userIds,
		Config: makeConnectorConfig(configOut),
	}

	return &out, configSet
}

func (r *IntegrationBitbucketResourceModel) populateConfig() map[string]*string {
     bitbucketConsumerKey := new(string)
if !r.BitbucketConsumerKey.IsUnknown() && !r.BitbucketConsumerKey.IsNull() {
*bitbucketConsumerKey = r.BitbucketConsumerKey.ValueString()
} else {
bitbucketConsumerKey = nil
}

        bitbucketConsumerSecret := new(string)
if !r.BitbucketConsumerSecret.IsUnknown() && !r.BitbucketConsumerSecret.IsNull() {
*bitbucketConsumerSecret = r.BitbucketConsumerSecret.ValueString()
} else {
bitbucketConsumerSecret = nil
}

        

    	configValues := map[string]*string{
    	"bitbucket_consumer_key": bitbucketConsumerKey,
"bitbucket_consumer_secret": bitbucketConsumerSecret,

    	}

    	return configValues
}

func (r *IntegrationBitbucketResourceModel) getConfig() (map[string]string, bool) {
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

func (r *IntegrationBitbucketResourceModel) ToGetSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationBitbucketResourceModel) ToDeleteSDKType() *shared.ConnectorServiceCreateDelegatedRequest {
	out := r.ToCreateDelegatedSDKType()
	return out
}

func (r *IntegrationBitbucketResourceModel) RefreshFromGetResponse(resp *shared.Connector) {
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
               if v, ok := values["bitbucket_consumer_key"]; ok {
r.BitbucketConsumerKey = types.StringValue(v.(string))
}

               
               
           }
       }
    }
}

func (r *IntegrationBitbucketResourceModel) RefreshFromUpdateResponse(resp *shared.Connector) {
	r.RefreshFromGetResponse(resp)
}

func (r *IntegrationBitbucketResourceModel) RefreshFromCreateResponse(resp *shared.Connector) {
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
                  if v, ok := values["bitbucket_consumer_key"]; ok {
r.BitbucketConsumerKey = types.StringValue(v.(string))
}

                  
                  
              }
          }
       }
}
