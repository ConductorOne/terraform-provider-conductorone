// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIWebhooksV1WebhooksServiceDeleteRequest struct {
	ID                           string                               `pathParam:"style=simple,explode=false,name=id"`
	WebhooksServiceDeleteRequest *shared.WebhooksServiceDeleteRequest `request:"mediaType=application/json"`
}

func (o *C1APIWebhooksV1WebhooksServiceDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIWebhooksV1WebhooksServiceDeleteRequest) GetWebhooksServiceDeleteRequest() *shared.WebhooksServiceDeleteRequest {
	if o == nil {
		return nil
	}
	return o.WebhooksServiceDeleteRequest
}

type C1APIWebhooksV1WebhooksServiceDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Empty response body. Status code indicates success.
	WebhooksServiceDeleteResponse *shared.WebhooksServiceDeleteResponse
}

func (o *C1APIWebhooksV1WebhooksServiceDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIWebhooksV1WebhooksServiceDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIWebhooksV1WebhooksServiceDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIWebhooksV1WebhooksServiceDeleteResponse) GetWebhooksServiceDeleteResponse() *shared.WebhooksServiceDeleteResponse {
	if o == nil {
		return nil
	}
	return o.WebhooksServiceDeleteResponse
}