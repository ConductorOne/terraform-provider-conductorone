package provider

import (
	"encoding/json"

	"github.com/conductorone/terraform-provider-conductorone/internal/provider/typeconvert"
	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *AppEntitlementResourceModel) ToUpdateSDKType() *shared.AppEntitlementInput {
	alias := new(string)
	if !r.Alias.IsUnknown() && !r.Alias.IsNull() {
		*alias = r.Alias.ValueString()
	} else {
		alias = nil
	}
	appID := new(string)
	if !r.AppID.IsUnknown() && !r.AppID.IsNull() {
		*appID = r.AppID.ValueString()
	} else {
		appID = nil
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
	var complianceFrameworkValueIds []string
	if r.ComplianceFrameworkValueIds != nil {
		complianceFrameworkValueIds = make([]string, 0, len(r.ComplianceFrameworkValueIds))
		for _, complianceFrameworkValueIdsItem := range r.ComplianceFrameworkValueIds {
			complianceFrameworkValueIds = append(complianceFrameworkValueIds, complianceFrameworkValueIdsItem.ValueString())
		}
	}
	defaultValuesApplied := new(bool)
	if !r.DefaultValuesApplied.IsUnknown() && !r.DefaultValuesApplied.IsNull() {
		*defaultValuesApplied = r.DefaultValuesApplied.ValueBool()
	} else {
		defaultValuesApplied = nil
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
	durationGrant := new(string)
	if !r.DurationGrant.IsUnknown() && !r.DurationGrant.IsNull() {
		*durationGrant = r.DurationGrant.ValueString()
	} else {
		durationGrant = nil
	}
	var durationUnset *shared.AppEntitlementDurationUnset
	if r.DurationUnset != nil {
		durationUnset = &shared.AppEntitlementDurationUnset{}
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
	isManuallyManaged := new(bool)
	if !r.IsManuallyManaged.IsUnknown() && !r.IsManuallyManaged.IsNull() {
		*isManuallyManaged = r.IsManuallyManaged.ValueBool()
	} else {
		isManuallyManaged = nil
	}
	matchBatonID := new(string)
	if !r.MatchBatonID.IsUnknown() && !r.MatchBatonID.IsNull() {
		*matchBatonID = r.MatchBatonID.ValueString()
	} else {
		matchBatonID = nil
	}
	overrideAccessRequestsDefaults := new(bool)
	if !r.OverrideAccessRequestsDefaults.IsUnknown() && !r.OverrideAccessRequestsDefaults.IsNull() {
		*overrideAccessRequestsDefaults = r.OverrideAccessRequestsDefaults.ValueBool()
	} else {
		overrideAccessRequestsDefaults = nil
	}
	var provisionPolicy *shared.ProvisionPolicy
	if r.ProvisionPolicy != nil {
		var actionProvision *shared.ActionProvision
		if r.ProvisionPolicy.ActionProvision != nil {
			actionName := new(string)
			if !r.ProvisionPolicy.ActionProvision.ActionName.IsUnknown() && !r.ProvisionPolicy.ActionProvision.ActionName.IsNull() {
				*actionName = r.ProvisionPolicy.ActionProvision.ActionName.ValueString()
			} else {
				actionName = nil
			}
			appId1 := new(string)
			if !r.ProvisionPolicy.ActionProvision.AppID.IsUnknown() && !r.ProvisionPolicy.ActionProvision.AppID.IsNull() {
				*appId1 = r.ProvisionPolicy.ActionProvision.AppID.ValueString()
			} else {
				appId1 = nil
			}
			connectorID := new(string)
			if !r.ProvisionPolicy.ActionProvision.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ActionProvision.ConnectorID.IsNull() {
				*connectorID = r.ProvisionPolicy.ActionProvision.ConnectorID.ValueString()
			} else {
				connectorID = nil
			}
			actionProvision = &shared.ActionProvision{
				ActionName:  actionName,
				AppID:       appId1,
				ConnectorID: connectorID,
			}
		}
		var connectorProvision *shared.ConnectorProvision
		if r.ProvisionPolicy.ConnectorProvision != nil {
			var accountProvision *shared.AccountProvision
			if r.ProvisionPolicy.ConnectorProvision.AccountProvision != nil {
				var config any
				if !r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config.IsNull() {
					_ = json.Unmarshal([]byte(r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config.ValueString()), &config)
				}
				connectorId1 := new(string)
				if !r.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsNull() {
					*connectorId1 = r.ProvisionPolicy.ConnectorProvision.AccountProvision.ConnectorID.ValueString()
				} else {
					connectorId1 = nil
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
					ConnectorID: connectorId1,
					DoNotSave:   doNotSave,
					SaveToVault: saveToVault,
					SchemaID:    schemaID,
				}
			}
			var defaultBehavior *shared.DefaultBehavior
			if r.ProvisionPolicy.ConnectorProvision.DefaultBehavior != nil {
				connectorId2 := new(string)
				if !r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsNull() {
					*connectorId2 = r.ProvisionPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.ValueString()
				} else {
					connectorId2 = nil
				}
				defaultBehavior = &shared.DefaultBehavior{
					ConnectorID: connectorId2,
				}
			}
			var deleteAccount *shared.DeleteAccount
			if r.ProvisionPolicy.ConnectorProvision.DeleteAccount != nil {
				connectorId3 := new(string)
				if !r.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsNull() {
					*connectorId3 = r.ProvisionPolicy.ConnectorProvision.DeleteAccount.ConnectorID.ValueString()
				} else {
					connectorId3 = nil
				}
				deleteAccount = &shared.DeleteAccount{
					ConnectorID: connectorId3,
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
			appId2 := new(string)
			if !r.ProvisionPolicy.DelegatedProvision.AppID.IsUnknown() && !r.ProvisionPolicy.DelegatedProvision.AppID.IsNull() {
				*appId2 = r.ProvisionPolicy.DelegatedProvision.AppID.ValueString()
			} else {
				appId2 = nil
			}
			entitlementID := new(string)
			if !r.ProvisionPolicy.DelegatedProvision.EntitlementID.IsUnknown() && !r.ProvisionPolicy.DelegatedProvision.EntitlementID.IsNull() {
				*entitlementID = r.ProvisionPolicy.DelegatedProvision.EntitlementID.ValueString()
			} else {
				entitlementID = nil
			}
			delegatedProvision = &shared.DelegatedProvision{
				AppID:         appId2,
				EntitlementID: entitlementID,
			}
		}
		var externalTicketProvision *shared.ExternalTicketProvision
		if r.ProvisionPolicy.ExternalTicketProvision != nil {
			appId3 := new(string)
			if !r.ProvisionPolicy.ExternalTicketProvision.AppID.IsUnknown() && !r.ProvisionPolicy.ExternalTicketProvision.AppID.IsNull() {
				*appId3 = r.ProvisionPolicy.ExternalTicketProvision.AppID.ValueString()
			} else {
				appId3 = nil
			}
			connectorId4 := new(string)
			if !r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.IsUnknown() && !r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.IsNull() {
				*connectorId4 = r.ProvisionPolicy.ExternalTicketProvision.ConnectorID.ValueString()
			} else {
				connectorId4 = nil
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
				AppID:                             appId3,
				ConnectorID:                       connectorId4,
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
			var provisionerAssignment *shared.ProvisionerAssignment
			if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment != nil {
				var appOwnerProvisioner *shared.AppOwnerProvisioner
				if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner != nil {
					allowReassignment := new(bool)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment.IsNull() {
						*allowReassignment = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment = nil
					}
					var fallbackUserIds []string
					if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds != nil {
						fallbackUserIds = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds))
						for _, v := range r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds {
							fallbackUserIds = append(fallbackUserIds, v.ValueString())
						}
					}
					appOwnerProvisioner = &shared.AppOwnerProvisioner{
						AllowReassignment: allowReassignment,
						FallbackUserIds:   fallbackUserIds,
					}
				}
				var entitlementOwnerProvisioner *shared.EntitlementOwnerProvisioner
				if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner != nil {
					allowReassignment1 := new(bool)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment.IsNull() {
						*allowReassignment1 = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment1 = nil
					}
					var fallbackUserIds1 []string
					if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds != nil {
						fallbackUserIds1 = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds))
						for _, v := range r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds {
							fallbackUserIds1 = append(fallbackUserIds1, v.ValueString())
						}
					}
					entitlementOwnerProvisioner = &shared.EntitlementOwnerProvisioner{
						AllowReassignment: allowReassignment1,
						FallbackUserIds:   fallbackUserIds1,
					}
				}
				var expressionProvisioner *shared.ExpressionProvisioner
				if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner != nil {
					allowReassignment2 := new(bool)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment.IsNull() {
						*allowReassignment2 = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment2 = nil
					}
					var expressions []string
					if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions != nil {
						expressions = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions))
						for _, v := range r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions {
							expressions = append(expressions, v.ValueString())
						}
					}
					var fallbackUserIds2 []string
					if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds != nil {
						fallbackUserIds2 = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds))
						for _, v := range r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds {
							fallbackUserIds2 = append(fallbackUserIds2, v.ValueString())
						}
					}
					expressionProvisioner = &shared.ExpressionProvisioner{
						AllowReassignment: allowReassignment2,
						Expressions:       expressions,
						FallbackUserIds:   fallbackUserIds2,
					}
				}
				var groupProvisioner *shared.GroupProvisioner
				if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner != nil {
					allowReassignment3 := new(bool)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment.IsNull() {
						*allowReassignment3 = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment3 = nil
					}
					appGroupID := new(string)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID.IsNull() {
						*appGroupID = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID.ValueString()
					} else {
						appGroupID = nil
					}
					appId1 := new(string)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID.IsNull() {
						*appId1 = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID.ValueString()
					} else {
						appId1 = nil
					}
					var fallbackUserIds3 []string
					if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds != nil {
						fallbackUserIds3 = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds))
						for _, v := range r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds {
							fallbackUserIds3 = append(fallbackUserIds3, v.ValueString())
						}
					}
					groupProvisioner = &shared.GroupProvisioner{
						AllowReassignment: allowReassignment3,
						AppGroupID:        appGroupID,
						AppID:             appId1,
						FallbackUserIds:   fallbackUserIds3,
					}
				}
				var managerProvisioner *shared.ManagerProvisioner
				if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner != nil {
					allowReassignment4 := new(bool)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment.IsNull() {
						*allowReassignment4 = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment4 = nil
					}
					var fallbackUserIds4 []string
					if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds != nil {
						fallbackUserIds4 = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds))
						for _, v := range r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds {
							fallbackUserIds4 = append(fallbackUserIds4, v.ValueString())
						}
					}
					managerProvisioner = &shared.ManagerProvisioner{
						AllowReassignment: allowReassignment4,
						FallbackUserIds:   fallbackUserIds4,
					}
				}
				var userProvisioner *shared.UserProvisioner
				if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner != nil {
					allowReassignment5 := new(bool)
					if !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment.IsUnknown() && !r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment.IsNull() {
						*allowReassignment5 = r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment5 = nil
					}
					var userProvisionerUserIds []string
					if r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds != nil {
						userProvisionerUserIds = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds))
						for _, v := range r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds {
							userProvisionerUserIds = append(userProvisionerUserIds, v.ValueString())
						}
					}
					userProvisioner = &shared.UserProvisioner{
						AllowReassignment: allowReassignment5,
						UserIds:           userProvisionerUserIds,
					}
				}
				provisionerAssignment = &shared.ProvisionerAssignment{
					AppOwnerProvisioner:         appOwnerProvisioner,
					EntitlementOwnerProvisioner: entitlementOwnerProvisioner,
					ExpressionProvisioner:       expressionProvisioner,
					GroupProvisioner:            groupProvisioner,
					ManagerProvisioner:          managerProvisioner,
					UserProvisioner:             userProvisioner,
				}
			}
			var userIds []string
			if r.ProvisionPolicy.ManualProvision.UserIds != nil {
				userIds = make([]string, 0, len(r.ProvisionPolicy.ManualProvision.UserIds))
				for _, userIdsItem := range r.ProvisionPolicy.ManualProvision.UserIds {
					userIds = append(userIds, userIdsItem.ValueString())
				}
			}
			manualProvision = &shared.ManualProvision{
				Instructions:          instructions1,
				ProvisionerAssignment: provisionerAssignment,
				UserIds:               userIds,
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
			ActionProvision:         actionProvision,
			ConnectorProvision:      connectorProvision,
			DelegatedProvision:      delegatedProvision,
			ExternalTicketProvision: externalTicketProvision,
			ManualProvision:         manualProvision,
			MultiStep:               multiStep,
			UnconfiguredProvision:   unconfiguredProvision,
			WebhookProvision:        webhookProvision,
		}
	}
	purpose := new(shared.Purpose)
	if !r.Purpose.IsUnknown() && !r.Purpose.IsNull() {
		*purpose = shared.Purpose(r.Purpose.ValueString())
	} else {
		purpose = nil
	}
	requestSchemaID := new(string)
	if !r.RequestSchemaID.IsUnknown() && !r.RequestSchemaID.IsNull() {
		*requestSchemaID = r.RequestSchemaID.ValueString()
	} else {
		requestSchemaID = nil
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
	sourceConnectorIds := make(map[string]string)
	for sourceConnectorIdsKey, sourceConnectorIdsValue := range r.SourceConnectorIds {
		sourceConnectorIdsInst := sourceConnectorIdsValue.ValueString()

		sourceConnectorIds[sourceConnectorIdsKey] = sourceConnectorIdsInst
	}
	var deprovisionerPolicy *shared.DeprovisionerPolicy
	if r.DeprovisionerPolicy != nil {
		var actionProvision1 *shared.ActionProvision
		if r.DeprovisionerPolicy.ActionProvision != nil {
			actionName1 := new(string)
			if !r.DeprovisionerPolicy.ActionProvision.ActionName.IsUnknown() && !r.DeprovisionerPolicy.ActionProvision.ActionName.IsNull() {
				*actionName1 = r.DeprovisionerPolicy.ActionProvision.ActionName.ValueString()
			} else {
				actionName1 = nil
			}
			appId4 := new(string)
			if !r.DeprovisionerPolicy.ActionProvision.AppID.IsUnknown() && !r.DeprovisionerPolicy.ActionProvision.AppID.IsNull() {
				*appId4 = r.DeprovisionerPolicy.ActionProvision.AppID.ValueString()
			} else {
				appId4 = nil
			}
			connectorId5 := new(string)
			if !r.DeprovisionerPolicy.ActionProvision.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ActionProvision.ConnectorID.IsNull() {
				*connectorId5 = r.DeprovisionerPolicy.ActionProvision.ConnectorID.ValueString()
			} else {
				connectorId5 = nil
			}
			actionProvision1 = &shared.ActionProvision{
				ActionName:  actionName1,
				AppID:       appId4,
				ConnectorID: connectorId5,
			}
		}
		var connectorProvision1 *shared.ConnectorProvision
		if r.DeprovisionerPolicy.ConnectorProvision != nil {
			var accountProvision1 *shared.AccountProvision
			if r.DeprovisionerPolicy.ConnectorProvision.AccountProvision != nil {
				var config1 any
				if !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config.IsNull() {
					_ = json.Unmarshal([]byte(r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config.ValueString()), &config1)
				}
				connectorId6 := new(string)
				if !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID.IsNull() {
					*connectorId6 = r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.ConnectorID.ValueString()
				} else {
					connectorId6 = nil
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
					ConnectorID: connectorId6,
					DoNotSave:   doNotSave1,
					SaveToVault: saveToVault1,
					SchemaID:    schemaId1,
				}
			}
			var defaultBehavior1 *shared.DefaultBehavior
			if r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior != nil {
				connectorId7 := new(string)
				if !r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.IsNull() {
					*connectorId7 = r.DeprovisionerPolicy.ConnectorProvision.DefaultBehavior.ConnectorID.ValueString()
				} else {
					connectorId7 = nil
				}
				defaultBehavior1 = &shared.DefaultBehavior{
					ConnectorID: connectorId7,
				}
			}
			var deleteAccount1 *shared.DeleteAccount
			if r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount != nil {
				connectorId8 := new(string)
				if !r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID.IsNull() {
					*connectorId8 = r.DeprovisionerPolicy.ConnectorProvision.DeleteAccount.ConnectorID.ValueString()
				} else {
					connectorId8 = nil
				}
				deleteAccount1 = &shared.DeleteAccount{
					ConnectorID: connectorId8,
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
			appId5 := new(string)
			if !r.DeprovisionerPolicy.DelegatedProvision.AppID.IsUnknown() && !r.DeprovisionerPolicy.DelegatedProvision.AppID.IsNull() {
				*appId5 = r.DeprovisionerPolicy.DelegatedProvision.AppID.ValueString()
			} else {
				appId5 = nil
			}
			entitlementId1 := new(string)
			if !r.DeprovisionerPolicy.DelegatedProvision.EntitlementID.IsUnknown() && !r.DeprovisionerPolicy.DelegatedProvision.EntitlementID.IsNull() {
				*entitlementId1 = r.DeprovisionerPolicy.DelegatedProvision.EntitlementID.ValueString()
			} else {
				entitlementId1 = nil
			}
			delegatedProvision1 = &shared.DelegatedProvision{
				AppID:         appId5,
				EntitlementID: entitlementId1,
			}
		}
		var externalTicketProvision1 *shared.ExternalTicketProvision
		if r.DeprovisionerPolicy.ExternalTicketProvision != nil {
			appId6 := new(string)
			if !r.DeprovisionerPolicy.ExternalTicketProvision.AppID.IsUnknown() && !r.DeprovisionerPolicy.ExternalTicketProvision.AppID.IsNull() {
				*appId6 = r.DeprovisionerPolicy.ExternalTicketProvision.AppID.ValueString()
			} else {
				appId6 = nil
			}
			connectorId9 := new(string)
			if !r.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID.IsUnknown() && !r.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID.IsNull() {
				*connectorId9 = r.DeprovisionerPolicy.ExternalTicketProvision.ConnectorID.ValueString()
			} else {
				connectorId9 = nil
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
				AppID:                             appId6,
				ConnectorID:                       connectorId9,
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
			var provisionerAssignment1 *shared.ProvisionerAssignment
			if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment != nil {
				var appOwnerProvisioner1 *shared.AppOwnerProvisioner
				if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner != nil {
					allowReassignment6 := new(bool)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment.IsNull() {
						*allowReassignment6 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment6 = nil
					}
					var fallbackUserIds5 []string
					if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds != nil {
						fallbackUserIds5 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds))
						for _, v := range r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds {
							fallbackUserIds5 = append(fallbackUserIds5, v.ValueString())
						}
					}
					appOwnerProvisioner1 = &shared.AppOwnerProvisioner{
						AllowReassignment: allowReassignment6,
						FallbackUserIds:   fallbackUserIds5,
					}
				}
				var entitlementOwnerProvisioner1 *shared.EntitlementOwnerProvisioner
				if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner != nil {
					allowReassignment7 := new(bool)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment.IsNull() {
						*allowReassignment7 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment7 = nil
					}
					var fallbackUserIds6 []string
					if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds != nil {
						fallbackUserIds6 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds))
						for _, v := range r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds {
							fallbackUserIds6 = append(fallbackUserIds6, v.ValueString())
						}
					}
					entitlementOwnerProvisioner1 = &shared.EntitlementOwnerProvisioner{
						AllowReassignment: allowReassignment7,
						FallbackUserIds:   fallbackUserIds6,
					}
				}
				var expressionProvisioner1 *shared.ExpressionProvisioner
				if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner != nil {
					allowReassignment8 := new(bool)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment.IsNull() {
						*allowReassignment8 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment8 = nil
					}
					var expressions1 []string
					if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions != nil {
						expressions1 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions))
						for _, v := range r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions {
							expressions1 = append(expressions1, v.ValueString())
						}
					}
					var fallbackUserIds7 []string
					if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds != nil {
						fallbackUserIds7 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds))
						for _, v := range r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds {
							fallbackUserIds7 = append(fallbackUserIds7, v.ValueString())
						}
					}
					expressionProvisioner1 = &shared.ExpressionProvisioner{
						AllowReassignment: allowReassignment8,
						Expressions:       expressions1,
						FallbackUserIds:   fallbackUserIds7,
					}
				}
				var groupProvisioner1 *shared.GroupProvisioner
				if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner != nil {
					allowReassignment9 := new(bool)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment.IsNull() {
						*allowReassignment9 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment9 = nil
					}
					appGroupId1 := new(string)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID.IsNull() {
						*appGroupId1 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID.ValueString()
					} else {
						appGroupId1 = nil
					}
					appId5 := new(string)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID.IsNull() {
						*appId5 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID.ValueString()
					} else {
						appId5 = nil
					}
					var fallbackUserIds8 []string
					if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds != nil {
						fallbackUserIds8 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds))
						for _, v := range r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds {
							fallbackUserIds8 = append(fallbackUserIds8, v.ValueString())
						}
					}
					groupProvisioner1 = &shared.GroupProvisioner{
						AllowReassignment: allowReassignment9,
						AppGroupID:        appGroupId1,
						AppID:             appId5,
						FallbackUserIds:   fallbackUserIds8,
					}
				}
				var managerProvisioner1 *shared.ManagerProvisioner
				if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner != nil {
					allowReassignment10 := new(bool)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment.IsNull() {
						*allowReassignment10 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment10 = nil
					}
					var fallbackUserIds9 []string
					if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds != nil {
						fallbackUserIds9 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds))
						for _, v := range r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds {
							fallbackUserIds9 = append(fallbackUserIds9, v.ValueString())
						}
					}
					managerProvisioner1 = &shared.ManagerProvisioner{
						AllowReassignment: allowReassignment10,
						FallbackUserIds:   fallbackUserIds9,
					}
				}
				var userProvisioner1 *shared.UserProvisioner
				if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner != nil {
					allowReassignment11 := new(bool)
					if !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment.IsUnknown() && !r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment.IsNull() {
						*allowReassignment11 = r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment.ValueBool()
					} else {
						allowReassignment11 = nil
					}
					var userProvisionerUserIds1 []string
					if r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds != nil {
						userProvisionerUserIds1 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds))
						for _, v := range r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds {
							userProvisionerUserIds1 = append(userProvisionerUserIds1, v.ValueString())
						}
					}
					userProvisioner1 = &shared.UserProvisioner{
						AllowReassignment: allowReassignment11,
						UserIds:           userProvisionerUserIds1,
					}
				}
				provisionerAssignment1 = &shared.ProvisionerAssignment{
					AppOwnerProvisioner:         appOwnerProvisioner1,
					EntitlementOwnerProvisioner: entitlementOwnerProvisioner1,
					ExpressionProvisioner:       expressionProvisioner1,
					GroupProvisioner:            groupProvisioner1,
					ManagerProvisioner:          managerProvisioner1,
					UserProvisioner:             userProvisioner1,
				}
			}
			var userIds1 []string
			if r.DeprovisionerPolicy.ManualProvision.UserIds != nil {
				userIds1 = make([]string, 0, len(r.DeprovisionerPolicy.ManualProvision.UserIds))
				for _, userIdsItem1 := range r.DeprovisionerPolicy.ManualProvision.UserIds {
					userIds1 = append(userIds1, userIdsItem1.ValueString())
				}
			}
			manualProvision1 = &shared.ManualProvision{
				Instructions:          instructions3,
				ProvisionerAssignment: provisionerAssignment1,
				UserIds:               userIds1,
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
			ActionProvision:         actionProvision1,
			ConnectorProvision:      connectorProvision1,
			DelegatedProvision:      delegatedProvision1,
			ExternalTicketProvision: externalTicketProvision1,
			ManualProvision:         manualProvision1,
			MultiStep:               multiStep1,
			UnconfiguredProvision:   unconfiguredProvision1,
			WebhookProvision:        webhookProvision1,
		}
	}
	out := shared.AppEntitlementInput{
		Alias:                          alias,
		AppID:                          appID,
		AppResourceID:                  appResourceID,
		AppResourceTypeID:              appResourceTypeID,
		CertifyPolicyID:                certifyPolicyID,
		ComplianceFrameworkValueIds:    complianceFrameworkValueIds,
		DefaultValuesApplied:           defaultValuesApplied,
		Description:                    description,
		DisplayName:                    displayName,
		DurationGrant:                  durationGrant,
		DurationUnset:                  durationUnset,
		EmergencyGrantEnabled:          emergencyGrantEnabled,
		EmergencyGrantPolicyID:         emergencyGrantPolicyID,
		GrantPolicyID:                  grantPolicyID,
		IsManuallyManaged:              isManuallyManaged,
		MatchBatonID:                   matchBatonID,
		OverrideAccessRequestsDefaults: overrideAccessRequestsDefaults,
		ProvisionPolicy:                provisionPolicy,
		Purpose:                        purpose,
		RequestSchemaID:                requestSchemaID,
		RevokePolicyID:                 revokePolicyID,
		RiskLevelValueID:               riskLevelValueID,
		Slug:                           slug,
		SourceConnectorIds:             sourceConnectorIds,
		DeprovisionerPolicy:            deprovisionerPolicy,
	}

	return &out
}

