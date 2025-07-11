// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsProxyCreateRequest struct {
	SrcAppID                         string                                   `pathParam:"style=simple,explode=false,name=src_app_id"`
	SrcAppEntitlementID              string                                   `pathParam:"style=simple,explode=false,name=src_app_entitlement_id"`
	DstAppID                         string                                   `pathParam:"style=simple,explode=false,name=dst_app_id"`
	DstAppEntitlementID              string                                   `pathParam:"style=simple,explode=false,name=dst_app_entitlement_id"`
	CreateAppEntitlementProxyRequest *shared.CreateAppEntitlementProxyRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppEntitlementsProxyCreateRequest) GetSrcAppID() string {
	if o == nil {
		return ""
	}
	return o.SrcAppID
}

func (o *C1APIAppV1AppEntitlementsProxyCreateRequest) GetSrcAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.SrcAppEntitlementID
}

func (o *C1APIAppV1AppEntitlementsProxyCreateRequest) GetDstAppID() string {
	if o == nil {
		return ""
	}
	return o.DstAppID
}

func (o *C1APIAppV1AppEntitlementsProxyCreateRequest) GetDstAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.DstAppEntitlementID
}

func (o *C1APIAppV1AppEntitlementsProxyCreateRequest) GetCreateAppEntitlementProxyRequest() *shared.CreateAppEntitlementProxyRequest {
	if o == nil {
		return nil
	}
	return o.CreateAppEntitlementProxyRequest
}

type C1APIAppV1AppEntitlementsProxyCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	CreateAppEntitlementProxyResponse *shared.CreateAppEntitlementProxyResponse
}

func (o *C1APIAppV1AppEntitlementsProxyCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsProxyCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsProxyCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppEntitlementsProxyCreateResponse) GetCreateAppEntitlementProxyResponse() *shared.CreateAppEntitlementProxyResponse {
	if o == nil {
		return nil
	}
	return o.CreateAppEntitlementProxyResponse
}
