// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APISettingsV1SessionSettingsServiceGetResponse struct {
	ContentType string
	// Successful response
	GetSessionSettingsResponse *shared.GetSessionSettingsResponse
	StatusCode                 int
	RawResponse                *http.Response
}

func (o *C1APISettingsV1SessionSettingsServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APISettingsV1SessionSettingsServiceGetResponse) GetGetSessionSettingsResponse() *shared.GetSessionSettingsResponse {
	if o == nil {
		return nil
	}
	return o.GetSessionSettingsResponse
}

func (o *C1APISettingsV1SessionSettingsServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APISettingsV1SessionSettingsServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
