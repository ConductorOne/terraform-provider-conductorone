// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The UpdateAppEntitlementRequest message contains the app entitlement and the fields to be updated.
type UpdateAppEntitlementRequest struct {
	// The app entitlement represents one permission in a downstream App (SAAS) that can be granted. For example, GitHub Read vs GitHub Write.
	//
	// This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
	//   - durationUnset
	//   - durationGrant
	//
	AppEntitlement *AppEntitlementInput `json:"entitlement,omitempty"`
	// The app entitlement expand mask allows the user to get additional information when getting responses containing app entitlement views.
	AppEntitlementExpandMask *AppEntitlementExpandMask `json:"expandMask,omitempty"`
}

func (o *UpdateAppEntitlementRequest) GetAppEntitlement() *AppEntitlementInput {
	if o == nil {
		return nil
	}
	return o.AppEntitlement
}

func (o *UpdateAppEntitlementRequest) GetAppEntitlementExpandMask() *AppEntitlementExpandMask {
	if o == nil {
		return nil
	}
	return o.AppEntitlementExpandMask
}
