// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request catalog management service get response returns a request catalog view with the expanded items in the expanded array indicated by the expand mask in the request.
	RequestCatalogManagementServiceGetResponse *shared.RequestCatalogManagementServiceGetResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetRequestCatalogManagementServiceGetResponse() *shared.RequestCatalogManagementServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceGetResponse
}
