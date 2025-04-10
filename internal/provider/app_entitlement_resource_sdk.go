package provider

import (
	"encoding/json"
	"time"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppEntitlementResourceModel) ToUpdateSDKType() *shared.AppEntitlementInput {
	var provisionPolicy *shared.ProvisionPolicy
	if r.ProvisionPolicy != nil {
		var connectorProvision *shared.ConnectorProvision
		if r.ProvisionPolicy.ConnectorProvision != nil {
			var accountProvision *shared.AccountProvision
			if r.ProvisionPolicy.ConnectorProvision.AccountProvision != nil {
				var config *shared.AccountProvisionConfig
				if r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config != nil {
					config = &shared.AccountProvisionConfig{}
				}
				connectorID := new(string)
				if !r.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsNull() {
					*connectorID = r.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID.ValueString()
				} else {
					connectorID = nil
				}
				schemaID := new(string)
				if !r.ProvisionPolicy.ConnectorProvision.AccountProvision.SchemaID.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.AccountProvision.SchemaID.IsNull() {
					*schemaID = r.ProvisionPolicy.ConnectorProvision.AccountProvision.SchemaID.ValueString()
				} else {
					schemaID = nil
				}
				accountProvision = &shared.AccountProvision{
					Config:      config,
					ConnectorID: connectorID,
					SchemaID:    schemaID,
				}
			}
			var defaultBehavior *shared.DefaultBehavior
			if r.ProvisionPolicy.ConnectorProvision.DefaultBehavior != nil {
				connectorId1 := new(string)
				if !r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsNull() {
					*connectorId1 = r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.ValueString()
				} else {
					connectorId1 = nil
				}
				defaultBehavior = &shared.DefaultBehavior{
					ConnectorID: connectorId1,
				}
			}
			connectorProvision = &shared.ConnectorProvision{
				AccountProvision: accountProvision,
				DefaultBehavior:  defaultBehavior,
			}
		}
		var delegatedProvision *shared.DelegatedProvision
		if r.ProvisionPolicy.DelegatedProvision != nil {
			appId1 := new(string)
			if !r.ProvisionPolicy.DelegatedProvision.AppID.IsUnknown() && !r.ProvisionPolicy.DelegatedProvision.AppID.IsNull() {
				*appId1 = r.ProvisionPolicy.DelegatedProvision.AppID.ValueString()
			} else {
				appId1 = nil
			}
			entitlementID := new(string)
			if !r.ProvisionPolicy.DelegatedProvision.EntitlementID.IsUnknown() && !r.ProvisionPolicy.DelegatedProvision.EntitlementID.IsNull() {
				*entitlementID = r.ProvisionPolicy.DelegatedProvision.EntitlementID.ValueString()
			} else {
				entitlementID = nil
			}
			delegatedProvision = &shared.DelegatedProvision{
				AppID:         appId1,
				EntitlementID: entitlementID,
			}
		}
		var externalTicketProvision *shared.ExternalTicketProvision
		if r.ProvisionPolicy.ExternalTicketProvision != nil {
			appId2 := new(string)
			if !r.ProvisionPolicy.ExternalTicketProvision.AppID.IsUnknown() && !r.ProvisionPolicy.ExternalTicketProvision.AppID.IsNull() {
				*appId2 = r.ProvisionPolicy.ExternalTicketProvision.AppID.ValueString()
			} else {
				appId2 = nil
			}
			connectorId2 := new(string)
			if !r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.IsNull() {
				*connectorId2 = r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.ValueString()
			} else {
				connectorId2 = nil
			}
			externalTicketProvisionerConfigID := new(string)
			if !r.ProvisionPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID.IsUnknown() && !r.ProvisionPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID.IsNull() {
				*externalTicketProvisionerConfigID = r.ProvisionPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID.ValueString()
			} else {
				externalTicketProvisionerConfigID = nil
			}
			instructions := new(string)
			if !r.ProvisionPolicy.ExternalTicketProvision.Instructions.IsUnknown() && !r.ProvisionPolicy.ExternalTicketProvision.Instructions.IsNull() {
				*instructions = r.ProvisionPolicy.ExternalTicketProvision.Instructions.ValueString()
			} else {
				instructions = nil
			}
			externalTicketProvision = &shared.ExternalTicketProvision{
				AppID:                             appId2,
				ConnectorID:                       connectorId2,
				ExternalTicketProvisionerConfigID: externalTicketProvisionerConfigID,
				Instructions:                      instructions,
			}
		}
		var manualProvision *shared.ManualProvision
		if r.ProvisionPolicy.ManualProvision != nil {
			instructions1 := new(string)
			if !r.ProvisionPolicy.ManualProvision.Instructions.IsUnknown() && !r.ProvisionPolicy.ManualProvision.Instructions.IsNull() {
				*instructions1 = r.ProvisionPolicy.ManualProvision.Instructions.ValueString()
			} else {
				instructions1 = nil
			}
			var userIds []string = []string{}
			for _, userIdsItem := range r.ProvisionPolicy.ManualProvision.UserIds {
				userIds = append(userIds, userIdsItem.ValueString())
			}
			manualProvision = &shared.ManualProvision{
				Instructions: instructions1,
				UserIds:      userIds,
			}
		}
		var multiStep interface{}
		if !r.ProvisionPolicy.MultiStep.IsUnknown() && !r.ProvisionPolicy.MultiStep.IsNull() {
			_ = json.Unmarshal([]byte(r.ProvisionPolicy.MultiStep.ValueString()), &multiStep)
		}
		var webhookProvision *shared.WebhookProvision
		if r.ProvisionPolicy.WebhookProvision != nil {
			webhookID := new(string)
			if !r.ProvisionPolicy.WebhookProvision.WebhookID.IsUnknown() && !r.ProvisionPolicy.WebhookProvision.WebhookID.IsNull() {
				*webhookID = r.ProvisionPolicy.WebhookProvision.WebhookID.ValueString()
			} else {
				webhookID = nil
			}
			webhookProvision = &shared.WebhookProvision{
				WebhookID: webhookID,
			}
		}
		provisionPolicy = &shared.ProvisionPolicy{
			ConnectorProvision:      connectorProvision,
			DelegatedProvision:      delegatedProvision,
			ExternalTicketProvision: externalTicketProvision,
			ManualProvision:         manualProvision,
			MultiStep:               multiStep,
			WebhookProvision:        webhookProvision,
		}
	}
	alias := new(string)
	if !r.Alias.IsUnknown() && !r.Alias.IsNull() {
		*alias = r.Alias.ValueString()
	} else {
		alias = nil
	}
	appId1 := new(string)
	if !r.AppID.IsUnknown() && !r.AppID.IsNull() {
		*appId1 = r.AppID.ValueString()
	} else {
		appId1 = nil
	}
	appResourceID := new(string)
	if !r.AppResourceID.IsUnknown() && !r.AppResourceID.IsNull() {
		*appResourceID = r.AppResourceID.ValueString()
	} else {
		appResourceID = nil
	}
	appResourceTypeID := new(string)
	if !r.AppResourceTypeID.IsUnknown() && !r.AppResourceTypeID.IsNull() {
		*appResourceTypeID = r.AppResourceTypeID.ValueString()
	} else {
		appResourceTypeID = nil
	}
	certifyPolicyID := new(string)
	if !r.CertifyPolicyID.IsUnknown() && !r.CertifyPolicyID.IsNull() {
		*certifyPolicyID = r.CertifyPolicyID.ValueString()
	} else {
		certifyPolicyID = nil
	}
	complianceFrameworkValueIds := make([]string, 0)
	for _, complianceFrameworkValueIdsItem := range r.ComplianceFrameworkValueIds {
		complianceFrameworkValueIds = append(complianceFrameworkValueIds, complianceFrameworkValueIdsItem.ValueString())
	}
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
	var durationUnset *shared.AppEntitlementDurationUnset
	if r.DurationUnset != nil {
		durationUnset = &shared.AppEntitlementDurationUnset{}
	}
	durationGrant := new(string)
	if !r.DurationGrant.IsUnknown() && !r.DurationGrant.IsNull() {
		*durationGrant = r.DurationGrant.ValueString()
	} else {
		durationGrant = nil
	}
	emergencyGrantEnabled := new(bool)
	if !r.EmergencyGrantEnabled.IsUnknown() && !r.EmergencyGrantEnabled.IsNull() {
		*emergencyGrantEnabled = r.EmergencyGrantEnabled.ValueBool()
	} else {
		emergencyGrantEnabled = nil
	}
	emergencyGrantPolicyID := new(string)
	if !r.EmergencyGrantPolicyID.IsUnknown() && !r.EmergencyGrantPolicyID.IsNull() {
		*emergencyGrantPolicyID = r.EmergencyGrantPolicyID.ValueString()
	} else {
		emergencyGrantPolicyID = nil
	}
	grantPolicyID := new(string)
	if !r.GrantPolicyID.IsUnknown() && !r.GrantPolicyID.IsNull() {
		*grantPolicyID = r.GrantPolicyID.ValueString()
	} else {
		grantPolicyID = nil
	}
	revokePolicyID := new(string)
	if !r.RevokePolicyID.IsUnknown() && !r.RevokePolicyID.IsNull() {
		*revokePolicyID = r.RevokePolicyID.ValueString()
	} else {
		revokePolicyID = nil
	}
	riskLevelValueID := new(string)
	if !r.RiskLevelValueID.IsUnknown() && !r.RiskLevelValueID.IsNull() {
		*riskLevelValueID = r.RiskLevelValueID.ValueString()
	} else {
		riskLevelValueID = nil
	}
	slug := new(string)
	if !r.Slug.IsUnknown() && !r.Slug.IsNull() {
		*slug = r.Slug.ValueString()
	} else {
		slug = nil
	}
	overrideAccessRequestsDefaults := new(bool)
	if !r.OverrideAccessRequestsDefaults.IsUnknown() && !r.OverrideAccessRequestsDefaults.IsNull() {
		*overrideAccessRequestsDefaults = r.OverrideAccessRequestsDefaults.ValueBool()
	} else {
		overrideAccessRequestsDefaults = nil
	}

	out := shared.AppEntitlementInput{
		ProvisionPolicy:                provisionPolicy,
		Alias:                          alias,
		AppID:                          appId1,
		AppResourceID:                  appResourceID,
		AppResourceTypeID:              appResourceTypeID,
		CertifyPolicyID:                certifyPolicyID,
		ComplianceFrameworkValueIds:    complianceFrameworkValueIds,
		Description:                    description,
		DisplayName:                    displayName,
		DurationGrant:                  durationGrant,
		DurationUnset:                  durationUnset,
		EmergencyGrantEnabled:          emergencyGrantEnabled,
		EmergencyGrantPolicyID:         emergencyGrantPolicyID,
		GrantPolicyID:                  grantPolicyID,
		RevokePolicyID:                 revokePolicyID,
		RiskLevelValueID:               riskLevelValueID,
		Slug:                           slug,
		OverrideAccessRequestsDefaults: overrideAccessRequestsDefaults,
	}
	return &out
}

