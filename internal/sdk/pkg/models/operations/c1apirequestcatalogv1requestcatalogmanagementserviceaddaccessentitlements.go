// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest struct {
	RequestCatalogManagementServiceAddAccessEntitlementsRequest *shared.RequestCatalogManagementServiceAddAccessEntitlementsRequest `request:"mediaType=application/json"`
	CatalogID                                                   string                                                              `pathParam:"style=simple,explode=false,name=catalog_id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceAddAccessEntitlementsResponse *shared.RequestCatalogManagementServiceAddAccessEntitlementsResponse
	StatusCode                                                   int
	RawResponse                                                  *http.Response
}