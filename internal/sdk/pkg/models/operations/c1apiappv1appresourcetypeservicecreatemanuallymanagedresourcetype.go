// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeRequest struct {
	AppID                                    string                                           `pathParam:"style=simple,explode=false,name=app_id"`
	CreateManuallyManagedResourceTypeRequest *shared.CreateManuallyManagedResourceTypeRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeRequest) GetCreateManuallyManagedResourceTypeRequest() *shared.CreateManuallyManagedResourceTypeRequest {
	if o == nil {
		return nil
	}
	return o.CreateManuallyManagedResourceTypeRequest
}

type C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	CreateManuallyManagedResourceTypeResponse *shared.CreateManuallyManagedResourceTypeResponse
}

func (o *C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeResponse) GetCreateManuallyManagedResourceTypeResponse() *shared.CreateManuallyManagedResourceTypeResponse {
	if o == nil {
		return nil
	}
	return o.CreateManuallyManagedResourceTypeResponse
}
