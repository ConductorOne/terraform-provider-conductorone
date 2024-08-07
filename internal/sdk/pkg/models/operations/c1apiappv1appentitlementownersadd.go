// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementOwnersAddRequest struct {
	AddAppEntitlementOwnerRequest *shared.AddAppEntitlementOwnerRequest `request:"mediaType=application/json"`
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
	EntitlementID                 string                                `pathParam:"style=simple,explode=false,name=entitlement_id"`
}

type C1APIAppV1AppEntitlementOwnersAddResponse struct {
	// Successful response
	AddAppEntitlementOwnerResponse *shared.AddAppEntitlementOwnerResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}
