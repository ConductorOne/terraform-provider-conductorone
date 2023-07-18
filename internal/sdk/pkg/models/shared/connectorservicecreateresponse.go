// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// ConnectorServiceCreateResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type ConnectorServiceCreateResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _ConnectorServiceCreateResponseExpanded ConnectorServiceCreateResponseExpanded

func (c *ConnectorServiceCreateResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _ConnectorServiceCreateResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = ConnectorServiceCreateResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c ConnectorServiceCreateResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_ConnectorServiceCreateResponseExpanded(c))
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

// ConnectorServiceCreateResponse - The ConnectorServiceCreateResponse message.
type ConnectorServiceCreateResponse struct {
	// The ConnectorView message.
	ConnectorView *ConnectorView `json:"connectorView,omitempty"`
	// The expanded field.
	Expanded []ConnectorServiceCreateResponseExpanded `json:"expanded,omitempty"`
}
