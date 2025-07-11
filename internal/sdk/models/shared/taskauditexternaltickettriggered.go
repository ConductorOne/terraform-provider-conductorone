// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskAuditExternalTicketTriggered message.
type TaskAuditExternalTicketTriggered struct {
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The connectorId field.
	ConnectorID *string `json:"connectorId,omitempty"`
	// The externalTicketId field.
	ExternalTicketID *string `json:"externalTicketId,omitempty"`
	// The externalTicketProvisionerConfigId field.
	ExternalTicketProvisionerConfigID *string `json:"externalTicketProvisionerConfigId,omitempty"`
	// The externalTicketProvisionerConfigName field.
	ExternalTicketProvisionerConfigName *string `json:"externalTicketProvisionerConfigName,omitempty"`
}

func (o *TaskAuditExternalTicketTriggered) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *TaskAuditExternalTicketTriggered) GetConnectorID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectorID
}

func (o *TaskAuditExternalTicketTriggered) GetExternalTicketID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalTicketID
}

func (o *TaskAuditExternalTicketTriggered) GetExternalTicketProvisionerConfigID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalTicketProvisionerConfigID
}

func (o *TaskAuditExternalTicketTriggered) GetExternalTicketProvisionerConfigName() *string {
	if o == nil {
		return nil
	}
	return o.ExternalTicketProvisionerConfigName
}
