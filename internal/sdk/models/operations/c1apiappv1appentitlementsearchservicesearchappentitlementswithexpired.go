// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest struct {
	AppID            string  `pathParam:"style=simple,explode=false,name=app_id"`
	AppEntitlementID string  `pathParam:"style=simple,explode=false,name=app_entitlement_id"`
	PageSize         *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken        *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest) GetAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.AppEntitlementID
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The SearchAppEntitlementsWithExpiredResponse message contains a list of results and a nextPageToken if applicable.
	SearchAppEntitlementsWithExpiredResponse *shared.SearchAppEntitlementsWithExpiredResponse
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse) GetSearchAppEntitlementsWithExpiredResponse() *shared.SearchAppEntitlementsWithExpiredResponse {
	if o == nil {
		return nil
	}
	return o.SearchAppEntitlementsWithExpiredResponse
}
