package provider

import (
	"context"
	"reflect"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestAppOwnerEntitlementResource_SchemaCoversAppEntitlementStruct is the
// IGA-1774 regression test for the sibling conductorone_app_owner_entitlement
// resource. See the doc comment on
// TestAppEntitlementOwnerEntitlementResource_SchemaCoversAppEntitlementStruct
// for the full story; the bug applied identically to both schemas.
func TestAppOwnerEntitlementResource_SchemaCoversAppEntitlementStruct(t *testing.T) {
	t.Parallel()
	schemaAttrs := appOwnerEntitlementSchemaAttrs(t)
	assertSchemaCoversStruct(t, schemaAttrs, reflect.TypeOf(tfTypes.AppEntitlement{}))
}

// TestAppOwnerEntitlementResource_SchemaHasDeletedAt is the narrow direct
// regression assertion.
func TestAppOwnerEntitlementResource_SchemaHasDeletedAt(t *testing.T) {
	t.Parallel()
	schemaAttrs := appOwnerEntitlementSchemaAttrs(t)
	if _, ok := schemaAttrs["deleted_at"]; !ok {
		t.Fatalf("conductorone_app_owner_entitlement schema is missing nested app_entitlement.deleted_at attribute (IGA-1774 regression); attributes present: %s", sortedSchemaKeys(schemaAttrs))
	}
}

func appOwnerEntitlementSchemaAttrs(t *testing.T) map[string]schema.Attribute {
	t.Helper()
	r := NewAppOwnerEntitlementResource()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Schema() returned diagnostics: %v", resp.Diagnostics)
	}
	return nestedAppEntitlementAttrs(t, resp.Schema)
}
