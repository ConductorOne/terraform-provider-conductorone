package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func requestSchemaTestObject(t testing.TB, objectType types.ObjectType, overrides map[string]attr.Value) types.Object {
	t.Helper()
	values := make(map[string]attr.Value)
	for name, attributeType := range objectType.AttrTypes {
		value, err := attributeType.ValueFromTerraform(context.Background(), tftypes.NewValue(attributeType.TerraformType(context.Background()), nil))
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

func TestRequestSchemaFieldTypeValidation(t *testing.T) {
	ctx := context.Background()
	fields := requestSchemaTestType[schema.ListNestedAttribute](t, requestSchemaTestSchema(t).Attributes["fields"])
	fieldType := requestSchemaTestType[types.ObjectType](t, requestSchemaTestType[types.ListType](t, fields.GetType()).ElemType)
	memberType := func(name string) types.ObjectType {
		return requestSchemaTestType[types.ObjectType](t, fieldType.AttrTypes[name])
	}
	empty := func(name string) types.Object { return requestSchemaTestObject(t, memberType(name), nil) }
	unknown := types.ObjectUnknown(memberType("bool_field").AttrTypes)
	cases := []struct {
		name      string
		value     types.Object
		wantError bool
	}{
		{name: "unknown whole field", value: types.ObjectUnknown(fieldType.AttrTypes)},
		{name: "missing type", value: requestSchemaTestObject(t, fieldType, nil), wantError: true},
		{name: "unknown label does not defer missing type", value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{"display_name": types.StringUnknown()}), wantError: true},
		{name: "null type", value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{"bool_field": types.ObjectNull(memberType("bool_field").AttrTypes)}), wantError: true},
		{name: "unknown type defers absence", value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{"bool_field": unknown})},
		{name: "unknown type may resolve null beside known type", value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{"bool_field": unknown, "file_field": empty("file_field")})},
		{name: "unknown cannot fix known conflict", value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{"bool_field": unknown, "file_field": empty("file_field"), "int64_field": empty("int64_field")}), wantError: true},
		{name: "false default is present", value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{"bool_field": requestSchemaTestObject(t, memberType("bool_field"), map[string]attr.Value{"default_value": types.BoolValue(false)})})},
	}
	for i, first := range requestSchemaFieldTypes {
		cases = append(cases, struct {
			name      string
			value     types.Object
			wantError bool
		}{name: "present " + first, value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{first: empty(first)})})
		for _, second := range requestSchemaFieldTypes[i+1:] {
			cases = append(cases, struct {
				name      string
				value     types.Object
				wantError bool
			}{name: first + " conflicts with " + second, value: requestSchemaTestObject(t, fieldType, map[string]attr.Value{first: empty(first), second: empty(second)}), wantError: true})
		}
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var response validator.ObjectResponse
			for _, check := range fields.NestedObject.Validators {
				check.ValidateObject(ctx, validator.ObjectRequest{Path: path.Root("fields").AtListIndex(1), ConfigValue: tc.value}, &response)
			}
			if response.Diagnostics.HasError() != tc.wantError {
				t.Fatalf("diagnostics=%v, wantError=%t", response.Diagnostics, tc.wantError)
			}
		})
	}
}

func TestRequestSchemaFieldTypeDiagnosticPath(t *testing.T) {
	ctx := context.Background()
	for _, index := range []int{0, 1} {
		for _, conflict := range []bool{false, true} {
			t.Run(fmt.Sprintf("index%d/conflict%t", index, conflict), func(t *testing.T) {
				model := &RequestSchemaResourceModel{Name: types.StringValue("Site")}
				for i := 0; i <= index; i++ {
					model.Fields = append(model.Fields, tfTypes.FormField{DisplayName: types.StringValue("Approved"), BoolField: &tfTypes.BoolField{DefaultValue: types.BoolValue(false)}})
				}
				if conflict {
					model.Fields[index].StringField = &tfTypes.FormStringField{TextField: &tfTypes.TextField{Multiline: types.BoolValue(false)}}
				} else {
					model.Fields[index].BoolField = nil
				}
				plan := requestSchemaTestPlan(t, model)
				config, err := tfprotov6.NewDynamicValue(plan.Schema.Type().TerraformType(ctx), plan.Raw)
				if err != nil {
					t.Fatal(err)
				}
				server, err := testAccProviderFactories["conductorone"]()
				if err != nil {
					t.Fatal(err)
				}
				response, err := server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{TypeName: "conductorone_request_schema", Config: &config})
				if err != nil {
					t.Fatal(err)
				}
				var found bool
				for _, diagnostic := range response.Diagnostics {
					if diagnostic.Severity != tfprotov6.DiagnosticSeverityError {
						continue
					}
					if diagnostic.Summary != "Invalid request-schema field type" || !strings.Contains(diagnostic.Detail, fmt.Sprintf("fields[%d]", index)) || !strings.Contains(diagnostic.Detail, "exactly one") {
						t.Fatalf("unexpected diagnostic: %v", diagnostic)
					}
					if diagnostic.Attribute == nil || diagnostic.Attribute.String() != tftypes.NewAttributePath().WithAttributeName("fields").WithElementKeyInt(index).String() {
						t.Fatalf("unexpected diagnostic path: %v", diagnostic)
					}
					found = true
				}
				if !found {
					t.Fatalf("missing field-type diagnostic: %v", response.Diagnostics)
				}
			})
		}
	}
}

