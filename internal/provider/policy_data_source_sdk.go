package provider

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *PolicyDataSourceModel) RefreshFromGetResponse(resp *shared.Policy) {
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
	if r.PolicySteps == nil && len(resp.PolicySteps) > 0 {
		r.PolicySteps = make(map[string]PolicySteps)
		for policyStepsKey, policyStepsValue := range resp.PolicySteps {
			var policyStepsResult PolicySteps
			policyStepsResult.Steps = nil
			for _, stepsItem := range policyStepsValue.Steps {
				var steps1 PolicyStep
				if steps1.Approval == nil {
					steps1.Approval = &ApprovalInput{}
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
					if steps1.Approval.AppGroupApproval == nil {
						steps1.Approval.AppGroupApproval = &AppGroupApprovalInput{}
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
					if steps1.Approval.AppOwnerApproval == nil {
						steps1.Approval.AppOwnerApproval = &AppOwnerApprovalInput{}
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
					if steps1.Approval.EntitlementOwnerApproval == nil {
						steps1.Approval.EntitlementOwnerApproval = &EntitlementOwnerApprovalInput{}
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
					if steps1.Approval.ManagerApproval == nil {
						steps1.Approval.ManagerApproval = &ManagerApprovalInput{}
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
					if steps1.Approval.SelfApproval == nil {
						steps1.Approval.SelfApproval = &SelfApprovalInput{}
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
					if steps1.Approval.UserApproval == nil {
						steps1.Approval.UserApproval = &UserApprovalInput{}
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
				if steps1.Provision == nil {
					steps1.Provision = &Accept{}
				}
				if stepsItem.Provision == nil {
					steps1.Provision = nil
				} else {
					steps1.Provision = &Accept{}
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
	r.PostActions = nil
	for _, postActionsItem := range resp.PostActions {
		var postActions1 PolicyPostActions
		if postActionsItem.CertifyRemediateImmediately != nil {
			postActions1.CertifyRemediateImmediately = types.BoolValue(*postActionsItem.CertifyRemediateImmediately)
		} else {
			postActions1.CertifyRemediateImmediately = types.BoolNull()
		}
		r.PostActions = append(r.PostActions, postActions1)
	}
	if resp.ReassignTasksToDelegates != nil {
		r.ReassignTasksToDelegates = types.BoolValue(*resp.ReassignTasksToDelegates)
	} else {
		r.ReassignTasksToDelegates = types.BoolNull()
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
