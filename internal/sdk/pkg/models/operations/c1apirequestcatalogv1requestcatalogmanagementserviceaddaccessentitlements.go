// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest struct {
	CatalogID                                                   string                                                              `pathParam:"style=simple,explode=false,name=catalog_id"`
	RequestCatalogManagementServiceAddAccessEntitlementsRequest *shared.RequestCatalogManagementServiceAddAccessEntitlementsRequest `request:"mediaType=application/json"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest) GetRequestCatalogManagementServiceAddAccessEntitlementsRequest() *shared.RequestCatalogManagementServiceAddAccessEntitlementsRequest {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceAddAccessEntitlementsRequest
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Empty response with a status code indicating success.
	RequestCatalogManagementServiceAddAccessEntitlementsResponse *shared.RequestCatalogManagementServiceAddAccessEntitlementsResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse) GetRequestCatalogManagementServiceAddAccessEntitlementsResponse() *shared.RequestCatalogManagementServiceAddAccessEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceAddAccessEntitlementsResponse
}