func TestRequestSchemaWriteValidation(t *testing.T) {
	ctx := context.Background()
	for _, operation := range []string{"create", "update"} {
		for _, shape := range []string{"missing type", "conflicting types", "deferred type resolves null", "deferred type resolves conflict", "deferred view resolves null", "deferred view resolves conflict", "conflicting views", "null field", "valid false", "valid string", "empty fields", "omitted fields"} {
			t.Run(operation+"/"+shape, func(t *testing.T) {
				var calls atomic.Int32
				api := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
					calls.Add(1)
					var body struct {
						Fields        []shared.FormField    `json:"fields"`
						RequestSchema *shared.RequestSchema `json:"requestSchema"`
					}
					if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
						t.Errorf("decode request: %v", err)
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					fields := body.Fields
					if operation == "update" {
						fields = body.RequestSchema.GetForm().GetFields()
					}
					for index := range fields {
						name := fmt.Sprintf("Approved_%d", index)
						fields[index].Name = &name
					}
					w.Header().Set("Content-Type", "application/json")
					if err := json.NewEncoder(w).Encode(map[string]any{"requestSchema": map[string]any{
						"id": "schema-id", "createdAt": "2026-11-13T00:00:00Z",
						"justificationVisibility": "JUSTIFICATION_VISIBILITY_UNSPECIFIED",
						"form":                    map[string]any{"displayName": "Site", "fields": fields, "fieldGroups": []any{}, "fieldRelationships": []any{}},
					}}); err != nil {
						t.Errorf("encode response: %v", err)
					}
				}))
				defer api.Close()
				model := &RequestSchemaResourceModel{ID: types.StringValue("schema-id"), Name: types.StringValue("Site")}
				for range 2 {
					model.Fields = append(model.Fields, tfTypes.FormField{DisplayName: types.StringValue("Approved"), BoolField: &tfTypes.BoolField{DefaultValue: types.BoolValue(false)}})
				}
				prior := requestSchemaTestPlan(t, model)
				switch shape {
				case "missing type", "deferred type resolves null":
					model.Fields[1].BoolField = nil
				case "conflicting types", "deferred type resolves conflict":
					model.Fields[1].StringField = &tfTypes.FormStringField{TextField: &tfTypes.TextField{Multiline: types.BoolValue(false)}}
				case "deferred view resolves null", "deferred view resolves conflict", "conflicting views", "valid string":
					model.Fields[1].BoolField = nil
					model.Fields[1].StringField = &tfTypes.FormStringField{}
					if shape != "deferred view resolves null" {
						model.Fields[1].StringField.TextField = &tfTypes.TextField{Multiline: types.BoolValue(false)}
					}
				case "empty fields":
					model.Fields = []tfTypes.FormField{}
				case "omitted fields":
					model.Fields = nil
				}
				plan := requestSchemaTestPlan(t, model)
				fieldPath := path.Root("fields").AtListIndex(1)
				if shape == "conflicting views" || shape == "deferred view resolves conflict" {
					viewPath := fieldPath.AtName("string_field").AtName("date_field")
					var view types.Object
					if diagnostics := plan.GetAttribute(ctx, viewPath, &view); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					if diagnostics := plan.SetAttribute(ctx, viewPath, requestSchemaTestObject(t, types.ObjectType{AttrTypes: view.AttributeTypes(ctx)}, nil)); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
				}
				if strings.HasPrefix(shape, "deferred") {
					deferred := plan
					memberPath := fieldPath.AtName("bool_field")
					if strings.HasPrefix(shape, "deferred view") {
						memberPath = fieldPath.AtName("string_field").AtName("text_field")
					}
					var value types.Object
					if diagnostics := plan.GetAttribute(ctx, memberPath, &value); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					if diagnostics := deferred.SetAttribute(ctx, memberPath, types.ObjectUnknown(value.AttributeTypes(ctx))); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					if diagnostics := deferred.SetAttribute(ctx, path.Root("id"), types.StringNull()); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					config, err := tfprotov6.NewDynamicValue(plan.Schema.Type().TerraformType(ctx), deferred.Raw)
					if err != nil {
						t.Fatal(err)
					}
					server, err := testAccProviderFactories["conductorone"]()
					if err != nil {
						t.Fatal(err)
					}
					response, err := server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{TypeName: "conductorone_request_schema", Config: &config})
					if err != nil || len(response.Diagnostics) != 0 {
						t.Fatalf("unknown expression must defer validation: response=%v error=%v", response, err)
					}
				}
				if shape == "null field" {
					var field types.Object
					if diagnostics := plan.GetAttribute(ctx, fieldPath, &field); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
					if diagnostics := plan.SetAttribute(ctx, fieldPath, types.ObjectNull(field.AttributeTypes(ctx))); diagnostics.HasError() {
						t.Fatal(diagnostics)
					}
				}
				r := requestSchemaTestType[*requestSchemaResource](t, NewRequestSchemaResource())
				r.client = sdk.New(sdk.WithServerURL(api.URL))
				var diagnostics diag.Diagnostics
				if operation == "create" {
					response := resource.CreateResponse{State: tfsdk.State{Schema: plan.Schema}}
					r.Create(ctx, resource.CreateRequest{Plan: plan}, &response)
					diagnostics = response.Diagnostics
				} else {
					response := resource.UpdateResponse{State: tfsdk.State(prior)}
					r.Update(ctx, resource.UpdateRequest{Plan: plan, State: tfsdk.State(prior)}, &response)
					diagnostics = response.Diagnostics
				}
				valid := strings.HasPrefix(shape, "valid") || shape == "empty fields" || shape == "omitted fields"
				if valid {
					if diagnostics.HasError() || calls.Load() != 1 {
						t.Fatalf("valid write: calls=%d diagnostics=%v", calls.Load(), diagnostics)
					}
				} else {
					if !diagnostics.HasError() || calls.Load() != 0 {
						t.Fatalf("invalid write must stop before HTTP: calls=%d diagnostics=%v", calls.Load(), diagnostics)
					}
					first := diagnostics.Errors()[0]
					if !strings.Contains(first.Detail(), "fields[1]") || !strings.Contains(first.Detail(), "exactly one") {
						t.Fatalf("missing indexed write diagnostic: %v", diagnostics)
					}
				}
			})
		}
	}
}

