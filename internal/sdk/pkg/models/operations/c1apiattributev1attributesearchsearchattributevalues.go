// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse struct {
	ContentType string
	// SearchAttributeValuesResponse is the response for searching AttributeValues.
	SearchAttributeValuesResponse *shared.SearchAttributeValuesResponse
	StatusCode                    int
	RawResponse                   *http.Response
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetSearchAttributeValuesResponse() *shared.SearchAttributeValuesResponse {
	if o == nil {
		return nil
	}
	return o.SearchAttributeValuesResponse
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAttributeV1AttributeSearchSearchAttributeValuesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
