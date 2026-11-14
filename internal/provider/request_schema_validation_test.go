package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func requestSchemaTestType[T any](t *testing.T, value any) T {
	t.Helper()
	typed, ok := value.(T)
	if !ok {
		t.Fatalf("unexpected type %T; want %T", value, typed)
	}
	return typed
}

func requestSchemaTestSchema(t *testing.T) schema.Schema {
	t.Helper()
	var resp resource.SchemaResponse
	NewRequestSchemaResource().Schema(context.Background(), resource.SchemaRequest{}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}
	return resp.Schema
}

func TestRequestSchemaStringViewValidation(t *testing.T) {
	ctx := context.Background()
	fields := requestSchemaTestType[schema.ListNestedAttribute](t, requestSchemaTestSchema(t).Attributes["fields"])
	stringField := requestSchemaTestType[schema.SingleNestedAttribute](t, fields.NestedObject.Attributes["string_field"])
	objectType := requestSchemaTestType[types.ObjectType](t, stringField.GetType())
	known := func(overrides map[string]attr.Value) types.Object {
		values := make(map[string]attr.Value)
		for name, attributeType := range objectType.AttrTypes {
			value, err := attributeType.ValueFromTerraform(ctx, tftypes.NewValue(attributeType.TerraformType(ctx), nil))
			if err != nil {
				t.Fatal(err)
			}
			values[name] = value
		}
		for name, value := range overrides {
			values[name] = value
		}
		return types.ObjectValueMust(objectType.AttrTypes, values)
	}
	emptyView := func(name string) types.Object {
		viewType := requestSchemaTestType[types.ObjectType](t, objectType.AttrTypes[name])
		values := make(map[string]attr.Value)
		for key, attributeType := range viewType.AttrTypes {
			value, err := attributeType.ValueFromTerraform(ctx, tftypes.NewValue(attributeType.TerraformType(ctx), nil))
			if err != nil {
				t.Fatal(err)
			}
			values[key] = value
		}
		return types.ObjectValueMust(viewType.AttrTypes, values)
	}
	unknownText := types.ObjectUnknown(requestSchemaTestType[types.ObjectType](t, objectType.AttrTypes["text_field"]).AttrTypes)
	cases := []struct {
		name      string
		value     types.Object
		wantError bool
	}{
		{name: "null string field", value: types.ObjectNull(objectType.AttrTypes)},
		{name: "unknown string field", value: types.ObjectUnknown(objectType.AttrTypes)},
		{name: "known zero views", value: known(nil), wantError: true},
		{name: "explicit null view", value: known(map[string]attr.Value{"text_field": types.ObjectNull(unknownText.AttributeTypes(ctx))}), wantError: true},
		{name: "unknown unrelated placeholder does not defer missing view", value: known(map[string]attr.Value{"placeholder": types.StringUnknown()}), wantError: true},
		{name: "unknown view defers zero", value: known(map[string]attr.Value{"text_field": unknownText})},
		{name: "unknown view defers one", value: known(map[string]attr.Value{"text_field": unknownText, "date_field": emptyView("date_field")})},
		{name: "unknown nested multiline is a present view", value: known(map[string]attr.Value{"text_field": types.ObjectValueMust(
			unknownText.AttributeTypes(ctx), map[string]attr.Value{"multiline": types.BoolUnknown(), "suffix": types.StringNull()},
		)})},
		{name: "false multiline is a present view", value: known(map[string]attr.Value{"text_field": types.ObjectValueMust(
			unknownText.AttributeTypes(ctx), map[string]attr.Value{"multiline": types.BoolValue(false), "suffix": types.StringNull()},
		)})},
	}
	for _, view := range requestSchemaStringViews {
		cases = append(cases, struct {
			name      string
			value     types.Object
			wantError bool
		}{name: "empty " + view, value: known(map[string]attr.Value{view: emptyView(view)})})
	}
	for i, first := range requestSchemaStringViews {
		for _, second := range requestSchemaStringViews[i+1:] {
			cases = append(cases, struct {
				name      string
				value     types.Object
				wantError bool
			}{
				name:  first + " conflicts with " + second,
				value: known(map[string]attr.Value{first: emptyView(first), second: emptyView(second)}), wantError: true,
			})
		}
	}
	cases = append(cases, struct {
		name      string
		value     types.Object
		wantError bool
	}{
		name:      "known conflict still fails with unknown view",
		value:     known(map[string]attr.Value{"text_field": unknownText, "date_field": emptyView("date_field"), "password_field": emptyView("password_field")}),
		wantError: true,
	})
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var resp validator.ObjectResponse
			for _, check := range stringField.Validators {
				check.ValidateObject(ctx, validator.ObjectRequest{Path: path.Root("fields").AtListIndex(0).AtName("string_field"), ConfigValue: tc.value}, &resp)
			}
			if resp.Diagnostics.HasError() != tc.wantError {
				t.Fatalf("diagnostics = %v, wantError = %t", resp.Diagnostics, tc.wantError)
			}
			if tc.wantError && !strings.Contains(resp.Diagnostics.Errors()[0].Detail(), "exactly one") {
				t.Fatalf("missing actionable diagnostic: %v", resp.Diagnostics)
			}
		})
	}
}

