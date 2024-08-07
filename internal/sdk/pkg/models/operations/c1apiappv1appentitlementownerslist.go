// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementOwnersListRequest struct {
	AppID         string   `pathParam:"style=simple,explode=false,name=app_id"`
	EntitlementID string   `pathParam:"style=simple,explode=false,name=entitlement_id"`
	PageSize      *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken     *string  `queryParam:"style=form,explode=true,name=page_token"`
}

type C1APIAppV1AppEntitlementOwnersListResponse struct {
	ContentType string
	// Successful response
	ListAppEntitlementOwnersResponse *shared.ListAppEntitlementOwnersResponse
	StatusCode                       int
	RawResponse                      *http.Response
}
