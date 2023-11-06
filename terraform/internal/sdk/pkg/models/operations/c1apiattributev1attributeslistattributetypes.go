// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesListAttributeTypesRequest struct {
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAttributeV1AttributesListAttributeTypesRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAttributeV1AttributesListAttributeTypesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAttributeV1AttributesListAttributeTypesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// ListAttributeTypesResponse is the response for listing attribute types.
	ListAttributeTypesResponse *shared.ListAttributeTypesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAttributeV1AttributesListAttributeTypesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAttributeV1AttributesListAttributeTypesResponse) GetListAttributeTypesResponse() *shared.ListAttributeTypesResponse {
	if o == nil {
		return nil
	}
	return o.ListAttributeTypesResponse
}

func (o *C1APIAttributeV1AttributesListAttributeTypesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAttributeV1AttributesListAttributeTypesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
