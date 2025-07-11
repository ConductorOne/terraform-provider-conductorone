// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// DirectoryServiceCreateRequest - Uplevel an app into a full directory.
type DirectoryServiceCreateRequest struct {
	// The AppID to make into a directory, providing identities and more for the C1 app.
	AppID *string `json:"appId,omitempty"`
}

func (o *DirectoryServiceCreateRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}
