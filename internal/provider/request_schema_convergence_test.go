package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// suppressDescFragment is the human-readable description emitted by the in-repo
// SuppressDiff plan modifiers (internal/planmodifiers/*/suppress_diff.go):
//
//	"Once set, the value of this attribute in state will not change."
//
// We match on a stable fragment so the assertion survives minor wording changes.
const suppressDescFragment = "will not change"

// TestRequestSchemaComputedAttrsArePlanStable is the IGA-743 *convergence*
// tripwire (distinct from the crash tripwires in
// request_schema_schema_symmetry_test.go / request_schema_roundtrip_test.go).
//
// The crash fix on PR #232 stopped the "Value Conversion Error", but a
// conductorone_request_schema with a bool_field still did NOT reach an empty
// plan: several server-derived computed attributes re-planned as
// "(known after apply)" on every plan after the first apply, so a config that
// already matched the read-back never converged.
//
// The fix marks each drifting attribute plan-stable via
// x-speakeasy-param-suppress-computed-diff in overlay.yaml — exactly the
// treatment already applied to RequestSchema.id. Speakeasy regenerates that
// into a SuppressDiff(ExplicitSuppress) plan modifier, which propagates the
// prior state value into the plan when the planned value is unknown (the
// framework's UseStateForUnknown behavior) and is a no-op when the user
// explicitly sets the attribute.
//
// This test FAILS on the broken HEAD (these attributes carry no plan modifier)
// and PASSES once the overlay change is regenerated.
func TestRequestSchemaComputedAttrsArePlanStable(t *testing.T) {
	ctx := context.Background()
	resp := &resource.SchemaResponse{}
	NewRequestSchemaResource().Schema(ctx, resource.SchemaRequest{}, resp)
	attrs := resp.Schema.Attributes

	// Top-level server-derived computed attributes.
	assertStringAttrStable(t, attrs, "created_at")
	assertStringAttrStable(t, attrs, "justification_visibility")
	assertListAttrStable(t, attrs, "field_groups")
	assertListAttrStable(t, attrs, "field_relationships")

	// fields[].name — the server derives an internal field name, so the nested
	// "name" attribute churns unless it is plan-stable.
	fields, ok := attrs["fields"].(schema.ListNestedAttribute)
	if !ok {
		t.Fatalf(`"fields" is %T, want schema.ListNestedAttribute`, attrs["fields"])
	}
	nameAttr, ok := fields.NestedObject.Attributes["name"]
	if !ok {
		t.Fatal(`"fields" nested object has no "name" attribute`)
	}
	strName, ok := nameAttr.(schema.StringAttribute)
	if !ok {
		t.Fatalf(`fields[].name is %T, want schema.StringAttribute`, nameAttr)
	}
	assertStringModifiersStable(t, "fields[].name", strName.PlanModifiers)
}

func assertStringAttrStable(t *testing.T, attrs map[string]schema.Attribute, name string) {
	t.Helper()
	a, ok := attrs[name].(schema.StringAttribute)
	if !ok {
		t.Fatalf("%q is %T, want schema.StringAttribute", name, attrs[name])
	}
	assertStringModifiersStable(t, name, a.PlanModifiers)
}

func assertStringModifiersStable(t *testing.T, name string, pms []planmodifier.String) {
	t.Helper()
	for _, pm := range pms {
		if strings.Contains(pm.Description(context.Background()), suppressDescFragment) {
			return
		}
	}
	t.Errorf("attribute %q has no state-stability plan modifier (got %d); it re-plans as "+
		"(known after apply) — IGA-743 convergence regression", name, len(pms))
}

func assertListAttrStable(t *testing.T, attrs map[string]schema.Attribute, name string) {
	t.Helper()
	a, ok := attrs[name].(schema.ListNestedAttribute)
	if !ok {
		t.Fatalf("%q is %T, want schema.ListNestedAttribute", name, attrs[name])
	}
	for _, pm := range a.PlanModifiers {
		if strings.Contains(pm.Description(context.Background()), suppressDescFragment) {
			return
		}
	}
	t.Errorf("attribute %q has no state-stability plan modifier (got %d); it re-plans as "+
		"(known after apply) — IGA-743 convergence regression", name, len(a.PlanModifiers))
}