func TestRequestSchemaStringViewDiagnosticPath(t *testing.T) {
	ctx := context.Background()
	s := requestSchemaTestSchema(t)
	fields := requestSchemaTestType[schema.ListNestedAttribute](t, s.Attributes["fields"])
	stringType := requestSchemaTestType[types.ObjectType](t, fields.NestedObject.Attributes["string_field"].GetType())
	dateType := requestSchemaTestType[types.ObjectType](t, stringType.AttrTypes["date_field"])
	dateValues := make(map[string]attr.Value)
	for name, attributeType := range dateType.AttrTypes {
		value, err := attributeType.ValueFromTerraform(ctx, tftypes.NewValue(attributeType.TerraformType(ctx), nil))
		if err != nil {
			t.Fatal(err)
		}
		dateValues[name] = value
	}
	for _, tc := range []struct {
		name      string
		index     int
		missing   bool
		conflict  bool
		wantError bool
	}{
		{name: "missing first view", missing: true, wantError: true},
		{name: "conflicting first views", conflict: true, wantError: true},
		{name: "missing second view", index: 1, missing: true, wantError: true},
		{name: "conflicting second views", index: 1, conflict: true, wantError: true},
		{name: "valid text view"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			model := &RequestSchemaResourceModel{Name: types.StringValue("Site")}
			for index := 0; index <= tc.index; index++ {
				model.Fields = append(model.Fields, tfTypes.FormField{
					DisplayName: types.StringValue(fmt.Sprintf("URL %d", index)),
					StringField: &tfTypes.FormStringField{TextField: &tfTypes.TextField{Multiline: types.BoolValue(false)}},
				})
			}
			if tc.missing {
				model.Fields[tc.index].StringField = &tfTypes.FormStringField{}
			}
			plan := requestSchemaTestPlan(t, model)
			fieldPath := path.Root("fields").AtListIndex(tc.index).AtName("string_field")
			if tc.conflict {
				if diags := plan.SetAttribute(ctx, fieldPath.AtName("date_field"), types.ObjectValueMust(dateType.AttrTypes, dateValues)); diags.HasError() {
					t.Fatal(diags)
				}
			}
			config, err := tfprotov6.NewDynamicValue(s.Type().TerraformType(ctx), plan.Raw)
			if err != nil {
				t.Fatal(err)
			}
			server, err := testAccProviderFactories["conductorone"]()
			if err != nil {
				t.Fatal(err)
			}
			resp, err := server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{TypeName: "conductorone_request_schema", Config: &config})
			if err != nil {
				t.Fatal(err)
			}
			var found bool
			for _, diagnostic := range resp.Diagnostics {
				if diagnostic.Severity != tfprotov6.DiagnosticSeverityError {
					continue
				}
				if diagnostic.Summary != "Invalid request-schema string view" {
					t.Fatalf("unexpected diagnostic: %v", diagnostic)
				}
				found = true
				if !strings.Contains(diagnostic.Detail, fieldPath.String()) || !strings.Contains(diagnostic.Detail, "exactly one") {
					t.Fatalf("missing field path or guidance: %v", diagnostic)
				}
				for _, view := range requestSchemaStringViews {
					if !strings.Contains(diagnostic.Detail, view) {
						t.Fatalf("missing allowed view %q: %v", view, diagnostic)
					}
				}
			}
			if found != tc.wantError {
				t.Fatalf("view error=%t, want %t; diagnostics=%v", found, tc.wantError, resp.Diagnostics)
			}
		})
	}
}

