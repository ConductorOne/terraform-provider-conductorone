// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest struct {
	RequestCatalogManagementServiceDeleteRequest *shared.RequestCatalogManagementServiceDeleteRequest `request:"mediaType=application/json"`
	ID                                           string                                               `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceDeleteResponse *shared.RequestCatalogManagementServiceDeleteResponse
	StatusCode                                    int
	RawResponse                                   *http.Response
}
