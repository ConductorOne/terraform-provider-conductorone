// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The GetStepUpProviderResponse message.
type GetStepUpProviderResponse struct {
	// The StepUpProvider message.
	//
	// This message contains a oneof named settings. Only a single field of the following list may be set at a time:
	//   - oauth2
	//   - microsoft
	//
	StepUpProvider *StepUpProvider `json:"stepUpProvider,omitempty"`
}

func (o *GetStepUpProviderResponse) GetStepUpProvider() *StepUpProvider {
	if o == nil {
		return nil
	}
	return o.StepUpProvider
}
