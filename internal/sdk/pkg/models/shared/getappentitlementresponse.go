// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// GetAppEntitlementResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type GetAppEntitlementResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _GetAppEntitlementResponseExpanded GetAppEntitlementResponseExpanded

func (c *GetAppEntitlementResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _GetAppEntitlementResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = GetAppEntitlementResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c GetAppEntitlementResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_GetAppEntitlementResponseExpanded(c))
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

// GetAppEntitlementResponse - The GetAppEntitlementResponse message.
type GetAppEntitlementResponse struct {
	// The AppEntitlementView message.
	AppEntitlementView *AppEntitlementView `json:"appEntitlementView,omitempty"`
	// The expanded field.
	Expanded []GetAppEntitlementResponseExpanded `json:"expanded,omitempty"`
}