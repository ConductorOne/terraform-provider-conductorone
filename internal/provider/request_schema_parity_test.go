package provider

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestRequestSchemaResourceParity asserts that the generated model struct's
// tfsdk tags and the resource Schema's attribute keys describe the same shape,
// recursively, for conductorone_request_schema.
//
// A mismatch is exactly the IGA-743 class of bug: the terraform-plugin-framework
// errors with "mismatch between struct and object" on refresh when a model field
// has no matching schema attribute (or vice versa). Two ways this drifts here:
//   - Speakeasy regenerates the model (internal/provider/types/*) with a new
//     Field variant but no matching schema attribute is wired up (the v1.0.30
//     orphan-field regression).
//   - request_schema_resource.go / _sdk.go were once .genignore'd
//     (hand-maintained) and drifted from the generated model (the renamed-type
//     string_field vs form_string_field half of IGA-743).
//
// This test guards both. It would have failed on commit 59b78696.
func TestRequestSchemaResourceParity(t *testing.T) {
	var resp resource.SchemaResponse
	NewRequestSchemaResource().Schema(context.Background(), resource.SchemaRequest{}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Schema() returned diagnostics: %v", resp.Diagnostics)
	}
	if len(resp.Schema.Blocks) != 0 {
		t.Fatalf("resource Schema uses Blocks; this parity test only walks Attributes and must be extended")
	}

	compareLevel(t, "", reflect.TypeOf(RequestSchemaResourceModel{}), resp.Schema.Attributes)
}

var attrValueType = reflect.TypeOf((*attr.Value)(nil)).Elem()

// compareLevel asserts the tfsdk tags of structType and the keys of attrs are
// the same set, then recurses into every nested key.
func compareLevel(t *testing.T, path string, structType reflect.Type, attrs map[string]schema.Attribute) {
	t.Helper()

	model := modelChildren(structType)
	modelKeys := sortedKeys(model)
	schemaKeys := sortedAttrKeys(attrs)

	if !equalStrings(modelKeys, schemaKeys) {
		t.Errorf("parity mismatch at %s:\n  struct-only tfsdk tags:    %v\n  schema-only attribute keys: %v",
			displayPath(path), difference(modelKeys, schemaKeys), difference(schemaKeys, modelKeys))
		return
	}

	for name, nested := range model {
		childAttrs, childNested := schemaChild(attrs[name])
		modelNested := nested != nil
		switch {
		case modelNested && childNested:
			compareLevel(t, join(path, name), nested, childAttrs)
		case modelNested != childNested:
			t.Errorf("nesting mismatch at %s: model-nested=%v schema-nested=%v",
				join(path, name), modelNested, childNested)
		}
	}
}

// modelChildren maps each tfsdk tag of a struct to its nested struct type, or
// nil when the field is a leaf (a terraform attr.Value such as types.String).
func modelChildren(t reflect.Type) map[string]reflect.Type {
	out := map[string]reflect.Type{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		tag := f.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}
		out[tag] = nestedStruct(f.Type)
	}
	return out
}

// nestedStruct strips pointers/slices and returns the underlying struct type
// when it is a nested model object, or nil for leaves (attr.Value primitives).
func nestedStruct(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer || t.Kind() == reflect.Slice {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	if t.Implements(attrValueType) || reflect.PointerTo(t).Implements(attrValueType) {
		return nil
	}
	return t
}

// schemaChild returns the nested attribute map of a nested attribute, or nil
// for a leaf attribute.
func schemaChild(a schema.Attribute) (map[string]schema.Attribute, bool) {
	switch v := a.(type) {
	case schema.SingleNestedAttribute:
		return v.Attributes, true
	case schema.ListNestedAttribute:
		return v.NestedObject.Attributes, true
	case schema.SetNestedAttribute:
		return v.NestedObject.Attributes, true
	case schema.MapNestedAttribute:
		return v.NestedObject.Attributes, true
	default:
		return nil, false
	}
}

func sortedKeys(m map[string]reflect.Type) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func sortedAttrKeys(m map[string]schema.Attribute) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func difference(a, b []string) []string {
	set := map[string]struct{}{}
	for _, s := range b {
		set[s] = struct{}{}
	}
	var out []string
	for _, s := range a {
		if _, ok := set[s]; !ok {
			out = append(out, s)
		}
	}
	return out
}

func join(path, name string) string {
	if path == "" {
		return name
	}
	return path + "." + name
}

func displayPath(path string) string {
	if path == "" {
		return "<root>"
	}
	return path
}
