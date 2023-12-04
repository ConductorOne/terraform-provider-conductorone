// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIUserV1UserServiceGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIUserV1UserServiceGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIUserV1UserServiceGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The UserServiceGetResponse returns a user view which has a user including JSONPATHs to the expanded items in the expanded array.
	UserServiceGetResponse *shared.UserServiceGetResponse
}

func (o *C1APIUserV1UserServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIUserV1UserServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIUserV1UserServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIUserV1UserServiceGetResponse) GetUserServiceGetResponse() *shared.UserServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.UserServiceGetResponse
}
