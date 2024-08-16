// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// ExporterState - The state field.
type ExporterState string

const (
	ExporterStateExportStateUnspecified ExporterState = "EXPORT_STATE_UNSPECIFIED"
	ExporterStateExportStateExporting   ExporterState = "EXPORT_STATE_EXPORTING"
	ExporterStateExportStateWaiting     ExporterState = "EXPORT_STATE_WAITING"
	ExporterStateExportStateError       ExporterState = "EXPORT_STATE_ERROR"
)

func (e ExporterState) ToPointer() *ExporterState {
	return &e
}

func (e *ExporterState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EXPORT_STATE_UNSPECIFIED":
		fallthrough
	case "EXPORT_STATE_EXPORTING":
		fallthrough
	case "EXPORT_STATE_WAITING":
		fallthrough
	case "EXPORT_STATE_ERROR":
		*e = ExporterState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExporterState: %v", v)
	}
}

// Exporter - The Exporter message.
//
// This message contains a oneof named export_to. Only a single field of the following list may be set at a time:
//   - datasource
type Exporter struct {
	// The ExportToDatasource message.
	ExportToDatasource *ExportToDatasource `json:"datasource,omitempty"`
	CreatedAt          *time.Time          `json:"createdAt,omitempty"`
	DeletedAt          *time.Time          `json:"deletedAt,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The exportId field.
	ExportID *string `json:"exportId,omitempty"`
	// The state field.
	State     *ExporterState `json:"state,omitempty"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty"`
	// we've synchorized this far
	WatermarkEventID *string `json:"watermarkEventId,omitempty"`
}

func (o *Exporter) GetExportToDatasource() *ExportToDatasource {
	if o == nil {
		return nil
	}
	return o.ExportToDatasource
}

func (o *Exporter) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Exporter) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Exporter) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Exporter) GetExportID() *string {
	if o == nil {
		return nil
	}
	return o.ExportID
}

func (o *Exporter) GetState() *ExporterState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Exporter) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Exporter) GetWatermarkEventID() *string {
	if o == nil {
		return nil
	}
	return o.WatermarkEventID
}

// ExporterInput - The Exporter message.
//
// This message contains a oneof named export_to. Only a single field of the following list may be set at a time:
//   - datasource
type ExporterInput struct {
	// The ExportToDatasource message.
	ExportToDatasource *ExportToDatasource `json:"datasource,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
}

func (o *ExporterInput) GetExportToDatasource() *ExportToDatasource {
	if o == nil {
		return nil
	}
	return o.ExportToDatasource
}

func (o *ExporterInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}
