// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APIAppV1ConnectorServiceListRequest struct {
	AppID     string  `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1ConnectorServiceListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1ConnectorServiceListRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1ConnectorServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1ConnectorServiceListResponse struct {
	// The ConnectorServiceListResponse message contains a list of results and a nextPageToken if applicable
	ConnectorServiceListResponse *shared.ConnectorServiceListResponse
	ContentType                  string
	StatusCode                   int
	RawResponse                  *http.Response
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetConnectorServiceListResponse() *shared.ConnectorServiceListResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceListResponse
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
