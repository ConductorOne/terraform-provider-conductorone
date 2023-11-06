// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementSearchServiceSearchResponse struct {
	// Successful response
	AppEntitlementSearchServiceSearchResponse *shared.AppEntitlementSearchServiceSearchResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchResponse) GetAppEntitlementSearchServiceSearchResponse() *shared.AppEntitlementSearchServiceSearchResponse {
	if o == nil {
		return nil
	}
	return o.AppEntitlementSearchServiceSearchResponse
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementSearchServiceSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
