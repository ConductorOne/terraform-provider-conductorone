// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIIamV1PersonalClientServiceCreateResponse struct {
	ContentType string
	// Successful response
	PersonalClientServiceCreateResponse *shared.PersonalClientServiceCreateResponse
	StatusCode                          int
	RawResponse                         *http.Response
}
