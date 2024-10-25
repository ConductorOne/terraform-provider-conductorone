// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsDeleteRequest struct {
	AppID                       string                              `pathParam:"style=simple,explode=false,name=app_id"`
	ID                          string                              `pathParam:"style=simple,explode=false,name=id"`
	DeleteAppEntitlementRequest *shared.DeleteAppEntitlementRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppEntitlementsDeleteRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIAppV1AppEntitlementsDeleteRequest) GetDeleteAppEntitlementRequest() *shared.DeleteAppEntitlementRequest {
	if o == nil {
		return nil
	}
	return o.DeleteAppEntitlementRequest
}

type C1APIAppV1AppEntitlementsDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	DeleteAppEntitlementResponse *shared.DeleteAppEntitlementResponse
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetDeleteAppEntitlementResponse() *shared.DeleteAppEntitlementResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAppEntitlementResponse
}