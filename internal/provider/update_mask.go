package provider

import (
	"strings"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type updateMaskField struct {
	terraformName string
	apiPath       string
	serialized    bool
}

// updateMaskForChanges builds a mask from fields which both changed in the
// Terraform plan and are present in the JSON payload. Update APIs reject an
// empty mask, while including an omitted field in a mask can overwrite it.
func updateMaskForChanges(state, plan types.Object, fields []updateMaskField) (*string, diag.Diagnostics) {
	var diags diag.Diagnostics
	stateAttributes := state.Attributes()
	planAttributes := plan.Attributes()
	paths := make([]string, 0, len(fields))

	for _, field := range fields {
		stateValue, stateOK := stateAttributes[field.terraformName]
		planValue, planOK := planAttributes[field.terraformName]
		if !stateOK || !planOK || planValue.Equal(stateValue) {
			continue
		}
		if field.serialized {
			paths = append(paths, field.apiPath)
		}
	}

	if len(paths) == 0 {
		diags.AddError(
			"Unable to build update mask",
			"The planned changes do not include a field that this provider can safely serialize for this update. No request was sent.",
		)
		return nil, diags
	}

	mask := strings.Join(paths, ",")
	return &mask, diags
}

func appResourceUpdateMaskForChanges(state, plan types.Object, input *shared.AppResourceInput) (*string, diag.Diagnostics) {
	if input == nil {
		return updateMaskForChanges(state, plan, nil)
	}

	return updateMaskForChanges(state, plan, []updateMaskField{
		{terraformName: "display_name", apiPath: "displayName", serialized: input.DisplayName != nil},
		{terraformName: "description", apiPath: "description", serialized: input.Description != nil},
		{terraformName: "annotations", apiPath: "annotations", serialized: plannedValueIsSet(plan, "annotations")},
	})
}

// plannedValueIsSet reports whether the practitioner configured a concrete
// value for the attribute. Collections need this instead of a length check on
// the payload: an emptied map is a real update, and omitempty drops it from the
// JSON body, so the mask is the only thing that still carries the intent to
// clear it.
func plannedValueIsSet(plan types.Object, terraformName string) bool {
	value, ok := plan.Attributes()[terraformName]
	return ok && !value.IsNull() && !value.IsUnknown()
}

// vaultUpdateMask is deliberately static rather than change-aware. The vault
// update RPC patches only display_name and description, and it mirrors both onto
// the vault's C1 app resource and owner entitlement through a mask filtered to
// those two paths — a mask that omits them clears unrelated owner-entitlement
// policy bindings. credential_expiration_duration is computed-only in the
// schema, and the vault-type oneof plans as a permanent diff because Read never
// populates it, so a plan/state comparison would routinely produce a mask
// without display_name or description.
func vaultUpdateMask(input *shared.VaultInput) (*string, diag.Diagnostics) {
	if input == nil {
		return updateMaskForSerializedFields(nil)
	}

	return updateMaskForSerializedFields([]updateMaskField{
		{apiPath: "displayName", serialized: input.DisplayName != nil},
		{apiPath: "description", serialized: input.Description != nil},
	})
}

func updateMaskForSerializedFields(fields []updateMaskField) (*string, diag.Diagnostics) {
	var diags diag.Diagnostics
	paths := make([]string, 0, len(fields))
	for _, field := range fields {
		if field.serialized {
			paths = append(paths, field.apiPath)
		}
	}
	if len(paths) == 0 {
		diags.AddError(
			"Unable to build update mask",
			"The update payload has no maskable fields. No request was sent.",
		)
		return nil, diags
	}
	mask := strings.Join(paths, ",")
	return &mask, diags
}
