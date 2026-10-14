package provider

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func aliasSchema(t *testing.T) schema.Schema {
	t.Helper()
	resp := &resource.SchemaResponse{}
	NewCustomAppEntitlementAliasResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("schema: %v", resp.Diagnostics)
	}
	if diags := resp.Schema.ValidateImplementation(context.Background()); diags.HasError() {
		t.Fatalf("schema does not validate: %v", diags)
	}
	return resp.Schema
}

func TestProvisionerPolicyIsDeprecatedComputedMirror(t *testing.T) {
	s := aliasSchema(t)
	policy := s.Attributes[provisionPolicyAttr].(schema.SingleNestedAttribute)
	mirror, ok := s.Attributes[provisionerPolicyAttr].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("%s missing or not a nested attribute", provisionerPolicyAttr)
	}
	if !mirror.Computed || mirror.Optional || mirror.Required {
		t.Errorf("provisioner_policy must be computed only: %+v", mirror)
	}
	if mirror.DeprecationMessage == "" {
		t.Error("provisioner_policy has no deprecation message")
	}
	if len(mirror.PlanModifiers) != 0 || len(mirror.Validators) != 0 {
		t.Error("provisioner_policy must not carry plan modifiers or validators")
	}
	// Same type means the mirror can hold provision_policy's value unchanged and
	// state written by earlier provider versions keeps decoding.
	if !mirror.GetType().Equal(policy.GetType()) {
		t.Errorf("provisioner_policy type differs from provision_policy:\n%s\n%s", mirror.GetType(), policy.GetType())
	}
	if !policy.Optional || !policy.Computed {
		t.Errorf("provision_policy must stay writable (optional+computed): %+v", policy)
	}
}

// The generated resource's schema must be untouched, so its decode structs keep working.
func TestAliasWrapperLeavesGeneratedSchemaUnchanged(t *testing.T) {
	inner := &resource.SchemaResponse{}
	NewCustomAppEntitlementResource().Schema(context.Background(), resource.SchemaRequest{}, inner)
	if _, ok := inner.Schema.Attributes[provisionerPolicyAttr]; ok {
		t.Fatal("generated schema already defines provisioner_policy; the wrapper is redundant")
	}
	outer := aliasSchema(t)
	if got, want := len(outer.Attributes), len(inner.Schema.Attributes)+1; got != want {
		t.Fatalf("outer has %d attributes, want %d", got, want)
	}
}

func TestStateConversionRoundTrip(t *testing.T) {
	ctx := context.Background()
	outer := aliasSchema(t)
	innerResp := &resource.SchemaResponse{}
	NewCustomAppEntitlementResource().Schema(ctx, resource.SchemaRequest{}, innerResp)
	inner := innerResp.Schema

	// Outer null stays null both ways.
	got, diags := toInner(ctx, inner, tftypes.NewValue(outer.Type().TerraformType(ctx), nil))
	if diags.HasError() || !got.IsNull() {
		t.Fatalf("null outer: %v %v", got, diags)
	}

	// Inner value with a provision_policy gains an identical provisioner_policy.
	innerType := inner.Type().TerraformType(ctx).(tftypes.Object)
	vals := map[string]tftypes.Value{}
	for name, typ := range innerType.AttributeTypes {
		vals[name] = tftypes.NewValue(typ, nil)
	}
	policyType := innerType.AttributeTypes[provisionPolicyAttr].(tftypes.Object)
	pvals := map[string]tftypes.Value{}
	for name, typ := range policyType.AttributeTypes {
		pvals[name] = tftypes.NewValue(typ, nil)
	}
	vals[provisionPolicyAttr] = tftypes.NewValue(policyType, pvals)

	state := tfsdk.State{Schema: outer}
	if diags := setOuterState(ctx, &state, tftypes.NewValue(innerType, vals)); diags.HasError() {
		t.Fatal(diags)
	}
	var attrs map[string]tftypes.Value
	if err := state.Raw.As(&attrs); err != nil {
		t.Fatal(err)
	}
	if !attrs[provisionerPolicyAttr].Equal(attrs[provisionPolicyAttr]) {
		t.Error("provisioner_policy does not mirror provision_policy")
	}

	back, diags := toInner(ctx, inner, state.Raw)
	if diags.HasError() {
		t.Fatal(diags)
	}
	if !back.Equal(tftypes.NewValue(innerType, vals)) {
		t.Error("round trip changed the inner value")
	}
}

