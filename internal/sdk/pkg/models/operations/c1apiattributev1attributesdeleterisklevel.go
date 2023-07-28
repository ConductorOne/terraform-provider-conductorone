// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesDeleteRiskLevelRequest struct {
	DeleteRiskLevelRequest *shared.DeleteRiskLevelRequest `request:"mediaType=application/json"`
	ID                     string                         `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAttributeV1AttributesDeleteRiskLevelResponse struct {
	ContentType string
	// Successful response
	DeleteRiskLevelResponse *shared.DeleteRiskLevelResponse
	StatusCode              int
	RawResponse             *http.Response
}
