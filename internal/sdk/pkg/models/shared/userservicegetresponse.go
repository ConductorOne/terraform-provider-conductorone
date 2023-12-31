// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// UserServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type UserServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _UserServiceGetResponseExpanded UserServiceGetResponseExpanded

func (c *UserServiceGetResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _UserServiceGetResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = UserServiceGetResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c UserServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_UserServiceGetResponseExpanded(c))
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

// UserServiceGetResponse - The UserServiceGetResponse message.
type UserServiceGetResponse struct {
	// The UserView message.
	UserView *UserView `json:"userView,omitempty"`
	// The expanded field.
	Expanded []UserServiceGetResponseExpanded `json:"expanded,omitempty"`
}
