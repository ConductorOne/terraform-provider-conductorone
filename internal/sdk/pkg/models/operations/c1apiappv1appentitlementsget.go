// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APIAppV1AppEntitlementsGetRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
	ID    string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppEntitlementsGetRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppEntitlementsGetResponse struct {
	ContentType string
	// The get app entitlement response returns an entitlement view containing paths in the expanded array for the objects expanded as indicated by the expand mask in the request.
	GetAppEntitlementResponse *shared.GetAppEntitlementResponse
	StatusCode                int
	RawResponse               *http.Response
}

func (o *C1APIAppV1AppEntitlementsGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsGetResponse) GetGetAppEntitlementResponse() *shared.GetAppEntitlementResponse {
	if o == nil {
		return nil
	}
	return o.GetAppEntitlementResponse
}

func (o *C1APIAppV1AppEntitlementsGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
