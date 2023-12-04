// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementOwnersRemoveRequest struct {
	RemoveAppEntitlementOwnerRequest *shared.RemoveAppEntitlementOwnerRequest `request:"mediaType=application/json"`
	AppID                            string                                   `pathParam:"style=simple,explode=false,name=app_id"`
	EntitlementID                    string                                   `pathParam:"style=simple,explode=false,name=entitlement_id"`
	UserID                           string                                   `pathParam:"style=simple,explode=false,name=user_id"`
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveRequest) GetRemoveAppEntitlementOwnerRequest() *shared.RemoveAppEntitlementOwnerRequest {
	if o == nil {
		return nil
	}
	return o.RemoveAppEntitlementOwnerRequest
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveRequest) GetEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.EntitlementID
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

type C1APIAppV1AppEntitlementOwnersRemoveResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The empty response message for removing an app entitlement owner.
	RemoveAppEntitlementOwnerResponse *shared.RemoveAppEntitlementOwnerResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveResponse) GetRemoveAppEntitlementOwnerResponse() *shared.RemoveAppEntitlementOwnerResponse {
	if o == nil {
		return nil
	}
	return o.RemoveAppEntitlementOwnerResponse
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementOwnersRemoveResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
