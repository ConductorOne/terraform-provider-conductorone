// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponsePolicyApprovalStepReassign - The ResponsePolicyApprovalStepReassign message.
type ResponsePolicyApprovalStepReassign struct {
	// optional comment
	Comment *string `json:"comment,omitempty"`
	// The newStepUserIds field.
	NewStepUserIds []string `json:"newStepUserIds,omitempty"`
}

func (o *ResponsePolicyApprovalStepReassign) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *ResponsePolicyApprovalStepReassign) GetNewStepUserIds() []string {
	if o == nil {
		return nil
	}
	return o.NewStepUserIds
}