func (r *AppEntitlementResourceModel) RefreshFromGetResponse(resp *shared.AppEntitlement) {
	if resp.Alias != nil {
		r.Alias = types.StringValue(*resp.Alias)
	} else {
		r.Alias = types.StringNull()
	}
	if resp.AppID != nil {
		r.AppID = types.StringValue(*resp.AppID)
	} else {
		r.AppID = types.StringNull()
	}
	if resp.AppResourceID != nil {
		r.AppResourceID = types.StringValue(*resp.AppResourceID)
	} else {
		r.AppResourceID = types.StringNull()
	}
	if resp.AppResourceTypeID != nil {
		r.AppResourceTypeID = types.StringValue(*resp.AppResourceTypeID)
	} else {
		r.AppResourceTypeID = types.StringNull()
	}
	if resp.CertifyPolicyID != nil {
		r.CertifyPolicyID = types.StringValue(*resp.CertifyPolicyID)
	} else {
		r.CertifyPolicyID = types.StringNull()
	}
	r.ComplianceFrameworkValueIds = nil
	for _, v := range resp.ComplianceFrameworkValueIds {
		r.ComplianceFrameworkValueIds = append(r.ComplianceFrameworkValueIds, types.StringValue(v))
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
	if resp.DurationGrant != nil {
		r.DurationGrant = types.StringValue(*resp.DurationGrant)
	} else {
		r.DurationGrant = types.StringNull()
	}
	if resp.DurationUnset == nil {
		r.DurationUnset = nil
	} else {
		r.DurationUnset = &tfTypes.CreateAppEntitlementRequestDurationUnset{}
	}
	if resp.EmergencyGrantEnabled != nil {
		r.EmergencyGrantEnabled = types.BoolValue(*resp.EmergencyGrantEnabled)
	} else {
		r.EmergencyGrantEnabled = types.BoolNull()
	}
	if resp.EmergencyGrantPolicyID != nil {
		r.EmergencyGrantPolicyID = types.StringValue(*resp.EmergencyGrantPolicyID)
	} else {
		r.EmergencyGrantPolicyID = types.StringNull()
	}
	if resp.GrantPolicyID != nil {
		r.GrantPolicyID = types.StringValue(*resp.GrantPolicyID)
	} else {
		r.GrantPolicyID = types.StringNull()
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}

	if resp.ProvisionPolicy == nil {
		r.ProvisionPolicy = nil
	} else {
		r.ProvisionPolicy = &tfTypes.ProvisionPolicy{}
		if resp.ProvisionPolicy.ConnectorProvision == nil {
			r.ProvisionPolicy.ConnectorProvision = nil
		} else {
			r.ProvisionPolicy.ConnectorProvision = &tfTypes.ConnectorProvision{}
			if resp.ProvisionPolicy.ConnectorProvision.AccountProvision == nil {
				r.ProvisionPolicy.ConnectorProvision.AccountProvision = nil
			} else {
				r.ProvisionPolicy.ConnectorProvision.AccountProvision = &tfTypes.AccountProvision{}
				if resp.ProvisionPolicy.ConnectorProvision.AccountProvision.Config == nil {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config = nil
				} else {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config = &tfTypes.AccountProvisionConfig{}
				}
				r.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID = types.StringPointerValue(resp.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID)
				r.ProvisionPolicy.ConnectorProvision.AccountProvision.SchemaID = types.StringPointerValue(resp.ProvisionPolicy.ConnectorProvision.AccountProvision.SchemaID)
			}
			if resp.ProvisionPolicy.ConnectorProvision.DefaultBehavior == nil {
				r.ProvisionPolicy.ConnectorProvision.DefaultBehavior = nil
			} else {
				r.ProvisionPolicy.ConnectorProvision.DefaultBehavior = &tfTypes.DefaultBehavior{}
				r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID = types.StringPointerValue(resp.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID)
			}
		}
		if resp.ProvisionPolicy.DelegatedProvision == nil {
			r.ProvisionPolicy.DelegatedProvision = nil
		} else {
			r.ProvisionPolicy.DelegatedProvision = &tfTypes.DelegatedProvision{}
			r.ProvisionPolicy.DelegatedProvision.AppID = types.StringPointerValue(resp.ProvisionPolicy.DelegatedProvision.AppID)
			r.ProvisionPolicy.DelegatedProvision.EntitlementID = types.StringPointerValue(resp.ProvisionPolicy.DelegatedProvision.EntitlementID)
		}
		if resp.ProvisionPolicy.ExternalTicketProvision == nil {
			r.ProvisionPolicy.ExternalTicketProvision = nil
		} else {
			r.ProvisionPolicy.ExternalTicketProvision = &tfTypes.ExternalTicketProvision{}
			r.ProvisionPolicy.ExternalTicketProvision.AppID = types.StringPointerValue(resp.ProvisionPolicy.ExternalTicketProvision.AppID)
			r.ProvisionPolicy.ExternalTicketProvision.ConnectorID = types.StringPointerValue(resp.ProvisionPolicy.ExternalTicketProvision.ConnectorID)
			r.ProvisionPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID = types.StringPointerValue(resp.ProvisionPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID)
			r.ProvisionPolicy.ExternalTicketProvision.Instructions = types.StringPointerValue(resp.ProvisionPolicy.ExternalTicketProvision.Instructions)
		}
		if resp.ProvisionPolicy.ManualProvision == nil {
			r.ProvisionPolicy.ManualProvision = nil
		} else {
			r.ProvisionPolicy.ManualProvision = &tfTypes.ManualProvision{}
			r.ProvisionPolicy.ManualProvision.Instructions = types.StringPointerValue(resp.ProvisionPolicy.ManualProvision.Instructions)
			if resp.ProvisionPolicy.ManualProvision.UserIds != nil {
				r.ProvisionPolicy.ManualProvision.UserIds = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.UserIds))
				for _, v := range resp.ProvisionPolicy.ManualProvision.UserIds {
					r.ProvisionPolicy.ManualProvision.UserIds = append(r.ProvisionPolicy.ManualProvision.UserIds, types.StringValue(v))
				}
			}
		}
		if resp.ProvisionPolicy.MultiStep == nil {
			r.ProvisionPolicy.MultiStep = types.StringNull()
		} else {
			multiStepResult1, _ := json.Marshal(resp.ProvisionPolicy.MultiStep)
			r.ProvisionPolicy.MultiStep = types.StringValue(string(multiStepResult1))
		}
		if resp.ProvisionPolicy.WebhookProvision == nil {
			r.ProvisionPolicy.WebhookProvision = nil
		} else {
			r.ProvisionPolicy.WebhookProvision = &tfTypes.WebhookProvision{}
			r.ProvisionPolicy.WebhookProvision.WebhookID = types.StringPointerValue(resp.ProvisionPolicy.WebhookProvision.WebhookID)
		}
	}
	if resp.RevokePolicyID != nil {
		r.RevokePolicyID = types.StringValue(*resp.RevokePolicyID)
	} else {
		r.RevokePolicyID = types.StringNull()
	}
	if resp.RiskLevelValueID != nil {
		r.RiskLevelValueID = types.StringValue(*resp.RiskLevelValueID)
	} else {
		r.RiskLevelValueID = types.StringNull()
	}
	if resp.Slug != nil {
		r.Slug = types.StringValue(*resp.Slug)
	} else {
		r.Slug = types.StringNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
	if resp.OverrideAccessRequestsDefaults != nil {
		r.OverrideAccessRequestsDefaults = types.BoolValue(*resp.OverrideAccessRequestsDefaults)
	} else {
		r.OverrideAccessRequestsDefaults = types.BoolNull()
	}
}

func (r *AppEntitlementResourceModel) RefreshFromCreateResponse(resp *shared.AppEntitlement) {
	r.RefreshFromGetResponse(resp)
}

func (r *AppEntitlementResourceModel) RefreshFromUpdateResponse(resp *shared.AppEntitlement) {
	r.RefreshFromGetResponse(resp)
}
