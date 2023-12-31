// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppOwnersSetRequest struct {
	SetAppOwnersRequest *shared.SetAppOwnersRequest `request:"mediaType=application/json"`
	AppID               string                      `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppOwnersSetResponse struct {
	ContentType string
	// The empty response message for setting the app owners.
	SetAppOwnersResponse *shared.SetAppOwnersResponse
	StatusCode           int
	RawResponse          *http.Response
}
