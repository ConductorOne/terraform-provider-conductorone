// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceRequest struct {
	AppID                                   string                                          `pathParam:"style=simple,explode=false,name=app_id"`
	AppResourceTypeID                       string                                          `pathParam:"style=simple,explode=false,name=app_resource_type_id"`
	ID                                      string                                          `pathParam:"style=simple,explode=false,name=id"`
	DeleteManuallyManagedAppResourceRequest *shared.DeleteManuallyManagedAppResourceRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceRequest) GetAppResourceTypeID() string {
	if o == nil {
		return ""
	}
	return o.AppResourceTypeID
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceRequest) GetDeleteManuallyManagedAppResourceRequest() *shared.DeleteManuallyManagedAppResourceRequest {
	if o == nil {
		return nil
	}
	return o.DeleteManuallyManagedAppResourceRequest
}

type C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	DeleteManuallyManagedAppResourceResponse *shared.DeleteManuallyManagedAppResourceResponse
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppResourceServiceDeleteManuallyManagedAppResourceResponse) GetDeleteManuallyManagedAppResourceResponse() *shared.DeleteManuallyManagedAppResourceResponse {
	if o == nil {
		return nil
	}
	return o.DeleteManuallyManagedAppResourceResponse
}
