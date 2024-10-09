// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ProvisionPolicy struct {
	ConnectorProvision      *ConnectorProvision      `tfsdk:"connector_provision"`
	DelegatedProvision      *DelegatedProvision      `tfsdk:"delegated_provision"`
	ExternalTicketProvision *ExternalTicketProvision `tfsdk:"external_ticket_provision"`
	ManualProvision         *ManualProvision         `tfsdk:"manual_provision"`
	MultiStep               types.String             `tfsdk:"multi_step"`
	WebhookProvision        *WebhookProvision        `tfsdk:"webhook_provision"`
}
