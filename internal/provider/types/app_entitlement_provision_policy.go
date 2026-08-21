package types

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AppEntitlementProvisionPolicy struct {
	ActionProvision         *ActionProvision                  `tfsdk:"action_provision"`
	ConnectorProvision      *AppEntitlementConnectorProvision `tfsdk:"connector_provision"`
	DelegatedProvision      *DelegatedProvision               `tfsdk:"delegated_provision"`
	ExternalTicketProvision *ExternalTicketProvision          `tfsdk:"external_ticket_provision"`
	ManualProvision         *AppEntitlementManualProvision    `tfsdk:"manual_provision"`
	MultiStep               jsontypes.Normalized              `tfsdk:"multi_step"`
	UnconfiguredProvision   *UnconfiguredProvision            `tfsdk:"unconfigured_provision"`
	WebhookProvision        *WebhookProvision                 `tfsdk:"webhook_provision"`
}

type AppEntitlementConnectorProvision struct {
	AccountProvision *AccountProvision `tfsdk:"account_provision"`
	DefaultBehavior  *DefaultBehavior  `tfsdk:"default_behavior"`
	DeleteAccount    *DeleteAccount    `tfsdk:"delete_account"`
}

type AppEntitlementManualProvision struct {
	ProvisionerAssignment *AppEntitlementProvisionerAssignment `tfsdk:"provisioner_assignment"`
	Instructions          types.String                         `tfsdk:"instructions"`
	UserIds               []types.String                       `tfsdk:"user_ids"`
}

type AppEntitlementProvisionerAssignment struct {
	AppOwnerProvisioner         *AppOwnerProvisioner         `tfsdk:"app_owner_provisioner"`
	EntitlementOwnerProvisioner *EntitlementOwnerProvisioner `tfsdk:"entitlement_owner_provisioner"`
	ExpressionProvisioner       *ExpressionProvisioner       `tfsdk:"expression_provisioner"`
	GroupProvisioner            *GroupProvisioner            `tfsdk:"group_provisioner"`
	ManagerProvisioner          *ManagerProvisioner          `tfsdk:"manager_provisioner"`
	UserProvisioner             *UserProvisioner             `tfsdk:"user_provisioner"`
}
