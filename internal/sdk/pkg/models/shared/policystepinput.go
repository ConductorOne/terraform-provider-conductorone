// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PolicyStepInput - The PolicyStep message.
//
// This message contains a oneof named step. Only a single field of the following list may be set at a time:
//   - approval
//   - provision
//   - accept
//   - reject
type PolicyStepInput struct {
	// This policy step indicates that a ticket should have an approved outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.
	Accept *Accept `json:"accept,omitempty"`
	// The Approval message.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - users
	//   - manager
	//   - appOwners
	//   - group
	//   - self
	//   - entitlementOwners
	//   - expression
	//
	Approval *ApprovalInput `json:"approval,omitempty"`
	// The provision step references a provision policy for this step.
	Provision *Provision `json:"provision,omitempty"`
	// This policy step indicates that a ticket should have a denied outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.
	Reject *Reject `json:"reject,omitempty"`
}

func (o *PolicyStepInput) GetAccept() *Accept {
	if o == nil {
		return nil
	}
	return o.Accept
}

func (o *PolicyStepInput) GetApproval() *ApprovalInput {
	if o == nil {
		return nil
	}
	return o.Approval
}

func (o *PolicyStepInput) GetProvision() *Provision {
	if o == nil {
		return nil
	}
	return o.Provision
}

func (o *PolicyStepInput) GetReject() *Reject {
	if o == nil {
		return nil
	}
	return o.Reject
}
