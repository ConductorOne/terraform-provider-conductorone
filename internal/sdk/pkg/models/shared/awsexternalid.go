// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The AWSExternalID message.
type AWSExternalID struct {
	// The externalId field.
	ExternalID *string `json:"externalId,omitempty"`
}

func (o *AWSExternalID) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}