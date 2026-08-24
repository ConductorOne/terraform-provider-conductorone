package provider

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNormalizeOktaDomainPlanModifier(t *testing.T) {
	mod := normalizeOktaDomainPlanModifier()

	tests := []struct {
		name string
		plan types.String
		want types.String
	}{
		{"admin form normalized", types.StringValue("integrator-3535680-admin.okta.com"), types.StringValue("integrator-3535680.okta.com")},
		{"already normalized unchanged", types.StringValue("integrator-3535680.okta.com"), types.StringValue("integrator-3535680.okta.com")},
		{"null unchanged", types.StringNull(), types.StringNull()},
		{"unknown unchanged", types.StringUnknown(), types.StringUnknown()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := planmodifier.StringRequest{PlanValue: tt.plan}
			resp := &planmodifier.StringResponse{}
			mod.PlanModifyString(context.Background(), req, resp)
			if resp.PlanValue != tt.want {
				t.Errorf("PlanModifyString(%q) = %q, want %q", tt.plan, resp.PlanValue, tt.want)
			}
		})
	}
}

func TestOktaDomainPlanModifierWired(t *testing.T) {
	cases := []struct {
		name     string
		newRes   func() resource.Resource
		attrName string
	}{
		{"okta", NewIntegrationOktaResource, "okta_domain"},
		{"okta_v2", NewIntegrationOktaV2Resource, "okta_v2_domain"},
		{"okta_ciam", NewIntegrationOktaCiamResource, "okta_ciam_domain"},
		{"okta_aws_federation", NewIntegrationOktaAwsFederationResource, "okta_aws_federation_domain"},
	}
	want := reflect.TypeOf(normalizeOktaDomainPlanModifier())
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &resource.SchemaResponse{}
			tc.newRes().Schema(context.Background(), resource.SchemaRequest{}, resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("building %s schema: %v", tc.name, resp.Diagnostics)
			}
			attr, ok := resp.Schema.Attributes[tc.attrName].(*schema.StringAttribute)
			if !ok {
				t.Fatalf("%s has type %T, want *schema.StringAttribute", tc.attrName, resp.Schema.Attributes[tc.attrName])
			}
			for _, modifier := range attr.PlanModifiers {
				if reflect.TypeOf(modifier) == want {
					return
				}
			}
			t.Fatalf("%s does not have normalizeOktaDomainPlanModifier: %#v", tc.attrName, attr.PlanModifiers)
		})
	}
}
