// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIWebhooksV1WebhooksSearchSearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	WebhooksSearchResponse *shared.WebhooksSearchResponse
}

func (o *C1APIWebhooksV1WebhooksSearchSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIWebhooksV1WebhooksSearchSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIWebhooksV1WebhooksSearchSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIWebhooksV1WebhooksSearchSearchResponse) GetWebhooksSearchResponse() *shared.WebhooksSearchResponse {
	if o == nil {
		return nil
	}
	return o.WebhooksSearchResponse
}
