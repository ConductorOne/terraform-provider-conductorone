// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateSessionSettingsResponse - The UpdateSessionSettingsResponse message.
type UpdateSessionSettingsResponse struct {
	// The SessionSettings message.
	SessionSettings *SessionSettings `json:"sessionSettings,omitempty"`
}

func (o *UpdateSessionSettingsResponse) GetSessionSettings() *SessionSettings {
	if o == nil {
		return nil
	}
	return o.SessionSettings
}
