// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// DirectoryServiceListResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type DirectoryServiceListResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _DirectoryServiceListResponseExpanded DirectoryServiceListResponseExpanded

func (c *DirectoryServiceListResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _DirectoryServiceListResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = DirectoryServiceListResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c DirectoryServiceListResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_DirectoryServiceListResponseExpanded(c))
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

// DirectoryServiceListResponse - The DirectoryServiceListResponse message.
type DirectoryServiceListResponse struct {
	// The expanded field.
	Expanded []DirectoryServiceListResponseExpanded `json:"expanded,omitempty"`
	// The list field.
	List []DirectoryView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}
