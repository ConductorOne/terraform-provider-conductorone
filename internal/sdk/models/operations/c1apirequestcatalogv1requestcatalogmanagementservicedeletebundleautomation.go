// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationRequest struct {
	RequestCatalogID              string                                `pathParam:"style=simple,explode=false,name=request_catalog_id"`
	DeleteBundleAutomationRequest *shared.DeleteBundleAutomationRequest `request:"mediaType=application/json"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationRequest) GetRequestCatalogID() string {
	if o == nil {
		return ""
	}
	return o.RequestCatalogID
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationRequest) GetDeleteBundleAutomationRequest() *shared.DeleteBundleAutomationRequest {
	if o == nil {
		return nil
	}
	return o.DeleteBundleAutomationRequest
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	DeleteBundleAutomationResponse *shared.DeleteBundleAutomationResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationResponse) GetDeleteBundleAutomationResponse() *shared.DeleteBundleAutomationResponse {
	if o == nil {
		return nil
	}
	return o.DeleteBundleAutomationResponse
}
