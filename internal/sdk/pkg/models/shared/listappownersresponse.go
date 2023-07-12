// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListAppOwnersResponse - The ListAppOwnersResponse message.
type ListAppOwnersResponse struct {
	// The list field.
	List []User `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The notificationToken field.
	NotificationToken *string `json:"notificationToken,omitempty"`
}