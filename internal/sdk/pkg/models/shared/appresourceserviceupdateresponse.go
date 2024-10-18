// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/utils"
)

// AppResourceServiceUpdateResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type AppResourceServiceUpdateResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string `json:"@type,omitempty"`
	AdditionalProperties any     `additionalProperties:"true" json:"-"`
}

func (a AppResourceServiceUpdateResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppResourceServiceUpdateResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppResourceServiceUpdateResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *AppResourceServiceUpdateResponseExpanded) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The AppResourceServiceUpdateResponse message.
type AppResourceServiceUpdateResponse struct {
	// The app resource view returns an app resource with paths for items in the expand mask filled in when this response is returned and a request expand mask has "*" or "app_id" or "resource_type_id".
	AppResourceView *AppResourceView `json:"appResourceView,omitempty"`
	// The expanded field.
	Expanded []AppResourceServiceUpdateResponseExpanded `json:"expanded,omitempty"`
}

func (o *AppResourceServiceUpdateResponse) GetAppResourceView() *AppResourceView {
	if o == nil {
		return nil
	}
	return o.AppResourceView
}

func (o *AppResourceServiceUpdateResponse) GetExpanded() []AppResourceServiceUpdateResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}