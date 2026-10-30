package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
)

// Destroy semantics for rows that may be backed by a connector (for example,
// rows adopted via match_baton_id or later merged with a connector-synced row).
//
// ConductorOne refuses to delete connector-synced rows, so a blind DELETE makes
// `terraform destroy` fail. Before deleting, the generated Delete functions
// call the helpers below, which re-read the row and decide:
//
//   - 404 or deleted_at set: already gone; skip DELETE, succeed.
//   - is_manually_managed == false: connector-managed; RELEASE it (skip DELETE,
//     warn, succeed so the framework drops it from state). The row in
//     ConductorOne is left untouched.
//   - is_manually_managed == true, or absent from the response: DELETE as before.
//   - any other read failure: surface the error.
//
// Hooked into the generated code by patches/13-release-connector-managed-on-destroy.patch.

const (
	connectorManagedReleaseSummary = "Resource released, not deleted"
	connectorManagedReleaseDetail  = "%s %q is managed by a connector in ConductorOne, so it cannot be deleted. " +
		"It was removed from Terraform state only; the row in ConductorOne was left unchanged and the connector keeps it in sync. " +
		"Applying this configuration again adopts the same ID."
)

func connectorManagedReleaseWarning(kind, id string) diag.Diagnostics {
	var diags diag.Diagnostics
	diags.AddWarning(connectorManagedReleaseSummary, fmt.Sprintf(connectorManagedReleaseDetail, kind, id))
	return diags
}

// releaseConnectorManagedAppResource returns skipDelete=true when the DELETE
// call must not be made (row already gone, or connector-managed and released).
func releaseConnectorManagedAppResource(ctx context.Context, client *sdk.ConductoroneAPI, data *AppResourceResourceModel) (skipDelete bool, diags diag.Diagnostics) {
	request, requestDiags := data.ToOperationsC1APIAppV1AppResourceServiceGetRequest(ctx)
	diags.Append(requestDiags...)
	if diags.HasError() {
		return false, diags
	}
	res, err := client.AppResource.Get(ctx, *request)
	if err != nil {
		diags.AddError("failure to read app resource before delete", err.Error())
		if res != nil && res.RawResponse != nil {
			diags.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return false, diags
	}
	if res == nil {
		diags.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return false, diags
	}
	if res.StatusCode == 404 {
		return true, diags
	}
	if res.StatusCode != 200 {
		diags.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return false, diags
	}
	row := res.GetAppResourceServiceGetResponse().GetAppResourceView().GetAppResource()
	if row == nil {
		diags.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return false, diags
	}
	if row.GetDeletedAt() != nil {
		return true, diags
	}
	if mm := row.GetIsManuallyManaged(); mm != nil && !*mm {
		diags.Append(connectorManagedReleaseWarning("App resource", data.ID.ValueString())...)
		return true, diags
	}
	return false, diags
}

// releaseConnectorManagedAppEntitlement is the conductorone_custom_app_entitlement
// counterpart of releaseConnectorManagedAppResource.
func releaseConnectorManagedAppEntitlement(ctx context.Context, client *sdk.ConductoroneAPI, data *CustomAppEntitlementResourceModel) (skipDelete bool, diags diag.Diagnostics) {
	request, requestDiags := data.ToOperationsC1APIAppV1AppEntitlementsGetRequest(ctx)
	diags.Append(requestDiags...)
	if diags.HasError() {
		return false, diags
	}
	res, err := client.AppEntitlements.Get(ctx, *request)
	if err != nil {
		diags.AddError("failure to read app entitlement before delete", err.Error())
		if res != nil && res.RawResponse != nil {
			diags.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return false, diags
	}
	if res == nil {
		diags.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return false, diags
	}
	if res.StatusCode == 404 {
		return true, diags
	}
	if res.StatusCode != 200 {
		diags.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return false, diags
	}
	row := res.GetAppEntitlementResponse.GetAppEntitlementView().GetAppEntitlement()
	if row == nil {
		diags.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return false, diags
	}
	if row.GetDeletedAt() != nil {
		return true, diags
	}
	if mm := row.GetIsManuallyManaged(); mm != nil && !*mm {
		diags.Append(connectorManagedReleaseWarning("App entitlement", data.ID.ValueString())...)
		return true, diags
	}
	return false, diags
}
