// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/internal/utils"
	"time"
)

// The AppEntitlementWithExpired message.
type AppEntitlementWithExpired struct {
	// Application User that represents an account in the application.
	AppUser    *AppUser   `json:"appUser,omitempty"`
	Discovered *time.Time `json:"discovered,omitempty"`
	Expired    *time.Time `json:"expired,omitempty"`
	// The User object provides all of the details for an user, as well as some configuration.
	User *User `json:"user,omitempty"`
}

func (a AppEntitlementWithExpired) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppEntitlementWithExpired) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppEntitlementWithExpired) GetAppUser() *AppUser {
	if o == nil {
		return nil
	}
	return o.AppUser
}

func (o *AppEntitlementWithExpired) GetDiscovered() *time.Time {
	if o == nil {
		return nil
	}
	return o.Discovered
}

func (o *AppEntitlementWithExpired) GetExpired() *time.Time {
	if o == nil {
		return nil
	}
	return o.Expired
}

func (o *AppEntitlementWithExpired) GetUser() *User {
	if o == nil {
		return nil
	}
	return o.User
}