func (r *AppEntitlementResourceModel) RefreshFromGetResponse(resp *shared.AppEntitlement) {
	if resp != nil {
		r.Alias = types.StringPointerValue(resp.Alias)
		r.AppID = types.StringPointerValue(resp.AppID)
		r.AppResourceID = types.StringPointerValue(resp.AppResourceID)
		r.AppResourceTypeID = types.StringPointerValue(resp.AppResourceTypeID)
		r.CertifyPolicyID = types.StringPointerValue(resp.CertifyPolicyID)
		if resp.ComplianceFrameworkValueIds != nil {
			r.ComplianceFrameworkValueIds = make([]types.String, 0, len(resp.ComplianceFrameworkValueIds))
			for _, v := range resp.ComplianceFrameworkValueIds {
				r.ComplianceFrameworkValueIds = append(r.ComplianceFrameworkValueIds, types.StringValue(v))
			}
		}
		r.CreatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreatedAt))
		r.DefaultValuesApplied = types.BoolPointerValue(resp.DefaultValuesApplied)
		r.DeletedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.DeletedAt))
		if resp.DeprovisionerPolicy == nil {
			r.DeprovisionerPolicy = nil
		} else {
			r.DeprovisionerPolicy = &tfTypes.DeprovisionerPolicy{}
			if resp.DeprovisionerPolicy.ActionProvision == nil {
				r.DeprovisionerPolicy.ActionProvision = nil
			} else {
				r.DeprovisionerPolicy.ActionProvision = &tfTypes.ActionProvision{}
				r.DeprovisionerPolicy.ActionProvision.ActionName = types.StringPointerValue(resp.DeprovisionerPolicy.ActionProvision.ActionName)
				r.DeprovisionerPolicy.ActionProvision.AppID = types.StringPointerValue(resp.DeprovisionerPolicy.ActionProvision.AppID)
				r.DeprovisionerPolicy.ActionProvision.ConnectorID = types.StringPointerValue(resp.DeprovisionerPolicy.ActionProvision.ConnectorID)
			}
			if resp.DeprovisionerPolicy.ConnectorProvision == nil {
				r.DeprovisionerPolicy.ConnectorProvision = nil
			} else {
				r.DeprovisionerPolicy.ConnectorProvision = &tfTypes.ConnectorProvision{}
				if resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision == nil {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision = nil
				} else {
					r.DeprovisionerPolicy.ConnectorProvision.AccountProvision = &tfTypes.AccountProvision{}
					if resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config == nil {
						r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config = jsontypes.NewNormalizedNull()
					} else {
						configResult, _ := json.Marshal(resp.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config)
						r.DeprovisionerPolicy.ConnectorProvision.AccountProvision.Config = jsontypes.NewNormalizedValue(string(configResult))
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
				if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment == nil {
					r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment = nil
				} else {
					r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment = &tfTypes.ProvisionerAssignment{}
					if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner == nil {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner = nil
					} else {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner = &tfTypes.AppOwnerProvisioner{}
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment = types.BoolPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment)
						if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds != nil {
							r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds))
							for _, v := range resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds {
								r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds = append(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner == nil {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner = nil
					} else {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner = &tfTypes.EntitlementOwnerProvisioner{}
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment = types.BoolPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment)
						if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds != nil {
							r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds))
							for _, v := range resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds {
								r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds = append(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner == nil {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner = nil
					} else {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner = &tfTypes.ExpressionProvisioner{}
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment = types.BoolPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment)
						if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions != nil {
							r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions))
							for _, v := range resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions {
								r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions = append(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions, types.StringValue(v))
							}
						}
						if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds != nil {
							r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds))
							for _, v := range resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds {
								r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds = append(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner == nil {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner = nil
					} else {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner = &tfTypes.GroupProvisioner{}
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment = types.BoolPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment)
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID = types.StringPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID)
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID = types.StringPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID)
						if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds != nil {
							r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds))
							for _, v := range resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds {
								r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds = append(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner == nil {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner = nil
					} else {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner = &tfTypes.ManagerProvisioner{}
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment = types.BoolPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment)
						if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds != nil {
							r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds))
							for _, v := range resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds {
								r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds = append(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner == nil {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner = nil
					} else {
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner = &tfTypes.UserProvisioner{}
						r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment = types.BoolPointerValue(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment)
						if resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds != nil {
							r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds = make([]types.String, 0, len(resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds))
							for _, v := range resp.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds {
								r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds = append(r.DeprovisionerPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds, types.StringValue(v))
							}
						}
					}
				}
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
		r.Description = types.StringPointerValue(resp.Description)
		r.DisplayName = types.StringPointerValue(resp.DisplayName)
		r.DurationGrant = types.StringPointerValue(resp.DurationGrant)
		if resp.DurationUnset == nil {
			r.DurationUnset = nil
		} else {
			r.DurationUnset = &tfTypes.CreateAppEntitlementRequestDurationUnset{}
		}
		r.EmergencyGrantEnabled = types.BoolPointerValue(resp.EmergencyGrantEnabled)
		r.EmergencyGrantPolicyID = types.StringPointerValue(resp.EmergencyGrantPolicyID)
		r.GrantCount = types.StringPointerValue(resp.GrantCount)
		r.GrantPolicyID = types.StringPointerValue(resp.GrantPolicyID)
		r.ID = types.StringPointerValue(resp.ID)
		r.IsAutomationEnabled = types.BoolPointerValue(resp.IsAutomationEnabled)
		r.IsManuallyManaged = types.BoolPointerValue(resp.IsManuallyManaged)
		r.MatchBatonID = types.StringPointerValue(resp.MatchBatonID)
		r.OverrideAccessRequestsDefaults = types.BoolPointerValue(resp.OverrideAccessRequestsDefaults)
		if resp.ProvisionPolicy == nil {
			r.ProvisionPolicy = nil
		} else {
			r.ProvisionPolicy = &tfTypes.ProvisionPolicy{}
			if resp.ProvisionPolicy.ActionProvision == nil {
				r.ProvisionPolicy.ActionProvision = nil
			} else {
				r.ProvisionPolicy.ActionProvision = &tfTypes.ActionProvision{}
				r.ProvisionPolicy.ActionProvision.ActionName = types.StringPointerValue(resp.ProvisionPolicy.ActionProvision.ActionName)
				r.ProvisionPolicy.ActionProvision.AppID = types.StringPointerValue(resp.ProvisionPolicy.ActionProvision.AppID)
				r.ProvisionPolicy.ActionProvision.ConnectorID = types.StringPointerValue(resp.ProvisionPolicy.ActionProvision.ConnectorID)
			}
			if resp.ProvisionPolicy.ConnectorProvision == nil {
				r.ProvisionPolicy.ConnectorProvision = nil
			} else {
				r.ProvisionPolicy.ConnectorProvision = &tfTypes.ConnectorProvision{}
				if resp.ProvisionPolicy.ConnectorProvision.AccountProvision == nil {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision = nil
				} else {
					r.ProvisionPolicy.ConnectorProvision.AccountProvision = &tfTypes.AccountProvision{}
					if resp.ProvisionPolicy.ConnectorProvision.AccountProvision.Config == nil {
						r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config = jsontypes.NewNormalizedNull()
					} else {
						configResult1, _ := json.Marshal(resp.ProvisionPolicy.ConnectorProvision.AccountProvision.Config)
						r.ProvisionPolicy.ConnectorProvision.AccountProvision.Config = jsontypes.NewNormalizedValue(string(configResult1))
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
				if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment == nil {
					r.ProvisionPolicy.ManualProvision.ProvisionerAssignment = nil
				} else {
					r.ProvisionPolicy.ManualProvision.ProvisionerAssignment = &tfTypes.ProvisionerAssignment{}
					if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner == nil {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner = nil
					} else {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner = &tfTypes.AppOwnerProvisioner{}
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment = types.BoolPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.AllowReassignment)
						if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds != nil {
							r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds))
							for _, v := range resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds {
								r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds = append(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.AppOwnerProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner == nil {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner = nil
					} else {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner = &tfTypes.EntitlementOwnerProvisioner{}
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment = types.BoolPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.AllowReassignment)
						if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds != nil {
							r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds))
							for _, v := range resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds {
								r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds = append(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.EntitlementOwnerProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner == nil {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner = nil
					} else {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner = &tfTypes.ExpressionProvisioner{}
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment = types.BoolPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.AllowReassignment)
						if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions != nil {
							r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions))
							for _, v := range resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions {
								r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions = append(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.Expressions, types.StringValue(v))
							}
						}
						if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds != nil {
							r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds))
							for _, v := range resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds {
								r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds = append(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ExpressionProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner == nil {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner = nil
					} else {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner = &tfTypes.GroupProvisioner{}
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment = types.BoolPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AllowReassignment)
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID = types.StringPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppGroupID)
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID = types.StringPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.AppID)
						if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds != nil {
							r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds))
							for _, v := range resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds {
								r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds = append(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.GroupProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner == nil {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner = nil
					} else {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner = &tfTypes.ManagerProvisioner{}
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment = types.BoolPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.AllowReassignment)
						if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds != nil {
							r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds))
							for _, v := range resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds {
								r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds = append(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.ManagerProvisioner.FallbackUserIds, types.StringValue(v))
							}
						}
					}
					if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner == nil {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner = nil
					} else {
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner = &tfTypes.UserProvisioner{}
						r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment = types.BoolPointerValue(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.AllowReassignment)
						if resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds != nil {
							r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds = make([]types.String, 0, len(resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds))
							for _, v := range resp.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds {
								r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds = append(r.ProvisionPolicy.ManualProvision.ProvisionerAssignment.UserProvisioner.UserIds, types.StringValue(v))
							}
						}
					}
				}
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
		if resp.Purpose != nil {
			r.Purpose = types.StringValue(string(*resp.Purpose))
		} else {
			r.Purpose = types.StringNull()
		}
		r.RequestSchemaID = types.StringPointerValue(resp.RequestSchemaID)
		r.RevokePolicyID = types.StringPointerValue(resp.RevokePolicyID)
		r.RiskLevelValueID = types.StringPointerValue(resp.RiskLevelValueID)
		r.Slug = types.StringPointerValue(resp.Slug)
		if len(resp.SourceConnectorIds) > 0 {
			r.SourceConnectorIds = make(map[string]types.String, len(resp.SourceConnectorIds))
			for key, value := range resp.SourceConnectorIds {
				r.SourceConnectorIds[key] = types.StringValue(value)
			}
		}
		r.SystemBuiltin = types.BoolPointerValue(resp.SystemBuiltin)
		r.UpdatedAt = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.UpdatedAt))
	}
}

func (r *AppEntitlementResourceModel) RefreshFromCreateResponse(resp *shared.AppEntitlement) {
	r.RefreshFromGetResponse(resp)
}

func (r *AppEntitlementResourceModel) RefreshFromUpdateResponse(resp *shared.AppEntitlement) {
	r.RefreshFromGetResponse(resp)
}
