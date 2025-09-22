package provider

import (
	"encoding/json"
	"time"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppEntitlementResourceModel) ToUpdateSDKType() *shared.AppEntitlementInput {
	var provisionPolicy *shared.ProvisionPolicy
	if r.ProvisionPolicy != nil {
		var connectorProvision *shared.ConnectorProvision
		if r.ProvisionPolicy.Connector != nil {
			var accountProvision *shared.AccountProvision
			if r.ProvisionPolicy.Connector.Account != nil {
				var config *shared.Config
				if r.ProvisionPolicy.Connector.Account.Config != nil {
					config = &shared.Config{}
				}
				connectorID := new(string)
				if !r.ProvisionPolicy.Connector.Account.ConnectorID.IsUnknown() && !r.ProvisionPolicy.Connector.Account.ConnectorID.IsNull() {
					*connectorID = r.ProvisionPolicy.Connector.Account.ConnectorID.ValueString()
				} else {
					connectorID = nil
				}
				schemaID := new(string)
				if !r.ProvisionPolicy.Connector.Account.SchemaID.IsUnknown() && !r.ProvisionPolicy.Connector.Account.SchemaID.IsNull() {
					*schemaID = r.ProvisionPolicy.Connector.Account.SchemaID.ValueString()
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
			if r.ProvisionPolicy.Connector.DefaultBehavior != nil {
				connectorId1 := new(string)
				if !r.ProvisionPolicy.Connector.DefaultBehavior.ConnectorID.IsUnknown() && !r.ProvisionPolicy.Connector.DefaultBehavior.ConnectorID.IsNull() {
					*connectorId1 = r.ProvisionPolicy.Connector.DefaultBehavior.ConnectorID.ValueString()
				} else {
					connectorId1 = nil
				}
				defaultBehavior = &shared.DefaultBehavior{
					ConnectorID: connectorId1,
				}
			}
			var deleteAccount *shared.DeleteAccount
			if r.ProvisionPolicy.Connector.DeleteAccount != nil {
				connectorId2 := new(string)
				if !r.ProvisionPolicy.Connector.DeleteAccount.ConnectorID.IsUnknown() && !r.ProvisionPolicy.Connector.DeleteAccount.ConnectorID.IsNull() {
					*connectorId2 = r.ProvisionPolicy.Connector.DeleteAccount.ConnectorID.ValueString()
				} else {
					connectorId2 = nil
				}
				deleteAccount = &shared.DeleteAccount{
					ConnectorID: connectorId2,
				}
			}
			connectorProvision = &shared.ConnectorProvision{
				Account: accountProvision,
				DefaultBehavior:  defaultBehavior,
				DeleteAccount:    deleteAccount,
			}
		}
		var delegatedProvision *shared.DelegatedProvision
		if r.ProvisionPolicy.Delegated != nil {
			appId1 := new(string)
			if !r.ProvisionPolicy.Delegated.AppID.IsUnknown() && !r.ProvisionPolicy.Delegated.AppID.IsNull() {
				*appId1 = r.ProvisionPolicy.Delegated.AppID.ValueString()
			} else {
				appId1 = nil
			}
			entitlementID := new(string)
			if !r.ProvisionPolicy.Delegated.EntitlementID.IsUnknown() && !r.ProvisionPolicy.Delegated.EntitlementID.IsNull() {
				*entitlementID = r.ProvisionPolicy.Delegated.EntitlementID.ValueString()
			} else {
				entitlementID = nil
			}
			delegatedProvision = &shared.DelegatedProvision{
				AppID:         appId1,
				EntitlementID: entitlementID,
			}
		}
		var externalTicketProvision *shared.ExternalTicketProvision
		if r.ProvisionPolicy.ExternalTicket != nil {
			appId2 := new(string)
			if !r.ProvisionPolicy.ExternalTicket.AppID.IsUnknown() && !r.ProvisionPolicy.ExternalTicket.AppID.IsNull() {
				*appId2 = r.ProvisionPolicy.ExternalTicket.AppID.ValueString()
			} else {
				appId2 = nil
			}
			connectorId2 := new(string)
			if !r.ProvisionPolicy.ExternalTicket.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ExternalTicket.ConnectorID.IsNull() {
				*connectorId2 = r.ProvisionPolicy.ExternalTicket.ConnectorID.ValueString()
			} else {
				connectorId2 = nil
			}
			externalTicketProvisionerConfigID := new(string)
			if !r.ProvisionPolicy.ExternalTicket.ExternalTicketProvisionerConfigID.IsUnknown() && !r.ProvisionPolicy.ExternalTicket.ExternalTicketProvisionerConfigID.IsNull() {
				*externalTicketProvisionerConfigID = r.ProvisionPolicy.ExternalTicket.ExternalTicketProvisionerConfigID.ValueString()
			} else {
				externalTicketProvisionerConfigID = nil
			}
			instructions := new(string)
			if !r.ProvisionPolicy.ExternalTicket.Instructions.IsUnknown() && !r.ProvisionPolicy.ExternalTicket.Instructions.IsNull() {
				*instructions = r.ProvisionPolicy.ExternalTicket.Instructions.ValueString()
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
		if r.ProvisionPolicy.Manual != nil {
			instructions1 := new(string)
			if !r.ProvisionPolicy.Manual.Instructions.IsUnknown() && !r.ProvisionPolicy.Manual.Instructions.IsNull() {
				*instructions1 = r.ProvisionPolicy.Manual.Instructions.ValueString()
			} else {
				instructions1 = nil
			}
			var userIds []string = []string{}
			for _, userIdsItem := range r.ProvisionPolicy.Manual.UserIds {
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
		var unconfiguredProvision *shared.UnconfiguredProvision
		if r.ProvisionPolicy.Unconfigured != nil {
			unconfiguredProvision = (*shared.UnconfiguredProvision)(r.ProvisionPolicy.Unconfigured)
		}
		var webhookProvision *shared.WebhookProvision
		if r.ProvisionPolicy.Webhook != nil {
			webhookID := new(string)
			if !r.ProvisionPolicy.Webhook.WebhookID.IsUnknown() && !r.ProvisionPolicy.Webhook.WebhookID.IsNull() {
				*webhookID = r.ProvisionPolicy.Webhook.WebhookID.ValueString()
			} else {
				webhookID = nil
			}
			webhookProvision = &shared.WebhookProvision{
				WebhookID: webhookID,
			}
		}
		provisionPolicy = &shared.ProvisionPolicy{
			Connector:      connectorProvision,
			Delegated:      delegatedProvision,
			ExternalTicket: externalTicketProvision,
			Manual:         manualProvision,
			MultiStep:      multiStep,
			Unconfigured:   unconfiguredProvision,
			Webhook:        webhookProvision,
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
		ProvisionerPolicy:                provisionPolicy,
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

	if resp.ProvisionerPolicy == nil {
		r.ProvisionPolicy = nil
	} else {
		r.ProvisionPolicy = &tfTypes.ProvisionPolicy{}
		if resp.ProvisionerPolicy.Connector == nil {
			r.ProvisionPolicy.Connector = nil
		} else {
			r.ProvisionPolicy.Connector = &tfTypes.ConnectorProvision{}
			if resp.ProvisionerPolicy.Connector.Account == nil {
				r.ProvisionPolicy.Connector.Account = nil
			} else {
				r.ProvisionPolicy.Connector.Account = &tfTypes.AccountProvision{}
				if resp.ProvisionerPolicy.Connector.Account.Config == nil {
					r.ProvisionPolicy.Connector.Account.Config = nil
				} else {
					r.ProvisionPolicy.Connector.Account.Config = &tfTypes.Config{}
				}
				r.ProvisionPolicy.Connector.Account.ConnectorID = types.StringPointerValue(resp.ProvisionerPolicy.Connector.Account.ConnectorID)
				r.ProvisionPolicy.Connector.Account.SchemaID = types.StringPointerValue(resp.ProvisionerPolicy.Connector.Account.SchemaID)
			}
			if resp.ProvisionerPolicy.Connector.DefaultBehavior == nil {
				r.ProvisionPolicy.Connector.DefaultBehavior = nil
			} else {
				r.ProvisionPolicy.Connector.DefaultBehavior = &tfTypes.DefaultBehavior{}
				r.ProvisionPolicy.Connector.DefaultBehavior.ConnectorID = types.StringPointerValue(resp.ProvisionerPolicy.Connector.DefaultBehavior.ConnectorID)
			}
			if resp.ProvisionerPolicy.Connector.DeleteAccount == nil {
				r.ProvisionPolicy.Connector.DeleteAccount = nil
			} else {
				r.ProvisionPolicy.Connector.DeleteAccount = &tfTypes.DeleteAccount{}
				r.ProvisionPolicy.Connector.DeleteAccount.ConnectorID = types.StringPointerValue(resp.ProvisionerPolicy.Connector.DeleteAccount.ConnectorID)
			}
		}
		if resp.ProvisionerPolicy.Delegated == nil {
			r.ProvisionPolicy.Delegated = nil
		} else {
			r.ProvisionPolicy.Delegated = &tfTypes.DelegatedProvision{}
			r.ProvisionPolicy.Delegated.AppID = types.StringPointerValue(resp.ProvisionerPolicy.Delegated.AppID)
			r.ProvisionPolicy.Delegated.EntitlementID = types.StringPointerValue(resp.ProvisionerPolicy.Delegated.EntitlementID)
		}
		if resp.ProvisionerPolicy.ExternalTicket == nil {
			r.ProvisionPolicy.ExternalTicket = nil
		} else {
			r.ProvisionPolicy.ExternalTicket = &tfTypes.ExternalTicketProvision{}
			r.ProvisionPolicy.ExternalTicket.AppID = types.StringPointerValue(resp.ProvisionerPolicy.ExternalTicket.AppID)
			r.ProvisionPolicy.ExternalTicket.ConnectorID = types.StringPointerValue(resp.ProvisionerPolicy.ExternalTicket.ConnectorID)
			r.ProvisionPolicy.ExternalTicket.ExternalTicketProvisionerConfigID = types.StringPointerValue(resp.ProvisionerPolicy.ExternalTicket.ExternalTicketProvisionerConfigID)
			r.ProvisionPolicy.ExternalTicket.Instructions = types.StringPointerValue(resp.ProvisionerPolicy.ExternalTicket.Instructions)
		}
		if resp.ProvisionerPolicy.Manual == nil {
			r.ProvisionPolicy.Manual = nil
		} else {
			r.ProvisionPolicy.Manual = &tfTypes.ManualProvision{}
			r.ProvisionPolicy.Manual.Instructions = types.StringPointerValue(resp.ProvisionerPolicy.Manual.Instructions)
			if resp.ProvisionerPolicy.Manual.UserIds != nil {
				r.ProvisionPolicy.Manual.UserIds = make([]types.String, 0, len(resp.ProvisionerPolicy.Manual.UserIds))
				for _, v := range resp.ProvisionerPolicy.Manual.UserIds {
					r.ProvisionPolicy.Manual.UserIds = append(r.ProvisionPolicy.Manual.UserIds, types.StringValue(v))
				}
			}
		}
		if resp.ProvisionerPolicy.MultiStep == nil {
			r.ProvisionPolicy.MultiStep = jsontypes.NewNormalizedNull()
		} else {
			multiStepResult1, _ := json.Marshal(resp.ProvisionerPolicy.MultiStep)
			r.ProvisionPolicy.MultiStep = jsontypes.NewNormalizedValue((string(multiStepResult1)))
		}
		if resp.ProvisionerPolicy.Unconfigured == nil {
			r.ProvisionPolicy.Unconfigured = nil
		} else {
			r.ProvisionPolicy.Unconfigured = (*tfTypes.UnconfiguredProvision)(resp.ProvisionerPolicy.Unconfigured)
		}
		if resp.ProvisionerPolicy.Webhook == nil {
			r.ProvisionPolicy.Webhook = nil
		} else {
			r.ProvisionPolicy.Webhook = &tfTypes.WebhookProvision{}
			r.ProvisionPolicy.Webhook.WebhookID = types.StringPointerValue(resp.ProvisionerPolicy.Webhook.WebhookID)
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
