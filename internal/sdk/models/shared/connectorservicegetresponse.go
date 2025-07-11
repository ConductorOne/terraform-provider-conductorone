// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ConnectorServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type ConnectorServiceGetResponseExpanded struct {
}

// The ConnectorServiceGetResponse message contains the connectorView, and an expand mask.
type ConnectorServiceGetResponse struct {
	// The ConnectorView object provides a connector response object, as well as JSONPATHs to related objects provided by expanders.
	ConnectorView *ConnectorView `json:"connectorView,omitempty"`
	// The array of expanded items indicated by the request.
	Expanded []ConnectorServiceGetResponseExpanded `json:"expanded,omitempty"`
}

func (o *ConnectorServiceGetResponse) GetConnectorView() *ConnectorView {
	if o == nil {
		return nil
	}
	return o.ConnectorView
}

func (o *ConnectorServiceGetResponse) GetExpanded() []ConnectorServiceGetResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
