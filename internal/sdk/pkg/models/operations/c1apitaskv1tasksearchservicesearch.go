// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type C1APITaskV1TaskSearchServiceSearchResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	TaskSearchResponse *shared.TaskSearchResponse
}
