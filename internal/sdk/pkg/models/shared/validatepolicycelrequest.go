// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ValidatePolicyCELRequest - The ValidatePolicyCELRequest message.
type ValidatePolicyCELRequest struct {
	// The cel field.
	Cel *string `json:"cel,omitempty"`
}

func (o *ValidatePolicyCELRequest) GetCel() *string {
	if o == nil {
		return nil
	}
	return o.Cel
}
