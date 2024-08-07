// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsUpdateRequest struct {
	UpdateAppEntitlementRequest *shared.UpdateAppEntitlementRequest `request:"mediaType=application/json"`
	AppID                       string                              `pathParam:"style=simple,explode=false,name=app_id"`
	ID                          string                              `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1AppEntitlementsUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	UpdateAppEntitlementResponse *shared.UpdateAppEntitlementResponse
}
