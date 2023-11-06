// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"conductorone/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

// CurrentStep - Search tasks that have this type of step as the current step.
type CurrentStep string

const (
	CurrentStepTaskSearchCurrentStepUnspecified CurrentStep = "TASK_SEARCH_CURRENT_STEP_UNSPECIFIED"
	CurrentStepTaskSearchCurrentStepApproval    CurrentStep = "TASK_SEARCH_CURRENT_STEP_APPROVAL"
	CurrentStepTaskSearchCurrentStepProvision   CurrentStep = "TASK_SEARCH_CURRENT_STEP_PROVISION"
)

func (e CurrentStep) ToPointer() *CurrentStep {
	return &e
}

func (e *CurrentStep) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_SEARCH_CURRENT_STEP_UNSPECIFIED":
		fallthrough
	case "TASK_SEARCH_CURRENT_STEP_APPROVAL":
		fallthrough
	case "TASK_SEARCH_CURRENT_STEP_PROVISION":
		*e = CurrentStep(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CurrentStep: %v", v)
	}
}

// EmergencyStatus - Search tasks that are or are not emergency access.
type EmergencyStatus string

const (
	EmergencyStatusUnspecified  EmergencyStatus = "UNSPECIFIED"
	EmergencyStatusAll          EmergencyStatus = "ALL"
	EmergencyStatusNonEmergency EmergencyStatus = "NON_EMERGENCY"
	EmergencyStatusEmergency    EmergencyStatus = "EMERGENCY"
)

func (e EmergencyStatus) ToPointer() *EmergencyStatus {
	return &e
}

func (e *EmergencyStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNSPECIFIED":
		fallthrough
	case "ALL":
		fallthrough
	case "NON_EMERGENCY":
		fallthrough
	case "EMERGENCY":
		*e = EmergencyStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmergencyStatus: %v", v)
	}
}

// SortBy - Sort tasks in a specific order.
type SortBy string

const (
	SortByTaskSearchSortByUnspecified     SortBy = "TASK_SEARCH_SORT_BY_UNSPECIFIED"
	SortByTaskSearchSortByAccount         SortBy = "TASK_SEARCH_SORT_BY_ACCOUNT"
	SortByTaskSearchSortByResource        SortBy = "TASK_SEARCH_SORT_BY_RESOURCE"
	SortByTaskSearchSortByAccountOwner    SortBy = "TASK_SEARCH_SORT_BY_ACCOUNT_OWNER"
	SortByTaskSearchSortByReverseTicketID SortBy = "TASK_SEARCH_SORT_BY_REVERSE_TICKET_ID"
)

func (e SortBy) ToPointer() *SortBy {
	return &e
}

func (e *SortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_SEARCH_SORT_BY_UNSPECIFIED":
		fallthrough
	case "TASK_SEARCH_SORT_BY_ACCOUNT":
		fallthrough
	case "TASK_SEARCH_SORT_BY_RESOURCE":
		fallthrough
	case "TASK_SEARCH_SORT_BY_ACCOUNT_OWNER":
		fallthrough
	case "TASK_SEARCH_SORT_BY_REVERSE_TICKET_ID":
		*e = SortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SortBy: %v", v)
	}
}

type TaskStates string

const (
	TaskStatesTaskStateUnspecified TaskStates = "TASK_STATE_UNSPECIFIED"
	TaskStatesTaskStateOpen        TaskStates = "TASK_STATE_OPEN"
	TaskStatesTaskStateClosed      TaskStates = "TASK_STATE_CLOSED"
)

func (e TaskStates) ToPointer() *TaskStates {
	return &e
}

func (e *TaskStates) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TASK_STATE_UNSPECIFIED":
		fallthrough
	case "TASK_STATE_OPEN":
		fallthrough
	case "TASK_STATE_CLOSED":
		*e = TaskStates(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TaskStates: %v", v)
	}
}

