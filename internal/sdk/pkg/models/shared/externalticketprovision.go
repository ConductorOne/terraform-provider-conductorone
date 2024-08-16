// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ExternalTicketProvision - This provision step indicates that we should check an external ticket to provision this entitlement
type ExternalTicketProvision struct {
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The connectorId field.
	ConnectorID *string `json:"connectorId,omitempty"`
	// The externalTicketProvisionerConfigId field.
	ExternalTicketProvisionerConfigID *string `json:"externalTicketProvisionerConfigId,omitempty"`
	// This field indicates a text body of instructions for the provisioner to indicate.
	Instructions *string `json:"instructions,omitempty"`
}

func (o *ExternalTicketProvision) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *ExternalTicketProvision) GetConnectorID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectorID
}

func (o *ExternalTicketProvision) GetExternalTicketProvisionerConfigID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalTicketProvisionerConfigID
}

func (o *ExternalTicketProvision) GetInstructions() *string {
	if o == nil {
		return nil
	}
	return o.Instructions
}
