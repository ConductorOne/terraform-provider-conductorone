// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppEntitlementAutomationRuleCEL message.
type AppEntitlementAutomationRuleCEL struct {
	// The expression field.
	Expression *string `json:"expression,omitempty"`
}

func (o *AppEntitlementAutomationRuleCEL) GetExpression() *string {
	if o == nil {
		return nil
	}
	return o.Expression
}
