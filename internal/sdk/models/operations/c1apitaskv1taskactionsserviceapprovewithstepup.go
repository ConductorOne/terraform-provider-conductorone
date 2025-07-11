// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceApproveWithStepUpRequest struct {
	TaskID                                     string                                             `pathParam:"style=simple,explode=false,name=task_id"`
	TaskActionsServiceApproveWithStepUpRequest *shared.TaskActionsServiceApproveWithStepUpRequest `request:"mediaType=application/json"`
}

func (o *C1APITaskV1TaskActionsServiceApproveWithStepUpRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *C1APITaskV1TaskActionsServiceApproveWithStepUpRequest) GetTaskActionsServiceApproveWithStepUpRequest() *shared.TaskActionsServiceApproveWithStepUpRequest {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceApproveWithStepUpRequest
}

type C1APITaskV1TaskActionsServiceApproveWithStepUpResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// TaskActionsServiceApproveWithStepUpResponse is the response for approving a task with step-up authentication
	TaskActionsServiceApproveWithStepUpResponse *shared.TaskActionsServiceApproveWithStepUpResponse
}

func (o *C1APITaskV1TaskActionsServiceApproveWithStepUpResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskActionsServiceApproveWithStepUpResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskActionsServiceApproveWithStepUpResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskActionsServiceApproveWithStepUpResponse) GetTaskActionsServiceApproveWithStepUpResponse() *shared.TaskActionsServiceApproveWithStepUpResponse {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceApproveWithStepUpResponse
}
