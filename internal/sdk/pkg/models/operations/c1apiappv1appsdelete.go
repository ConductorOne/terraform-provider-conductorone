// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsDeleteRequest struct {
	DeleteAppRequest *shared.DeleteAppRequest `request:"mediaType=application/json"`
	ID               string                   `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1AppsDeleteResponse struct {
	ContentType string
	// Successful response
	DeleteAppResponse *shared.DeleteAppResponse
	StatusCode        int
	RawResponse       *http.Response
}
