package provider

import (
	"context"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestUpdateMaskRejectsChangedFieldMissingFromPayload(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{"description": types.StringValue("before")},
		map[string]attr.Value{"description": types.StringNull()},
	)

	mask, diags := appResourceUpdateMaskForChanges(state, plan, &shared.AppResourceInput{})
	if mask != nil {
		t.Errorf("mask = %q, want nil", *mask)
	}
	if !diags.HasError() {
		t.Fatal("expected diagnostics when a changed field is omitted from the payload")
	}
}

func TestPayloadMasksContainOnlySerializedBackendPaths(t *testing.T) {
	mask, diags := vaultUpdateMask(&shared.VaultInput{
		DisplayName: stringPointer("vault"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	// description is absent from the payload, so it must stay out of the mask
	// even though vaultUpdateMask lists it.
	if got, want := *mask, "displayName"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func TestAppResourceUpdateMaskPreservesFieldsTheSchemaDoesNotModel(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{
			"display_name": types.StringValue("before"),
			"description":  types.StringValue("unchanged"),
			"annotations":  types.MapNull(types.StringType),
		},
		map[string]attr.Value{
			"display_name": types.StringValue("after"),
			"description":  types.StringValue("unchanged"),
			"annotations":  types.MapNull(types.StringType),
		},
	)

	mask, diags := appResourceUpdateMaskForChanges(state, plan, &shared.AppResourceInput{
		DisplayName: stringPointer("after"),
		Description: stringPointer("unchanged"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	// customDescription and accessConfigId must stay out of the mask. The
	// backend assigns custom_description unconditionally from the merged
	// payload, so a mask that omits it keeps the stored value while a nil mask
	// clears it; access_config_id is computed-only in the schema.
	if got, want := *mask, "displayName"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func TestVaultUpdateMaskAlwaysCarriesTheMirroredPaths(t *testing.T) {
	mask, diags := vaultUpdateMask(&shared.VaultInput{
		DisplayName: stringPointer("vault"),
		Description: stringPointer("secrets for the deploy pipeline"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	// The vault update RPC copies display_name and description onto the vault's
	// app resource and owner entitlement through a mask filtered to exactly
	// these two paths. Dropping either one from the request mask leaves that
	// filtered mask empty, which full-replaces the owner entitlement from a
	// two-field stub and clears its policy bindings.
	if got, want := *mask, "displayName,description"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func TestUpdateMaskTreatsUnknownPlannedValueAsAChange(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{
			"display_name": types.StringValue("unchanged"),
			"description":  types.StringValue("stored"),
		},
		map[string]attr.Value{
			"display_name": types.StringValue("unchanged"),
			"description":  types.StringUnknown(),
		},
	)

	mask, diags := updateMaskForChanges(state, plan, []updateMaskField{
		{terraformName: "display_name", apiPath: "displayName", serialized: true},
		{terraformName: "description", apiPath: "description", serialized: true},
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	// An unknown planned value is not equal to the stored value, so it lands in
	// the mask. Terraform marks computed attributes unknown on every plan even
	// when the practitioner changed nothing, so callers must not list
	// computed-only attributes: doing so produces a mask on every apply and can
	// crowd out the paths that actually needed to be there.
	if got, want := *mask, "description"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func updateMaskObjects(state, plan map[string]attr.Value) (types.Object, types.Object) {
	attributeTypes := make(map[string]attr.Type, len(state))
	for name, value := range state {
		attributeTypes[name] = value.Type(context.Background())
	}
	return types.ObjectValueMust(attributeTypes, state), types.ObjectValueMust(attributeTypes, plan)
}

func stringPointer(value string) *string {
	return &value
}

func TestAppResourceUpdateMaskClearsAnnotations(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{
			"display_name": types.StringValue("unchanged"),
			"annotations":  types.MapValueMust(types.StringType, map[string]attr.Value{"managed_by": types.StringValue("terraform")}),
		},
		map[string]attr.Value{
			"display_name": types.StringValue("unchanged"),
			"annotations":  types.MapValueMust(types.StringType, map[string]attr.Value{}),
		},
	)

	mask, diags := appResourceUpdateMaskForChanges(state, plan, &shared.AppResourceInput{
		DisplayName: stringPointer("unchanged"),
		Annotations: map[string]string{},
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	// Emptying the map is the practitioner clearing it. omitempty drops the
	// field from the JSON body, so the mask is the only thing carrying that
	// intent; without the path the stored annotations survive the update and
	// the next refresh drifts them back into state.
	if got, want := *mask, "annotations"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func TestAppResourceUpdateMaskClearsAnnotationsAlongsideAnotherChange(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{
			"display_name": types.StringValue("before"),
			"annotations":  types.MapValueMust(types.StringType, map[string]attr.Value{"managed_by": types.StringValue("terraform")}),
		},
		map[string]attr.Value{
			"display_name": types.StringValue("after"),
			"annotations":  types.MapValueMust(types.StringType, map[string]attr.Value{}),
		},
	)

	mask, diags := appResourceUpdateMaskForChanges(state, plan, &shared.AppResourceInput{
		DisplayName: stringPointer("after"),
		Annotations: map[string]string{},
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got, want := *mask, "displayName,annotations"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func TestAppResourceUpdateMaskLeavesUnknownAnnotationsAlone(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{
			"display_name": types.StringValue("before"),
			"annotations":  types.MapNull(types.StringType),
		},
		map[string]attr.Value{
			"display_name": types.StringValue("after"),
			"annotations":  types.MapUnknown(types.StringType),
		},
	)

	mask, diags := appResourceUpdateMaskForChanges(state, plan, &shared.AppResourceInput{
		DisplayName: stringPointer("after"),
		Annotations: map[string]string{},
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	// An unconfigured Computed attribute plans as unknown. Masking it would
	// clear whatever the remote holds, which the practitioner never asked for.
	if got, want := *mask, "displayName"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}
