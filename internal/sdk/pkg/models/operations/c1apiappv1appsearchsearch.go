// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppSearchSearchResponse struct {
	ContentType string
	// Successful response
	SearchAppsResponse *shared.SearchAppsResponse
	StatusCode         int
	RawResponse        *http.Response
}
