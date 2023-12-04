// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIDirectoryV1DirectoryServiceCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The DirectoryServiceCreateResponse message.
	DirectoryServiceCreateResponse *shared.DirectoryServiceCreateResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIDirectoryV1DirectoryServiceCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIDirectoryV1DirectoryServiceCreateResponse) GetDirectoryServiceCreateResponse() *shared.DirectoryServiceCreateResponse {
	if o == nil {
		return nil
	}
	return o.DirectoryServiceCreateResponse
}

func (o *C1APIDirectoryV1DirectoryServiceCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIDirectoryV1DirectoryServiceCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
