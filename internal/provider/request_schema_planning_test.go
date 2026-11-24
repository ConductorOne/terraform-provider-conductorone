package provider

import (
	"context"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestRequestSchemaNamePlanningProtocol(t *testing.T) {
	ctx := context.Background()
	for _, rename := range []bool{false, true} {
		t.Run(map[bool]string{false: "metadata change preserves name", true: "label change regenerates name"}[rename], func(t *testing.T) {
			model := &RequestSchemaResourceModel{
				ID: types.StringValue("schema-id"), Name: types.StringValue("Site"),
				Fields: []tfTypes.FormField{{
					Name: types.StringValue("URL_0"), DisplayName: types.StringValue("URL"),
					StringField: &tfTypes.FormStringField{TextField: &tfTypes.TextField{Multiline: types.BoolValue(false)}},
				}},
			}
			prior := requestSchemaTestPlan(t, model)
			model.Description = types.StringValue("Updated form description")
			model.Fields[0].Name = types.StringUnknown()
			if rename {
				model.Fields[0].DisplayName = types.StringValue("New URL")
			}
			proposed := requestSchemaTestPlan(t, model)
			model.ID = types.StringNull()
			model.Fields[0].Name = types.StringNull()
			configured := requestSchemaTestPlan(t, model)
			dynamic := func(plan tfsdk.Plan) *tfprotov6.DynamicValue {
				value, err := tfprotov6.NewDynamicValue(plan.Schema.Type().TerraformType(ctx), plan.Raw)
				if err != nil {
					t.Fatal(err)
				}
				return &value
			}
			server, err := testAccProviderFactories["conductorone"]()
			if err != nil {
				t.Fatal(err)
			}
			response, err := server.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{
				TypeName: "conductorone_request_schema", Config: dynamic(configured),
				PriorState: dynamic(prior), ProposedNewState: dynamic(proposed),
			})
			if err != nil {
				t.Fatal(err)
			}
			for _, diagnostic := range response.Diagnostics {
				if diagnostic.Severity == tfprotov6.DiagnosticSeverityError {
					t.Fatal(diagnostic)
				}
			}
			if len(response.RequiresReplace) != 0 || response.PlannedState == nil {
				t.Fatalf("unexpected replacement or missing state: %v", response)
			}
			raw, err := response.PlannedState.Unmarshal(proposed.Schema.Type().TerraformType(ctx))
			if err != nil {
				t.Fatal(err)
			}
			result := tfsdk.Plan{Schema: proposed.Schema, Raw: raw}
			var name types.String
			if diagnostics := result.GetAttribute(ctx, path.Root("fields").AtListIndex(0).AtName("name"), &name); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			if rename {
				if !name.IsUnknown() {
					t.Fatalf("renamed field must await server-generated name: %v", name)
				}
			} else if name.ValueString() != "URL_0" || name.IsUnknown() || name.IsNull() {
				t.Fatalf("unchanged label must preserve its known name: %v", name)
			}
		})
	}
}

func TestRequestSchemaNamePlanningBoundaries(t *testing.T) {
	ctx := context.Background()
	for _, boundary := range []string{"create", "destroy", "null prior name", "unknown prior name", "known planned name", "unknown configured fields", "null configured label", "unknown stored fields", "unknown planned fields"} {
		t.Run(boundary, func(t *testing.T) {
			model := &RequestSchemaResourceModel{Fields: []tfTypes.FormField{{Name: types.StringNull(), DisplayName: types.StringValue("URL")}}}
			configured := requestSchemaTestPlan(t, model)
			model.Fields[0].Name = types.StringValue("URL_0")
			stored := requestSchemaTestPlan(t, model)
			model.Fields[0].Name = types.StringUnknown()
			planned := requestSchemaTestPlan(t, model)
			set := func(plan *tfsdk.Plan, target path.Path, value any) {
				if diagnostics := plan.SetAttribute(ctx, target, value); diagnostics.HasError() {
					t.Fatal(diagnostics)
				}
			}
			namePath := path.Root("fields").AtListIndex(0).AtName("name")
			var fields types.List
			if diagnostics := configured.GetAttribute(ctx, path.Root("fields"), &fields); diagnostics.HasError() {
				t.Fatal(diagnostics)
			}
			switch boundary {
			case "create":
				stored.Raw = tftypes.NewValue(stored.Schema.Type().TerraformType(ctx), nil)
			case "destroy":
				planned.Raw = tftypes.NewValue(planned.Schema.Type().TerraformType(ctx), nil)
			case "null prior name":
				set(&stored, namePath, types.StringNull())
			case "unknown prior name":
				set(&stored, namePath, types.StringUnknown())
			case "known planned name":
				set(&planned, namePath, types.StringValue("Known_0"))
			case "unknown configured fields":
				set(&configured, path.Root("fields"), types.ListUnknown(fields.ElementType(ctx)))
			case "null configured label":
				set(&configured, path.Root("fields").AtListIndex(0).AtName("display_name"), types.StringNull())
			case "unknown stored fields":
				set(&stored, path.Root("fields"), types.ListUnknown(fields.ElementType(ctx)))
			case "unknown planned fields":
				set(&planned, path.Root("fields"), types.ListUnknown(fields.ElementType(ctx)))
			}
			response := resource.ModifyPlanResponse{Plan: planned}
			requestSchemaTestType[resource.ResourceWithModifyPlan](t, NewRequestSchemaResource()).ModifyPlan(ctx, resource.ModifyPlanRequest{
				Config: tfsdk.Config(configured), State: tfsdk.State(stored), Plan: planned,
			}, &response)
			if response.Diagnostics.HasError() || !response.Plan.Raw.Equal(planned.Raw) || len(response.RequiresReplace) != 0 {
				t.Fatalf("%s must not rewrite the plan: diagnostics=%v replacement=%v", boundary, response.Diagnostics, response.RequiresReplace)
			}
		})
	}
}
