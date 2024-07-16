// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceTypeServiceListRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

type C1APIAppV1AppResourceTypeServiceListResponse struct {
	// Successful response
	AppResourceTypeServiceListResponse *shared.AppResourceTypeServiceListResponse
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
}
