// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APITaskV1TaskActionsServiceCommentRequest struct {
	TaskActionsServiceCommentRequest *shared.TaskActionsServiceCommentRequest `request:"mediaType=application/json"`
	TaskID                           string                                   `pathParam:"style=simple,explode=false,name=task_id"`
}

func (o *C1APITaskV1TaskActionsServiceCommentRequest) GetTaskActionsServiceCommentRequest() *shared.TaskActionsServiceCommentRequest {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceCommentRequest
}

func (o *C1APITaskV1TaskActionsServiceCommentRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

type C1APITaskV1TaskActionsServiceCommentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Task actions service comment response returns the task view inluding the expanded array of items that are indicated by the expand mask on the request.
	TaskActionsServiceCommentResponse *shared.TaskActionsServiceCommentResponse
}

func (o *C1APITaskV1TaskActionsServiceCommentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskActionsServiceCommentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskActionsServiceCommentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskActionsServiceCommentResponse) GetTaskActionsServiceCommentResponse() *shared.TaskActionsServiceCommentResponse {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceCommentResponse
}
