// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"conductorone/internal/sdk/pkg/utils"
	"time"
)

// DeniedAction - The denied action indicates that the c1.api.policy.v1.ApprovalInstance had an outcome of denied.
type DeniedAction struct {
	DeniedAt *time.Time `json:"deniedAt,omitempty"`
	// The UserID that denied this step.
	UserID *string `json:"userId,omitempty"`
}

func (d DeniedAction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeniedAction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeniedAction) GetDeniedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeniedAt
}

func (o *DeniedAction) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
