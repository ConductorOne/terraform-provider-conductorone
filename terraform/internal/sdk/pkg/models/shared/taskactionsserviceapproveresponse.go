// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"conductorone/internal/sdk/pkg/utils"
)

// TaskActionsServiceApproveResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskActionsServiceApproveResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string     `json:"@type,omitempty"`
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
}

func (t TaskActionsServiceApproveResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskActionsServiceApproveResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskActionsServiceApproveResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *TaskActionsServiceApproveResponseExpanded) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The TaskActionsServiceApproveResponse returns a task view with paths indicating the location of expanded items in the array.
type TaskActionsServiceApproveResponse struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// List of serialized related objects.
	Expanded []TaskActionsServiceApproveResponseExpanded `json:"expanded,omitempty"`
	// The ID of the ticket (task) approve action created by this request.
	TicketActionID *string `json:"ticketActionId,omitempty"`
}

func (o *TaskActionsServiceApproveResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskActionsServiceApproveResponse) GetExpanded() []TaskActionsServiceApproveResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *TaskActionsServiceApproveResponse) GetTicketActionID() *string {
	if o == nil {
		return nil
	}
	return o.TicketActionID
}
