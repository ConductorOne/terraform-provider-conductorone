// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest struct {
	RequestCatalogManagementServiceUpdateRequest *shared.RequestCatalogManagementServiceUpdateRequest `request:"mediaType=application/json"`
	ID                                           string                                               `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceGetResponse *shared.RequestCatalogManagementServiceGetResponse
	StatusCode                                 int
	RawResponse                                *http.Response
}
