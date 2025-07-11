// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AgentMode - The mode of the agent, full control, change policy only, or comment only.
type AgentMode string

const (
	AgentModeApprovalAgentModeUnspecified      AgentMode = "APPROVAL_AGENT_MODE_UNSPECIFIED"
	AgentModeApprovalAgentModeFullControl      AgentMode = "APPROVAL_AGENT_MODE_FULL_CONTROL"
	AgentModeApprovalAgentModeChangePolicyOnly AgentMode = "APPROVAL_AGENT_MODE_CHANGE_POLICY_ONLY"
	AgentModeApprovalAgentModeCommentOnly      AgentMode = "APPROVAL_AGENT_MODE_COMMENT_ONLY"
)

func (e AgentMode) ToPointer() *AgentMode {
	return &e
}
func (e *AgentMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "APPROVAL_AGENT_MODE_UNSPECIFIED":
		fallthrough
	case "APPROVAL_AGENT_MODE_FULL_CONTROL":
		fallthrough
	case "APPROVAL_AGENT_MODE_CHANGE_POLICY_ONLY":
		fallthrough
	case "APPROVAL_AGENT_MODE_COMMENT_ONLY":
		*e = AgentMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AgentMode: %v", v)
	}
}

// AgentApproval - The agent to assign the task to.
type AgentApproval struct {
	// The mode of the agent, full control, change policy only, or comment only.
	AgentMode *AgentMode `json:"agentMode,omitempty"`
	// The agent user ID to assign the task to.
	AgentUserID *string `json:"agentUserId,omitempty"`
	// Instructions for the agent.
	Instructions *string `json:"instructions,omitempty"`
	// The allow list of policy IDs to re-route the task to.
	PolicyIds []string `json:"policyIds,omitempty"`
}

func (o *AgentApproval) GetAgentMode() *AgentMode {
	if o == nil {
		return nil
	}
	return o.AgentMode
}

func (o *AgentApproval) GetAgentUserID() *string {
	if o == nil {
		return nil
	}
	return o.AgentUserID
}

func (o *AgentApproval) GetInstructions() *string {
	if o == nil {
		return nil
	}
	return o.Instructions
}

func (o *AgentApproval) GetPolicyIds() []string {
	if o == nil {
		return nil
	}
	return o.PolicyIds
}
