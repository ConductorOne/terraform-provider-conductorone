// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type PolicyStep struct {
	Accept    *Accept `tfsdk:"accept"`
	Approval  *ApprovalInput `tfsdk:"approval"`
	Provision *Provision     `tfsdk:"provision"`
	Reject    *Reject `tfsdk:"reject"`
}
