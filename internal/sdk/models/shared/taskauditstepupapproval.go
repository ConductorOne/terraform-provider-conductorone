// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskAuditStepUpApproval message.
type TaskAuditStepUpApproval struct {
	// The stepUpTransactionId field.
	StepUpTransactionID *string `json:"stepUpTransactionId,omitempty"`
}

func (o *TaskAuditStepUpApproval) GetStepUpTransactionID() *string {
	if o == nil {
		return nil
	}
	return o.StepUpTransactionID
}
