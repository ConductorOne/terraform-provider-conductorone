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
		{"double admin suffix idempotent", "x-admin-admin.okta.com", "x.okta.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeOktaDomain(tt.input); got != tt.want {
				t.Errorf("normalizeOktaDomain(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNormalizeOktaDomainPlanModifier(t *testing.T) {
	mod := normalizeOktaDomainPlanModifier()

	// A planned value with a trailing -admin is normalized so it matches the
	// normalized value the API stores and returns, avoiding a perpetual diff.
	req := planmodifier.StringRequest{
		PlanValue: types.StringValue("tenant-admin.okta.com"),
	}
	resp := &planmodifier.StringResponse{}
	mod.PlanModifyString(context.Background(), req, resp)
	if got, want := resp.PlanValue.ValueString(), "tenant.okta.com"; got != want {
		t.Errorf("plan modifier normalized %q, want %q", got, want)
	}

	// An already-normalized planned value passes through unchanged.
	req = planmodifier.StringRequest{
		PlanValue: types.StringValue("tenant.okta.com"),
	}
	resp = &planmodifier.StringResponse{}
	mod.PlanModifyString(context.Background(), req, resp)
	if got, want := resp.PlanValue.ValueString(), "tenant.okta.com"; got != want {
		t.Errorf("plan modifier changed %q, want %q", got, want)
	}

	// A null planned value is left untouched.
	req = planmodifier.StringRequest{
		PlanValue: types.StringNull(),
	}
	resp = &planmodifier.StringResponse{}
	mod.PlanModifyString(context.Background(), req, resp)
	if !resp.PlanValue.IsNull() {
		t.Errorf("plan modifier changed a null plan value to %q", resp.PlanValue.ValueString())
	}
}
