// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// TaskTypeOffboardingOutcome - The outcome field.
type TaskTypeOffboardingOutcome string

const (
	TaskTypeOffboardingOutcomeOffboardingOutcomeUnspecified TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_UNSPECIFIED"
	TaskTypeOffboardingOutcomeOffboardingOutcomeInProgress  TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_IN_PROGRESS"
	TaskTypeOffboardingOutcomeOffboardingOutcomeDone        TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_DONE"
	TaskTypeOffboardingOutcomeOffboardingOutcomeError       TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_ERROR"
	TaskTypeOffboardingOutcomeOffboardingOutcomeCancelled   TaskTypeOffboardingOutcome = "OFFBOARDING_OUTCOME_CANCELLED"
)

func (e TaskTypeOffboardingOutcome) ToPointer() *TaskTypeOffboardingOutcome {
	return &e
}

func (e *TaskTypeOffboardingOutcome) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OFFBOARDING_OUTCOME_UNSPECIFIED":
		fallthrough
	case "OFFBOARDING_OUTCOME_IN_PROGRESS":
		fallthrough
	case "OFFBOARDING_OUTCOME_DONE":
		fallthrough
	case "OFFBOARDING_OUTCOME_ERROR":
		fallthrough
	case "OFFBOARDING_OUTCOME_CANCELLED":
		*e = TaskTypeOffboardingOutcome(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskTypeOffboardingOutcome: %v", v)
	}
}

// TaskTypeOffboarding1 - The TaskTypeOffboarding message.
type TaskTypeOffboarding1 struct {
	// The outcome field.
	Outcome     *TaskTypeOffboardingOutcome `json:"outcome,omitempty"`
	OutcomeTime *time.Time                  `json:"outcomeTime,omitempty"`
	// The subjectUserId field.
	SubjectUserID *string `json:"subjectUserId,omitempty"`
}

func (o *TaskTypeOffboarding1) GetOutcome() *TaskTypeOffboardingOutcome {
	if o == nil {
		return nil
	}
	return o.Outcome
}

func (o *TaskTypeOffboarding1) GetOutcomeTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.OutcomeTime
}

func (o *TaskTypeOffboarding1) GetSubjectUserID() *string {
	if o == nil {
		return nil
	}
	return o.SubjectUserID
}
