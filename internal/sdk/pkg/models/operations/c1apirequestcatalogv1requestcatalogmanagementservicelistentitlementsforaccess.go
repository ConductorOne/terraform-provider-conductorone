// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalog_id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceListEntitlementsForAccessResponse *shared.RequestCatalogManagementServiceListEntitlementsForAccessResponse
	StatusCode                                                       int
	RawResponse                                                      *http.Response
}
