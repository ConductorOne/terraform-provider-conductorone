// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppUsageControlsServiceGetRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppUsageControlsServiceGetResponse struct {
	ContentType string
	// Successful response
	GetAppUsageControlsResponse *shared.GetAppUsageControlsResponse
	StatusCode                  int
	RawResponse                 *http.Response
}
