package provider

import (
	"context"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestRequestableEntriesRefreshPreservesResourceOwnership(t *testing.T) {
	model := AccessProfileRequestableEntriesResourceModel{
		AppEntitlements: []tfTypes.AppEntitlementRef{
			{AppID: types.StringValue("app-a"), ID: types.StringValue("entitlement-a")},
		},
	}
	appA, entitlementA := "app-a", "entitlement-a"
	appB, entitlementB := "app-b", "entitlement-b"

	diags := model.RefreshFromSharedRequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse(context.Background(), &shared.RequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse{
		AppEntitlements: []shared.AppEntitlementRef{
			{AppID: &appA, ID: &entitlementA},
			{AppID: &appB, ID: &entitlementB},
		},
	})
	if diags.HasError() {
		t.Fatalf("refresh returned diagnostics: %v", diags)
	}
	if len(model.AppEntitlements) != 1 {
		t.Fatalf("refresh adopted %d catalog entries; expected only this resource's 1 managed entry", len(model.AppEntitlements))
	}
	if got := model.AppEntitlements[0]; got.AppID.ValueString() != appA || got.ID.ValueString() != entitlementA {
		t.Fatalf("refresh preserved %+v, expected %s/%s", got, appA, entitlementA)
	}
}

func TestRequestableEntriesDiffOnlyMutatesOwnedEntries(t *testing.T) {
	current := []tfTypes.AppEntitlementRef{
		{AppID: types.StringValue("app-a"), ID: types.StringValue("entitlement-a")},
	}
	desired := []tfTypes.AppEntitlementRef{
		{AppID: types.StringValue("app-a"), ID: types.StringValue("entitlement-c")},
	}

	added, removed := appEntitlementRefDifference(current, desired)

	if len(added) != 1 || added[0].ID.ValueString() != "entitlement-c" {
		t.Fatalf("added = %#v, expected only entitlement-c", added)
	}
	if len(removed) != 1 || removed[0].ID.ValueString() != "entitlement-a" {
		t.Fatalf("removed = %#v, expected only entitlement-a", removed)
	}
}

func TestRequestableEntriesUpgradesLegacyState(t *testing.T) {
	ctx := context.Background()
	r := &AccessProfileRequestableEntriesResource{}
	var schemaResponse resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	if schemaResponse.Schema.Version != 1 {
		t.Fatalf("schema version = %d, expected 1", schemaResponse.Schema.Version)
	}

	upgrade, ok := r.UpgradeState(ctx)[0]
	if !ok {
		t.Fatal("missing upgrade path for legacy replace-all state")
	}
	legacyState := tfsdk.State{Schema: *upgrade.PriorSchema}
	legacy := &AccessProfileRequestableEntriesResourceModel{
		AppEntitlements: []tfTypes.AppEntitlementRef{
			{AppID: types.StringValue("app-a"), ID: types.StringValue("entitlement-a")},
		},
		CatalogID:      types.StringValue("catalog-a"),
		CreateRequests: types.BoolValue(true),
	}
	if diags := legacyState.Set(ctx, legacy); diags.HasError() {
		t.Fatalf("setting legacy state: %v", diags)
	}

	response := &resource.UpgradeStateResponse{State: tfsdk.State{Schema: schemaResponse.Schema}}
	upgrade.StateUpgrader(ctx, resource.UpgradeStateRequest{State: &legacyState}, response)
	if response.Diagnostics.HasError() {
		t.Fatalf("upgrading legacy state: %v", response.Diagnostics)
	}

	var migrated AccessProfileRequestableEntriesResourceModel
	if diags := response.State.Get(ctx, &migrated); diags.HasError() {
		t.Fatalf("reading migrated state: %v", diags)
	}
	if migrated.CatalogID.ValueString() != "catalog-a" || !migrated.CreateRequests.ValueBool() {
		t.Fatalf("migrated state lost stable fields: %#v", migrated)
	}
	if migrated.AppEntitlements != nil {
		t.Fatalf("migrated state retained %d ambiguous entries", len(migrated.AppEntitlements))
	}
}

func TestRequestableEntriesSchemaRequiresStableOwnership(t *testing.T) {
	r := &AccessProfileRequestableEntriesResource{}
	var response resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &response)

	catalog, ok := response.Schema.Attributes["catalog_id"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("catalog_id is %T, expected schema.StringAttribute", response.Schema.Attributes["catalog_id"])
	}
	if len(catalog.StringPlanModifiers()) == 0 {
		t.Fatal("catalog_id must require replacement when changed")
	}

	entries, ok := response.Schema.Attributes["app_entitlements"].(schema.ListNestedAttribute)
	if !ok {
		t.Fatalf("app_entitlements is %T, expected schema.ListNestedAttribute", response.Schema.Attributes["app_entitlements"])
	}
	for _, name := range []string{"app_id", "id"} {
		entry, ok := entries.NestedObject.Attributes[name].(schema.StringAttribute)
		if !ok {
			t.Fatalf("app_entitlements.%s is %T, expected schema.StringAttribute", name, entries.NestedObject.Attributes[name])
		}
		if !entry.IsRequired() {
			t.Errorf("app_entitlements.%s must be required", name)
		}
	}
}
