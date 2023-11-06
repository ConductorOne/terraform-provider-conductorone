// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"conductorone/internal/sdk/pkg/utils"
	"time"
)

// The AppEntitlementUserBinding represents the relationship that gives an app user access to an app entitlement
type AppEntitlementUserBinding struct {
	// The ID of the app entitlement that the app user has access to
	AppEntitlementID *string `json:"appEntitlementId,omitempty"`
	// The ID of the app associated with the app entitlement
	AppID *string `json:"appId,omitempty"`
	// The ID of the app user that has access to the app entitlement
	AppUserID     *string    `json:"appUserId,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
	DeprovisionAt *time.Time `json:"deprovisionAt,omitempty"`
}

func (a AppEntitlementUserBinding) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppEntitlementUserBinding) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppEntitlementUserBinding) GetAppEntitlementID() *string {
	if o == nil {
		return nil
	}
	return o.AppEntitlementID
}

func (o *AppEntitlementUserBinding) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppEntitlementUserBinding) GetAppUserID() *string {
	if o == nil {
		return nil
	}
	return o.AppUserID
}

func (o *AppEntitlementUserBinding) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppEntitlementUserBinding) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppEntitlementUserBinding) GetDeprovisionAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeprovisionAt
}
