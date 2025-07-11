// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIAutomationsV1AutomationServiceDeleteAutomationRequest struct {
	ID                      string                          `pathParam:"style=simple,explode=false,name=id"`
	DeleteAutomationRequest *shared.DeleteAutomationRequest `request:"mediaType=application/json"`
}

func (o *C1APIAutomationsV1AutomationServiceDeleteAutomationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIAutomationsV1AutomationServiceDeleteAutomationRequest) GetDeleteAutomationRequest() *shared.DeleteAutomationRequest {
	if o == nil {
		return nil
	}
	return o.DeleteAutomationRequest
}

type C1APIAutomationsV1AutomationServiceDeleteAutomationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	DeleteAutomationResponse *shared.DeleteAutomationResponse
}

func (o *C1APIAutomationsV1AutomationServiceDeleteAutomationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAutomationsV1AutomationServiceDeleteAutomationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAutomationsV1AutomationServiceDeleteAutomationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAutomationsV1AutomationServiceDeleteAutomationResponse) GetDeleteAutomationResponse() *shared.DeleteAutomationResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAutomationResponse
}
