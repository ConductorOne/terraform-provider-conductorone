// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppEntitlementAutomationRuleBasic message.
type AppEntitlementAutomationRuleBasic struct {
	// The expression field.
	Expression *string `json:"expression,omitempty"`
}

func (o *AppEntitlementAutomationRuleBasic) GetExpression() *string {
	if o == nil {
		return nil
	}
	return o.Expression
}
