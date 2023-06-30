// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppOwnersListRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppOwnersListResponse struct {
	ContentType string
	// Successful response
	ListAppOwnersResponse *shared.ListAppOwnersResponse
	StatusCode            int
	RawResponse           *http.Response
}
