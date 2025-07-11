// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceCreateRequest struct {
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
	ConnectorServiceCreateRequest *shared.ConnectorServiceCreateRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1ConnectorServiceCreateRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1ConnectorServiceCreateRequest) GetConnectorServiceCreateRequest() *shared.ConnectorServiceCreateRequest {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceCreateRequest
}

type C1APIAppV1ConnectorServiceCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The ConnectorServiceCreateResponse is the response returned from creating a connector.
	ConnectorServiceCreateResponse *shared.ConnectorServiceCreateResponse
}

func (o *C1APIAppV1ConnectorServiceCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1ConnectorServiceCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1ConnectorServiceCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1ConnectorServiceCreateResponse) GetConnectorServiceCreateResponse() *shared.ConnectorServiceCreateResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceCreateResponse
}
