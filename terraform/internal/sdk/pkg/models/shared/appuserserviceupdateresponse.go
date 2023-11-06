// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"conductorone/internal/sdk/pkg/utils"
)

// AppUserServiceUpdateResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type AppUserServiceUpdateResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string     `json:"@type,omitempty"`
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
}

func (a AppUserServiceUpdateResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppUserServiceUpdateResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppUserServiceUpdateResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *AppUserServiceUpdateResponseExpanded) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The AppUserServiceUpdateResponse message.
type AppUserServiceUpdateResponse struct {
	// The AppUserView contains an app user as well as paths for apps, identity users, and last usage in expanded arrays.
	AppUserView *AppUserView `json:"appUserView,omitempty"`
	// The expanded field.
	Expanded []AppUserServiceUpdateResponseExpanded `json:"expanded,omitempty"`
}

func (o *AppUserServiceUpdateResponse) GetAppUserView() *AppUserView {
	if o == nil {
		return nil
	}
	return o.AppUserView
}

func (o *AppUserServiceUpdateResponse) GetExpanded() []AppUserServiceUpdateResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
