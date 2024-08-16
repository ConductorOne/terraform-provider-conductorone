// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// WaitInstanceState - The state field.
type WaitInstanceState string

const (
	WaitInstanceStateWaitInstanceStateUnspecified WaitInstanceState = "WAIT_INSTANCE_STATE_UNSPECIFIED"
	WaitInstanceStateWaitInstanceStateWaiting     WaitInstanceState = "WAIT_INSTANCE_STATE_WAITING"
	WaitInstanceStateWaitInstanceStateCompleted   WaitInstanceState = "WAIT_INSTANCE_STATE_COMPLETED"
	WaitInstanceStateWaitInstanceStateTimedOut    WaitInstanceState = "WAIT_INSTANCE_STATE_TIMED_OUT"
)

func (e WaitInstanceState) ToPointer() *WaitInstanceState {
	return &e
}

func (e *WaitInstanceState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WAIT_INSTANCE_STATE_UNSPECIFIED":
		fallthrough
	case "WAIT_INSTANCE_STATE_WAITING":
		fallthrough
	case "WAIT_INSTANCE_STATE_COMPLETED":
		fallthrough
	case "WAIT_INSTANCE_STATE_TIMED_OUT":
		*e = WaitInstanceState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WaitInstanceState: %v", v)
	}
}

// WaitInstance - Used by the policy engine to describe an instantiated wait step.
//
// This message contains a oneof named until. Only a single field of the following list may be set at a time:
//   - condition
//
// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
//   - succeeded
//   - timedOut
type WaitInstance struct {
	// The ConditionSucceeded message.
	ConditionSucceeded *ConditionSucceeded `json:"succeeded,omitempty"`
	// The ConditionTimedOut message.
	ConditionTimedOut *ConditionTimedOut `json:"timedOut,omitempty"`
	// Used by the policy engine to describe an instantiated condition to wait on.
	WaitConditionInstance *WaitConditionInstance `json:"condition,omitempty"`
	// The comment to post on first failed check.
	CommentOnFirstWait *string `json:"commentOnFirstWait,omitempty"`
	// The comment to post if we timeout.
	CommentOnTimeout *string `json:"commentOnTimeout,omitempty"`
	// The name field.
	Name             *string    `json:"name,omitempty"`
	StartedWaitingAt *time.Time `json:"startedWaitingAt,omitempty"`
	// The state field.
	State           *WaitInstanceState `json:"state,omitempty"`
	Timeout         *time.Time         `json:"timeout,omitempty"`
	TimeoutDuration *string            `json:"timeoutDuration,omitempty"`
}

func (o *WaitInstance) GetConditionSucceeded() *ConditionSucceeded {
	if o == nil {
		return nil
	}
	return o.ConditionSucceeded
}

func (o *WaitInstance) GetConditionTimedOut() *ConditionTimedOut {
	if o == nil {
		return nil
	}
	return o.ConditionTimedOut
}

func (o *WaitInstance) GetWaitConditionInstance() *WaitConditionInstance {
	if o == nil {
		return nil
	}
	return o.WaitConditionInstance
}

func (o *WaitInstance) GetCommentOnFirstWait() *string {
	if o == nil {
		return nil
	}
	return o.CommentOnFirstWait
}

func (o *WaitInstance) GetCommentOnTimeout() *string {
	if o == nil {
		return nil
	}
	return o.CommentOnTimeout
}

func (o *WaitInstance) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WaitInstance) GetStartedWaitingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedWaitingAt
}

func (o *WaitInstance) GetState() *WaitInstanceState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WaitInstance) GetTimeout() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *WaitInstance) GetTimeoutDuration() *string {
	if o == nil {
		return nil
	}
	return o.TimeoutDuration
}
