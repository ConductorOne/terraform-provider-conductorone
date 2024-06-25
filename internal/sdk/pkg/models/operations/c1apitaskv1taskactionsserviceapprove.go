// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceApproveRequest struct {
	TaskActionsServiceApproveRequest *shared.TaskActionsServiceApproveRequest `request:"mediaType=application/json"`
	TaskID                           string                                   `pathParam:"style=simple,explode=false,name=task_id"`
}

type C1APITaskV1TaskActionsServiceApproveResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	TaskActionsServiceApproveResponse *shared.TaskActionsServiceApproveResponse
}
