// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	RequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse *shared.RequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppResponse) GetRequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse() *shared.RequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse
}
