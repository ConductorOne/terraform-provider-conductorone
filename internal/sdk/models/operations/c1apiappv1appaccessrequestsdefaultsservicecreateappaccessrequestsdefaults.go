// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest struct {
	AppID                              string                                     `pathParam:"style=simple,explode=false,name=app_id"`
	CreateAccessRequestDefaultsRequest *shared.CreateAccessRequestDefaultsRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest) GetCreateAccessRequestDefaultsRequest() *shared.CreateAccessRequestDefaultsRequest {
	if o == nil {
		return nil
	}
	return o.CreateAccessRequestDefaultsRequest
}

type C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	AppAccessRequestDefaults *shared.AppAccessRequestDefaults
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse) GetAppAccessRequestDefaults() *shared.AppAccessRequestDefaults {
	if o == nil {
		return nil
	}
	return o.AppAccessRequestDefaults
}
