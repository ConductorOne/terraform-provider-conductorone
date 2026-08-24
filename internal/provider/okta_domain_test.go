package provider

import (
	"context"
	"testing"

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
