// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceGetResponse *shared.RequestCatalogManagementServiceGetResponse
	StatusCode                                 int
	RawResponse                                *http.Response
}