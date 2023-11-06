// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AppUserStatusInput - The satus of the applicaiton user.
type AppUserStatusInput struct {
}

// Status - The application user status field.
type Status string

const (
	StatusStatusUnspecified Status = "STATUS_UNSPECIFIED"
	StatusStatusEnabled     Status = "STATUS_ENABLED"
	StatusStatusDisabled    Status = "STATUS_DISABLED"
	StatusStatusDeleted     Status = "STATUS_DELETED"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STATUS_UNSPECIFIED":
		fallthrough
	case "STATUS_ENABLED":
		fallthrough
	case "STATUS_DISABLED":
		fallthrough
	case "STATUS_DELETED":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// AppUserStatus - The satus of the applicaiton user.
type AppUserStatus struct {
	// The details of applicaiton user status.
	Details *string `json:"details,omitempty"`
	// The application user status field.
	Status *Status `json:"status,omitempty"`
}

func (o *AppUserStatus) GetDetails() *string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *AppUserStatus) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}
