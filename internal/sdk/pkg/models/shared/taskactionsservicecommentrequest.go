// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskActionsServiceCommentRequest - The TaskActionsServiceCommentRequest message.
type TaskActionsServiceCommentRequest struct {
	// The comment field.
	Comment *string `json:"comment,omitempty"`
	//  Make sure to update the TicketExpandMask
	//
	ExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
}
