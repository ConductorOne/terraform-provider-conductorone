// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The WebhookSpec message.
type WebhookSpec struct {
	// The destination field.
	Destination *string `json:"destination,omitempty"`
}

func (o *WebhookSpec) GetDestination() *string {
	if o == nil {
		return nil
	}
	return o.Destination
}
