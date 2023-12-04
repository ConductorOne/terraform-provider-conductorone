// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppOwnersSetRequest struct {
	SetAppOwnersRequest *shared.SetAppOwnersRequest `request:"mediaType=application/json"`
	AppID               string                      `pathParam:"style=simple,explode=false,name=app_id"`
}

func (o *C1APIAppV1AppOwnersSetRequest) GetSetAppOwnersRequest() *shared.SetAppOwnersRequest {
	if o == nil {
		return nil
	}
	return o.SetAppOwnersRequest
}

func (o *C1APIAppV1AppOwnersSetRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type C1APIAppV1AppOwnersSetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The empty response message for setting the app owners.
	SetAppOwnersResponse *shared.SetAppOwnersResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppOwnersSetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppOwnersSetResponse) GetSetAppOwnersResponse() *shared.SetAppOwnersResponse {
	if o == nil {
		return nil
	}
	return o.SetAppOwnersResponse
}

func (o *C1APIAppV1AppOwnersSetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppOwnersSetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
