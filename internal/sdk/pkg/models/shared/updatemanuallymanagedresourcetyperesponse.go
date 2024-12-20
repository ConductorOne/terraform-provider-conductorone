// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/utils"
)

// UpdateManuallyManagedResourceTypeResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type UpdateManuallyManagedResourceTypeResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string `json:"@type,omitempty"`
	AdditionalProperties any     `additionalProperties:"true" json:"-"`
}

func (u UpdateManuallyManagedResourceTypeResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateManuallyManagedResourceTypeResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateManuallyManagedResourceTypeResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *UpdateManuallyManagedResourceTypeResponseExpanded) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The UpdateManuallyManagedResourceTypeResponse message.
type UpdateManuallyManagedResourceTypeResponse struct {
	// The AppResourceType is referenced by an app entitlement defining its resource types. Commonly things like Group or Role.
	AppResourceType *AppResourceType `json:"appResourceType,omitempty"`
	// The expanded field.
	Expanded []UpdateManuallyManagedResourceTypeResponseExpanded `json:"expanded,omitempty"`
}

func (o *UpdateManuallyManagedResourceTypeResponse) GetAppResourceType() *AppResourceType {
	if o == nil {
		return nil
	}
	return o.AppResourceType
}

func (o *UpdateManuallyManagedResourceTypeResponse) GetExpanded() []UpdateManuallyManagedResourceTypeResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
