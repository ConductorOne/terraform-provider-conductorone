// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// TaskServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _TaskServiceGetResponseExpanded TaskServiceGetResponseExpanded

func (c *TaskServiceGetResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _TaskServiceGetResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = TaskServiceGetResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c TaskServiceGetResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_TaskServiceGetResponseExpanded(c))
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

// TaskServiceGetResponse - The TaskServiceGetResponse message.
type TaskServiceGetResponse struct {
	// The TaskView message.
	TaskView *TaskView `json:"taskView,omitempty"`
	// The expanded field.
	Expanded []TaskServiceGetResponseExpanded `json:"expanded,omitempty"`
}
