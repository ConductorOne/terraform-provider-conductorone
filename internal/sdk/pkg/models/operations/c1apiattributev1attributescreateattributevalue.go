// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesCreateAttributeValueResponse struct {
	ContentType string
	// Successful response
	CreateAttributeValueResponse *shared.CreateAttributeValueResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
