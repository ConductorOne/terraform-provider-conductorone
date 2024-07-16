// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIIamV1RolesUpdateRequest struct {
	UpdateRoleRequest *shared.UpdateRoleRequest `request:"mediaType=application/json"`
	RoleID            string                    `pathParam:"style=simple,explode=false,name=role_id"`
}

type C1APIIamV1RolesUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	UpdateRolesResponse *shared.UpdateRolesResponse
}
