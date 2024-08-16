// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type C1APIPolicyV1PoliciesUpdateRequest struct {
	UpdatePolicyRequestInput *shared.UpdatePolicyRequestInput `request:"mediaType=application/json"`
	ID                       string                           `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIPolicyV1PoliciesUpdateRequest) GetUpdatePolicyRequestInput() *shared.UpdatePolicyRequestInput {
	if o == nil {
		return nil
	}
	return o.UpdatePolicyRequestInput
}

func (o *C1APIPolicyV1PoliciesUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIPolicyV1PoliciesUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The UpdatePolicyResponse message contains the updated policy object.
	UpdatePolicyResponse *shared.UpdatePolicyResponse
}

func (o *C1APIPolicyV1PoliciesUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIPolicyV1PoliciesUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIPolicyV1PoliciesUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIPolicyV1PoliciesUpdateResponse) GetUpdatePolicyResponse() *shared.UpdatePolicyResponse {
	if o == nil {
		return nil
	}
	return o.UpdatePolicyResponse
}
