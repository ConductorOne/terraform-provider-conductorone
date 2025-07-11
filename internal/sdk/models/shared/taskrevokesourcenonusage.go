// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/internal/utils"
	"time"
)

// The TaskRevokeSourceNonUsage message indicates that the source of the revoke task is due to the grant not being used.
type TaskRevokeSourceNonUsage struct {
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	LastLogin *time.Time `json:"lastLogin,omitempty"`
}

func (t TaskRevokeSourceNonUsage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskRevokeSourceNonUsage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskRevokeSourceNonUsage) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *TaskRevokeSourceNonUsage) GetLastLogin() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastLogin
}
