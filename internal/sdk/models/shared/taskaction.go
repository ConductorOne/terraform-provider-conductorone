// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TaskTypes string

const (
	TaskTypesTaskTypeUnspecified TaskTypes = "TASK_TYPE_UNSPECIFIED"
	TaskTypesTaskTypeRequest     TaskTypes = "TASK_TYPE_REQUEST"
	TaskTypesTaskTypeRevoke      TaskTypes = "TASK_TYPE_REVOKE"
	TaskTypesTaskTypeReview      TaskTypes = "TASK_TYPE_REVIEW"
)

func (e TaskTypes) ToPointer() *TaskTypes {
	return &e
}
func (e *TaskTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_TYPE_UNSPECIFIED":
		fallthrough
	case "TASK_TYPE_REQUEST":
		fallthrough
	case "TASK_TYPE_REVOKE":
		fallthrough
	case "TASK_TYPE_REVIEW":
		*e = TaskTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskTypes: %v", v)
	}
}

// The TaskAction message.
//
// This message contains a oneof named action. Only a single field of the following list may be set at a time:
//   - close
//   - reassign
type TaskAction struct {
	// The CloseAction message.
	//
	// This message contains a oneof named user_identifier. Only a single field of the following list may be set at a time:
	//   - userIdCel
	//   - userRef
	//
	CloseAction *CloseAction `json:"close,omitempty"`
	// The ReassignAction message.
	//
	// This message contains a oneof named assignee_user_identifier. Only a single field of the following list may be set at a time:
	//   - assigneeUserIdCel
	//   - assigneeUserRef
	//
	//
	// This message contains a oneof named subject_user_identifier. Only a single field of the following list may be set at a time:
	//   - subjectUserIdCel
	//   - subjectUserRef
	//
	ReassignAction *ReassignAction `json:"reassign,omitempty"`
	// The taskTypes field.
	TaskTypes []TaskTypes `json:"taskTypes,omitempty"`
}

func (o *TaskAction) GetCloseAction() *CloseAction {
	if o == nil {
		return nil
	}
	return o.CloseAction
}

func (o *TaskAction) GetReassignAction() *ReassignAction {
	if o == nil {
		return nil
	}
	return o.ReassignAction
}

func (o *TaskAction) GetTaskTypes() []TaskTypes {
	if o == nil {
		return nil
	}
	return o.TaskTypes
}