// TaskSearchRequest - Search for tasks based on a plethora filters.
type TaskSearchRequest struct {
	// The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	TaskExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
	// Search tasks that belong to any of the access reviews included in this list.
	AccessReviewIds []string `json:"accessReviewIds,omitempty"`
	// Search tasks that have any of these account owners.
	AccountOwnerIds []string `json:"accountOwnerIds,omitempty"`
	// Search tasks that have this actor ID.
	ActorID *string `json:"actorId,omitempty"`
	// Search tasks that have any of these app entitlement IDs.
	AppEntitlementIds []string `json:"appEntitlementIds,omitempty"`
	// Search tasks that have any of these app resource IDs.
	AppResourceIds []string `json:"appResourceIds,omitempty"`
	// Search tasks that have any of these app resource type IDs.
	AppResourceTypeIds []string `json:"appResourceTypeIds,omitempty"`
	// Search tasks that have any of these app users as subjects.
	AppUserSubjectIds []string `json:"appUserSubjectIds,omitempty"`
	// Search tasks that have any of these apps as targets.
	ApplicationIds []string `json:"applicationIds,omitempty"`
	// Search tasks by  List of UserIDs which are currently assigned these Tasks
	AssigneesInIds []string   `json:"assigneesInIds,omitempty"`
	CreatedAfter   *time.Time `json:"createdAfter,omitempty"`
	CreatedBefore  *time.Time `json:"createdBefore,omitempty"`
	// Search tasks that have this type of step as the current step.
	CurrentStep *CurrentStep `json:"currentStep,omitempty"`
	// Search tasks that are or are not emergency access.
	EmergencyStatus *EmergencyStatus `json:"emergencyStatus,omitempty"`
	// Search tasks that do not have any of these app entitlement IDs.
	ExcludeAppEntitlementIds []string `json:"excludeAppEntitlementIds,omitempty"`
	// Exclude Specific TaskIDs from this serach result.
	ExcludeIds []string `json:"excludeIds,omitempty"`
	// Whether or not to include deleted tasks.
	IncludeDeleted *bool `json:"includeDeleted,omitempty"`
	// Search tasks where the user would see this task in the My Work section
	MyWorkUserIds []string `json:"myWorkUserIds,omitempty"`
	// Search tasks that were created by any of the users in this array.
	OpenerIds []string `json:"openerIds,omitempty"`
	// The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	PageSize *float64 `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// Search tasks that were acted on by any of these users.
	PreviouslyActedOnIds []string `json:"previouslyActedOnIds,omitempty"`
	// Fuzzy search tasks by display name or description. Also can search by numeric ID.
	Query *string `json:"query,omitempty"`
	// Query tasks by display name, description, or numeric ID.
	Refs []TaskRef `json:"refs,omitempty"`
	// Sort tasks in a specific order.
	SortBy *SortBy `json:"sortBy,omitempty"`
	// Search tasks where these users are the subject.
	SubjectIds []string `json:"subjectIds,omitempty"`
	// Search tasks with this task state.
	TaskStates []TaskStates `json:"taskStates,omitempty"`
	// Search tasks with this task type. This is a oneOf, and needs an object, which can be empty, to sort.
	TaskTypes []TaskTypeInput `json:"taskTypes,omitempty"`
}

func (t TaskSearchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskSearchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskSearchRequest) GetTaskExpandMask() *TaskExpandMask {
	if o == nil {
		return nil
	}
	return o.TaskExpandMask
}

func (o *TaskSearchRequest) GetAccessReviewIds() []string {
	if o == nil {
		return nil
	}
	return o.AccessReviewIds
}

func (o *TaskSearchRequest) GetAccountOwnerIds() []string {
	if o == nil {
		return nil
	}
	return o.AccountOwnerIds
}

func (o *TaskSearchRequest) GetActorID() *string {
	if o == nil {
		return nil
	}
	return o.ActorID
}

func (o *TaskSearchRequest) GetAppEntitlementIds() []string {
	if o == nil {
		return nil
	}
	return o.AppEntitlementIds
}

func (o *TaskSearchRequest) GetAppResourceIds() []string {
	if o == nil {
		return nil
	}
	return o.AppResourceIds
}

func (o *TaskSearchRequest) GetAppResourceTypeIds() []string {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeIds
}

func (o *TaskSearchRequest) GetAppUserSubjectIds() []string {
	if o == nil {
		return nil
	}
	return o.AppUserSubjectIds
}

func (o *TaskSearchRequest) GetApplicationIds() []string {
	if o == nil {
		return nil
	}
	return o.ApplicationIds
}

func (o *TaskSearchRequest) GetAssigneesInIds() []string {
	if o == nil {
		return nil
	}
	return o.AssigneesInIds
}

func (o *TaskSearchRequest) GetCreatedAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAfter
}

func (o *TaskSearchRequest) GetCreatedBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedBefore
}

func (o *TaskSearchRequest) GetCurrentStep() *CurrentStep {
	if o == nil {
		return nil
	}
	return o.CurrentStep
}

func (o *TaskSearchRequest) GetEmergencyStatus() *EmergencyStatus {
	if o == nil {
		return nil
	}
	return o.EmergencyStatus
}

func (o *TaskSearchRequest) GetExcludeAppEntitlementIds() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeAppEntitlementIds
}

func (o *TaskSearchRequest) GetExcludeIds() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeIds
}

func (o *TaskSearchRequest) GetIncludeDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDeleted
}

func (o *TaskSearchRequest) GetMyWorkUserIds() []string {
	if o == nil {
		return nil
	}
	return o.MyWorkUserIds
}

func (o *TaskSearchRequest) GetOpenerIds() []string {
	if o == nil {
		return nil
	}
	return o.OpenerIds
}

func (o *TaskSearchRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *TaskSearchRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *TaskSearchRequest) GetPreviouslyActedOnIds() []string {
	if o == nil {
		return nil
	}
	return o.PreviouslyActedOnIds
}

func (o *TaskSearchRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *TaskSearchRequest) GetRefs() []TaskRef {
	if o == nil {
		return nil
	}
	return o.Refs
}

func (o *TaskSearchRequest) GetSortBy() *SortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *TaskSearchRequest) GetSubjectIds() []string {
	if o == nil {
		return nil
	}
	return o.SubjectIds
}

func (o *TaskSearchRequest) GetTaskStates() []TaskStates {
	if o == nil {
		return nil
	}
	return o.TaskStates
}

func (o *TaskSearchRequest) GetTaskTypes() []TaskTypeInput {
	if o == nil {
		return nil
	}
	return o.TaskTypes
}
