// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIIamV1RolesListResponse struct {
	ContentType string
	// Successful response
	ListRolesResponse *shared.ListRolesResponse
	StatusCode        int
	RawResponse       *http.Response
}