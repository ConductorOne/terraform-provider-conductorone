// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

// TaskActionsServiceApproveResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskActionsServiceApproveResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _TaskActionsServiceApproveResponseExpanded TaskActionsServiceApproveResponseExpanded

func (c *TaskActionsServiceApproveResponseExpanded) UnmarshalJSON(bs []byte) error {
	data := _TaskActionsServiceApproveResponseExpanded{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = TaskActionsServiceApproveResponseExpanded(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c TaskActionsServiceApproveResponseExpanded) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_TaskActionsServiceApproveResponseExpanded(c))
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

// TaskActionsServiceApproveResponse - The TaskActionsServiceApproveResponse message.
type TaskActionsServiceApproveResponse struct {
	// The expanded field.
	Expanded []TaskActionsServiceApproveResponseExpanded `json:"expanded,omitempty"`
	// The TaskView message.
	TaskView *TaskView `json:"taskView,omitempty"`
}
