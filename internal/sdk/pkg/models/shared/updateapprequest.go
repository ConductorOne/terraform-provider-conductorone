// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateAppRequest - The UpdateAppRequest message.
type UpdateAppRequest struct {
	// The App message.
	App *App `json:"app,omitempty"`
	UpdateMask *string `json:"updateMask,omitempty"`
}