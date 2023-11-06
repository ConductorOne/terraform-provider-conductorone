// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIDirectoryV1DirectoryServiceListRequest struct {
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIDirectoryV1DirectoryServiceListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIDirectoryV1DirectoryServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIDirectoryV1DirectoryServiceListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The DirectoryServiceListResponse message contains a list of results and a nextPageToken if applicable.
	DirectoryServiceListResponse *shared.DirectoryServiceListResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIDirectoryV1DirectoryServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIDirectoryV1DirectoryServiceListResponse) GetDirectoryServiceListResponse() *shared.DirectoryServiceListResponse {
	if o == nil {
		return nil
	}
	return o.DirectoryServiceListResponse
}

func (o *C1APIDirectoryV1DirectoryServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIDirectoryV1DirectoryServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
