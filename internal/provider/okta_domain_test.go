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

func TestNormalizeOktaDomain(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"admin domain with suffix", "integrator-3535680-admin.okta.com", "integrator-3535680.okta.com"},
		{"admin domain no suffix", "integrator-3535680-admin", "integrator-3535680"},
		{"already normalized", "integrator-3535680.okta.com", "integrator-3535680.okta.com"},
		{"mycompany admin", "mycompany-admin.okta.com", "mycompany.okta.com"},
		{"mycompany normalized", "mycompany.okta.com", "mycompany.okta.com"},
		{"empty string", "", ""},
		{"no admin in hostname", "already-normal.okta.com", "already-normal.okta.com"},
		{"admin in middle of hostname unchanged", "foo-admin-bar.okta.com", "foo-admin-bar.okta.com"},
		{"admin suffix oktapreview tld", "tenant-admin.oktapreview.com", "tenant.oktapreview.com"},
		{"admin suffix okta-emea tld", "tenant-admin.okta-emea.com", "tenant.okta-emea.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeOktaDomain(tt.input); got != tt.want {
				t.Errorf("normalizeOktaDomain(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestOktaDomainPlanModifier verifies the plan modifier normalizes the planned
// value so it matches the server-normalized state (the fix for the
// "inconsistent result after apply" error), and leaves the plan untouched when
// the config value is null or unknown.
func TestOktaDomainPlanModifier(t *testing.T) {
	m := oktaDomainPlanModifier{}

	tests := []struct {
		name string
		cfg  types.String
		plan types.String // initial plan value
		want types.String
	}{
		{"admin domain normalized", types.StringValue("integrator-3535680-admin.okta.com"), types.StringUnknown(), types.StringValue("integrator-3535680.okta.com")},
		{"bare admin hostname normalized", types.StringValue("integrator-3535680-admin"), types.StringUnknown(), types.StringValue("integrator-3535680")},
		{"already normalized unchanged", types.StringValue("integrator-3535680.okta.com"), types.StringUnknown(), types.StringValue("integrator-3535680.okta.com")},
		{"null config leaves plan unchanged", types.StringNull(), types.StringValue("prior"), types.StringValue("prior")},
		{"unknown config leaves plan unchanged", types.StringUnknown(), types.StringValue("prior"), types.StringValue("prior")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := planmodifier.StringRequest{ConfigValue: tt.cfg, PlanValue: tt.plan}
			resp := &planmodifier.StringResponse{PlanValue: tt.plan}
			m.PlanModifyString(context.Background(), req, resp)
			got := resp.PlanValue
			if got.IsNull() != tt.want.IsNull() || got.IsUnknown() != tt.want.IsUnknown() {
				t.Errorf("plan value null/unknown mismatch: got null=%v unknown=%v, want null=%v unknown=%v",
					got.IsNull(), got.IsUnknown(), tt.want.IsNull(), tt.want.IsUnknown())
				return
			}
			if !got.IsNull() && !got.IsUnknown() && got.ValueString() != tt.want.ValueString() {
				t.Errorf("PlanModifyString(%q) = %q, want %q", tt.cfg.ValueString(), got.ValueString(), tt.want.ValueString())
			}
		})
	}
}

// TestPopulateConfigNormalizesDomain verifies the production create/update
// path emits the normalized domain into the config map sent to the API.
func TestPopulateConfigNormalizesDomain(t *testing.T) {
	r := IntegrationOktaV2ResourceModel{
		OktaV2Domain: types.StringValue("integrator-3535680-admin.okta.com"),
	}
	config := r.populateConfig()
	got, ok := config["okta_v2_domain"].(*string)
	if !ok || got == nil {
		t.Fatalf("config[okta_v2_domain] = %#v, want *string", config["okta_v2_domain"])
	}
	if *got != "integrator-3535680.okta.com" {
		t.Errorf("populateConfig domain = %q, want %q", *got, "integrator-3535680.okta.com")
	}
}

// TestOktaDomainAttributesHavePlanModifier guards the schema wiring: if the
// oktaDomainPlanModifier is ever dropped from one of the 4 domain attributes,
// the "inconsistent result after apply" bug silently returns. This test fails
// if any attribute loses the modifier.
func TestOktaDomainAttributesHavePlanModifier(t *testing.T) {
	resources := []struct {
		name string
		new  func() resource.Resource
		attr string
	}{
		{"okta", NewIntegrationOktaResource, "okta_domain"},
		{"okta_v2", NewIntegrationOktaV2Resource, "okta_v2_domain"},
		{"okta_ciam", NewIntegrationOktaCiamResource, "okta_ciam_domain"},
		{"okta_aws_federation", NewIntegrationOktaAwsFederationResource, "okta_aws_federation_domain"},
	}
	want := reflect.TypeOf(oktaDomainPlanModifier{})
	for _, rc := range resources {
		t.Run(rc.name, func(t *testing.T) {
			resp := &resource.SchemaResponse{}
			rc.new().Schema(context.Background(), resource.SchemaRequest{}, resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("building %s schema: %v", rc.name, resp.Diagnostics)
			}
			attr, ok := resp.Schema.Attributes[rc.attr].(*schema.StringAttribute)
			if !ok {
				t.Fatalf("%s has type %T, want *schema.StringAttribute", rc.attr, resp.Schema.Attributes[rc.attr])
			}
			for _, modifier := range attr.PlanModifiers {
				if reflect.TypeOf(modifier) == want {
					return
				}
			}
			t.Fatalf("%s does not have oktaDomainPlanModifier: %#v", rc.attr, attr.PlanModifiers)
		})
	}
}

// TestOktaDomainPlanMatchesState asserts the plan-vs-state invariant: the
// planned value (plan modifier) and the outbound value (populateConfig, which
// the server stores and returns as state) must agree, so Terraform's strict
// plan-vs-state equality check passes for an admin-domain config.
func TestOktaDomainPlanMatchesState(t *testing.T) {
	cfg := types.StringValue("integrator-3535680-admin.okta.com")

	req := planmodifier.StringRequest{ConfigValue: cfg}
	resp := &planmodifier.StringResponse{}
	oktaDomainPlanModifier{}.PlanModifyString(context.Background(), req, resp)
	planned := resp.PlanValue.ValueString()

	r := IntegrationOktaV2ResourceModel{OktaV2Domain: cfg}
	config := r.populateConfig()
	outbound, ok := config["okta_v2_domain"].(*string)
	if !ok || outbound == nil {
		t.Fatalf("config[okta_v2_domain] = %#v, want *string", config["okta_v2_domain"])
	}

	if planned != *outbound {
		t.Errorf("planned %q != outbound %q; plan-vs-state would mismatch", planned, *outbound)
	}
}

// TestOktaDomainReplanNoPerpetualDiff models the existing-resource re-plan
// scenario: config holds the raw admin domain, state holds the server-
// normalized value. Re-planning must produce a plan equal to state, otherwise
// Terraform would show a perpetual diff on every plan.
func TestOktaDomainReplanNoPerpetualDiff(t *testing.T) {
	cfg := types.StringValue("integrator-3535680-admin.okta.com")
	state := types.StringValue(normalizeOktaDomain(cfg.ValueString())) // server-normalized state

	req := planmodifier.StringRequest{ConfigValue: cfg}
	resp := &planmodifier.StringResponse{}
	oktaDomainPlanModifier{}.PlanModifyString(context.Background(), req, resp)

	if !resp.PlanValue.Equal(state) {
		t.Errorf("re-plan %q != state %q; perpetual diff", resp.PlanValue.ValueString(), state.ValueString())
	}
}
