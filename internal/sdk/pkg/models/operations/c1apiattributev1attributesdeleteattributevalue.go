// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIAttributeV1AttributesDeleteAttributeValueRequest struct {
	DeleteAttributeValueRequest *shared.DeleteAttributeValueRequest `request:"mediaType=application/json"`
	ID                          string                              `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAttributeV1AttributesDeleteAttributeValueResponse struct {
	ContentType string
	// Successful response
	DeleteAttributeValueResponse *shared.DeleteAttributeValueResponse
	StatusCode                   int
	RawResponse                  *http.Response
}
