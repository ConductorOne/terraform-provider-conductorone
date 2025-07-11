// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// DirectoryServiceCreateResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type DirectoryServiceCreateResponseExpanded struct {
}

// The DirectoryServiceCreateResponse message.
type DirectoryServiceCreateResponse struct {
	// The directory view contains a directory and an app_path which is a JSONPATH set to the location in the expand mask that the expanded app will live if requested by the expander.
	DirectoryView *DirectoryView `json:"directoryView,omitempty"`
	// List of serialized related objects.
	Expanded []DirectoryServiceCreateResponseExpanded `json:"expanded,omitempty"`
}

func (o *DirectoryServiceCreateResponse) GetDirectoryView() *DirectoryView {
	if o == nil {
		return nil
	}
	return o.DirectoryView
}

func (o *DirectoryServiceCreateResponse) GetExpanded() []DirectoryServiceCreateResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
