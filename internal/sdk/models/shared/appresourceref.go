// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppResourceRef message.
type AppResourceRef struct {
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The appResourceTypeId field.
	AppResourceTypeID *string `json:"appResourceTypeId,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
}

func (o *AppResourceRef) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppResourceRef) GetAppResourceTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeID
}

func (o *AppResourceRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
