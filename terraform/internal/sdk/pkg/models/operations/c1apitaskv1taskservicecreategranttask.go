// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskServiceCreateGrantTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The TaskServiceCreateGrantResponse returns a task view which has a task including JSONPATHs to the expanded items in the expanded array.
	TaskServiceCreateGrantResponse *shared.TaskServiceCreateGrantResponse
}

func (o *C1APITaskV1TaskServiceCreateGrantTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskServiceCreateGrantTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskServiceCreateGrantTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskServiceCreateGrantTaskResponse) GetTaskServiceCreateGrantResponse() *shared.TaskServiceCreateGrantResponse {
	if o == nil {
		return nil
	}
	return o.TaskServiceCreateGrantResponse
}
