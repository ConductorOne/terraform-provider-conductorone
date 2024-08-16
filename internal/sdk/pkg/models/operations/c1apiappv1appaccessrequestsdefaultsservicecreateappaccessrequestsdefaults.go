// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest struct {
	AppAccessRequestDefaults1 *shared.AppAccessRequestDefaults1 `request:"mediaType=application/json"`
	AppID                     string                            `pathParam:"style=simple,explode=false,name=app_id"`
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest) GetAppAccessRequestDefaults1() *shared.AppAccessRequestDefaults1 {
	if o == nil {
		return nil
	}
	return o.AppAccessRequestDefaults1
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse struct {
	// Successful response
	AppAccessRequestDefaults *shared.AppAccessRequestDefaults
	ContentType              string
	StatusCode               int
	RawResponse              *http.Response
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse) GetAppAccessRequestDefaults() *shared.AppAccessRequestDefaults {
	if o == nil {
		return nil
	}
	return o.AppAccessRequestDefaults
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
