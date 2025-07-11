// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// DirectoryView - The directory view contains a directory and an app_path which is a JSONPATH set to the location in the expand mask that the expanded app will live if requested by the expander.
type DirectoryView struct {
	// JSONPATH expression indicating the location of the App object in the  array.
	AppPath *string `json:"appPath,omitempty"`
	// This object indicates that an app is also a directory.
	Directory *Directory `json:"directory,omitempty"`
}

func (o *DirectoryView) GetAppPath() *string {
	if o == nil {
		return nil
	}
	return o.AppPath
}

func (o *DirectoryView) GetDirectory() *Directory {
	if o == nil {
		return nil
	}
	return o.Directory
}
