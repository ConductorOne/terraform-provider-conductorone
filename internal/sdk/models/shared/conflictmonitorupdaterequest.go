// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ConflictMonitorUpdateRequest message.
type ConflictMonitorUpdateRequest struct {
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	// The NotificationConfig message.
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`
}

func (o *ConflictMonitorUpdateRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ConflictMonitorUpdateRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *ConflictMonitorUpdateRequest) GetNotificationConfig() *NotificationConfig {
	if o == nil {
		return nil
	}
	return o.NotificationConfig
}
