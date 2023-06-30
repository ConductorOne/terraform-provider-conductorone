// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Provision - The Provision message.
type Provision struct {
	// The assigned field.
	Assigned *bool `json:"assigned,omitempty"`
	// The ProvisionPolicy message.
	//
	// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
	//   - connector
	//   - manual
	//   - delegated
	//
	ProvisionPolicy *ProvisionPolicy `json:"provisionPolicy,omitempty"`
}
