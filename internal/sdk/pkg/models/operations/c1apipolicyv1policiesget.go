// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIPolicyV1PoliciesGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIPolicyV1PoliciesGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The GetPolicyResponse message contains the policy object.
	GetPolicyResponse *shared.GetPolicyResponse
}

func (o *C1APIPolicyV1PoliciesGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIPolicyV1PoliciesGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIPolicyV1PoliciesGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIPolicyV1PoliciesGetResponse) GetGetPolicyResponse() *shared.GetPolicyResponse {
	if o == nil {
		return nil
	}
	return o.GetPolicyResponse
}
