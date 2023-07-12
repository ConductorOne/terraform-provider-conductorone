// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// UserServiceListResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type UserServiceListResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _UserServiceListResponseExpanded UserServiceListResponseExpanded

func (c *UserServiceListResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _UserServiceListResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = UserServiceListResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c UserServiceListResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_UserServiceListResponseExpanded(c))
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

// UserServiceListResponse - The UserServiceListResponse message.
type UserServiceListResponse struct {
	// The expanded field.
	Expanded []UserServiceListResponseExpanded `json:"expanded,omitempty"`
	// The list field.
	List []UserView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The notificationToken field.
	NotificationToken *string `json:"notificationToken,omitempty"`
}