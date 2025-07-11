// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceSkipStepRequest struct {
	TaskID                            string                                    `pathParam:"style=simple,explode=false,name=task_id"`
	TaskActionsServiceSkipStepRequest *shared.TaskActionsServiceSkipStepRequest `request:"mediaType=application/json"`
}

func (o *C1APITaskV1TaskActionsServiceSkipStepRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *C1APITaskV1TaskActionsServiceSkipStepRequest) GetTaskActionsServiceSkipStepRequest() *shared.TaskActionsServiceSkipStepRequest {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceSkipStepRequest
}

type C1APITaskV1TaskActionsServiceSkipStepResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	TaskServiceActionResponse *shared.TaskServiceActionResponse
}

func (o *C1APITaskV1TaskActionsServiceSkipStepResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskActionsServiceSkipStepResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskActionsServiceSkipStepResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskActionsServiceSkipStepResponse) GetTaskServiceActionResponse() *shared.TaskServiceActionResponse {
	if o == nil {
		return nil
	}
	return o.TaskServiceActionResponse
}
