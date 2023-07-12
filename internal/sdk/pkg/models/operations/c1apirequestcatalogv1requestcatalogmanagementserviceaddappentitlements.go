// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest struct {
	RequestCatalogManagementServiceAddAppEntitlementsRequest *shared.RequestCatalogManagementServiceAddAppEntitlementsRequest `request:"mediaType=application/json"`
	CatalogID                                                string                                                           `pathParam:"style=simple,explode=false,name=catalog_id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceAddAppEntitlementsResponse *shared.RequestCatalogManagementServiceAddAppEntitlementsResponse
	StatusCode                                                int
	RawResponse                                               *http.Response
}
