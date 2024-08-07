// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceOwnersListRequest struct {
	AppID          string `pathParam:"style=simple,explode=false,name=app_id"`
	ResourceID     string `pathParam:"style=simple,explode=false,name=resource_id"`
	ResourceTypeID string `pathParam:"style=simple,explode=false,name=resource_type_id"`
}

type C1APIAppV1AppResourceOwnersListResponse struct {
	ContentType string
	// Successful response
	ListAppResourceOwnersResponse *shared.ListAppResourceOwnersResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