func BenchmarkRequestSchemaFieldNamePlanning(b *testing.B) {
	ctx := context.Background()
	for _, count := range []int{1, 10, 100} {
		b.Run(fmt.Sprintf("fields%d", count), func(b *testing.B) {
			configModel := &RequestSchemaResourceModel{}
			stateModel := &RequestSchemaResourceModel{}
			planModel := &RequestSchemaResourceModel{}
			for index := range count {
				configModel.Fields = append(configModel.Fields, tfTypes.FormField{DisplayName: types.StringValue("URL")})
				stateModel.Fields = append(stateModel.Fields, tfTypes.FormField{DisplayName: types.StringValue("URL"), Name: types.StringValue(fmt.Sprintf("URL_%d", index))})
				planModel.Fields = append(planModel.Fields, tfTypes.FormField{DisplayName: types.StringValue("URL"), Name: types.StringUnknown()})
			}
			planned := requestSchemaTestPlan(b, planModel)
			request := resource.ModifyPlanRequest{
				Config: tfsdk.Config(requestSchemaTestPlan(b, configModel)),
				State:  tfsdk.State(requestSchemaTestPlan(b, stateModel)), Plan: planned,
			}
			r := requestSchemaTestType[resource.ResourceWithModifyPlan](b, NewRequestSchemaResource())
			b.ReportAllocs()
			for b.Loop() {
				response := resource.ModifyPlanResponse{Plan: request.Plan}
				r.ModifyPlan(ctx, request, &response)
				if response.Diagnostics.HasError() {
					b.Fatal(response.Diagnostics)
				}
				var last types.String
				if diagnostics := response.Plan.GetAttribute(ctx, path.Root("fields").AtListIndex(count-1).AtName("name"), &last); diagnostics.HasError() || !last.Equal(stateModel.Fields[count-1].Name) {
					b.Fatalf("name not preserved: name=%v diagnostics=%v", last, diagnostics)
				}
			}
		})
	}
}