// policyValue builds a value of the inner schema with the named policy holding
// the given arm values (arm name to value); every other attribute is null.
func policyValue(t *testing.T, policy string, arms map[string]string) tftypes.Value {
	t.Helper()
	ctx := context.Background()
	innerResp := &resource.SchemaResponse{}
	NewCustomAppEntitlementResource().Schema(ctx, resource.SchemaRequest{}, innerResp)
	innerType := innerResp.Schema.Type().TerraformType(ctx).(tftypes.Object)

	vals := map[string]tftypes.Value{}
	for name, typ := range innerType.AttributeTypes {
		vals[name] = tftypes.NewValue(typ, nil)
	}
	policyType := innerType.AttributeTypes[policy].(tftypes.Object)
	armVals := map[string]tftypes.Value{}
	for name, typ := range policyType.AttributeTypes {
		armVals[name] = tftypes.NewValue(typ, nil)
	}
	for name, marker := range arms {
		// Only nullness matters to dropStaleArms; a string arm carries the marker,
		// an object arm is an object with all-null attributes.
		switch typ := policyType.AttributeTypes[name].(type) {
		case tftypes.Object:
			attrs := map[string]tftypes.Value{}
			for n, at := range typ.AttributeTypes {
				attrs[n] = tftypes.NewValue(at, nil)
			}
			armVals[name] = tftypes.NewValue(typ, attrs)
		default:
			armVals[name] = tftypes.NewValue(typ, marker)
		}
	}
	vals[policy] = tftypes.NewValue(policyType, armVals)
	return tftypes.NewValue(innerType, vals)
}

func nonNullArms(t *testing.T, v tftypes.Value, policy string) []string {
	t.Helper()
	attrs, err := objectAttrs(v)
	if err != nil {
		t.Fatal(err)
	}
	arms, err := objectAttrs(attrs[policy])
	if err != nil {
		t.Fatal(err)
	}
	var out []string
	for name, arm := range arms {
		if !arm.IsNull() {
			out = append(out, name)
		}
	}
	sort.Strings(out)
	return out
}

func TestDropStaleArms(t *testing.T) {
	for _, policy := range []string{provisionPolicyAttr, deprovisionerPolicyAttr} {
		t.Run(policy, func(t *testing.T) {
			stored := policyValue(t, policy, map[string]string{"multi_step": `{"provisionSteps":[]}`, "delegated": ""})

			got, err := dropStaleArms(policyValue(t, policy, map[string]string{"delegated": ""}), stored, policy)
			if err != nil {
				t.Fatal(err)
			}
			if arms := nonNullArms(t, got, policy); strings.Join(arms, ",") != "delegated" {
				t.Errorf("config selects delegated; arms left in state = %v", arms)
			}

			got, err = dropStaleArms(policyValue(t, policy, map[string]string{"multi_step": `{"provisionSteps":[]}`}), stored, policy)
			if err != nil {
				t.Fatal(err)
			}
			if arms := nonNullArms(t, got, policy); strings.Join(arms, ",") != "multi_step" {
				t.Errorf("config selects multi_step; arms left in state = %v", arms)
			}

			// No arm in configuration: nothing was chosen, the stored arms stay.
			got, err = dropStaleArms(policyValue(t, policy, nil), stored, policy)
			if err != nil {
				t.Fatal(err)
			}
			if !got.Equal(stored) {
				t.Error("state changed although configuration names no arm")
			}
		})
	}
}
