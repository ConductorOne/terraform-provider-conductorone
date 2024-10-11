package operations

import (
	"net/http"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppEntitlementsProxyDeleteRequest struct {
	SrcAppID                         string                                   `pathParam:"style=simple,explode=false,name=src_app_id"`
	SrcAppEntitlementID              string                                   `pathParam:"style=simple,explode=false,name=src_app_entitlement_id"`
	DstAppID                         string                                   `pathParam:"style=simple,explode=false,name=dst_app_id"`
	DstAppEntitlementID              string                                   `pathParam:"style=simple,explode=false,name=dst_app_entitlement_id"`
	DeleteAppEntitlementProxyRequest *shared.DeleteAppEntitlementProxyRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteRequest) GetSrcAppID() string {
	if o == nil {
		return ""
	}
	return o.SrcAppID
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteRequest) GetSrcAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.SrcAppEntitlementID
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteRequest) GetDstAppID() string {
	if o == nil {
		return ""
	}
	return o.DstAppID
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteRequest) GetDstAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.DstAppEntitlementID
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteRequest) GetDeleteAppEntitlementProxyRequest() *shared.DeleteAppEntitlementProxyRequest {
	if o == nil {
		return nil
	}
	return o.DeleteAppEntitlementProxyRequest
}

type C1APIAppV1AppEntitlementsProxyDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	DeleteAppEntitlementProxyResponse *shared.DeleteAppEntitlementProxyResponse
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppEntitlementsProxyDeleteResponse) GetDeleteAppEntitlementProxyResponse() *shared.DeleteAppEntitlementProxyResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAppEntitlementProxyResponse
}
