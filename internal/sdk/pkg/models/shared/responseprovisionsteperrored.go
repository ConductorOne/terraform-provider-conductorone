// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ResponseProvisionStepErrored - The ResponseProvisionStepErrored message.
type ResponseProvisionStepErrored struct {
	// optional comment
	Comment *string `json:"comment,omitempty"`
}

func (o *ResponseProvisionStepErrored) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}
