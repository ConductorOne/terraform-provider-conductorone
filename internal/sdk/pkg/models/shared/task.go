// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type TaskActions string

const (
	TaskActionsTaskActionTypeUnspecified                   TaskActions = "TASK_ACTION_TYPE_UNSPECIFIED"
	TaskActionsTaskActionTypeClose                         TaskActions = "TASK_ACTION_TYPE_CLOSE"
	TaskActionsTaskActionTypeApprove                       TaskActions = "TASK_ACTION_TYPE_APPROVE"
	TaskActionsTaskActionTypeDeny                          TaskActions = "TASK_ACTION_TYPE_DENY"
	TaskActionsTaskActionTypeComment                       TaskActions = "TASK_ACTION_TYPE_COMMENT"
	TaskActionsTaskActionTypeDelete                        TaskActions = "TASK_ACTION_TYPE_DELETE"
	TaskActionsTaskActionTypeReassign                      TaskActions = "TASK_ACTION_TYPE_REASSIGN"
	TaskActionsTaskActionTypeRestart                       TaskActions = "TASK_ACTION_TYPE_RESTART"
	TaskActionsTaskActionTypeSendReminder                  TaskActions = "TASK_ACTION_TYPE_SEND_REMINDER"
	TaskActionsTaskActionTypeProvisionComplete             TaskActions = "TASK_ACTION_TYPE_PROVISION_COMPLETE"
	TaskActionsTaskActionTypeProvisionCancelled            TaskActions = "TASK_ACTION_TYPE_PROVISION_CANCELLED"
	TaskActionsTaskActionTypeProvisionErrored              TaskActions = "TASK_ACTION_TYPE_PROVISION_ERRORED"
	TaskActionsTaskActionTypeProvisionAppUserTargetCreated TaskActions = "TASK_ACTION_TYPE_PROVISION_APP_USER_TARGET_CREATED"
	TaskActionsTaskActionTypeRollbackSkipped               TaskActions = "TASK_ACTION_TYPE_ROLLBACK_SKIPPED"
	TaskActionsTaskActionTypeHardReset                     TaskActions = "TASK_ACTION_TYPE_HARD_RESET"
)

func (e TaskActions) ToPointer() *TaskActions {
	return &e
}

func (e *TaskActions) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_ACTION_TYPE_UNSPECIFIED":
		fallthrough
	case "TASK_ACTION_TYPE_CLOSE":
		fallthrough
	case "TASK_ACTION_TYPE_APPROVE":
		fallthrough
	case "TASK_ACTION_TYPE_DENY":
		fallthrough
	case "TASK_ACTION_TYPE_COMMENT":
		fallthrough
	case "TASK_ACTION_TYPE_DELETE":
		fallthrough
	case "TASK_ACTION_TYPE_REASSIGN":
		fallthrough
	case "TASK_ACTION_TYPE_RESTART":
		fallthrough
	case "TASK_ACTION_TYPE_SEND_REMINDER":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_COMPLETE":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_CANCELLED":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_ERRORED":
		fallthrough
	case "TASK_ACTION_TYPE_PROVISION_APP_USER_TARGET_CREATED":
		fallthrough
	case "TASK_ACTION_TYPE_ROLLBACK_SKIPPED":
		fallthrough
	case "TASK_ACTION_TYPE_HARD_RESET":
		*e = TaskActions(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskActions: %v", v)
	}
}

// TaskAnnotations - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskAnnotations struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _TaskAnnotations TaskAnnotations

func (c *TaskAnnotations) UnmarshalJSON(bs []byte) error {
	data := _TaskAnnotations{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = TaskAnnotations(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "@type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c TaskAnnotations) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_TaskAnnotations(c))
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

// TaskProcessing - The processing field.
type TaskProcessing string

const (
	TaskProcessingTaskProcessingTypeUnspecified TaskProcessing = "TASK_PROCESSING_TYPE_UNSPECIFIED"
	TaskProcessingTaskProcessingTypeProcessing  TaskProcessing = "TASK_PROCESSING_TYPE_PROCESSING"
	TaskProcessingTaskProcessingTypeWaiting     TaskProcessing = "TASK_PROCESSING_TYPE_WAITING"
	TaskProcessingTaskProcessingTypeDone        TaskProcessing = "TASK_PROCESSING_TYPE_DONE"
)

func (e TaskProcessing) ToPointer() *TaskProcessing {
	return &e
}

func (e *TaskProcessing) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_PROCESSING_TYPE_UNSPECIFIED":
		fallthrough
	case "TASK_PROCESSING_TYPE_PROCESSING":
		fallthrough
	case "TASK_PROCESSING_TYPE_WAITING":
		fallthrough
	case "TASK_PROCESSING_TYPE_DONE":
		*e = TaskProcessing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskProcessing: %v", v)
	}
}

// TaskState -  State
type TaskState string

const (
	TaskStateTaskStateUnspecified TaskState = "TASK_STATE_UNSPECIFIED"
	TaskStateTaskStateOpen        TaskState = "TASK_STATE_OPEN"
	TaskStateTaskStateClosed      TaskState = "TASK_STATE_CLOSED"
)

func (e TaskState) ToPointer() *TaskState {
	return &e
}

func (e *TaskState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_STATE_UNSPECIFIED":
		fallthrough
	case "TASK_STATE_OPEN":
		fallthrough
	case "TASK_STATE_CLOSED":
		*e = TaskState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskState: %v", v)
	}
}

// Task - The Task message.
type Task struct {
	// The actions field.
	Actions []TaskActions `json:"actions,omitempty"`
	// The analysisId field.
	AnalysisID *string `json:"analysisId,omitempty"`
	// The annotations field.
	Annotations []TaskAnnotations `json:"annotations,omitempty"`
	// The commentCount field.
	CommentCount *float64   `json:"commentCount,omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	// The createdByUserId field.
	CreatedByUserID *string    `json:"createdByUserId,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The externalRefs field.
	ExternalRefs []ExternalRef `json:"externalRefs,omitempty"`
	//  General Metadata
	//
	ID *string `json:"id,omitempty"`
	// The numericId field.
	NumericID *string `json:"numericId,omitempty"`
	// The PolicyInstance message.
	Policy *PolicyInstance `json:"policy,omitempty"`
	// The processing field.
	Processing *TaskProcessing `json:"processing,omitempty"`
	//  State
	//
	State *TaskState `json:"state,omitempty"`
	// The stepApproverIds field.
	StepApproverIds []string `json:"stepApproverIds,omitempty"`
	// The TaskType message.
	//
	// This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
	//   - grant
	//   - revoke
	//   - certify
	//
	Type      *TaskType  `json:"type,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	//  External IDS
	//
	UserID *string `json:"userId,omitempty"`
}
