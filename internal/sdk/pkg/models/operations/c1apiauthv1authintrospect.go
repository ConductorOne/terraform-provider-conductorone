// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APIAuthV1AuthIntrospectResponse struct {
	ContentType string
	// IntrospectResponse contains information about the current user who is authenticated.
	IntrospectResponse *shared.IntrospectResponse
	StatusCode         int
	RawResponse        *http.Response
}

func (o *C1APIAuthV1AuthIntrospectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAuthV1AuthIntrospectResponse) GetIntrospectResponse() *shared.IntrospectResponse {
	if o == nil {
		return nil
	}
	return o.IntrospectResponse
}

func (o *C1APIAuthV1AuthIntrospectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAuthV1AuthIntrospectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
