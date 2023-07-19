// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceGetResponse *shared.RequestCatalogManagementServiceGetResponse
	StatusCode                                 int
	RawResponse                                *http.Response
}