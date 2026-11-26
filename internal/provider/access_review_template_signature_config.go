package provider

import (
	"context"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// addSignatureConfigRemovalToUpdateMask names signatureConfig in the update mask
// when the plan removes a block that prior state still holds, so the API clears
// it. The generated mask names signatureConfig only when the plan sets it, because
// naming it unconditionally trips the e-signature feature gate on tenants without
// that feature.
func addSignatureConfigRemovalToUpdateMask(
	ctx context.Context,
	state tfsdk.State,
	planned *AccessReviewTemplateResourceModel,
	request *operations.C1APIAccessreviewV1AccessReviewTemplateServiceUpdateRequest,
) diag.Diagnostics {
	if planned == nil || planned.SignatureConfig != nil || request == nil || request.AccessReviewTemplateServiceUpdateRequest == nil {
		return nil
	}

	var prior types.Object
	diags := state.GetAttribute(ctx, path.Root("signature_config"), &prior)
	if diags.HasError() || prior.IsNull() || prior.IsUnknown() {
		return diags
	}

	body := request.AccessReviewTemplateServiceUpdateRequest
	mask := "signatureConfig"
	if body.UpdateMask != nil && *body.UpdateMask != "" {
		mask = *body.UpdateMask + "," + mask
	}
	body.UpdateMask = &mask
	return diags
}
