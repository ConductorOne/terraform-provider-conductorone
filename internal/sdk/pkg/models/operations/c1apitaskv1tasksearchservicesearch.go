// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskSearchServiceSearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The TaskSearchResponse message contains a list of results and a nextPageToken if applicable.
	TaskSearchResponse *shared.TaskSearchResponse
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetTaskSearchResponse() *shared.TaskSearchResponse {
	if o == nil {
		return nil
	}
	return o.TaskSearchResponse
}
