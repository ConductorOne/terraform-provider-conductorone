package provider

import (
	"context"
	"reflect"
	"sort"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestRequestSchemaFieldsMatchFormFieldStruct guards against the IGA-743 drift.
//
// request_schema_resource.go and request_schema_resource_sdk.go are listed in
// .genignore, so Speakeasy does not regenerate them. The conversion type
// tfTypes.FormField (internal/provider/types/form_field.go) is generated and
// NOT genignored, so it moves forward on every regen. When the two diverge,
// the terraform-plugin-framework reflection that maps the schema's "fields"
// object onto []tfTypes.FormField fails at runtime with:
//
//	mismatch between struct and object: Struct defines fields not found in
//	object: ... Object defines fields not found in struct: ...
//
// This is exactly the customer-facing "Value Conversion Error" reported in
// IGA-743 / GitHub #196 (broken 1.0.30 .. 1.3.0). This test asserts the
// invariant directly: every tfsdk tag on FormField has a matching attribute in
// the resource schema's "fields" nested object, and vice versa.
func TestRequestSchemaFieldsMatchFormFieldStruct(t *testing.T) {
	ctx := context.Background()
	resp := &resource.SchemaResponse{}
	NewRequestSchemaResource().Schema(ctx, resource.SchemaRequest{}, resp)

	fieldsAttr, ok := resp.Schema.Attributes["fields"]
	if !ok {
		t.Fatal(`resource schema has no "fields" attribute`)
	}
	listNested, ok := fieldsAttr.(schema.ListNestedAttribute)
	if !ok {
		t.Fatalf(`"fields" is %T, want schema.ListNestedAttribute`, fieldsAttr)
	}

	schemaAttrs := map[string]bool{}
	for name := range listNested.NestedObject.Attributes {
		schemaAttrs[name] = true
	}

	structTags := map[string]bool{}
	rt := reflect.TypeOf(tfTypes.FormField{})
	for i := 0; i < rt.NumField(); i++ {
		if tag, ok := rt.Field(i).Tag.Lookup("tfsdk"); ok {
			structTags[tag] = true
		}
	}

	for tag := range structTags {
		if !schemaAttrs[tag] {
			t.Errorf("FormField struct defines tfsdk:%q with no matching attribute in the resource schema \"fields\" object", tag)
		}
	}
	for name := range schemaAttrs {
		if !structTags[name] {
			t.Errorf("resource schema \"fields\" object defines attribute %q with no matching field on the FormField struct", name)
		}
	}

	if t.Failed() {
		t.Logf("schema \"fields\" attributes: %v", sortedKeys(schemaAttrs))
		t.Logf("FormField struct tfsdk tags: %v", sortedKeys(structTags))
	}
}

func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
