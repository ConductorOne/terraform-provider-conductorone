// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestCatalogManagementServiceCreateRequest - The RequestCatalogManagementServiceCreateRequest message.
type RequestCatalogManagementServiceCreateRequest struct {
	// The RequestCatalogExpandMask message.
	RequestCatalogExpandMask *RequestCatalogExpandMask `json:"expandMask,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The ownerIds field.
	OwnerIds []string `json:"ownerIds,omitempty"`
	// The published field.
	Published *bool `json:"published,omitempty"`
	// The visibleToEveryone field.
	VisibleToEveryone *bool `json:"visibleToEveryone,omitempty"`
}
