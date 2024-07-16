// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest struct {
	RequestCatalogManagementServiceRemoveAccessEntitlementsRequest *shared.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest `request:"mediaType=application/json"`
	CatalogID                                                      string                                                                 `pathParam:"style=simple,explode=false,name=catalog_id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceRemoveAccessEntitlementsResponse *shared.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse
	StatusCode                                                      int
	RawResponse                                                     *http.Response
}
