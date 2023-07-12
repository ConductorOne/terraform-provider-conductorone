// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIPolicyV1PoliciesUpdateRequest struct {
	UpdatePolicyRequest *shared.UpdatePolicyRequest `request:"mediaType=application/json"`
	ID                  string                      `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIPolicyV1PoliciesUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	UpdatePolicyResponse *shared.UpdatePolicyResponse
}
