// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	RequestCatalogManagementServiceListResponse *shared.RequestCatalogManagementServiceListResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse) GetRequestCatalogManagementServiceListResponse() *shared.RequestCatalogManagementServiceListResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceListResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
