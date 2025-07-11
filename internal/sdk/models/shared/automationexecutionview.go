// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AutomationExecutionView message.
type AutomationExecutionView struct {
	// The AutomationExecution message.
	AutomationExecution *AutomationExecution `json:"automationExecution,omitempty"`
	// The automationExecutionTriggerPath field.
	AutomationExecutionTriggerPath *string `json:"automationExecutionTriggerPath,omitempty"`
	// The automationPath field.
	AutomationPath *string `json:"automationPath,omitempty"`
}

func (o *AutomationExecutionView) GetAutomationExecution() *AutomationExecution {
	if o == nil {
		return nil
	}
	return o.AutomationExecution
}

func (o *AutomationExecutionView) GetAutomationExecutionTriggerPath() *string {
	if o == nil {
		return nil
	}
	return o.AutomationExecutionTriggerPath
}

func (o *AutomationExecutionView) GetAutomationPath() *string {
	if o == nil {
		return nil
	}
	return o.AutomationPath
}
