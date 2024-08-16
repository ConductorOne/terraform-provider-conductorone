// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APIAppV1ConnectorServiceForceSyncRequest struct {
	ForceSyncRequest *shared.ForceSyncRequest `request:"mediaType=application/json"`
	AppID            string                   `pathParam:"style=simple,explode=false,name=app_id"`
	ConnectorID      string                   `pathParam:"style=simple,explode=false,name=connector_id"`
}

func (o *C1APIAppV1ConnectorServiceForceSyncRequest) GetForceSyncRequest() *shared.ForceSyncRequest {
	if o == nil {
		return nil
	}
	return o.ForceSyncRequest
}

func (o *C1APIAppV1ConnectorServiceForceSyncRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1ConnectorServiceForceSyncRequest) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

type C1APIAppV1ConnectorServiceForceSyncResponse struct {
	ContentType string
	// Successful response
	ForceSyncResponse *shared.ForceSyncResponse
	StatusCode        int
	RawResponse       *http.Response
}

func (o *C1APIAppV1ConnectorServiceForceSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1ConnectorServiceForceSyncResponse) GetForceSyncResponse() *shared.ForceSyncResponse {
	if o == nil {
		return nil
	}
	return o.ForceSyncResponse
}

func (o *C1APIAppV1ConnectorServiceForceSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1ConnectorServiceForceSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
