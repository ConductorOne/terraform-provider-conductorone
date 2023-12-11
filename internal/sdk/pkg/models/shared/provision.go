// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Provision - The provision step references a provision policy for this step.
type Provision struct {
	// Provision should be empty on the Policy struct, this field is populated by the app entitlement, it is not empty when getting Policies on Tickets
	// However, terraform does not support tickets, so this field can just point to an empty struct, type inference for empty structs are the same so we can just use the same struct as Accept
}
