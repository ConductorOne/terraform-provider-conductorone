// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ApprovalInput - The Approval message.
//
// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
//   - users
//   - manager
//   - appOwners
//   - group
//   - self
//   - entitlementOwners
//   - expression
type ApprovalInput struct {
	// App owner approval provides the configuration for an approval step when the app owner is the target.
	AppOwnerApproval *AppOwnerApprovalInput `json:"appOwners,omitempty"`
	// The entitlement owner approval allows configuration of the approval step when the target approvers are the entitlement owners.
	EntitlementOwnerApproval *EntitlementOwnerApprovalInput `json:"entitlementOwners,omitempty"`
	// The ExpressionApproval message.
	ExpressionApproval *ExpressionApprovalInput `json:"expression,omitempty"`
	// The AppGroupApproval object provides the configuration for setting a group as the approvers of an approval policy step.
	AppGroupApproval *AppGroupApprovalInput `json:"group,omitempty"`
	// The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task.
	ManagerApproval *ManagerApprovalInput `json:"manager,omitempty"`
	// The self approval object describes the configuration of a policy step that needs to be approved by the target of the request.
	SelfApproval *SelfApprovalInput `json:"self,omitempty"`
	// The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users.
	UserApproval *UserApprovalInput `json:"users,omitempty"`
}

func (o *ApprovalInput) GetAppOwnerApproval() *AppOwnerApprovalInput {
	if o == nil {
		return nil
	}
	return o.AppOwnerApproval
}

func (o *ApprovalInput) GetEntitlementOwnerApproval() *EntitlementOwnerApprovalInput {
	if o == nil {
		return nil
	}
	return o.EntitlementOwnerApproval
}

func (o *ApprovalInput) GetExpressionApproval() *ExpressionApprovalInput {
	if o == nil {
		return nil
	}
	return o.ExpressionApproval
}

func (o *ApprovalInput) GetAppGroupApproval() *AppGroupApprovalInput {
	if o == nil {
		return nil
	}
	return o.AppGroupApproval
}

func (o *ApprovalInput) GetManagerApproval() *ManagerApprovalInput {
	if o == nil {
		return nil
	}
	return o.ManagerApproval
}

func (o *ApprovalInput) GetSelfApproval() *SelfApprovalInput {
	if o == nil {
		return nil
	}
	return o.SelfApproval
}

func (o *ApprovalInput) GetUserApproval() *UserApprovalInput {
	if o == nil {
		return nil
	}
	return o.UserApproval
}