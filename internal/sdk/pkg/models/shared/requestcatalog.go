// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// RequestCatalog - The RequestCatalog message.
type RequestCatalog struct {
	// The accessEntitlements field.
	AccessEntitlements []AppEntitlement `json:"accessEntitlements,omitempty"`
	// The appIds field.
	AppIds    []string   `json:"appIds,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The createdByUserId field.
	CreatedByUserID *string `json:"createdByUserId,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The published field.
	Published *bool      `json:"published,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The visibleToEveryone field.
	VisibleToEveryone *bool `json:"visibleToEveryone,omitempty"`
}
