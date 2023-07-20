// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIPolicyV1PoliciesGetResponse struct {
	ContentType string
	// Successful response
	GetPolicyResponse *shared.GetPolicyResponse
	StatusCode        int
	RawResponse       *http.Response
}
