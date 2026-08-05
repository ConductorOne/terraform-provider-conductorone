package provider

import (
	"context"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestFunctionUpdateMaskForChangesIncludesOnlyChangedSerializedField(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{
			"display_name": types.StringValue("before"),
			"description":  types.StringValue("unchanged"),
		},
		map[string]attr.Value{
			"display_name": types.StringValue("after"),
			"description":  types.StringValue("unchanged"),
		},
	)

	mask, diags := functionUpdateMaskForChanges(state, plan, &shared.FunctionInput{
		DisplayName: stringPointer("after"),
		Description: stringPointer("unchanged"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got, want := *mask, "displayName"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func TestAccessReviewUpdateMaskExcludesAbsentAndUnsupportedFields(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{
			"description":         types.StringValue("before"),
			"notification_config": types.ObjectNull(map[string]attr.Type{}),
		},
		map[string]attr.Value{
			"description":         types.StringValue("after"),
			"notification_config": types.ObjectNull(map[string]attr.Type{}),
		},
	)

	mask, diags := accessReviewUpdateMaskForChanges(state, plan, &shared.AccessReviewInput{
		Description: stringPointer("after"),
	})
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got, want := *mask, "description"; got != want {
		t.Errorf("mask = %q, want %q", got, want)
	}
}

func TestUpdateMaskRejectsChangedFieldMissingFromPayload(t *testing.T) {
	state, plan := updateMaskObjects(
		map[string]attr.Value{"description": types.StringValue("before")},
		map[string]attr.Value{"description": types.StringNull()},
	)

	mask, diags := functionUpdateMaskForChanges(state, plan, &shared.FunctionInput{})
	if mask != nil {
		t.Errorf("mask = %q, want nil", *mask)
	}
	if !diags.HasError() {
		t.Fatal("expected diagnostics when a changed field is omitted from the payload")
	}
}

func TestPayloadMasksContainOnlySerializedBackendPaths(t *testing.T) {
	functionMask, functionDiags := functionUpdateMask(&shared.FunctionInput{
		DisplayName: stringPointer("function"),
	})
	if functionDiags.HasError() {
		t.Fatalf("unexpected function diagnostics: %v", functionDiags)
	}
	if got, want := *functionMask, "displayName"; got != want {
		t.Errorf("function mask = %q, want %q", got, want)
	}

	accessReviewMask, accessReviewDiags := accessReviewUpdateMask(&shared.AccessReviewInput{
		Description: stringPointer("campaign"),
	})
	if accessReviewDiags.HasError() {
		t.Fatalf("unexpected access review diagnostics: %v", accessReviewDiags)
	}
	if got, want := *accessReviewMask, "description"; got != want {
		t.Errorf("access review mask = %q, want %q", got, want)
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
