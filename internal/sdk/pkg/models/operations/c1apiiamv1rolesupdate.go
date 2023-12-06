// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIIamV1RolesUpdateRequest struct {
	RoleID            string                    `pathParam:"style=simple,explode=false,name=role_id"`
	UpdateRoleRequest *shared.UpdateRoleRequest `request:"mediaType=application/json"`
}

func (o *C1APIIamV1RolesUpdateRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

func (o *C1APIIamV1RolesUpdateRequest) GetUpdateRoleRequest() *shared.UpdateRoleRequest {
	if o == nil {
		return nil
	}
	return o.UpdateRoleRequest
}

type C1APIIamV1RolesUpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// UpdateRolesResponse is the response message containing the updated role.
	UpdateRolesResponse *shared.UpdateRolesResponse
}

func (o *C1APIIamV1RolesUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIIamV1RolesUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIIamV1RolesUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIIamV1RolesUpdateResponse) GetUpdateRolesResponse() *shared.UpdateRolesResponse {
	if o == nil {
		return nil
	}
	return o.UpdateRolesResponse
}
