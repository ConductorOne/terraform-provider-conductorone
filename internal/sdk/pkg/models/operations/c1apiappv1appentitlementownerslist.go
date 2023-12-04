// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementOwnersListRequest struct {
	AppID         string   `pathParam:"style=simple,explode=false,name=app_id"`
	EntitlementID string   `pathParam:"style=simple,explode=false,name=entitlement_id"`
	PageSize      *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken     *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppEntitlementOwnersListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementOwnersListRequest) GetEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.EntitlementID
}

func (o *C1APIAppV1AppEntitlementOwnersListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppEntitlementOwnersListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppEntitlementOwnersListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The response message for listing app entitlement owners.
	ListAppEntitlementOwnersResponse *shared.ListAppEntitlementOwnersResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppEntitlementOwnersListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementOwnersListResponse) GetListAppEntitlementOwnersResponse() *shared.ListAppEntitlementOwnersResponse {
	if o == nil {
		return nil
	}
	return o.ListAppEntitlementOwnersResponse
}

func (o *C1APIAppV1AppEntitlementOwnersListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementOwnersListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
