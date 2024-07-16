// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceUpdateRequest struct {
	ConnectorServiceUpdateRequest *shared.ConnectorServiceUpdateRequest `request:"mediaType=application/json"`
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
	ID                            string                                `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1ConnectorServiceUpdateResponse struct {
	// Successful response
	ConnectorServiceUpdateResponse *shared.ConnectorServiceUpdateResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}
