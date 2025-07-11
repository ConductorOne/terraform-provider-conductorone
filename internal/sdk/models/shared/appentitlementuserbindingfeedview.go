// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppEntitlementUserBindingFeedView message.
type AppEntitlementUserBindingFeedView struct {
	// The appPath field.
	AppPath *string `json:"appPath,omitempty"`
	// The appUserPath field.
	AppUserPath *string `json:"appUserPath,omitempty"`
	// The entitlementPath field.
	EntitlementPath *string `json:"entitlementPath,omitempty"`
	// The AppEntitlementUserBindingFeed message.
	AppEntitlementUserBindingFeed *AppEntitlementUserBindingFeed `json:"feed,omitempty"`
	// The ticketPath field.
	TicketPath *string `json:"ticketPath,omitempty"`
}

func (o *AppEntitlementUserBindingFeedView) GetAppPath() *string {
	if o == nil {
		return nil
	}
	return o.AppPath
}

func (o *AppEntitlementUserBindingFeedView) GetAppUserPath() *string {
	if o == nil {
		return nil
	}
	return o.AppUserPath
}

func (o *AppEntitlementUserBindingFeedView) GetEntitlementPath() *string {
	if o == nil {
		return nil
	}
	return o.EntitlementPath
}

func (o *AppEntitlementUserBindingFeedView) GetAppEntitlementUserBindingFeed() *AppEntitlementUserBindingFeed {
	if o == nil {
		return nil
	}
	return o.AppEntitlementUserBindingFeed
}

func (o *AppEntitlementUserBindingFeedView) GetTicketPath() *string {
	if o == nil {
		return nil
	}
	return o.TicketPath
}
