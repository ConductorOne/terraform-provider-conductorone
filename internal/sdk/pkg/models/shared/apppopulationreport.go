// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

// State - The state field tracks the state of the AppPopulationReport. This state field can be one of REPORT_STATE_PENDING, REPORT_STATE_UNSPECIFIED, REPORT_STATE_OK, REPORT_STATE_ERROR.
type State string

const (
	StateReportStateUnspecified State = "REPORT_STATE_UNSPECIFIED"
	StateReportStatePending     State = "REPORT_STATE_PENDING"
	StateReportStateOk          State = "REPORT_STATE_OK"
	StateReportStateError       State = "REPORT_STATE_ERROR"
)

func (e State) ToPointer() *State {
	return &e
}

func (e *State) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REPORT_STATE_UNSPECIFIED":
		fallthrough
	case "REPORT_STATE_PENDING":
		fallthrough
	case "REPORT_STATE_OK":
		fallthrough
	case "REPORT_STATE_ERROR":
		*e = State(v)
		return nil
	default:
		return fmt.Errorf("invalid value for State: %v", v)
	}
}

// The AppPopulationReport is a generated report for a specific app that gives details about the app's users. These details include what groups, roles, and other entitlements the users have access to.
type AppPopulationReport struct {
	// The appId is the Id of the app which the report is generated for.
	AppID     *string    `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The downloadUrl is the url used for downloading the AppPopulationReport.
	DownloadURL *string `json:"downloadUrl,omitempty"`
	// The hashes field contains the file hashes of the report.
	Hashes map[string]string `json:"hashes,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The state field tracks the state of the AppPopulationReport. This state field can be one of REPORT_STATE_PENDING, REPORT_STATE_UNSPECIFIED, REPORT_STATE_OK, REPORT_STATE_ERROR.
	State *State `json:"state,omitempty"`
}

func (a AppPopulationReport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppPopulationReport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppPopulationReport) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppPopulationReport) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppPopulationReport) GetDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.DownloadURL
}

func (o *AppPopulationReport) GetHashes() map[string]string {
	if o == nil {
		return nil
	}
	return o.Hashes
}

func (o *AppPopulationReport) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppPopulationReport) GetState() *State {
	if o == nil {
		return nil
	}
	return o.State
}
