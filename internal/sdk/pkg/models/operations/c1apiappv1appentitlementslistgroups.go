// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppEntitlementsListGroupsRequest struct {
	AppEntitlementID string `pathParam:"style=simple,explode=false,name=app_entitlement_id"`
	AppID            string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppEntitlementsListGroupsResponse struct {
	ContentType string
	// Successful response
	ListAppEntitlementGroupsResponse *shared.ListAppEntitlementGroupsResponse
	StatusCode                       int
	RawResponse                      *http.Response
}