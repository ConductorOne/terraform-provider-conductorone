// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/utils"
)

// TaskActionsServiceDenyResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskActionsServiceDenyResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string     `json:"@type,omitempty"`
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
}

func (t TaskActionsServiceDenyResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskActionsServiceDenyResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskActionsServiceDenyResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *TaskActionsServiceDenyResponseExpanded) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The TaskActionsServiceDenyResponse returns a task view with paths indicating the location of expanded items in the array.
type TaskActionsServiceDenyResponse struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// List of serialized related objects.
	Expanded []TaskActionsServiceDenyResponseExpanded `json:"expanded,omitempty"`
	// The ID of the ticket (task) deny action created by this request.
	TicketActionID *string `json:"ticketActionId,omitempty"`
}

func (o *TaskActionsServiceDenyResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskActionsServiceDenyResponse) GetExpanded() []TaskActionsServiceDenyResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *TaskActionsServiceDenyResponse) GetTicketActionID() *string {
	if o == nil {
		return nil
	}
	return o.TicketActionID
}
