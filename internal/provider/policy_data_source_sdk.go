// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-terraform/internal/provider/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"time"
)

func (r *PolicyDataSourceModel) RefreshFromSharedPolicy(resp *shared.Policy) {
	if resp != nil {
		if resp.CreatedAt != nil {
			r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		} else {
			r.CreatedAt = types.StringNull()
		}
		if resp.DeletedAt != nil {
			r.DeletedAt = types.StringValue(resp.DeletedAt.Format(time.RFC3339Nano))
		} else {
			r.DeletedAt = types.StringNull()
		}
		r.Description = types.StringPointerValue(resp.Description)
		r.DisplayName = types.StringPointerValue(resp.DisplayName)
		r.ID = types.StringPointerValue(resp.ID)
		if len(resp.PolicySteps) > 0 {
			r.PolicySteps = make(map[string]tfTypes.PolicySteps)
			for policyStepsKey, policyStepsValue := range resp.PolicySteps {
				var policyStepsResult tfTypes.PolicySteps
				for stepsCount, stepsItem := range policyStepsValue.Steps {
					var steps1 tfTypes.PolicyStep
					if stepsItem.Accept == nil {
						steps1.Accept = nil
					} else {
						steps1.Accept = &tfTypes.Accept{}
						steps1.Accept.AcceptMessage = types.StringPointerValue(stepsItem.Accept.AcceptMessage)
					}
					if stepsItem.Approval == nil {
						steps1.Approval = nil
					} else {
						steps1.Approval = &tfTypes.Approval{}
						steps1.Approval.AllowReassignment = types.BoolPointerValue(stepsItem.Approval.AllowReassignment)
						if stepsItem.Approval.AppGroupApproval == nil {
							steps1.Approval.AppGroupApproval = nil
						} else {
							steps1.Approval.AppGroupApproval = &tfTypes.AppGroupApproval{}
							steps1.Approval.AppGroupApproval.AllowSelfApproval = types.BoolPointerValue(stepsItem.Approval.AppGroupApproval.AllowSelfApproval)
							steps1.Approval.AppGroupApproval.AppGroupID = types.StringPointerValue(stepsItem.Approval.AppGroupApproval.AppGroupID)
							steps1.Approval.AppGroupApproval.AppID = types.StringPointerValue(stepsItem.Approval.AppGroupApproval.AppID)
							steps1.Approval.AppGroupApproval.Fallback = types.BoolPointerValue(stepsItem.Approval.AppGroupApproval.Fallback)
							steps1.Approval.AppGroupApproval.FallbackUserIds = []types.String{}
							for _, v := range stepsItem.Approval.AppGroupApproval.FallbackUserIds {
								steps1.Approval.AppGroupApproval.FallbackUserIds = append(steps1.Approval.AppGroupApproval.FallbackUserIds, types.StringValue(v))
							}
						}
						if stepsItem.Approval.AppOwnerApproval == nil {
							steps1.Approval.AppOwnerApproval = nil
						} else {
							steps1.Approval.AppOwnerApproval = &tfTypes.AppOwnerApproval{}
							steps1.Approval.AppOwnerApproval.AllowSelfApproval = types.BoolPointerValue(stepsItem.Approval.AppOwnerApproval.AllowSelfApproval)
						}
						steps1.Approval.Assigned = types.BoolPointerValue(stepsItem.Approval.Assigned)
						if stepsItem.Approval.EntitlementOwnerApproval == nil {
							steps1.Approval.EntitlementOwnerApproval = nil
						} else {
							steps1.Approval.EntitlementOwnerApproval = &tfTypes.EntitlementOwnerApproval{}
							steps1.Approval.EntitlementOwnerApproval.AllowSelfApproval = types.BoolPointerValue(stepsItem.Approval.EntitlementOwnerApproval.AllowSelfApproval)
							steps1.Approval.EntitlementOwnerApproval.Fallback = types.BoolPointerValue(stepsItem.Approval.EntitlementOwnerApproval.Fallback)
							steps1.Approval.EntitlementOwnerApproval.FallbackUserIds = []types.String{}
							for _, v := range stepsItem.Approval.EntitlementOwnerApproval.FallbackUserIds {
								steps1.Approval.EntitlementOwnerApproval.FallbackUserIds = append(steps1.Approval.EntitlementOwnerApproval.FallbackUserIds, types.StringValue(v))
							}
						}
						if stepsItem.Approval.ExpressionApproval == nil {
							steps1.Approval.ExpressionApproval = nil
						} else {
							steps1.Approval.ExpressionApproval = &tfTypes.ExpressionApproval{}
							steps1.Approval.ExpressionApproval.AllowSelfApproval = types.BoolPointerValue(stepsItem.Approval.ExpressionApproval.AllowSelfApproval)
							steps1.Approval.ExpressionApproval.AssignedUserIds = []types.String{}
							for _, v := range stepsItem.Approval.ExpressionApproval.AssignedUserIds {
								steps1.Approval.ExpressionApproval.AssignedUserIds = append(steps1.Approval.ExpressionApproval.AssignedUserIds, types.StringValue(v))
							}
							steps1.Approval.ExpressionApproval.Expressions = []types.String{}
							for _, v := range stepsItem.Approval.ExpressionApproval.Expressions {
								steps1.Approval.ExpressionApproval.Expressions = append(steps1.Approval.ExpressionApproval.Expressions, types.StringValue(v))
							}
							steps1.Approval.ExpressionApproval.Fallback = types.BoolPointerValue(stepsItem.Approval.ExpressionApproval.Fallback)
							steps1.Approval.ExpressionApproval.FallbackUserIds = []types.String{}
							for _, v := range stepsItem.Approval.ExpressionApproval.FallbackUserIds {
								steps1.Approval.ExpressionApproval.FallbackUserIds = append(steps1.Approval.ExpressionApproval.FallbackUserIds, types.StringValue(v))
							}
						}
						if stepsItem.Approval.ManagerApproval == nil {
							steps1.Approval.ManagerApproval = nil
						} else {
							steps1.Approval.ManagerApproval = &tfTypes.ManagerApproval{}
							steps1.Approval.ManagerApproval.AllowSelfApproval = types.BoolPointerValue(stepsItem.Approval.ManagerApproval.AllowSelfApproval)
							steps1.Approval.ManagerApproval.AssignedUserIds = []types.String{}
							for _, v := range stepsItem.Approval.ManagerApproval.AssignedUserIds {
								steps1.Approval.ManagerApproval.AssignedUserIds = append(steps1.Approval.ManagerApproval.AssignedUserIds, types.StringValue(v))
							}
							steps1.Approval.ManagerApproval.Fallback = types.BoolPointerValue(stepsItem.Approval.ManagerApproval.Fallback)
							steps1.Approval.ManagerApproval.FallbackUserIds = []types.String{}
							for _, v := range stepsItem.Approval.ManagerApproval.FallbackUserIds {
								steps1.Approval.ManagerApproval.FallbackUserIds = append(steps1.Approval.ManagerApproval.FallbackUserIds, types.StringValue(v))
							}
						}
						steps1.Approval.RequireApprovalReason = types.BoolPointerValue(stepsItem.Approval.RequireApprovalReason)
						steps1.Approval.RequireDenialReason = types.BoolPointerValue(stepsItem.Approval.RequireDenialReason)
						steps1.Approval.RequireReassignmentReason = types.BoolPointerValue(stepsItem.Approval.RequireReassignmentReason)
						if stepsItem.Approval.SelfApproval == nil {
							steps1.Approval.SelfApproval = nil
						} else {
							steps1.Approval.SelfApproval = &tfTypes.SelfApproval{}
							steps1.Approval.SelfApproval.AssignedUserIds = []types.String{}
							for _, v := range stepsItem.Approval.SelfApproval.AssignedUserIds {
								steps1.Approval.SelfApproval.AssignedUserIds = append(steps1.Approval.SelfApproval.AssignedUserIds, types.StringValue(v))
							}
							steps1.Approval.SelfApproval.Fallback = types.BoolPointerValue(stepsItem.Approval.SelfApproval.Fallback)
							steps1.Approval.SelfApproval.FallbackUserIds = []types.String{}
							for _, v := range stepsItem.Approval.SelfApproval.FallbackUserIds {
								steps1.Approval.SelfApproval.FallbackUserIds = append(steps1.Approval.SelfApproval.FallbackUserIds, types.StringValue(v))
							}
						}
						if stepsItem.Approval.UserApproval == nil {
							steps1.Approval.UserApproval = nil
						} else {
							steps1.Approval.UserApproval = &tfTypes.UserApproval{}
							steps1.Approval.UserApproval.AllowSelfApproval = types.BoolPointerValue(stepsItem.Approval.UserApproval.AllowSelfApproval)
							steps1.Approval.UserApproval.UserIds = []types.String{}
							for _, v := range stepsItem.Approval.UserApproval.UserIds {
								steps1.Approval.UserApproval.UserIds = append(steps1.Approval.UserApproval.UserIds, types.StringValue(v))
							}
						}
					}
					if stepsItem.Provision == nil {
						steps1.Provision = nil
					} else {
						steps1.Provision = &tfTypes.Provision{}
						steps1.Provision.Assigned = types.BoolPointerValue(stepsItem.Provision.Assigned)
						if stepsItem.Provision.ProvisionPolicy == nil {
							steps1.Provision.ProvisionPolicy = nil
						} else {
							steps1.Provision.ProvisionPolicy = &tfTypes.ProvisionPolicy{}
							if stepsItem.Provision.ProvisionPolicy.ConnectorProvision == nil {
								steps1.Provision.ProvisionPolicy.ConnectorProvision = nil
							} else {
								steps1.Provision.ProvisionPolicy.ConnectorProvision = &tfTypes.Three{}
							}
							if stepsItem.Provision.ProvisionPolicy.DelegatedProvision == nil {
								steps1.Provision.ProvisionPolicy.DelegatedProvision = nil
							} else {
								steps1.Provision.ProvisionPolicy.DelegatedProvision = &tfTypes.DelegatedProvision{}
								steps1.Provision.ProvisionPolicy.DelegatedProvision.AppID = types.StringPointerValue(stepsItem.Provision.ProvisionPolicy.DelegatedProvision.AppID)
								steps1.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID = types.StringPointerValue(stepsItem.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID)
								steps1.Provision.ProvisionPolicy.DelegatedProvision.Implicit = types.BoolPointerValue(stepsItem.Provision.ProvisionPolicy.DelegatedProvision.Implicit)
							}
							if stepsItem.Provision.ProvisionPolicy.ManualProvision == nil {
								steps1.Provision.ProvisionPolicy.ManualProvision = nil
							} else {
								steps1.Provision.ProvisionPolicy.ManualProvision = &tfTypes.ManualProvision{}
								steps1.Provision.ProvisionPolicy.ManualProvision.Instructions = types.StringPointerValue(stepsItem.Provision.ProvisionPolicy.ManualProvision.Instructions)
								steps1.Provision.ProvisionPolicy.ManualProvision.UserIds = []types.String{}
								for _, v := range stepsItem.Provision.ProvisionPolicy.ManualProvision.UserIds {
									steps1.Provision.ProvisionPolicy.ManualProvision.UserIds = append(steps1.Provision.ProvisionPolicy.ManualProvision.UserIds, types.StringValue(v))
								}
							}
							if stepsItem.Provision.ProvisionPolicy.WebhookProvision == nil {
								steps1.Provision.ProvisionPolicy.WebhookProvision = nil
							} else {
								steps1.Provision.ProvisionPolicy.WebhookProvision = &tfTypes.WebhookProvision{}
								steps1.Provision.ProvisionPolicy.WebhookProvision.WebhookID = types.StringPointerValue(stepsItem.Provision.ProvisionPolicy.WebhookProvision.WebhookID)
							}
						}
						if stepsItem.Provision.ProvisionTarget == nil {
							steps1.Provision.ProvisionTarget = nil
						} else {
							steps1.Provision.ProvisionTarget = &tfTypes.ProvisionTarget{}
							steps1.Provision.ProvisionTarget.AppEntitlementID = types.StringPointerValue(stepsItem.Provision.ProvisionTarget.AppEntitlementID)
							steps1.Provision.ProvisionTarget.AppID = types.StringPointerValue(stepsItem.Provision.ProvisionTarget.AppID)
							steps1.Provision.ProvisionTarget.AppUserID = types.StringPointerValue(stepsItem.Provision.ProvisionTarget.AppUserID)
							steps1.Provision.ProvisionTarget.GrantDuration = types.StringPointerValue(stepsItem.Provision.ProvisionTarget.GrantDuration)
						}
					}
					if stepsItem.Reject == nil {
						steps1.Reject = nil
					} else {
						steps1.Reject = &tfTypes.Reject{}
						steps1.Reject.RejectMessage = types.StringPointerValue(stepsItem.Reject.RejectMessage)
					}
					if stepsCount+1 > len(policyStepsResult.Steps) {
						policyStepsResult.Steps = append(policyStepsResult.Steps, steps1)
					} else {
						policyStepsResult.Steps[stepsCount].Accept = steps1.Accept
						policyStepsResult.Steps[stepsCount].Approval = steps1.Approval
						policyStepsResult.Steps[stepsCount].Provision = steps1.Provision
						policyStepsResult.Steps[stepsCount].Reject = steps1.Reject
					}
				}
				r.PolicySteps[policyStepsKey] = policyStepsResult
			}
		}
		if resp.PolicyType != nil {
			r.PolicyType = types.StringValue(string(*resp.PolicyType))
		} else {
			r.PolicyType = types.StringNull()
		}
		if len(r.PostActions) > len(resp.PostActions) {
			r.PostActions = r.PostActions[:len(resp.PostActions)]
		}
		for postActionsCount, postActionsItem := range resp.PostActions {
			var postActions1 tfTypes.PolicyPostActions
			postActions1.CertifyRemediateImmediately = types.BoolPointerValue(postActionsItem.CertifyRemediateImmediately)
			if postActionsCount+1 > len(r.PostActions) {
				r.PostActions = append(r.PostActions, postActions1)
			} else {
				r.PostActions[postActionsCount].CertifyRemediateImmediately = postActions1.CertifyRemediateImmediately
			}
		}
		r.ReassignTasksToDelegates = types.BoolPointerValue(resp.ReassignTasksToDelegates)
		if len(r.Rules) > len(resp.Rules) {
			r.Rules = r.Rules[:len(resp.Rules)]
		}
		for rulesCount, rulesItem := range resp.Rules {
			var rules1 tfTypes.Rule
			rules1.Condition = types.StringPointerValue(rulesItem.Condition)
			rules1.PolicyKey = types.StringPointerValue(rulesItem.PolicyKey)
			if rulesCount+1 > len(r.Rules) {
				r.Rules = append(r.Rules, rules1)
			} else {
				r.Rules[rulesCount].Condition = rules1.Condition
				r.Rules[rulesCount].PolicyKey = rules1.PolicyKey
			}
		}
		r.SystemBuiltin = types.BoolPointerValue(resp.SystemBuiltin)
		if resp.UpdatedAt != nil {
			r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		} else {
			r.UpdatedAt = types.StringNull()
		}
	}
}
