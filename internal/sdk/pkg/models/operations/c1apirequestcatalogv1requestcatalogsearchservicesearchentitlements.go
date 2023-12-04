// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The RequestCatalogSearchServiceSearchEntitlementsResponse message contains a list of results and a nextPageToken if applicable.
	RequestCatalogSearchServiceSearchEntitlementsResponse *shared.RequestCatalogSearchServiceSearchEntitlementsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse) GetRequestCatalogSearchServiceSearchEntitlementsResponse() *shared.RequestCatalogSearchServiceSearchEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogSearchServiceSearchEntitlementsResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
