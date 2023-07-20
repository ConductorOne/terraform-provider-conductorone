// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListForAppResourceRequest struct {
	AppID             string `pathParam:"style=simple,explode=false,name=app_id"`
	AppResourceID     string `pathParam:"style=simple,explode=false,name=app_resource_id"`
	AppResourceTypeID string `pathParam:"style=simple,explode=false,name=app_resource_type_id"`
}

type C1APIAppV1AppEntitlementsListForAppResourceResponse struct {
	ContentType string
	// Successful response
	ListAppEntitlementsResponse *shared.ListAppEntitlementsResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
