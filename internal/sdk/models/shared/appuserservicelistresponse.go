// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// Expanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type Expanded struct {
}

// The AppUserServiceListResponse message.
type AppUserServiceListResponse struct {
	// The expanded field.
	Expanded []Expanded `json:"expanded,omitempty"`
	// The list field.
	List []AppUserView `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *AppUserServiceListResponse) GetExpanded() []Expanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *AppUserServiceListResponse) GetList() []AppUserView {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *AppUserServiceListResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
