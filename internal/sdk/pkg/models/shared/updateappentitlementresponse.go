// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// UpdateAppEntitlementResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type UpdateAppEntitlementResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _UpdateAppEntitlementResponseExpanded UpdateAppEntitlementResponseExpanded

func (c *UpdateAppEntitlementResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _UpdateAppEntitlementResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = UpdateAppEntitlementResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c UpdateAppEntitlementResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_UpdateAppEntitlementResponseExpanded(c))
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

// UpdateAppEntitlementResponse - The UpdateAppEntitlementResponse message.
type UpdateAppEntitlementResponse struct {
	// The AppEntitlementView message.
	AppEntitlementView *AppEntitlementView `json:"appEntitlementView,omitempty"`
	// The expanded field.
	Expanded []UpdateAppEntitlementResponseExpanded `json:"expanded,omitempty"`
}