// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func BoolPointer(b bool) *bool {
	return &b
}

func (r *PolicyResourceModel) ToCreateSDKType() *shared.CreatePolicyRequest {
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
	policySteps := make(map[string]shared.PolicySteps)
	for policyStepsKey, policyStepsValue := range r.PolicySteps {
		var steps []shared.PolicyStepInput = nil
		for _, stepsItem := range policyStepsValue.Steps {
			var accept *shared.Accept
			if stepsItem.Accept != nil {
				accept = &shared.Accept{}
			}
			if step.Approval != nil {
				newPolicyStep := shared.PolicyStep{
					Approval: &shared.Approval{
						AllowReassignment:         step.Approval.AllowReassignment.ValueBoolPointer(),
						Assigned:                  step.Approval.Assigned.ValueBoolPointer(),
						RequireApprovalReason:     step.Approval.RequireApprovalReason.ValueBoolPointer(),
						RequireReassignmentReason: step.Approval.RequireReassignmentReason.ValueBoolPointer(),
					},
				}
				if step.Approval.AppOwnerApproval != nil {
					newPolicyStep.Approval.AppOwnerApproval = &shared.AppOwnerApproval{
						AllowSelfApproval: step.Approval.AppOwnerApproval.AllowSelfApproval.ValueBoolPointer(),
					}
				}
				if step.Approval.EntitlementOwnerApproval != nil {
					newPolicyStep.Approval.EntitlementOwnerApproval = &shared.EntitlementOwnerApproval{
						AllowSelfApproval: step.Approval.EntitlementOwnerApproval.AllowSelfApproval.ValueBoolPointer(),
						Fallback:          step.Approval.EntitlementOwnerApproval.Fallback.ValueBoolPointer(),
					}
					for _, v := range step.Approval.EntitlementOwnerApproval.FallbackUserIds {
						newPolicyStep.Approval.EntitlementOwnerApproval.FallbackUserIds = append(newPolicyStep.Approval.EntitlementOwnerApproval.FallbackUserIds, v.ValueString())
					}
				}
				if step.Approval.ManagerApproval != nil {
					newPolicyStep.Approval.ManagerApproval = &shared.ManagerApproval{
						AllowSelfApproval: step.Approval.ManagerApproval.AllowSelfApproval.ValueBoolPointer(),
						Fallback:          step.Approval.ManagerApproval.Fallback.ValueBoolPointer(),
					}
					for _, v := range step.Approval.ManagerApproval.FallbackUserIds {
						newPolicyStep.Approval.ManagerApproval.FallbackUserIds = append(newPolicyStep.Approval.ManagerApproval.FallbackUserIds, v.ValueString())
					}
					for _, v := range step.Approval.ManagerApproval.AssignedUserIds {
						newPolicyStep.Approval.ManagerApproval.AssignedUserIds = append(newPolicyStep.Approval.ManagerApproval.AssignedUserIds, v.ValueString())
					}
				}
				if step.Approval.SelfApproval != nil {
					newPolicyStep.Approval.SelfApproval = &shared.SelfApproval{
						Fallback: step.Approval.SelfApproval.Fallback.ValueBoolPointer(),
					}
					for _, v := range step.Approval.SelfApproval.FallbackUserIds {
						newPolicyStep.Approval.SelfApproval.FallbackUserIds = append(newPolicyStep.Approval.SelfApproval.FallbackUserIds, v.ValueString())
					}
					for _, v := range step.Approval.SelfApproval.AssignedUserIds {
						newPolicyStep.Approval.SelfApproval.AssignedUserIds = append(newPolicyStep.Approval.SelfApproval.AssignedUserIds, v.ValueString())
					}
				}
				if step.Approval.UserApproval != nil {
					newPolicyStep.Approval.UserApproval = &shared.UserApproval{
						AllowSelfApproval: step.Approval.UserApproval.AllowSelfApproval.ValueBoolPointer(),
					}
					for _, v := range step.Approval.UserApproval.UserIds {
						newPolicyStep.Approval.UserApproval.UserIds = append(newPolicyStep.Approval.UserApproval.UserIds, v.ValueString())
					}
				}
				steps = append(steps, newPolicyStep)
			}
			if step.Provision != nil {
				newPolicyStep := shared.PolicyStep{
					Provision: &shared.Provision{
						Assigned: step.Provision.Assigned.ValueBoolPointer(),
					},
				}
				if step.Provision.ProvisionPolicy != nil {
					if step.Provision.ProvisionPolicy.ConnectorProvision != nil {
						newPolicyStep.Provision.ProvisionPolicy = &shared.ProvisionPolicy{
							ConnectorProvision: &shared.ConnectorProvision{},
						}
					}
					if step.Provision.ProvisionPolicy.DelegatedProvision != nil {
						newPolicyStep.Provision.ProvisionPolicy = &shared.ProvisionPolicy{
							DelegatedProvision: &shared.DelegatedProvision{
								AppID:         step.Provision.ProvisionPolicy.DelegatedProvision.AppID.ValueStringPointer(),
								EntitlementID: step.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID.ValueStringPointer(),
							},
						}
					}
					if step.Provision.ProvisionPolicy.ManualProvision != nil {
						newPolicyStep.Provision.ProvisionPolicy = &shared.ProvisionPolicy{
							ManualProvision: &shared.ManualProvision{},
						}
					}
			var reject *shared.Reject
			if stepsItem.Reject != nil {
				reject = &shared.Reject{}
			}
			steps = append(steps, shared.PolicyStep{
				Accept:    accept,
				Approval:  approval,
				Provision: provision,
				Reject:    reject,
			})
		}
		policyStepsInst := shared.PolicyStepsInput{
			Steps: steps,
		}
		policySteps[policyStepsKey] = policyStepsInst
	}
	policyType := new(shared.CreatePolicyRequestPolicyType)
	if !r.PolicyType.IsUnknown() && !r.PolicyType.IsNull() {
		*policyType = shared.CreatePolicyRequestPolicyType(r.PolicyType.ValueString())
	} else {
		policyType = nil
	}
	var postActions []shared.PolicyPostActions = nil
	for _, postActionsItem := range r.PostActions {
		certifyRemediateImmediately := new(bool)
		if !postActionsItem.CertifyRemediateImmediately.IsUnknown() && !postActionsItem.CertifyRemediateImmediately.IsNull() {
			*certifyRemediateImmediately = postActionsItem.CertifyRemediateImmediately.ValueBool()
		} else {
			certifyRemediateImmediately = nil
		}
		postActions = append(postActions, shared.PolicyPostActions{
			CertifyRemediateImmediately: certifyRemediateImmediately,
		})
	}
	reassignTasksToDelegates := new(bool)
	if !r.ReassignTasksToDelegates.IsUnknown() && !r.ReassignTasksToDelegates.IsNull() {
		*reassignTasksToDelegates = r.ReassignTasksToDelegates.ValueBool()
	} else {
		reassignTasksToDelegates = nil
	}
	out := shared.CreatePolicyRequest{
		Description:              description,
		DisplayName:              displayName,
		PolicySteps:              policySteps,
		PolicyType:               policyType,
		PostActions:              postActions,
		ReassignTasksToDelegates: reassignTasksToDelegates,
	}
	return &out
}

func (r *PolicyResourceModel) ToGetSDKType() *shared.CreatePolicyRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *PolicyResourceModel) ToUpdateSDKType() *shared.PolicyInput {
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
	policySteps := make(map[string]shared.PolicyStepsInput)
	for policyStepsKey, policyStepsValue := range r.PolicySteps {
		var steps []shared.PolicyStepInput = nil
		for _, stepsItem := range policyStepsValue.Steps {
			var accept *shared.Accept
			if stepsItem.Accept != nil {
				accept = &shared.Accept{}
			}
			// This is all read-only, so we don't need to do anything with it.
			var approval *shared.ApprovalInput
			if stepsItem.Approval != nil {
				var appOwnerApproval *shared.AppOwnerApprovalInput
				if stepsItem.Approval.AppOwnerApproval != nil {
					appOwnerApproval = &shared.AppOwnerApprovalInput{}
				}
				var entitlementOwnerApproval *shared.EntitlementOwnerApprovalInput
				if stepsItem.Approval.EntitlementOwnerApproval != nil {
					entitlementOwnerApproval = &shared.EntitlementOwnerApprovalInput{}
				}
				var expressionApproval *shared.ExpressionApprovalInput
				if stepsItem.Approval.ExpressionApproval != nil {
					expressionApproval = &shared.ExpressionApprovalInput{}
				}
				var appGroupApproval *shared.AppGroupApprovalInput
				if stepsItem.Approval.AppGroupApproval != nil {
					appGroupApproval = &shared.AppGroupApprovalInput{}
				}
				var managerApproval *shared.ManagerApprovalInput
				if stepsItem.Approval.ManagerApproval != nil {
					managerApproval = &shared.ManagerApprovalInput{}
				}
				var selfApproval *shared.SelfApprovalInput
				if stepsItem.Approval.SelfApproval != nil {
					selfApproval = &shared.SelfApprovalInput{}
				}
				var userApproval *shared.UserApprovalInput
				if stepsItem.Approval.UserApproval != nil {
					userApproval = &shared.UserApprovalInput{}
				}
				approval = &shared.ApprovalInput{
					AppOwnerApproval:         appOwnerApproval,
					EntitlementOwnerApproval: entitlementOwnerApproval,
					ExpressionApproval:       expressionApproval,
					AppGroupApproval:         appGroupApproval,
					ManagerApproval:          managerApproval,
					SelfApproval:             selfApproval,
					UserApproval:             userApproval,
				}
			}
			var provision *shared.Provision
			if stepsItem.Provision != nil {
				assigned := new(bool)
				if !stepsItem.Provision.Assigned.IsUnknown() && !stepsItem.Provision.Assigned.IsNull() {
					*assigned = stepsItem.Provision.Assigned.ValueBool()
				} else {
					assigned = nil
				}
				var provisionPolicy *shared.ProvisionPolicy
				if stepsItem.Provision.ProvisionPolicy != nil {
					var connectorProvision *shared.ConnectorProvision
					if stepsItem.Provision.ProvisionPolicy.ConnectorProvision != nil {
						connectorProvision = &shared.ConnectorProvision{}
					}
					var delegatedProvision *shared.DelegatedProvision
					if stepsItem.Provision.ProvisionPolicy.DelegatedProvision != nil {
						appID := new(string)
						if !stepsItem.Provision.ProvisionPolicy.DelegatedProvision.AppID.IsUnknown() && !stepsItem.Provision.ProvisionPolicy.DelegatedProvision.AppID.IsNull() {
							*appID = stepsItem.Provision.ProvisionPolicy.DelegatedProvision.AppID.ValueString()
						} else {
							appID = nil
						}
						entitlementID := new(string)
						if !stepsItem.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID.IsUnknown() && !stepsItem.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID.IsNull() {
							*entitlementID = stepsItem.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID.ValueString()
						} else {
							entitlementID = nil
						}
						implicit := new(bool)
						if !stepsItem.Provision.ProvisionPolicy.DelegatedProvision.Implicit.IsUnknown() && !stepsItem.Provision.ProvisionPolicy.DelegatedProvision.Implicit.IsNull() {
							*implicit = stepsItem.Provision.ProvisionPolicy.DelegatedProvision.Implicit.ValueBool()
						} else {
							implicit = nil
						}
						delegatedProvision = &shared.DelegatedProvision{
							AppID:         appID,
							EntitlementID: entitlementID,
							Implicit:      implicit,
						}
					}
					var manualProvision *shared.ManualProvision
					if stepsItem.Provision.ProvisionPolicy.ManualProvision != nil {
						instructions := new(string)
						if !stepsItem.Provision.ProvisionPolicy.ManualProvision.Instructions.IsUnknown() && !stepsItem.Provision.ProvisionPolicy.ManualProvision.Instructions.IsNull() {
							*instructions = stepsItem.Provision.ProvisionPolicy.ManualProvision.Instructions.ValueString()
						} else {
							instructions = nil
						}
						var userIds []string = nil
						for _, userIdsItem := range stepsItem.Provision.ProvisionPolicy.ManualProvision.UserIds {
							userIds = append(userIds, userIdsItem.ValueString())
						}
						manualProvision = &shared.ManualProvision{
							Instructions: instructions,
							UserIds:      userIds,
						}
					}
					provisionPolicy = &shared.ProvisionPolicy{
						ConnectorProvision: connectorProvision,
						DelegatedProvision: delegatedProvision,
						ManualProvision:    manualProvision,
					}
				}
				var provisionTarget *shared.ProvisionTarget
				if stepsItem.Provision.ProvisionTarget != nil {
					provisionTarget = &shared.ProvisionTarget{
						AppEntitlementID: stepsItem.Provision.ProvisionTarget.AppEntitlementID.ValueStringPointer(),
						AppID:            stepsItem.Provision.ProvisionTarget.AppID.ValueStringPointer(),
						AppUserID:        stepsItem.Provision.ProvisionTarget.AppUserID.ValueStringPointer(),
						GrantDuration:    stepsItem.Provision.ProvisionTarget.GrantDuration.ValueStringPointer(),
					}
				}
				provision = &shared.Provision{
					Assigned:        assigned,
					ProvisionPolicy: provisionPolicy,
					ProvisionTarget: provisionTarget,
				}
			}

			var reject *shared.Reject
			if stepsItem.Reject != nil {
				reject = &shared.Reject{}
			}
			steps = append(steps, shared.PolicyStepInput{
				Accept:    accept,
				Approval:  approval,
				Provision: provision,
				Reject:    reject,
			})
		}
		policyStepsInst := shared.PolicyStepsInput{
			Steps: steps,
		}
		policySteps[policyStepsKey] = policyStepsInst
	}
	policyType := new(shared.PolicyType)
	if !r.PolicyType.IsUnknown() && !r.PolicyType.IsNull() {
		*policyType = shared.PolicyType(r.PolicyType.ValueString())
	} else {
		policyType = nil
	}
	var postActions []shared.PolicyPostActions = nil
	for _, postActionsItem := range r.PostActions {
		certifyRemediateImmediately := new(bool)
		if !postActionsItem.CertifyRemediateImmediately.IsUnknown() && !postActionsItem.CertifyRemediateImmediately.IsNull() {
			*certifyRemediateImmediately = postActionsItem.CertifyRemediateImmediately.ValueBool()
		} else {
			certifyRemediateImmediately = nil
		}
		postActions = append(postActions, shared.PolicyPostActions{
			CertifyRemediateImmediately: certifyRemediateImmediately,
		})
	}
	reassignTasksToDelegates := new(bool)
	if !r.ReassignTasksToDelegates.IsUnknown() && !r.ReassignTasksToDelegates.IsNull() {
		*reassignTasksToDelegates = r.ReassignTasksToDelegates.ValueBool()
	} else {
		reassignTasksToDelegates = nil
	}
	var rules []shared.Rule = nil
	for _, rulesItem := range r.Rules {
		condition := new(string)
		if !rulesItem.Condition.IsUnknown() && !rulesItem.Condition.IsNull() {
			*condition = rulesItem.Condition.ValueString()
		} else {
			condition = nil
		}
		policyKey := new(string)
		if !rulesItem.PolicyKey.IsUnknown() && !rulesItem.PolicyKey.IsNull() {
			*policyKey = rulesItem.PolicyKey.ValueString()
		} else {
			policyKey = nil
	}
	rules = append(rules, shared.Rule{
			Condition: condition,
			PolicyKey: policyKey,
		})
	}
	out := shared.PolicyInput{
		Description:              description,
		DisplayName:              displayName,
		PolicySteps:              policySteps,
		PolicyType:               policyType,
		PostActions:              postActions,
		ReassignTasksToDelegates: reassignTasksToDelegates,
		Rules:                    rules,
	}
	return &out
}

func (r *PolicyResourceModel) ToDeleteSDKType() *shared.DeletePolicyRequest {
	out := shared.DeletePolicyRequest{}
	return &out
}

func (r *PolicyResourceModel) RefreshFromGetResponse(resp *shared.Policy) {
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
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	r.PolicySteps = nil
	if len(resp.PolicySteps) > 0 {
		r.PolicySteps = make(map[string]PolicySteps)
		for policyStepsKey, policyStepsValue := range resp.PolicySteps {
			var policyStepsResult PolicySteps
			policyStepsResult.Steps = nil
			for _, stepsItem := range policyStepsValue.Steps {
				var steps1 PolicyStep
				if stepsItem.Accept == nil {
					steps1.Accept = nil
				} else {
					steps1.Accept = &Accept{}
				}
				if stepsItem.Approval == nil {
					steps1.Approval = nil
				} else {
					steps1.Approval = &ApprovalInput{}
					if stepsItem.Approval.AllowReassignment != nil {
						steps1.Approval.AllowReassignment = types.BoolValue(*stepsItem.Approval.AllowReassignment)
					} else {
						steps1.Approval.AllowReassignment = types.BoolNull()
					}
					if stepsItem.Approval.AppGroupApproval == nil {
						steps1.Approval.AppGroupApproval = nil
					} else {
						steps1.Approval.AppGroupApproval = &AppGroupApprovalInput{}
						if stepsItem.Approval.AppGroupApproval.AllowSelfApproval != nil {
							steps1.Approval.AppGroupApproval.AllowSelfApproval = types.BoolValue(*stepsItem.Approval.AppGroupApproval.AllowSelfApproval)
						} else {
							steps1.Approval.AppGroupApproval.AllowSelfApproval = types.BoolNull()
						}
						if stepsItem.Approval.AppGroupApproval.AppGroupID != nil {
							steps1.Approval.AppGroupApproval.AppGroupID = types.StringValue(*stepsItem.Approval.AppGroupApproval.AppGroupID)
						} else {
							steps1.Approval.AppGroupApproval.AppGroupID = types.StringNull()
						}
						if stepsItem.Approval.AppGroupApproval.AppID != nil {
							steps1.Approval.AppGroupApproval.AppID = types.StringValue(*stepsItem.Approval.AppGroupApproval.AppID)
						} else {
							steps1.Approval.AppGroupApproval.AppID = types.StringNull()
						}
						if stepsItem.Approval.AppGroupApproval.Fallback != nil {
							steps1.Approval.AppGroupApproval.Fallback = types.BoolValue(*stepsItem.Approval.AppGroupApproval.Fallback)
						} else {
							steps1.Approval.AppGroupApproval.Fallback = types.BoolNull()
						}
						steps1.Approval.AppGroupApproval.FallbackUserIds = nil
						for _, v := range stepsItem.Approval.AppGroupApproval.FallbackUserIds {
							steps1.Approval.AppGroupApproval.FallbackUserIds = append(steps1.Approval.AppGroupApproval.FallbackUserIds, types.StringValue(v))
						}
					}
					if stepsItem.Approval.AppOwnerApproval == nil {
						steps1.Approval.AppOwnerApproval = nil
					} else {
						steps1.Approval.AppOwnerApproval = &AppOwnerApprovalInput{}
						if stepsItem.Approval.AppOwnerApproval.AllowSelfApproval != nil {
							steps1.Approval.AppOwnerApproval.AllowSelfApproval = types.BoolValue(*stepsItem.Approval.AppOwnerApproval.AllowSelfApproval)
						} else {
							steps1.Approval.AppOwnerApproval.AllowSelfApproval = types.BoolNull()
						}
					}
					if stepsItem.Approval.Assigned != nil {
						steps1.Approval.Assigned = types.BoolValue(*stepsItem.Approval.Assigned)
					} else {
						steps1.Approval.Assigned = types.BoolNull()
					}
					if stepsItem.Approval.EntitlementOwnerApproval == nil {
						steps1.Approval.EntitlementOwnerApproval = nil
					} else {
						steps1.Approval.EntitlementOwnerApproval = &EntitlementOwnerApprovalInput{}
						if stepsItem.Approval.EntitlementOwnerApproval.AllowSelfApproval != nil {
							steps1.Approval.EntitlementOwnerApproval.AllowSelfApproval = types.BoolValue(*stepsItem.Approval.EntitlementOwnerApproval.AllowSelfApproval)
						} else {
							steps1.Approval.EntitlementOwnerApproval.AllowSelfApproval = types.BoolNull()
						}
						if stepsItem.Approval.EntitlementOwnerApproval.Fallback != nil {
							steps1.Approval.EntitlementOwnerApproval.Fallback = types.BoolValue(*stepsItem.Approval.EntitlementOwnerApproval.Fallback)
						} else {
							steps1.Approval.EntitlementOwnerApproval.Fallback = types.BoolNull()
						}
						steps1.Approval.EntitlementOwnerApproval.FallbackUserIds = nil
						for _, v := range stepsItem.Approval.EntitlementOwnerApproval.FallbackUserIds {
							steps1.Approval.EntitlementOwnerApproval.FallbackUserIds = append(steps1.Approval.EntitlementOwnerApproval.FallbackUserIds, types.StringValue(v))
						}
					}
					if stepsItem.Approval.ExpressionApproval == nil {
						steps1.Approval.ExpressionApproval = nil
					} else {
						steps1.Approval.ExpressionApproval = &ExpressionApprovalInput{}
						if stepsItem.Approval.ExpressionApproval.AllowSelfApproval != nil {
							steps1.Approval.ExpressionApproval.AllowSelfApproval = types.BoolValue(*stepsItem.Approval.ExpressionApproval.AllowSelfApproval)
						} else {
							steps1.Approval.ExpressionApproval.AllowSelfApproval = types.BoolNull()
						}
						steps1.Approval.ExpressionApproval.AssignedUserIds = nil
						for _, v := range stepsItem.Approval.ExpressionApproval.AssignedUserIds {
							steps1.Approval.ExpressionApproval.AssignedUserIds = append(steps1.Approval.ExpressionApproval.AssignedUserIds, types.StringValue(v))
						}
						steps1.Approval.ExpressionApproval.Expressions = nil
						for _, v := range stepsItem.Approval.ExpressionApproval.Expressions {
							steps1.Approval.ExpressionApproval.Expressions = append(steps1.Approval.ExpressionApproval.Expressions, types.StringValue(v))
						}
						if stepsItem.Approval.ExpressionApproval.Fallback != nil {
							steps1.Approval.ExpressionApproval.Fallback = types.BoolValue(*stepsItem.Approval.ExpressionApproval.Fallback)
						} else {
							steps1.Approval.ExpressionApproval.Fallback = types.BoolNull()
						}
						steps1.Approval.ExpressionApproval.FallbackUserIds = nil
						for _, v := range stepsItem.Approval.ExpressionApproval.FallbackUserIds {
							steps1.Approval.ExpressionApproval.FallbackUserIds = append(steps1.Approval.ExpressionApproval.FallbackUserIds, types.StringValue(v))
						}
					}
					if stepsItem.Approval.ManagerApproval == nil {
						steps1.Approval.ManagerApproval = nil
					} else {
						steps1.Approval.ManagerApproval = &ManagerApprovalInput{}
						if stepsItem.Approval.ManagerApproval.AllowSelfApproval != nil {
							steps1.Approval.ManagerApproval.AllowSelfApproval = types.BoolValue(*stepsItem.Approval.ManagerApproval.AllowSelfApproval)
						} else {
							steps1.Approval.ManagerApproval.AllowSelfApproval = types.BoolNull()
						}
						steps1.Approval.ManagerApproval.AssignedUserIds = nil
						for _, v := range stepsItem.Approval.ManagerApproval.AssignedUserIds {
							steps1.Approval.ManagerApproval.AssignedUserIds = append(steps1.Approval.ManagerApproval.AssignedUserIds, types.StringValue(v))
						}
						if stepsItem.Approval.ManagerApproval.Fallback != nil {
							steps1.Approval.ManagerApproval.Fallback = types.BoolValue(*stepsItem.Approval.ManagerApproval.Fallback)
						} else {
							steps1.Approval.ManagerApproval.Fallback = types.BoolNull()
						}
						steps1.Approval.ManagerApproval.FallbackUserIds = nil
						for _, v := range stepsItem.Approval.ManagerApproval.FallbackUserIds {
							steps1.Approval.ManagerApproval.FallbackUserIds = append(steps1.Approval.ManagerApproval.FallbackUserIds, types.StringValue(v))
						}
					}
					if stepsItem.Approval.RequireApprovalReason != nil {
						steps1.Approval.RequireApprovalReason = types.BoolValue(*stepsItem.Approval.RequireApprovalReason)
					} else {
						steps1.Approval.RequireApprovalReason = types.BoolNull()
					}
					if stepsItem.Approval.RequireReassignmentReason != nil {
						steps1.Approval.RequireReassignmentReason = types.BoolValue(*stepsItem.Approval.RequireReassignmentReason)
					} else {
						steps1.Approval.RequireReassignmentReason = types.BoolNull()
					}
					if stepsItem.Approval.SelfApproval == nil {
						steps1.Approval.SelfApproval = nil
					} else {
						steps1.Approval.SelfApproval = &SelfApprovalInput{}
						steps1.Approval.SelfApproval.AssignedUserIds = nil
						for _, v := range stepsItem.Approval.SelfApproval.AssignedUserIds {
							steps1.Approval.SelfApproval.AssignedUserIds = append(steps1.Approval.SelfApproval.AssignedUserIds, types.StringValue(v))
						}
						if stepsItem.Approval.SelfApproval.Fallback != nil {
							steps1.Approval.SelfApproval.Fallback = types.BoolValue(*stepsItem.Approval.SelfApproval.Fallback)
						} else {
							steps1.Approval.SelfApproval.Fallback = types.BoolNull()
						}
						steps1.Approval.SelfApproval.FallbackUserIds = nil
						for _, v := range stepsItem.Approval.SelfApproval.FallbackUserIds {
							steps1.Approval.SelfApproval.FallbackUserIds = append(steps1.Approval.SelfApproval.FallbackUserIds, types.StringValue(v))
						}
					}
					if stepsItem.Approval.UserApproval == nil {
						steps1.Approval.UserApproval = nil
					} else {
						steps1.Approval.UserApproval = &UserApprovalInput{}
						if stepsItem.Approval.UserApproval.AllowSelfApproval != nil {
							steps1.Approval.UserApproval.AllowSelfApproval = types.BoolValue(*stepsItem.Approval.UserApproval.AllowSelfApproval)
						} else {
							steps1.Approval.UserApproval.AllowSelfApproval = types.BoolNull()
						}
						steps1.Approval.UserApproval.UserIds = nil
						for _, v := range stepsItem.Approval.UserApproval.UserIds {
							steps1.Approval.UserApproval.UserIds = append(steps1.Approval.UserApproval.UserIds, types.StringValue(v))
						}
					}
				}
				if stepsItem.Provision == nil {
					steps1.Provision = nil
				} else {
					steps1.Provision = &Provision{}
					if stepsItem.Provision.Assigned != nil {
						steps1.Provision.Assigned = types.BoolValue(*stepsItem.Provision.Assigned)
					} else {
						steps1.Provision.Assigned = types.BoolNull()
					}
					if stepsItem.Provision.ProvisionPolicy == nil {
						steps1.Provision.ProvisionPolicy = nil
					} else {
						steps1.Provision.ProvisionPolicy = &ProvisionPolicy{}
						if stepsItem.Provision.ProvisionPolicy.ConnectorProvision == nil {
							steps1.Provision.ProvisionPolicy.ConnectorProvision = nil
						} else {
							steps1.Provision.ProvisionPolicy.ConnectorProvision = &ConnectorProvision{}
						}
						if stepsItem.Provision.ProvisionPolicy.DelegatedProvision == nil {
							steps1.Provision.ProvisionPolicy.DelegatedProvision = nil
						} else {
							steps1.Provision.ProvisionPolicy.DelegatedProvision = &DelegatedProvision{}
							if stepsItem.Provision.ProvisionPolicy.DelegatedProvision.AppID != nil {
								steps1.Provision.ProvisionPolicy.DelegatedProvision.AppID = types.StringValue(*stepsItem.Provision.ProvisionPolicy.DelegatedProvision.AppID)
							} else {
								steps1.Provision.ProvisionPolicy.DelegatedProvision.AppID = types.StringNull()
							}
							if stepsItem.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID != nil {
								steps1.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID = types.StringValue(*stepsItem.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID)
							} else {
								steps1.Provision.ProvisionPolicy.DelegatedProvision.EntitlementID = types.StringNull()
							}
							if stepsItem.Provision.ProvisionPolicy.DelegatedProvision.Implicit != nil {
								steps1.Provision.ProvisionPolicy.DelegatedProvision.Implicit = types.BoolValue(*stepsItem.Provision.ProvisionPolicy.DelegatedProvision.Implicit)
							} else {
								steps1.Provision.ProvisionPolicy.DelegatedProvision.Implicit = types.BoolNull()
							}
						}
						if stepsItem.Provision.ProvisionPolicy.ManualProvision == nil {
							steps1.Provision.ProvisionPolicy.ManualProvision = nil
						} else {
							steps1.Provision.ProvisionPolicy.ManualProvision = &ManualProvision{}
							if stepsItem.Provision.ProvisionPolicy.ManualProvision.Instructions != nil {
								steps1.Provision.ProvisionPolicy.ManualProvision.Instructions = types.StringValue(*stepsItem.Provision.ProvisionPolicy.ManualProvision.Instructions)
							} else {
								steps1.Provision.ProvisionPolicy.ManualProvision.Instructions = types.StringNull()
							}
							steps1.Provision.ProvisionPolicy.ManualProvision.UserIds = nil
							for _, v := range stepsItem.Provision.ProvisionPolicy.ManualProvision.UserIds {
								steps1.Provision.ProvisionPolicy.ManualProvision.UserIds = append(steps1.Provision.ProvisionPolicy.ManualProvision.UserIds, types.StringValue(v))
							}
						}
					}
					if stepsItem.Provision.ProvisionTarget == nil {
						steps1.Provision.ProvisionTarget = nil
					} else {
						steps1.Provision.ProvisionTarget = &ProvisionTarget{}
						if stepsItem.Provision.ProvisionTarget.AppEntitlementID != nil {
							steps1.Provision.ProvisionTarget.AppEntitlementID = types.StringValue(*stepsItem.Provision.ProvisionTarget.AppEntitlementID)
						} else {
							steps1.Provision.ProvisionTarget.AppEntitlementID = types.StringNull()
						}
						if stepsItem.Provision.ProvisionTarget.AppID != nil {
							steps1.Provision.ProvisionTarget.AppID = types.StringValue(*stepsItem.Provision.ProvisionTarget.AppID)
						} else {
							steps1.Provision.ProvisionTarget.AppID = types.StringNull()
						}
						if stepsItem.Provision.ProvisionTarget.AppUserID != nil {
							steps1.Provision.ProvisionTarget.AppUserID = types.StringValue(*stepsItem.Provision.ProvisionTarget.AppUserID)
						} else {
							steps1.Provision.ProvisionTarget.AppUserID = types.StringNull()
						}
						if stepsItem.Provision.ProvisionTarget.GrantDuration != nil {
							steps1.Provision.ProvisionTarget.GrantDuration = types.StringValue(*stepsItem.Provision.ProvisionTarget.GrantDuration)
						} else {
							steps1.Provision.ProvisionTarget.GrantDuration = types.StringNull()
						}
					}
				}
				if stepsItem.Reject == nil {
					steps1.Reject = nil
				} else {
					steps1.Reject = &Reject{}
				}
			policyStepsResult.Steps = append(policyStepsResult.Steps, steps1)
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
		var postActions1 PolicyPostActions
		if postActionsItem.CertifyRemediateImmediately != nil {
			postActions1.CertifyRemediateImmediately = types.BoolValue(*postActionsItem.CertifyRemediateImmediately)
		} else {
			postActions1.CertifyRemediateImmediately = types.BoolNull()
		}
	if postActionsCount+1 > len(r.PostActions) {
		r.PostActions = append(r.PostActions, postActions1)
	} else {
			r.PostActions[postActionsCount].CertifyRemediateImmediately = postActions1.CertifyRemediateImmediately
		}
	}
	if resp.ReassignTasksToDelegates != nil {
		r.ReassignTasksToDelegates = types.BoolValue(*resp.ReassignTasksToDelegates)
	} else {
		r.ReassignTasksToDelegates = types.BoolNull()
	}
	if len(r.Rules) > len(resp.Rules) {
		r.Rules = r.Rules[:len(resp.Rules)]
	}
	for rulesCount, rulesItem := range resp.Rules {
		var rules1 Rule
		if rulesItem.Condition != nil {
			rules1.Condition = types.StringValue(*rulesItem.Condition)
		} else {
			rules1.Condition = types.StringNull()
		}
		if rulesItem.PolicyKey != nil {
			rules1.PolicyKey = types.StringValue(*rulesItem.PolicyKey)
		} else {
			rules1.PolicyKey = types.StringNull()
		}
		if rulesCount+1 > len(r.Rules) {
			r.Rules = append(r.Rules, rules1)
		} else {
			r.Rules[rulesCount].Condition = rules1.Condition
			r.Rules[rulesCount].PolicyKey = rules1.PolicyKey
		}
	}
	if resp.SystemBuiltin != nil {
		r.SystemBuiltin = types.BoolValue(*resp.SystemBuiltin)
	} else {
		r.SystemBuiltin = types.BoolNull()
	}
	if resp.UpdatedAt != nil {
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	} else {
		r.UpdatedAt = types.StringNull()
	}
}

func (r *PolicyResourceModel) RefreshFromCreateResponse(resp *shared.Policy) {
	r.RefreshFromGetResponse(resp)
}

func (r *PolicyResourceModel) RefreshFromUpdateResponse(resp *shared.Policy) {
	r.RefreshFromGetResponse(resp)
}
