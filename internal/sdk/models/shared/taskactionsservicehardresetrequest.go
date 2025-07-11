// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskActionsServiceHardResetRequest object lets you reset a task and recalculate its policy.
type TaskActionsServiceHardResetRequest struct {
	// The comment attached to the request.
	Comment *string `json:"comment,omitempty"`
}

func (o *TaskActionsServiceHardResetRequest) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}
