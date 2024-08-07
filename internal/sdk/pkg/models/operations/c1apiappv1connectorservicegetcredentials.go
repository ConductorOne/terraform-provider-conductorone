// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceGetCredentialsRequest struct {
	AppID       string `pathParam:"style=simple,explode=false,name=app_id"`
	ConnectorID string `pathParam:"style=simple,explode=false,name=connector_id"`
	ID          string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1ConnectorServiceGetCredentialsResponse struct {
	// Successful response
	ConnectorServiceGetCredentialsResponse *shared.ConnectorServiceGetCredentialsResponse
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
}
