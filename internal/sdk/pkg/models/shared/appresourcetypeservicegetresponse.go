// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/utils"
)

// AppResourceTypeServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type AppResourceTypeServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string     `json:"@type,omitempty"`
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
}

func (a AppResourceTypeServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppResourceTypeServiceGetResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppResourceTypeServiceGetResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *AppResourceTypeServiceGetResponseExpanded) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The AppResourceTypeServiceGetResponse contains an expanded array containing the expanded values indicated by the expand mask
//
//	in the request and an app resource type view containing the resource type and JSONPATHs indicating which objects are where in the expand mask.
type AppResourceTypeServiceGetResponse struct {
	// The AppResourceTypeView message.
	AppResourceTypeView *AppResourceTypeView `json:"appResourceTypeView,omitempty"`
	// List of serialized related objects.
	Expanded []AppResourceTypeServiceGetResponseExpanded `json:"expanded,omitempty"`
}

func (o *AppResourceTypeServiceGetResponse) GetAppResourceTypeView() *AppResourceTypeView {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeView
}

func (o *AppResourceTypeServiceGetResponse) GetExpanded() []AppResourceTypeServiceGetResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
