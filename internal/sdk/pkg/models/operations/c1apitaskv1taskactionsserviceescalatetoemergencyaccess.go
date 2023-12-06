// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessRequest struct {
	TaskID                                             string                                                     `pathParam:"style=simple,explode=false,name=task_id"`
	TaskActionsServiceEscalateToEmergencyAccessRequest *shared.TaskActionsServiceEscalateToEmergencyAccessRequest `request:"mediaType=application/json"`
}

func (o *C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessRequest) GetTaskActionsServiceEscalateToEmergencyAccessRequest() *shared.TaskActionsServiceEscalateToEmergencyAccessRequest {
	if o == nil {
		return nil
	}
	return o.TaskActionsServiceEscalateToEmergencyAccessRequest
}

type C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	TaskServiceActionResponse *shared.TaskServiceActionResponse
}

func (o *C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessResponse) GetTaskServiceActionResponse() *shared.TaskServiceActionResponse {
	if o == nil {
		return nil
	}
	return o.TaskServiceActionResponse
}
