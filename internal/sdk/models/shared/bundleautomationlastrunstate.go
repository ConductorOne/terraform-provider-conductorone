// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/internal/utils"
	"time"
)

// BundleAutomationLastRunStateStatus - The status field.
type BundleAutomationLastRunStateStatus string

const (
	BundleAutomationLastRunStateStatusBundleAutomationRunStatusUnspecified BundleAutomationLastRunStateStatus = "BUNDLE_AUTOMATION_RUN_STATUS_UNSPECIFIED"
	BundleAutomationLastRunStateStatusBundleAutomationRunStatusSuccess     BundleAutomationLastRunStateStatus = "BUNDLE_AUTOMATION_RUN_STATUS_SUCCESS"
	BundleAutomationLastRunStateStatusBundleAutomationRunStatusFailure     BundleAutomationLastRunStateStatus = "BUNDLE_AUTOMATION_RUN_STATUS_FAILURE"
	BundleAutomationLastRunStateStatusBundleAutomationRunStatusInProgress  BundleAutomationLastRunStateStatus = "BUNDLE_AUTOMATION_RUN_STATUS_IN_PROGRESS"
)

func (e BundleAutomationLastRunStateStatus) ToPointer() *BundleAutomationLastRunStateStatus {
	return &e
}
func (e *BundleAutomationLastRunStateStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUNDLE_AUTOMATION_RUN_STATUS_UNSPECIFIED":
		fallthrough
	case "BUNDLE_AUTOMATION_RUN_STATUS_SUCCESS":
		fallthrough
	case "BUNDLE_AUTOMATION_RUN_STATUS_FAILURE":
		fallthrough
	case "BUNDLE_AUTOMATION_RUN_STATUS_IN_PROGRESS":
		*e = BundleAutomationLastRunStateStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BundleAutomationLastRunStateStatus: %v", v)
	}
}

// The BundleAutomationLastRunState message.
type BundleAutomationLastRunState struct {
	// The errorMessage field.
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	LastRunAt    *time.Time `json:"lastRunAt,omitempty"`
	// The status field.
	Status *BundleAutomationLastRunStateStatus `json:"status,omitempty"`
}

func (b BundleAutomationLastRunState) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BundleAutomationLastRunState) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BundleAutomationLastRunState) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *BundleAutomationLastRunState) GetLastRunAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastRunAt
}

func (o *BundleAutomationLastRunState) GetStatus() *BundleAutomationLastRunStateStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
