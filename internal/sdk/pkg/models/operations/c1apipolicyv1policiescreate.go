// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIPolicyV1PoliciesCreateResponse struct {
	ContentType string
	// Successful response
	CreatePolicyResponse *shared.CreatePolicyResponse
	StatusCode           int
	RawResponse          *http.Response
}
