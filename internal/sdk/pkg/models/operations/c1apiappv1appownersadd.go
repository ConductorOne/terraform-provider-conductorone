// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppOwnersAddRequest struct {
	AddAppOwnerRequest *shared.AddAppOwnerRequest `request:"mediaType=application/json"`
	AppID              string                     `pathParam:"style=simple,explode=false,name=app_id"`
	UserID             string                     `pathParam:"style=simple,explode=false,name=user_id"`
}

type C1APIAppV1AppOwnersAddResponse struct {
	// Successful response
	AddAppOwnerResponse *shared.AddAppOwnerResponse
	ContentType         string
	StatusCode          int
	RawResponse         *http.Response
}
