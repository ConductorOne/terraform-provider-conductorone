// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"conductorone/internal/sdk/pkg/utils"
	"time"
)

// The AppEntitlementUserView (aka grant view) describes the relationship between an app user and an entitlement. They have more recently been referred to as grants.
type AppEntitlementUserView struct {
	// The AppUserView contains an app user as well as paths for apps, identity users, and last usage in expanded arrays.
	AppUserView                            *AppUserView `json:"appUser,omitempty"`
	AppEntitlementUserBindingCreatedAt     *time.Time   `json:"appEntitlementUserBindingCreatedAt,omitempty"`
	AppEntitlementUserBindingDeprovisionAt *time.Time   `json:"appEntitlementUserBindingDeprovisionAt,omitempty"`
}

func (a AppEntitlementUserView) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppEntitlementUserView) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppEntitlementUserView) GetAppUserView() *AppUserView {
	if o == nil {
		return nil
	}
	return o.AppUserView
}

func (o *AppEntitlementUserView) GetAppEntitlementUserBindingCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AppEntitlementUserBindingCreatedAt
}

func (o *AppEntitlementUserView) GetAppEntitlementUserBindingDeprovisionAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AppEntitlementUserBindingDeprovisionAt
}
