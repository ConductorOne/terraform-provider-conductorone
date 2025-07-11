// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// StepUpProviderInput - The StepUpProvider message.
//
// This message contains a oneof named settings. Only a single field of the following list may be set at a time:
//   - oauth2
//   - microsoft
type StepUpProviderInput struct {
	// The clientId field.
	ClientID *string `json:"clientId,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The enabled field.
	Enabled *bool `json:"enabled,omitempty"`
	// The issuerUrl field.
	IssuerURL *string `json:"issuerUrl,omitempty"`
	// StepUpMicrosoftSettings represents a Microsoft Entra Provider using Conditional Access Policies to enforce step-up authentication.
	StepUpMicrosoftSettings *StepUpMicrosoftSettings `json:"microsoft,omitempty"`
	// StepUpOAuth2Settings repersents an OAuth2 provider that supports RFC 9470 <https://www.rfc-editor.org/rfc/rfc9470>
	//
	//  Common ACR values for OAuth2 providers include:
	//    - "urn:okta:loa:1fa:any" (okta)
	//    - "urn:okta:loa:1fa:pwd" (okta)
	//    - "urn:okta:loa:2fa:any" (okta)
	//    - "urn:okta:loa:2fa:any:ifpossible" (okta)
	//    - "phr" (okta)
	//    - "phrh" (okta)
	StepUpOAuth2Settings *StepUpOAuth2Settings `json:"oauth2,omitempty"`
}

func (o *StepUpProviderInput) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *StepUpProviderInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *StepUpProviderInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *StepUpProviderInput) GetIssuerURL() *string {
	if o == nil {
		return nil
	}
	return o.IssuerURL
}

func (o *StepUpProviderInput) GetStepUpMicrosoftSettings() *StepUpMicrosoftSettings {
	if o == nil {
		return nil
	}
	return o.StepUpMicrosoftSettings
}

func (o *StepUpProviderInput) GetStepUpOAuth2Settings() *StepUpOAuth2Settings {
	if o == nil {
		return nil
	}
	return o.StepUpOAuth2Settings
}
