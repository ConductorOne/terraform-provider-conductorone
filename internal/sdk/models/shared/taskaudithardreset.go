// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskAuditHardReset message.
type TaskAuditHardReset struct {
	// The oldPolicyStepId field.
	OldPolicyStepID *string `json:"oldPolicyStepId,omitempty"`
}

func (o *TaskAuditHardReset) GetOldPolicyStepID() *string {
	if o == nil {
		return nil
	}
	return o.OldPolicyStepID
}
