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
				var doNotSave *shared.DoNotSave
				if r.ProvisionPolicy.ConnectorProvision.AccountProvision.DoNotSave != nil {
					doNotSave = &shared.DoNotSave{}
				}
				var saveToVault *shared.SaveToVault
				if r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault != nil {
					var vaultIds []string
					if r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds != nil {
						vaultIds = make([]string, 0, len(r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds))
						for _, vaultIdsItem := range r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds {
							vaultIds = append(vaultIds, vaultIdsItem.ValueString())
						}
					}
					saveToVault = &shared.SaveToVault{
						VaultIds: vaultIds,
					}
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
					DoNotSave:   doNotSave,
					SaveToVault: saveToVault,
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
			var deleteAccount *shared.DeleteAccount
			if r.ProvisionPolicy.ConnectorProvision.DeleteAccount != nil {
				connectorId2 := new(string)
				if !r.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsNull() {
					*connectorId2 = r.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID.ValueString()
				} else {
					connectorId2 = nil
				}
				deleteAccount = &shared.DeleteAccount{
					ConnectorID: connectorId2,
				}
			}
			connectorProvision = &shared.ConnectorProvision{
				AccountProvision: accountProvision,
				DefaultBehavior:  defaultBehavior,
				DeleteAccount:    deleteAccount,
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
			connectorId3 := new(string)
			if !r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.IsNull() {
				*connectorId3 = r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.ValueString()
			} else {
				connectorId3 = nil
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
				ConnectorID:                       connectorId3,
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
			var userIds []string
			if r.ProvisionPolicy.ManualProvision.UserIds != nil {
				userIds = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.UserIds))
				for _, userIdsItem := range r.ProvisionPolicy.ManualProvision.UserIds {
					userIds = append(userIds, userIdsItem.ValueString())
				}
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
		if r.ProvisionPolicy.UnconfiguredProvision != nil {
			unconfiguredProvision = &shared.UnconfiguredProvision{}
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
			UnconfiguredProvision:   unconfiguredProvision,
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
	var deprovisionerPolicy *shared.DeprovisionerPolicy
	if r.DeprovisionerPolicy != nil {
		var connectorProvision1 *shared.ConnectorProvision
		if r.DeprovisionerPolicy.ConnectorProvision != nil {
			var accountProvision1 *shared.AccountProvision
			if r.DeprovisionerPolicy.ConnectorProvision.AccountProvision != nil {
				var config1 *shared.AccountProvisionConfig
				if r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config != nil {
					config1 = &shared.AccountProvisionConfig{}
				}
				connectorId4 := new(string)
				if !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsNull() {
					*connectorId4 = r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID.ValueString()
				} else {
					connectorId4 = nil
				}
				var doNotSave1 *shared.DoNotSave
				if r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.DoNotSave != nil {
					doNotSave1 = &shared.DoNotSave{}
				}
				var saveToVault1 *shared.SaveToVault
				if r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault != nil {
					var vaultIds1 []string
					if r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds != nil {
						vaultIds1 = make([]string, 0, len(r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds))
						for _, vaultIdsItem1 := range r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds {
							vaultIds1 = append(vaultIds1, vaultIdsItem1.ValueString())
						}
					}
					saveToVault1 = &shared.SaveToVault{
						VaultIds: vaultIds1,
					}
				}
				schemaId1 := new(string)
				if !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SchemaID.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SchemaID.IsNull() {
					*schemaId1 = r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SchemaID.ValueString()
				} else {
					schemaId1 = nil
				}
				accountProvision1 = &shared.AccountProvision{
					Config:      config1,
					ConnectorID: connectorId4,
					DoNotSave:   doNotSave1,
					SaveToVault: saveToVault1,
					SchemaID:    schemaId1,
				}
			}
			var defaultBehavior1 *shared.DefaultBehavior
			if r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior != nil {
				connectorId5 := new(string)
				if !r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsNull() {
					*connectorId5 = r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.ValueString()
				} else {
					connectorId5 = nil
				}
				defaultBehavior1 = &shared.DefaultBehavior{
					ConnectorID: connectorId5,
				}
			}
			var deleteAccount1 *shared.DeleteAccount
			if r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount != nil {
				connectorId6 := new(string)
				if !r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsNull() {
					*connectorId6 = r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID.ValueString()
				} else {
					connectorId6 = nil
				}
				deleteAccount1 = &shared.DeleteAccount{
					ConnectorID: connectorId6,
				}
			}
			connectorProvision1 = &shared.ConnectorProvision{
				AccountProvision: accountProvision1,
				DefaultBehavior:  defaultBehavior1,
				DeleteAccount:    deleteAccount1,
			}
		}
		var delegatedProvision1 *shared.DelegatedProvision
		if r.DeprovisionerPolicy.DelegatedProvision != nil {
			appId3 := new(string)
			if !r.DeprovisionerPolicy.DelegatedProvision.AppID.IsUnknown() && !r.DeprovisionerPolicy.DelegatedProvision.AppID.IsNull() {
				*appId3 = r.DeprovisionerPolicy.DelegatedProvision.AppID.ValueString()
			} else {
				appId3 = nil
			}
			entitlementId1 := new(string)
			if !r.DeprovisionerPolicy.DelegatedProvision.EntitlementID.IsUnknown() && !r.DeprovisionerPolicy.DelegatedProvision.EntitlementID.IsNull() {
				*entitlementId1 = r.DeprovisionerPolicy.DelegatedProvision.EntitlementID.ValueString()
			} else {
				entitlementId1 = nil
			}
			delegatedProvision1 = &shared.DelegatedProvision{
				AppID:         appId3,
				EntitlementID: entitlementId1,
			}
		}
		var externalTicketProvision1 *shared.ExternalTicketProvision
		if r.DeprovisionerPolicy.ExternalTicketProvision != nil {
			appId4 := new(string)
			if !r.DeprovisionerPolicy.ExternalTicketProvision.AppID.IsUnknown() && !r.DeprovisionerPolicy.ExternalTicketProvision.AppID.IsNull() {
				*appId4 = r.DeprovisionerPolicy.ExternalTicketProvision.AppID.ValueString()
			} else {
				appId4 = nil
			}
			connectorId7 := new(string)
			if !r.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID.IsNull() {
				*connectorId7 = r.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID.ValueString()
			} else {
				connectorId7 = nil
			}
			externalTicketProvisionerConfigId1 := new(string)
			if !r.DeprovisionerPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID.IsUnknown() && !r.DeprovisionerPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID.IsNull() {
				*externalTicketProvisionerConfigId1 = r.DeprovisionerPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID.ValueString()
			} else {
				externalTicketProvisionerConfigId1 = nil
			}
			instructions2 := new(string)
			if !r.DeprovisionerPolicy.ExternalTicketProvision.Instructions.IsUnknown() && !r.DeprovisionerPolicy.ExternalTicketProvision.Instructions.IsNull() {
				*instructions2 = r.DeprovisionerPolicy.ExternalTicketProvision.Instructions.ValueString()
			} else {
				instructions2 = nil
			}
			externalTicketProvision1 = &shared.ExternalTicketProvision{
				AppID:                             appId4,
				ConnectorID:                       connectorId7,
				ExternalTicketProvisionerConfigID: externalTicketProvisionerConfigId1,
				Instructions:                      instructions2,
			}
		}
		var manualProvision1 *shared.ManualProvision
		if r.DeprovisionerPolicy.ManualProvision != nil {
			instructions3 := new(string)
			if !r.DeprovisionerPolicy.ManualProvision.Instructions.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.Instructions.IsNull() {
				*instructions3 = r.DeprovisionerPolicy.ManualProvision.Instructions.ValueString()
			} else {
				instructions3 = nil
			}
			var userIds1 []string
			if r.DeprovisionerPolicy.ManualProvision.UserIds != nil {
				userIds1 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.UserIds))
				for _, userIdsItem1 := range r.DeprovisionerPolicy.ManualProvision.UserIds {
					userIds1 = append(userIds1, userIdsItem1.ValueString())
				}
			}
			manualProvision1 = &shared.ManualProvision{
				Instructions: instructions3,
				UserIds:      userIds1,
			}
		}
		var multiStep1 interface{}
		if !r.DeprovisionerPolicy.MultiStep.IsUnknown() && !r.DeprovisionerPolicy.MultiStep.IsNull() {
			_ = json.Unmarshal([]byte(r.DeprovisionerPolicy.MultiStep.ValueString()), &multiStep1)
		}
		var unconfiguredProvision1 *shared.UnconfiguredProvision
		if r.DeprovisionerPolicy.UnconfiguredProvision != nil {
			unconfiguredProvision1 = &shared.UnconfiguredProvision{}
		}
		var webhookProvision1 *shared.WebhookProvision
		if r.DeprovisionerPolicy.WebhookProvision != nil {
			webhookId1 := new(string)
			if !r.DeprovisionerPolicy.WebhookProvision.WebhookID.IsUnknown() && !r.DeprovisionerPolicy.WebhookProvision.WebhookID.IsNull() {
				*webhookId1 = r.DeprovisionerPolicy.WebhookProvision.WebhookID.ValueString()
			} else {
				webhookId1 = nil
			}
			webhookProvision1 = &shared.WebhookProvision{
				WebhookID: webhookId1,
			}
		}
		deprovisionerPolicy = &shared.DeprovisionerPolicy{
			ConnectorProvision:      connectorProvision1,
			DelegatedProvision:      delegatedProvision1,
			ExternalTicketProvision: externalTicketProvision1,
			ManualProvision:         manualProvision1,
			MultiStep:               multiStep1,
			UnconfiguredProvision:   unconfiguredProvision1,
			WebhookProvision:        webhookProvision1,
		}
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
		DeprovisionerPolicy:            deprovisionerPolicy,
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
	if resp.DeprovisionerPolicy == nil {
		r.DeprovisionerPolicy = nil
	} else {
		r.DeprovisionerPolicy = &tfTypes.DeprovisionerPolicy{}
		if resp.DeprovisionerPolicy.ConnectorProvision == nil {
			r.DeprovisionerPolicy.ConnectorProvision = nil
		} else {
			r.DeprovisionerPolicy.ConnectorProvision = &tfTypes.ConnectorProvision{}
			if resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision == nil {
				r.DeprovisionerPolicy.ConnectorProvision.AccountProvision = nil
			} else {
				r.DeprovisionerPolicy.ConnectorProvision.AccountProvision = &tfTypes.AccountProvision{}
				if resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config == nil {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config = nil
				} else {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config = &tfTypes.AccountProvisionConfig{}
				}
				r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID = types.StringPointerValue(resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID)
				if resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.DoNotSave == nil {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.DoNotSave = nil
				} else {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.DoNotSave = &tfTypes.DoNotSave{}
				}
				if resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault == nil {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault = nil
				} else {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault = &tfTypes.SaveToVault{}
					if resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds != nil {
						r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds))
						for _, v := range resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds {
							r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds = append(r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds, types.StringValue(v))
						}
					}
				}
				r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SchemaID = types.StringPointerValue(resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.SchemaID)
			}
			if resp.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior == nil {
				r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior = nil
			} else {
				r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior = &tfTypes.DefaultBehavior{}
				r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID = types.StringPointerValue(resp.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID)
			}
			if resp.DeprovisionerPolicy.ConnectorProvision.DeleteAccount == nil {
				r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount = nil
			} else {
				r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount = &tfTypes.DeleteAccount{}
				r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID = types.StringPointerValue(resp.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID)
			}
		}
		if resp.DeprovisionerPolicy.DelegatedProvision == nil {
			r.DeprovisionerPolicy.DelegatedProvision = nil
		} else {
			r.DeprovisionerPolicy.DelegatedProvision = &tfTypes.DelegatedProvision{}
			r.DeprovisionerPolicy.DelegatedProvision.AppID = types.StringPointerValue(resp.DeprovisionerPolicy.DelegatedProvision.AppID)
			r.DeprovisionerPolicy.DelegatedProvision.EntitlementID = types.StringPointerValue(resp.DeprovisionerPolicy.DelegatedProvision.EntitlementID)
		}
		if resp.DeprovisionerPolicy.ExternalTicketProvision == nil {
			r.DeprovisionerPolicy.ExternalTicketProvision = nil
		} else {
			r.DeprovisionerPolicy.ExternalTicketProvision = &tfTypes.ExternalTicketProvision{}
			r.DeprovisionerPolicy.ExternalTicketProvision.AppID = types.StringPointerValue(resp.DeprovisionerPolicy.ExternalTicketProvision.AppID)
			r.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID = types.StringPointerValue(resp.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID)
			r.DeprovisionerPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID = types.StringPointerValue(resp.DeprovisionerPolicy.ExternalTicketProvision.ExternalTicketProvisionerConfigID)
			r.DeprovisionerPolicy.ExternalTicketProvision.Instructions = types.StringPointerValue(resp.DeprovisionerPolicy.ExternalTicketProvision.Instructions)
		}
		if resp.DeprovisionerPolicy.ManualProvision == nil {
			r.DeprovisionerPolicy.ManualProvision = nil
		} else {
			r.DeprovisionerPolicy.ManualProvision = &tfTypes.ManualProvision{}
			r.DeprovisionerPolicy.ManualProvision.Instructions = types.StringPointerValue(resp.DeprovisionerPolicy.ManualProvision.Instructions)
			if resp.DeprovisionerPolicy.ManualProvision.UserIds != nil {
				r.DeprovisionerPolicy.ManualProvision.UserIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.UserIds))
				for _, v := range resp.DeprovisionerPolicy.ManualProvision.UserIds {
					r.DeprovisionerPolicy.ManualProvision.UserIds = append(r.DeprovisionerPolicy.ManualProvision.UserIds, types.StringValue(v))
				}
			}
		}
		if resp.DeprovisionerPolicy.MultiStep == nil {
			r.DeprovisionerPolicy.MultiStep = jsontypes.NewNormalizedNull()
		} else {
			multiStepResult, _ := json.Marshal(resp.DeprovisionerPolicy.MultiStep)
			r.DeprovisionerPolicy.MultiStep = jsontypes.NewNormalizedValue(string(multiStepResult))
		}
		if resp.DeprovisionerPolicy.UnconfiguredProvision == nil {
			r.DeprovisionerPolicy.UnconfiguredProvision = nil
		} else {
			r.DeprovisionerPolicy.UnconfiguredProvision = &tfTypes.UnconfiguredProvision{}
		}
		if resp.DeprovisionerPolicy.WebhookProvision == nil {
			r.DeprovisionerPolicy.WebhookProvision = nil
		} else {
			r.DeprovisionerPolicy.WebhookProvision = &tfTypes.WebhookProvision{}
			r.DeprovisionerPolicy.WebhookProvision.WebhookID = types.StringPointerValue(resp.DeprovisionerPolicy.WebhookProvision.WebhookID)
		}
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
				if resp.ProvisionPolicy.ConnectorProvision.AccountProvision.DoNotSave == nil {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision.DoNotSave = nil
				} else {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision.DoNotSave = &tfTypes.DoNotSave{}
				}
				if resp.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault == nil {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault = nil
				} else {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault = &tfTypes.SaveToVault{}
					if resp.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds != nil {
						r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds = make([]types.String, 0, len(resp.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds))
						for _, v := range resp.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds {
							r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds = append(r.ProvisionPolicy.ConnectorProvision.AccountProvision.SaveToVault.VaultIds, types.StringValue(v))
						}
					}
				}
				r.ProvisionPolicy.ConnectorProvision.AccountProvision.SchemaID = types.StringPointerValue(resp.ProvisionPolicy.ConnectorProvision.AccountProvision.SchemaID)
			}
			if resp.ProvisionPolicy.ConnectorProvision.DefaultBehavior == nil {
				r.ProvisionPolicy.ConnectorProvision.DefaultBehavior = nil
			} else {
				r.ProvisionPolicy.ConnectorProvision.DefaultBehavior = &tfTypes.DefaultBehavior{}
				r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID = types.StringPointerValue(resp.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID)
			}
			if resp.ProvisionPolicy.ConnectorProvision.DeleteAccount == nil {
				r.ProvisionPolicy.ConnectorProvision.DeleteAccount = nil
			} else {
				r.ProvisionPolicy.ConnectorProvision.DeleteAccount = &tfTypes.DeleteAccount{}
				r.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID = types.StringPointerValue(resp.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID)
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
			r.ProvisionPolicy.MultiStep = jsontypes.NewNormalizedNull()
		} else {
			multiStepResult1, _ := json.Marshal(resp.ProvisionPolicy.MultiStep)
			r.ProvisionPolicy.MultiStep = jsontypes.NewNormalizedValue(string(multiStepResult1))
		}
		if resp.ProvisionPolicy.UnconfiguredProvision == nil {
			r.ProvisionPolicy.UnconfiguredProvision = nil
		} else {
			r.ProvisionPolicy.UnconfiguredProvision = &tfTypes.UnconfiguredProvision{}
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
