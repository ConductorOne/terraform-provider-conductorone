// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListUsersRequest struct {
	AppEntitlementID string `pathParam:"style=simple,explode=false,name=app_entitlement_id"`
	AppID            string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppEntitlementsListUsersResponse struct {
	ContentType string
	// Successful response
	ListAppEntitlementUsersResponse *shared.ListAppEntitlementUsersResponse
	StatusCode                      int
	RawResponse                     *http.Response
}
