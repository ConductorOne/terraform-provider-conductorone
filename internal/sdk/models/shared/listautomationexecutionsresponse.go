// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ListAutomationExecutionsResponse message.
type ListAutomationExecutionsResponse struct {
	// The automationExecutions field.
	AutomationExecutions []AutomationExecution `json:"automationExecutions,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *ListAutomationExecutionsResponse) GetAutomationExecutions() []AutomationExecution {
	if o == nil {
		return nil
	}
	return o.AutomationExecutions
}

func (o *ListAutomationExecutionsResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
