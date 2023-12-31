// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateAppRequest - The CreateAppRequest message.
type CreateAppRequest struct {
	// The certifyPolicyId field.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The grantPolicyId field.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The monthlyCostUsd field.
	MonthlyCostUsd *float64 `json:"monthlyCostUsd,omitempty"`
	// The owners field.
	Owners []string `json:"owners,omitempty"`
	// The revokePolicyId field.
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
}
