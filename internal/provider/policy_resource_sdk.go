// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func BoolPointer(b bool) *bool {
	return &b
}

const prefix = "c1_prefixed_"

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func PrependPolicyStepId(id string) string {
	if isDigit(id[0]) {
		return prefix + id
	}
	return id
}

func RemovePolicyStepIdPrefix(id string) string {
	prefix_length := len(prefix)
	// We only want to remove the prefix if we appended it, this is not a foolproof way to check but the odds someone will have a policy step id that starts with c1_prefixed_\d{1} is low.
	if len(id) >= prefix_length+1 && id[:prefix_length] == prefix && isDigit(id[prefix_length]) {
		return strings.TrimPrefix(id, prefix)
	}
	return id
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
		var steps []shared.PolicyStep = nil
		for _, step := range policyStepsValue.Steps {
			if step.Accept != nil {
				newPolicyStep := shared.PolicyStep{
					Accept: &shared.Accept{},
				}
				steps = append(steps, newPolicyStep)
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
				if step.Approval.AppGroupApproval != nil {
					newPolicyStep.Approval.AppGroupApproval = &shared.AppGroupApproval{
						AllowSelfApproval: step.Approval.AppGroupApproval.AllowSelfApproval.ValueBoolPointer(),
						AppGroupID:        step.Approval.AppGroupApproval.AppGroupID.ValueStringPointer(),
						AppID:             step.Approval.AppGroupApproval.AppID.ValueStringPointer(),
						Fallback:          step.Approval.AppGroupApproval.Fallback.ValueBoolPointer(),
					}
					for _, v := range step.Approval.AppGroupApproval.FallbackUserIds {
						newPolicyStep.Approval.AppGroupApproval.FallbackUserIds = append(newPolicyStep.Approval.AppGroupApproval.FallbackUserIds, v.ValueString())
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
				if step.Approval.ExpressionApproval != nil {
					newPolicyStep.Approval.ExpressionApproval = &shared.ExpressionApproval{
						AllowSelfApproval: step.Approval.ExpressionApproval.AllowSelfApproval.ValueBoolPointer(),
						Fallback:          step.Approval.ExpressionApproval.Fallback.ValueBoolPointer(),
					}
					for _, v := range step.Approval.ExpressionApproval.AssignedUserIds {
						newPolicyStep.Approval.ExpressionApproval.AssignedUserIds = append(newPolicyStep.Approval.ExpressionApproval.AssignedUserIds, v.ValueString())
					}
					for _, v := range step.Approval.ExpressionApproval.Expressions {
						newPolicyStep.Approval.ExpressionApproval.Expressions = append(newPolicyStep.Approval.ExpressionApproval.Expressions, v.ValueString())
					}
					for _, v := range step.Approval.ExpressionApproval.FallbackUserIds {
						newPolicyStep.Approval.ExpressionApproval.FallbackUserIds = append(newPolicyStep.Approval.ExpressionApproval.FallbackUserIds, v.ValueString())
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
					Provision: &shared.Provision{},
				}
				steps = append(steps, newPolicyStep)
			}
			if step.Reject != nil {
				newPolicyStep := shared.PolicyStep{
					Reject: &shared.Reject{},
				}
				steps = append(steps, newPolicyStep)
			}
		}
		policyStepsInst := shared.PolicySteps{
			Steps: steps,
		}
		policySteps[PrependPolicyStepId(policyStepsKey)] = policyStepsInst
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
	out := shared.CreatePolicyRequest{
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
	policySteps := make(map[string]shared.PolicySteps)
	for policyStepsKey, policyStepsValue := range r.PolicySteps {
		var steps []shared.PolicyStep = nil
		for _, step := range policyStepsValue.Steps {
			if step.Accept != nil {
				newPolicyStep := shared.PolicyStep{
					Accept: &shared.Accept{},
				}
				steps = append(steps, newPolicyStep)
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
				if step.Approval.AppGroupApproval != nil {
					newPolicyStep.Approval.AppGroupApproval = &shared.AppGroupApproval{
						AllowSelfApproval: step.Approval.AppGroupApproval.AllowSelfApproval.ValueBoolPointer(),
						AppGroupID:        step.Approval.AppGroupApproval.AppGroupID.ValueStringPointer(),
						AppID:             step.Approval.AppGroupApproval.AppID.ValueStringPointer(),
						Fallback:          step.Approval.AppGroupApproval.Fallback.ValueBoolPointer(),
					}
					for _, v := range step.Approval.AppGroupApproval.FallbackUserIds {
						newPolicyStep.Approval.AppGroupApproval.FallbackUserIds = append(newPolicyStep.Approval.AppGroupApproval.FallbackUserIds, v.ValueString())
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
				if step.Approval.ExpressionApproval != nil {
					newPolicyStep.Approval.ExpressionApproval = &shared.ExpressionApproval{
						AllowSelfApproval: step.Approval.ExpressionApproval.AllowSelfApproval.ValueBoolPointer(),
						Fallback:          step.Approval.ExpressionApproval.Fallback.ValueBoolPointer(),
					}
					for _, v := range step.Approval.ExpressionApproval.AssignedUserIds {
						newPolicyStep.Approval.ExpressionApproval.AssignedUserIds = append(newPolicyStep.Approval.ExpressionApproval.AssignedUserIds, v.ValueString())
					}
					for _, v := range step.Approval.ExpressionApproval.Expressions {
						newPolicyStep.Approval.ExpressionApproval.Expressions = append(newPolicyStep.Approval.ExpressionApproval.Expressions, v.ValueString())
					}
					for _, v := range step.Approval.ExpressionApproval.FallbackUserIds {
						newPolicyStep.Approval.ExpressionApproval.FallbackUserIds = append(newPolicyStep.Approval.ExpressionApproval.FallbackUserIds, v.ValueString())
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
					Provision: &shared.Provision{},
				}
				steps = append(steps, newPolicyStep)
			}
			if step.Reject != nil {
				newPolicyStep := shared.PolicyStep{
					Reject: &shared.Reject{},
				}
				steps = append(steps, newPolicyStep)
			}
		}
		policyStepsInst := shared.PolicySteps{
			Steps: steps,
		}
		policySteps[RemovePolicyStepIdPrefix(policyStepsKey)] = policyStepsInst
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
					// Empty struct type inference is the same, so we can just use the same struct as Accept
					steps1.Provision = &Accept{}
				}
				if stepsItem.Reject == nil {
					steps1.Reject = nil
				} else {
					steps1.Reject = &Reject{}
				}
				policyStepsResult.Steps = append(policyStepsResult.Steps, steps1)
			}
			r.PolicySteps[PrependPolicyStepId(policyStepsKey)] = policyStepsResult
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
