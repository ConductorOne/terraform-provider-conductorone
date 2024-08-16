// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PayloadPolicyApprovalStep - The PayloadPolicyApprovalStep message.
type PayloadPolicyApprovalStep struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// List of serialized related objects.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *PayloadPolicyApprovalStep) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *PayloadPolicyApprovalStep) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
