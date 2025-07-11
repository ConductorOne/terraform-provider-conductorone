// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest struct {
	CatalogID                                                string                                                           `pathParam:"style=simple,explode=false,name=catalog_id"`
	RequestCatalogManagementServiceAddAppEntitlementsRequest *shared.RequestCatalogManagementServiceAddAppEntitlementsRequest `request:"mediaType=application/json"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest) GetRequestCatalogManagementServiceAddAppEntitlementsRequest() *shared.RequestCatalogManagementServiceAddAppEntitlementsRequest {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceAddAppEntitlementsRequest
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Empty response with a status code indicating success.
	RequestCatalogManagementServiceAddAppEntitlementsResponse *shared.RequestCatalogManagementServiceAddAppEntitlementsResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse) GetRequestCatalogManagementServiceAddAppEntitlementsResponse() *shared.RequestCatalogManagementServiceAddAppEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceAddAppEntitlementsResponse
}
