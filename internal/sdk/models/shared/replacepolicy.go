// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ReplacePolicy message.
type ReplacePolicy struct {
	// The policyId field.
	PolicyID *string `json:"policyId,omitempty"`
}

func (o *ReplacePolicy) GetPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.PolicyID
}
