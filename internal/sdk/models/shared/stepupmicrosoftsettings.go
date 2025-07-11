// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// StepUpMicrosoftSettings represents a Microsoft Entra Provider using Conditional Access Policies to enforce step-up authentication.
type StepUpMicrosoftSettings struct {
	// The conditionalAccessIds field.
	ConditionalAccessIds []string `json:"conditionalAccessIds,omitempty"`
	// The tenant field.
	Tenant *string `json:"tenant,omitempty"`
}

func (o *StepUpMicrosoftSettings) GetConditionalAccessIds() []string {
	if o == nil {
		return nil
	}
	return o.ConditionalAccessIds
}

func (o *StepUpMicrosoftSettings) GetTenant() *string {
	if o == nil {
		return nil
	}
	return o.Tenant
}
