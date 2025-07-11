// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceTypeServiceGetRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
	ID    string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppResourceTypeServiceGetRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceTypeServiceGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppResourceTypeServiceGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The AppResourceTypeServiceGetResponse contains an expanded array containing the expanded values indicated by the expand mask
	//  in the request and an app resource type view containing the resource type and JSONPATHs indicating which objects are where in the expand mask.
	AppResourceTypeServiceGetResponse *shared.AppResourceTypeServiceGetResponse
}

func (o *C1APIAppV1AppResourceTypeServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceTypeServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceTypeServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppResourceTypeServiceGetResponse) GetAppResourceTypeServiceGetResponse() *shared.AppResourceTypeServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeServiceGetResponse
}
