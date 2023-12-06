// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesDeleteRequest struct {
	ID                  string                      `pathParam:"style=simple,explode=false,name=id"`
	DeletePolicyRequest *shared.DeletePolicyRequest `request:"mediaType=application/json"`
}

func (o *C1APIPolicyV1PoliciesDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIPolicyV1PoliciesDeleteRequest) GetDeletePolicyRequest() *shared.DeletePolicyRequest {
	if o == nil {
		return nil
	}
	return o.DeletePolicyRequest
}

type C1APIPolicyV1PoliciesDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Empty response with a status code indicating success.
	DeletePolicyResponse *shared.DeletePolicyResponse
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetDeletePolicyResponse() *shared.DeletePolicyResponse {
	if o == nil {
		return nil
	}
	return o.DeletePolicyResponse
}
