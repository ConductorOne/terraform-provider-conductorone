// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskActionsServiceApproveRequest object lets you approve a task.
type TaskActionsServiceApproveRequest struct {
	// The comment attached to the request.
	Comment *string `json:"comment,omitempty"`
	// The ID of the policy step on the given task to approve.
	PolicyStepID string `json:"policyStepId"`
}

func (o *TaskActionsServiceApproveRequest) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *TaskActionsServiceApproveRequest) GetPolicyStepID() string {
	if o == nil {
		return ""
	}
	return o.PolicyStepID
}
