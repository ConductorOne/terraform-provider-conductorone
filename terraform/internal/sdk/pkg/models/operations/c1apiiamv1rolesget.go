// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIIamV1RolesGetRequest struct {
	RoleID string `pathParam:"style=simple,explode=false,name=role_id"`
}

func (o *C1APIIamV1RolesGetRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

type C1APIIamV1RolesGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The GetRolesResponse message contains the retrieved role.
	GetRolesResponse *shared.GetRolesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIIamV1RolesGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIIamV1RolesGetResponse) GetGetRolesResponse() *shared.GetRolesResponse {
	if o == nil {
		return nil
	}
	return o.GetRolesResponse
}

func (o *C1APIIamV1RolesGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIIamV1RolesGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
