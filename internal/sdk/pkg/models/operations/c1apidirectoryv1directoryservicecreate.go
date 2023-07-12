// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIDirectoryV1DirectoryServiceCreateResponse struct {
	ContentType string
	// Successful response
	DirectoryServiceCreateResponse *shared.DirectoryServiceCreateResponse
	StatusCode                     int
	RawResponse                    *http.Response
}