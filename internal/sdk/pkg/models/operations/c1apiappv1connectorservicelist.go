// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceListRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1ConnectorServiceListResponse struct {
	// Successful response
	ConnectorServiceListResponse *shared.ConnectorServiceListResponse
	ContentType                  string
	StatusCode                   int
	RawResponse                  *http.Response
}
