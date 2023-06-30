// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppReportServiceListRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppReportServiceListResponse struct {
	// Successful response
	AppReportServiceListResponse *shared.AppReportServiceListResponse
	ContentType                  string
	StatusCode                   int
	RawResponse                  *http.Response
}
