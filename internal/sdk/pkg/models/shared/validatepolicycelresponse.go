// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ValidatePolicyCELResponse - The ValidatePolicyCELResponse message.
type ValidatePolicyCELResponse struct {
	// The markers field.
	Markers []Marker `json:"markers,omitempty"`
}

func (o *ValidatePolicyCELResponse) GetMarkers() []Marker {
	if o == nil {
		return nil
	}
	return o.Markers
}
