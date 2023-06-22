// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppReportActionServiceGenerateReportRequest struct {
	AppActionsServiceGenerateReportRequest *shared.AppActionsServiceGenerateReportRequest `request:"mediaType=application/json"`
	AppID                                  string                                         `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppReportActionServiceGenerateReportResponse struct {
	// Successful response
	AppActionsServiceGenerateReportResponse *shared.AppActionsServiceGenerateReportResponse
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
}
