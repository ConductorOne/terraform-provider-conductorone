// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ConnectorView struct {
	Connector *Connector   `tfsdk:"connector"`
	AppPath   types.String `tfsdk:"app_path"`
	UsersPath types.String `tfsdk:"users_path"`
}
