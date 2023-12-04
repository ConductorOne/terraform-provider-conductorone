// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceDenyRequest struct {
	TaskActionsServiceDenyRequest *shared.TaskActionsServiceDenyRequest `request:"mediaType=application/json"`
	TaskID                        string                                `pathParam:"style=simple,explode=false,name=task_id"`
}

func (o *C1APITaskV1TaskActionsServiceDenyRequest) GetTaskActionsServiceDenyRequest() *shared.TaskActionsServiceDenyRequest {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceDenyRequest
}

func (o *C1APITaskV1TaskActionsServiceDenyRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

type C1APITaskV1TaskActionsServiceDenyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The TaskActionsServiceDenyResponse returns a task view with paths indicating the location of expanded items in the array.
	TaskActionsServiceDenyResponse *shared.TaskActionsServiceDenyResponse
}

func (o *C1APITaskV1TaskActionsServiceDenyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskActionsServiceDenyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskActionsServiceDenyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskActionsServiceDenyResponse) GetTaskActionsServiceDenyResponse() *shared.TaskActionsServiceDenyResponse {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceDenyResponse
}
