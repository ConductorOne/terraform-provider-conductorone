// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The WebhooksServiceCreateRequest message.
type WebhooksServiceCreateRequest struct {
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName string `json:"displayName"`
	// The url field.
	URL string `json:"url"`
}

func (o *WebhooksServiceCreateRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *WebhooksServiceCreateRequest) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *WebhooksServiceCreateRequest) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
