// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"conductorone/internal/sdk/pkg/models/shared"
)

type C1APIAppV1AppResourceServiceListRequest struct {
	AppID             string `pathParam:"style=simple,explode=false,name=app_id"`
	AppResourceTypeID string `pathParam:"style=simple,explode=false,name=app_resource_type_id"`
}

type C1APIAppV1AppResourceServiceListResponse struct {
	// Successful response
	AppResourceServiceListResponse *shared.AppResourceServiceListResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}
