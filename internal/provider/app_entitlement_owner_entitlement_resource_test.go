package provider

import (
	"context"
	"reflect"
	"sort"
	"strings"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestAppEntitlementOwnerEntitlementResource_SchemaCoversAppEntitlementStruct is a
// regression test for IGA-1774.
//
// The bug: types.AppEntitlement (a generated tfsdk-tagged struct) gained
// new fields (e.g. deleted_at) but the resource schema overlay for
// conductorone_app_entitlement_owner_entitlement was not regenerated to
// match. The struct/schema mismatch surfaced at runtime as
// "Value Conversion Error: Mismatch between struct and object type".
//
// This test pulls the resource's Schema(), walks the nested app_entitlement
// SingleNestedAttribute, and asserts every tfsdk field declared on
// types.AppEntitlement has a corresponding schema attribute. If a future
// regen drops a field from the schema while leaving it on the struct (or
// vice versa), this test fails before users hit the runtime error.
func TestAppEntitlementOwnerEntitlementResource_SchemaCoversAppEntitlementStruct(t *testing.T) {
	t.Parallel()
	schemaAttrs := appEntitlementOwnerEntitlementSchemaAttrs(t)
	assertSchemaCoversStruct(t, schemaAttrs, reflect.TypeOf(tfTypes.AppEntitlement{}))
}

// TestAppEntitlementOwnerEntitlementResource_SchemaHasDeletedAt is the narrow,
// direct regression assertion for IGA-1774. If this single attribute ever
// disappears from the nested app_entitlement block again, this test names it
// in the failure message so the cause is obvious.
func TestAppEntitlementOwnerEntitlementResource_SchemaHasDeletedAt(t *testing.T) {
	t.Parallel()
	schemaAttrs := appEntitlementOwnerEntitlementSchemaAttrs(t)
	if _, ok := schemaAttrs["deleted_at"]; !ok {
		t.Fatalf("conductorone_app_entitlement_owner_entitlement schema is missing nested app_entitlement.deleted_at attribute (IGA-1774 regression); attributes present: %s", sortedSchemaKeys(schemaAttrs))
	}
}

// appEntitlementOwnerEntitlementSchemaAttrs returns the map of attribute
// names under the nested app_entitlement block of the resource's schema.
func appEntitlementOwnerEntitlementSchemaAttrs(t *testing.T) map[string]schema.Attribute {
	t.Helper()
	r := NewAppEntitlementOwnerEntitlementResource()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Schema() returned diagnostics: %v", resp.Diagnostics)
	}
	return nestedAppEntitlementAttrs(t, resp.Schema)
}

// nestedAppEntitlementAttrs extracts the "app_entitlement"
// SingleNestedAttribute's inner attribute map from a resource schema.
func nestedAppEntitlementAttrs(t *testing.T, s schema.Schema) map[string]schema.Attribute {
	t.Helper()
	attr, ok := s.Attributes["app_entitlement"]
	if !ok {
		t.Fatal("schema has no top-level app_entitlement attribute")
	}
	nested, ok := attr.(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("app_entitlement is %T, want schema.SingleNestedAttribute", attr)
	}
	return nested.Attributes
}

// assertSchemaCoversStruct walks every tfsdk-tagged field on structType
// and fails the test for each field that is not present in schemaAttrs.
// This is the broader regression net: any struct field added later that
// the schema overlay forgets will trip the test.
func assertSchemaCoversStruct(t *testing.T, schemaAttrs map[string]schema.Attribute, structType reflect.Type) {
	t.Helper()
	var missing []string
	for i := 0; i < structType.NumField(); i++ {
		tag, ok := structType.Field(i).Tag.Lookup("tfsdk")
		if !ok || tag == "" || tag == "-" {
			continue
		}
		// Strip any tag options (e.g. ",optional"); the framework doesn't
		// use them on tfsdk tags today, but be defensive.
		name := strings.SplitN(tag, ",", 2)[0]
		if _, ok := schemaAttrs[name]; !ok {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		t.Fatalf("schema is missing attributes declared on %s: %v\n(schema attributes present: %s)",
			structType.Name(), missing, sortedSchemaKeys(schemaAttrs))
	}
}

func sortedSchemaKeys(m map[string]schema.Attribute) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
