// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// WebhooksServiceCreateRequest - The WebhooksServiceCreateRequest message.
type WebhooksServiceCreateRequest struct {
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The url field.
	URL *string `json:"url,omitempty"`
}

func (o *WebhooksServiceCreateRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *WebhooksServiceCreateRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *WebhooksServiceCreateRequest) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
