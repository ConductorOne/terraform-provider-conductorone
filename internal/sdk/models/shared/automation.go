// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/internal/utils"
	"time"
)

// PrimaryTriggerType - The primaryTriggerType field.
type PrimaryTriggerType string

const (
	PrimaryTriggerTypeTriggerTypeUnspecified       PrimaryTriggerType = "TRIGGER_TYPE_UNSPECIFIED"
	PrimaryTriggerTypeTriggerTypeUserProfileChange PrimaryTriggerType = "TRIGGER_TYPE_USER_PROFILE_CHANGE"
	PrimaryTriggerTypeTriggerTypeAppUserCreate     PrimaryTriggerType = "TRIGGER_TYPE_APP_USER_CREATE"
	PrimaryTriggerTypeTriggerTypeAppUserUpdate     PrimaryTriggerType = "TRIGGER_TYPE_APP_USER_UPDATE"
	PrimaryTriggerTypeTriggerTypeUnusedAccess      PrimaryTriggerType = "TRIGGER_TYPE_UNUSED_ACCESS"
	PrimaryTriggerTypeTriggerTypeUserCreated       PrimaryTriggerType = "TRIGGER_TYPE_USER_CREATED"
)

func (e PrimaryTriggerType) ToPointer() *PrimaryTriggerType {
	return &e
}
func (e *PrimaryTriggerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRIGGER_TYPE_UNSPECIFIED":
		fallthrough
	case "TRIGGER_TYPE_USER_PROFILE_CHANGE":
		fallthrough
	case "TRIGGER_TYPE_APP_USER_CREATE":
		fallthrough
	case "TRIGGER_TYPE_APP_USER_UPDATE":
		fallthrough
	case "TRIGGER_TYPE_UNUSED_ACCESS":
		fallthrough
	case "TRIGGER_TYPE_USER_CREATED":
		*e = PrimaryTriggerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PrimaryTriggerType: %v", v)
	}
}

// The Automation message.
type Automation struct {
	// the app id this workflow_template belongs to
	AppID *string `json:"appId,omitempty"`
	// The automationSteps field.
	AutomationSteps []AutomationStep `json:"automationSteps,omitempty"`
	// The AutomationContext message.
	AutomationContext *AutomationContext `json:"context,omitempty"`
	CreatedAt         *time.Time         `json:"createdAt,omitempty"`
	// The currentVersion field.
	CurrentVersion *string `json:"currentVersion,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The draftAutomationSteps field.
	DraftAutomationSteps []AutomationStep `json:"draftAutomationSteps,omitempty"`
	// The draftTriggers field.
	DraftTriggers []AutomationTrigger `json:"draftTriggers,omitempty"`
	// The enabled field.
	Enabled *bool `json:"enabled,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The isDraft field.
	IsDraft        *bool      `json:"isDraft,omitempty"`
	LastExecutedAt *time.Time `json:"lastExecutedAt,omitempty"`
	// The primaryTriggerType field.
	PrimaryTriggerType *PrimaryTriggerType `json:"primaryTriggerType,omitempty"`
	// The triggers field.
	Triggers []AutomationTrigger `json:"triggers,omitempty"`
}

func (a Automation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Automation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Automation) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *Automation) GetAutomationSteps() []AutomationStep {
	if o == nil {
		return nil
	}
	return o.AutomationSteps
}

func (o *Automation) GetAutomationContext() *AutomationContext {
	if o == nil {
		return nil
	}
	return o.AutomationContext
}

func (o *Automation) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Automation) GetCurrentVersion() *string {
	if o == nil {
		return nil
	}
	return o.CurrentVersion
}

func (o *Automation) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Automation) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Automation) GetDraftAutomationSteps() []AutomationStep {
	if o == nil {
		return nil
	}
	return o.DraftAutomationSteps
}

func (o *Automation) GetDraftTriggers() []AutomationTrigger {
	if o == nil {
		return nil
	}
	return o.DraftTriggers
}

func (o *Automation) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *Automation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Automation) GetIsDraft() *bool {
	if o == nil {
		return nil
	}
	return o.IsDraft
}

func (o *Automation) GetLastExecutedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastExecutedAt
}

func (o *Automation) GetPrimaryTriggerType() *PrimaryTriggerType {
	if o == nil {
		return nil
	}
	return o.PrimaryTriggerType
}

func (o *Automation) GetTriggers() []AutomationTrigger {
	if o == nil {
		return nil
	}
	return o.Triggers
}

// AutomationInput - The Automation message.
type AutomationInput struct {
	// the app id this workflow_template belongs to
	AppID *string `json:"appId,omitempty"`
	// The automationSteps field.
	AutomationSteps []AutomationStep `json:"automationSteps,omitempty"`
	// The AutomationContext message.
	AutomationContext *AutomationContext `json:"context,omitempty"`
	CreatedAt         *time.Time         `json:"createdAt,omitempty"`
	// The currentVersion field.
	CurrentVersion *string `json:"currentVersion,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The draftAutomationSteps field.
	DraftAutomationSteps []AutomationStep `json:"draftAutomationSteps,omitempty"`
	// The draftTriggers field.
	DraftTriggers []AutomationTrigger `json:"draftTriggers,omitempty"`
	// The enabled field.
	Enabled *bool `json:"enabled,omitempty"`
	// The isDraft field.
	IsDraft        *bool      `json:"isDraft,omitempty"`
	LastExecutedAt *time.Time `json:"lastExecutedAt,omitempty"`
	// The primaryTriggerType field.
	PrimaryTriggerType *PrimaryTriggerType `json:"primaryTriggerType,omitempty"`
	// The triggers field.
	Triggers []AutomationTrigger `json:"triggers,omitempty"`
}

func (a AutomationInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AutomationInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AutomationInput) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AutomationInput) GetAutomationSteps() []AutomationStep {
	if o == nil {
		return nil
	}
	return o.AutomationSteps
}

func (o *AutomationInput) GetAutomationContext() *AutomationContext {
	if o == nil {
		return nil
	}
	return o.AutomationContext
}

func (o *AutomationInput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AutomationInput) GetCurrentVersion() *string {
	if o == nil {
		return nil
	}
	return o.CurrentVersion
}

func (o *AutomationInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AutomationInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AutomationInput) GetDraftAutomationSteps() []AutomationStep {
	if o == nil {
		return nil
	}
	return o.DraftAutomationSteps
}

func (o *AutomationInput) GetDraftTriggers() []AutomationTrigger {
	if o == nil {
		return nil
	}
	return o.DraftTriggers
}

func (o *AutomationInput) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AutomationInput) GetIsDraft() *bool {
	if o == nil {
		return nil
	}
	return o.IsDraft
}

func (o *AutomationInput) GetLastExecutedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastExecutedAt
}

func (o *AutomationInput) GetPrimaryTriggerType() *PrimaryTriggerType {
	if o == nil {
		return nil
	}
	return o.PrimaryTriggerType
}

func (o *AutomationInput) GetTriggers() []AutomationTrigger {
	if o == nil {
		return nil
	}
	return o.Triggers
}
