// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// DirectoryServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type DirectoryServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _DirectoryServiceGetResponseExpanded DirectoryServiceGetResponseExpanded

func (c *DirectoryServiceGetResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _DirectoryServiceGetResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = DirectoryServiceGetResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c DirectoryServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_DirectoryServiceGetResponseExpanded(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

// DirectoryServiceGetResponse - The DirectoryServiceGetResponse message.
type DirectoryServiceGetResponse struct {
	// The DirectoryView message.
	DirectoryView *DirectoryView `json:"directoryView,omitempty"`
	// The expanded field.
	Expanded []DirectoryServiceGetResponseExpanded `json:"expanded,omitempty"`
}
