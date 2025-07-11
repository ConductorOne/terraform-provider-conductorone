// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ConnectorView object provides a connector response object, as well as JSONPATHs to related objects provided by expanders.
type ConnectorView struct {
	// JSONPATH expression indicating the location of the App object in the expanded array.
	AppPath *string `json:"appPath,omitempty"`
	// A Connector is used to sync objects into Apps
	Connector *Connector `json:"connector,omitempty"`
	// JSONPATH expression indicating the location of the User object in the expanded array. This is the user that is a direct target of the ticket without a specific relationship to a potentially non-existent app user.
	UsersPath *string `json:"usersPath,omitempty"`
}

func (o *ConnectorView) GetAppPath() *string {
	if o == nil {
		return nil
	}
	return o.AppPath
}

func (o *ConnectorView) GetConnector() *Connector {
	if o == nil {
		return nil
	}
	return o.Connector
}

func (o *ConnectorView) GetUsersPath() *string {
	if o == nil {
		return nil
	}
	return o.UsersPath
}
