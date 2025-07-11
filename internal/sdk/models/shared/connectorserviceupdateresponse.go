// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ConnectorServiceUpdateResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type ConnectorServiceUpdateResponseExpanded struct {
}

// ConnectorServiceUpdateResponse is the response returned by the update method.
type ConnectorServiceUpdateResponse struct {
	// The ConnectorView object provides a connector response object, as well as JSONPATHs to related objects provided by expanders.
	ConnectorView *ConnectorView `json:"connectorView,omitempty"`
	// The array of expanded items indicated by the request.
	Expanded []ConnectorServiceUpdateResponseExpanded `json:"expanded,omitempty"`
}

func (o *ConnectorServiceUpdateResponse) GetConnectorView() *ConnectorView {
	if o == nil {
		return nil
	}
	return o.ConnectorView
}

func (o *ConnectorServiceUpdateResponse) GetExpanded() []ConnectorServiceUpdateResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
