// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskServiceCreateGrantTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	TaskServiceCreateGrantResponse *shared.TaskServiceCreateGrantResponse
}