// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskServiceGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APITaskV1TaskServiceGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APITaskV1TaskServiceGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The TaskServiceGetResponse returns a task view which has a task including JSONPATHs to the expanded items in the expanded array.
	TaskServiceGetResponse *shared.TaskServiceGetResponse
}

func (o *C1APITaskV1TaskServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskServiceGetResponse) GetTaskServiceGetResponse() *shared.TaskServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.TaskServiceGetResponse
}
