// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BundleAutomationRuleEntitlement - The BundleAutomationRuleEntitlement message.
type BundleAutomationRuleEntitlement struct {
	// The entitlementRefs field.
	EntitlementRefs []AppEntitlementRef `json:"entitlementRefs,omitempty"`
}

func (o *BundleAutomationRuleEntitlement) GetEntitlementRefs() []AppEntitlementRef {
	if o == nil {
		return nil
	}
	return o.EntitlementRefs
}