func requestSchemaTestPlan(t *testing.T, model *RequestSchemaResourceModel) tfsdk.Plan {
	t.Helper()
	plan := tfsdk.Plan{Schema: requestSchemaTestSchema(t)}
	if diags := plan.Set(context.Background(), model); diags.HasError() {
		t.Fatal(diags)
	}
	return plan
}

func TestRequestSchemaNestedNameIsComputedOnly(t *testing.T) {
	ctx := context.Background()
	s := requestSchemaTestSchema(t)
	fields := requestSchemaTestType[schema.ListNestedAttribute](t, s.Attributes["fields"])
	name := requestSchemaTestType[schema.StringAttribute](t, fields.NestedObject.Attributes["name"])
	if name.Optional || name.Required || !name.Computed {
		t.Fatalf("nested name must be computed-only: %+v", name)
	}
	if !s.Attributes["name"].IsOptional() {
		t.Fatal("top-level resource name must remain configurable")
	}
	server, err := testAccProviderFactories["conductorone"]()
	if err != nil {
		t.Fatal(err)
	}
	for _, explicit := range []bool{false, true} {
		model := &RequestSchemaResourceModel{Name: types.StringValue("Site"), Fields: []tfTypes.FormField{{DisplayName: types.StringValue("URL")}}}
		if explicit {
			model.Fields[0].Name = types.StringValue("client_alias")
		}
		plan := requestSchemaTestPlan(t, model)
		config, err := tfprotov6.NewDynamicValue(s.Type().TerraformType(ctx), plan.Raw)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{TypeName: "conductorone_request_schema", Config: &config})
		if err != nil {
			t.Fatal(err)
		}
		var computedNameError, migrationGuidance bool
		for _, d := range resp.Diagnostics {
			if d.Severity == tfprotov6.DiagnosticSeverityError && d.Attribute != nil && strings.Contains(d.Attribute.String(), "name") {
				computedNameError = true
			}
			if strings.Contains(d.Detail, "fields[0].name") && strings.Contains(d.Detail, "top-level resource name") {
				migrationGuidance = true
			}
		}
		if computedNameError != explicit || migrationGuidance != explicit {
			t.Fatalf("explicit=%t diagnostics=%v", explicit, resp.Diagnostics)
		}
	}
}

func TestRequestSchemaFieldNamePlanStability(t *testing.T) {
	ctx := context.Background()
	for _, tc := range []struct {
		name       string
		configured []types.String
		stored     []types.String
		preserve   bool
	}{
		{name: "unchanged labels", configured: []types.String{types.StringValue("URL"), types.StringValue("URL")}, stored: []types.String{types.StringValue("URL"), types.StringValue("URL")}, preserve: true},
		{name: "changed duplicate predecessor", configured: []types.String{types.StringValue("URL"), types.StringValue("URL")}, stored: []types.String{types.StringValue("Other"), types.StringValue("URL")}},
		{name: "unknown label", configured: []types.String{types.StringValue("URL"), types.StringUnknown()}, stored: []types.String{types.StringValue("URL"), types.StringValue("URL")}},
		{name: "removed predecessor", configured: []types.String{types.StringValue("URL")}, stored: []types.String{types.StringValue("URL"), types.StringValue("URL")}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			makeModel := func(labels []types.String) *RequestSchemaResourceModel {
				m := &RequestSchemaResourceModel{}
				for _, label := range labels {
					m.Fields = append(m.Fields, tfTypes.FormField{DisplayName: label})
				}
				return m
			}
			configured, stored := requestSchemaTestPlan(t, makeModel(tc.configured)), requestSchemaTestPlan(t, makeModel(tc.stored))
			var resp planmodifier.StringResponse
			resp.PlanValue = types.StringUnknown()
			requestSchemaFieldName{}.PlanModifyString(ctx, planmodifier.StringRequest{
				Path: path.Root("fields").AtListIndex(0).AtName("name"), PlanValue: types.StringUnknown(), StateValue: types.StringValue("URL_0"),
				Config: tfsdk.Config(configured), State: tfsdk.State(stored),
			}, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatal(resp.Diagnostics)
			}
			if got := !resp.PlanValue.IsUnknown(); got != tc.preserve {
				t.Fatalf("preserve=%t, want %t", got, tc.preserve)
			}
			if resp.RequiresReplace {
				t.Fatal("computed names must not replace the resource")
			}
		})
	}
}

