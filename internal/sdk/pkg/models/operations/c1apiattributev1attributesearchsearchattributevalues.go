// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse struct {
	ContentType string
	// Successful response
	SearchAttributeValuesResponse *shared.SearchAttributeValuesResponse
	StatusCode                    int
	RawResponse                   *http.Response
}
