package provider

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestAccessProfileTypeSchema(t *testing.T) {
	resp := &resource.SchemaResponse{}
	NewAccessProfileResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("building access profile resource schema: %v", resp.Diagnostics)
	}

	attr, ok := resp.Schema.Attributes["type"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("type has schema type %T, want schema.StringAttribute", resp.Schema.Attributes["type"])
	}
	if !attr.Optional || !attr.Computed {
		t.Fatalf("type must be Optional and Computed, got Optional=%t Computed=%t", attr.Optional, attr.Computed)
	}

	want := reflect.TypeOf(stringplanmodifier.RequiresReplaceIfConfigured())
	hasRequiresReplace := false
	for _, modifier := range attr.PlanModifiers {
		if reflect.TypeOf(modifier) == want {
			hasRequiresReplace = true
			break
		}
	}
	if !hasRequiresReplace {
		t.Fatalf("type does not require replacement when configured: %#v", attr.PlanModifiers)
	}

	dataSourceResp := &datasource.SchemaResponse{}
	NewAccessProfileDataSource().Schema(context.Background(), datasource.SchemaRequest{}, dataSourceResp)
	if dataSourceResp.Diagnostics.HasError() {
		t.Fatalf("building access profile data source schema: %v", dataSourceResp.Diagnostics)
	}
	dataSourceAttr, ok := dataSourceResp.Schema.Attributes["type"].(datasourceschema.StringAttribute)
	if !ok {
		t.Fatalf("data source type has schema type %T, want schema.StringAttribute", dataSourceResp.Schema.Attributes["type"])
	}
	if !dataSourceAttr.Computed || dataSourceAttr.Optional {
		t.Fatalf("data source type must be Computed-only, got Optional=%t Computed=%t", dataSourceAttr.Optional, dataSourceAttr.Computed)
	}
}

func TestAccessProfileTypeReplacementBehavior(t *testing.T) {
	modifier := stringplanmodifier.RequiresReplaceIfConfigured()
	nonNullRaw := tftypes.NewValue(tftypes.String, "resource")

	t.Run("omitted type accepts backfill without replacement", func(t *testing.T) {
		resp := &planmodifier.StringResponse{}
		modifier.PlanModifyString(context.Background(), planmodifier.StringRequest{
			ConfigValue: types.StringNull(),
			Plan:        tfsdk.Plan{Raw: nonNullRaw},
			PlanValue:   types.StringValue("REQUEST_CATALOG_TYPE_CATALOG_AND_BUNDLE"),
			State:       tfsdk.State{Raw: nonNullRaw},
			StateValue:  types.StringValue("REQUEST_CATALOG_TYPE_UNSPECIFIED"),
		}, resp)

		if resp.RequiresReplace {
			t.Fatal("omitted type must not replace the resource when backfill changes its computed value")
		}
	})

	t.Run("configured type change requires replacement", func(t *testing.T) {
		resp := &planmodifier.StringResponse{}
		modifier.PlanModifyString(context.Background(), planmodifier.StringRequest{
			ConfigValue: types.StringValue("REQUEST_CATALOG_TYPE_BUNDLE"),
			Plan:        tfsdk.Plan{Raw: nonNullRaw},
			PlanValue:   types.StringValue("REQUEST_CATALOG_TYPE_BUNDLE"),
			State:       tfsdk.State{Raw: nonNullRaw},
			StateValue:  types.StringValue("REQUEST_CATALOG_TYPE_CATALOG"),
		}, resp)

		if !resp.RequiresReplace {
			t.Fatal("configured type change must replace the resource")
		}
	})
}

func TestAccessProfileTypeMappings(t *testing.T) {
	ctx := context.Background()

	t.Run("configured type is sent on create", func(t *testing.T) {
		model := &AccessProfileResourceModel{
			DisplayName: types.StringValue("requestable"),
			Type:        types.StringValue("REQUEST_CATALOG_TYPE_CATALOG"),
		}
		request, diags := model.ToSharedRequestCatalogManagementServiceCreateRequest(ctx)
		if diags.HasError() {
			t.Fatalf("mapping create request: %v", diags)
		}
		if request.Type == nil || string(*request.Type) != "REQUEST_CATALOG_TYPE_CATALOG" {
			t.Fatalf("create type = %v, want REQUEST_CATALOG_TYPE_CATALOG", request.Type)
		}
	})

	t.Run("omitted type stays absent on create", func(t *testing.T) {
		model := &AccessProfileResourceModel{
			DisplayName: types.StringValue("legacy-compatible"),
			Type:        types.StringNull(),
		}
		request, diags := model.ToSharedRequestCatalogManagementServiceCreateRequest(ctx)
		if diags.HasError() {
			t.Fatalf("mapping create request: %v", diags)
		}
		if request.Type != nil {
			t.Fatalf("create type = %v, want nil", *request.Type)
		}
	})

	t.Run("read refreshes stored type", func(t *testing.T) {
		catalogType := shared.RequestCatalogTypeRequestCatalogTypeBundle
		model := &AccessProfileResourceModel{}
		diags := model.RefreshFromSharedRequestCatalog(ctx, &shared.RequestCatalog{Type: &catalogType})
		if diags.HasError() {
			t.Fatalf("refreshing access profile: %v", diags)
		}
		if got := model.Type.ValueString(); got != "REQUEST_CATALOG_TYPE_BUNDLE" {
			t.Fatalf("refreshed type = %q, want REQUEST_CATALOG_TYPE_BUNDLE", got)
		}
	})

	t.Run("type is absent from update", func(t *testing.T) {
		model := &AccessProfileResourceModel{Type: types.StringValue("REQUEST_CATALOG_TYPE_BUNDLE")}
		request, diags := model.ToSharedRequestCatalogManagementServiceUpdateRequest(ctx)
		if diags.HasError() {
			t.Fatalf("mapping update request: %v", diags)
		}
		body, err := json.Marshal(request)
		if err != nil {
			t.Fatalf("marshalling update request: %v", err)
		}
		if strings.Contains(string(body), `"type"`) {
			t.Fatalf("update request unexpectedly contains type: %s", body)
		}
	})
}
