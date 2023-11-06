// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"conductorone/internal/sdk/pkg/utils"
)

// DirectoryServiceCreateResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type DirectoryServiceCreateResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string     `json:"@type,omitempty"`
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
}

func (d DirectoryServiceCreateResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DirectoryServiceCreateResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DirectoryServiceCreateResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *DirectoryServiceCreateResponseExpanded) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The DirectoryServiceCreateResponse message.
type DirectoryServiceCreateResponse struct {
	// The directory view contains a directory and an app_path which is a JSONPATH set to the location in the expand mask that the expanded app will live if requested by the expander.
	DirectoryView *DirectoryView `json:"directoryView,omitempty"`
	// List of serialized related objects.
	Expanded []DirectoryServiceCreateResponseExpanded `json:"expanded,omitempty"`
}

func (o *DirectoryServiceCreateResponse) GetDirectoryView() *DirectoryView {
	if o == nil {
		return nil
	}
	return o.DirectoryView
}

func (o *DirectoryServiceCreateResponse) GetExpanded() []DirectoryServiceCreateResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
