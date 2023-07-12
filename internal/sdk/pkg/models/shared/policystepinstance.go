// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PolicyStepInstanceState - The state field.
type PolicyStepInstanceState string

const (
	PolicyStepInstanceStatePolicyStepStateUnspecified PolicyStepInstanceState = "POLICY_STEP_STATE_UNSPECIFIED"
	PolicyStepInstanceStatePolicyStepStateActive      PolicyStepInstanceState = "POLICY_STEP_STATE_ACTIVE"
	PolicyStepInstanceStatePolicyStepStateDone        PolicyStepInstanceState = "POLICY_STEP_STATE_DONE"
)

func (e PolicyStepInstanceState) ToPointer() *PolicyStepInstanceState {
	return &e
}

func (e *PolicyStepInstanceState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POLICY_STEP_STATE_UNSPECIFIED":
		fallthrough
	case "POLICY_STEP_STATE_ACTIVE":
		fallthrough
	case "POLICY_STEP_STATE_DONE":
		*e = PolicyStepInstanceState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PolicyStepInstanceState: %v", v)
	}
}

// PolicyStepInstance - The PolicyStepInstance message.
//
// This message contains a oneof named instance. Only a single field of the following list may be set at a time:
//   - approval
//   - provision
type PolicyStepInstance struct {
	// The ApprovalInstance message.
	//
	// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
	//   - approved
	//   - denied
	//   - reassigned
	//   - restarted
	//   - reassignedByError
	//
	ApprovalInstance *ApprovalInstance `json:"approval,omitempty"`
	// The ProvisionInstance message.
	//
	// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
	//   - completed
	//   - cancelled
	//   - errored
	//   - reassignedByError
	//
	ProvisionInstance *ProvisionInstance `json:"provision,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The state field.
	State *PolicyStepInstanceState `json:"state,omitempty"`
}