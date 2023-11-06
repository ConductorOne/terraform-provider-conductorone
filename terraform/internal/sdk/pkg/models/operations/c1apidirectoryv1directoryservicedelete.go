// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIDirectoryV1DirectoryServiceDeleteRequest struct {
	DirectoryServiceDeleteRequest *shared.DirectoryServiceDeleteRequest `request:"mediaType=application/json"`
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
}

func (o *C1APIDirectoryV1DirectoryServiceDeleteRequest) GetDirectoryServiceDeleteRequest() *shared.DirectoryServiceDeleteRequest {
	if o == nil {
		return nil
	}
	return o.DirectoryServiceDeleteRequest
}

func (o *C1APIDirectoryV1DirectoryServiceDeleteRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type C1APIDirectoryV1DirectoryServiceDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Empty response with a status code indicating success.
	DirectoryServiceDeleteResponse *shared.DirectoryServiceDeleteResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIDirectoryV1DirectoryServiceDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIDirectoryV1DirectoryServiceDeleteResponse) GetDirectoryServiceDeleteResponse() *shared.DirectoryServiceDeleteResponse {
	if o == nil {
		return nil
	}
	return o.DirectoryServiceDeleteResponse
}

func (o *C1APIDirectoryV1DirectoryServiceDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIDirectoryV1DirectoryServiceDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