func TestRequestSchemaCRUDRetainsServerNames(t *testing.T) {
	ctx := context.Background()
	label, fieldName := "URL", "URL_0"
	api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprintf(w, `{"requestSchema":{"id":"schema-id","form":{"displayName":"Site","fields":[{"name":%q,"displayName":%q,"required":true,"stringField":{"textField":{"multiline":false}}}],"fieldGroups":[],"fieldRelationships":[]},"createdAt":"2026-11-13T00:00:00Z","justificationVisibility":"JUSTIFICATION_VISIBILITY_UNSPECIFIED"}}`, fieldName, label); err != nil {
			t.Errorf("write request-schema response: %v", err)
		}
	}))
	defer api.Close()
	r := requestSchemaTestType[*requestSchemaResource](t, NewRequestSchemaResource())
	r.client = sdk.New(sdk.WithServerURL(api.URL))
	model := &RequestSchemaResourceModel{Name: types.StringValue("Site"), Fields: []tfTypes.FormField{{
		DisplayName: types.StringValue(label), Name: types.StringUnknown(), Required: types.BoolValue(true),
		StringField: &tfTypes.FormStringField{TextField: &tfTypes.TextField{Multiline: types.BoolValue(false)}},
	}}}
	plan := requestSchemaTestPlan(t, model)
	created := resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
	r.Create(ctx, resource.CreateRequest{Plan: plan}, &created)
	assertName := func(state tfsdk.State, expected string) {
		t.Helper()
		var got types.String
		if diags := state.GetAttribute(ctx, path.Root("fields").AtListIndex(0).AtName("name"), &got); diags.HasError() {
			t.Fatal(diags)
		}
		if got.IsNull() || got.IsUnknown() || got.ValueString() != expected {
			t.Fatalf("name=%v, want %s", got, expected)
		}
	}
	if created.Diagnostics.HasError() {
		t.Fatal(created.Diagnostics)
	}
	assertName(created.State, "URL_0")

	// Existing state/import needs no schema-version migration or replacement.
	read := resource.ReadResponse{State: created.State}
	r.Read(ctx, resource.ReadRequest{State: created.State}, &read)
	if read.Diagnostics.HasError() {
		t.Fatal(read.Diagnostics)
	}
	assertName(read.State, "URL_0")

	empty := requestSchemaTestPlan(t, &RequestSchemaResourceModel{})
	imported := resource.ImportStateResponse{State: tfsdk.State(empty)}
	requestSchemaTestType[resource.ResourceWithImportState](t, r).ImportState(ctx, resource.ImportStateRequest{ID: "schema-id"}, &imported)
	if imported.Diagnostics.HasError() {
		t.Fatal(imported.Diagnostics)
	}
	importRead := resource.ReadResponse{State: imported.State}
	r.Read(ctx, resource.ReadRequest{State: imported.State}, &importRead)
	if importRead.Diagnostics.HasError() {
		t.Fatal(importRead.Diagnostics)
	}
	assertName(importRead.State, "URL_0")

	label, fieldName = "New URL", "New_URL_0"
	model.ID = types.StringValue("schema-id")
	model.Fields[0].DisplayName = types.StringValue(label)
	updatedPlan := requestSchemaTestPlan(t, model)
	updated := resource.UpdateResponse{State: read.State}
	r.Update(ctx, resource.UpdateRequest{Plan: updatedPlan, State: read.State}, &updated)
	if updated.Diagnostics.HasError() {
		t.Fatal(updated.Diagnostics)
	}
	assertName(updated.State, "New_URL_0")
}
