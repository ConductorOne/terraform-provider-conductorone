package provider

import (
	"context"
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
		{"custom domain admin unchanged", "login-admin.mycompany.com", "login-admin.mycompany.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeOktaDomain(tt.input); got != tt.want {
				t.Errorf("normalizeOktaDomain(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestNormalizeDomainModifier verifies the plan modifier normalizes the planned
// value so it matches the API-normalized state (the G4 contract: apply succeeds
// and no drift on subsequent plan). Without it, Terraform's post-apply
// consistency check fails with "Provider produced inconsistent result after
// apply".
func TestNormalizeDomainModifier(t *testing.T) {
	ctx := context.Background()
	m := normalizeDomainModifier{}

	tests := []struct {
		name   string
		config types.String
		want   string
	}{
		{"admin domain normalized", types.StringValue("integrator-3535680-admin.okta.com"), "integrator-3535680.okta.com"},
		{"already normalized unchanged", types.StringValue("integrator-3535680.okta.com"), "integrator-3535680.okta.com"},
		{"empty unchanged", types.StringValue(""), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := planmodifier.StringRequest{ConfigValue: tt.config}
			resp := &planmodifier.StringResponse{}
			m.PlanModifyString(ctx, req, resp)
			if resp.PlanValue.ValueString() != tt.want {
				t.Errorf("PlanModifyString(%q) = %q, want %q", tt.config.ValueString(), resp.PlanValue.ValueString(), tt.want)
			}
		})
	}

	// Null config: plan value must be left untouched (no panic, no change).
	t.Run("null config untouched", func(t *testing.T) {
		req := planmodifier.StringRequest{ConfigValue: types.StringNull()}
		resp := &planmodifier.StringResponse{}
		m.PlanModifyString(ctx, req, resp)
		if !resp.PlanValue.IsNull() {
			t.Errorf("null config: PlanValue = %q, want null", resp.PlanValue.ValueString())
		}
	})
}

// TestDomainAttributesHavePlanModifier verifies the normalizeDomainModifier is
// actually wired into the four Okta domain attributes' PlanModifiers. A
// modifier that is never wired into the schema does nothing — this catches the
// missing-wiring regression that would leave G4 (apply succeeds + no drift)
// unfixed while the modifier unit test still passes.
func TestDomainAttributesHavePlanModifier(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name     string
		resource resource.Resource
		attrName string
	}{
		{"okta", &IntegrationOktaResource{}, "okta_domain"},
		{"okta_v2", &IntegrationOktaV2Resource{}, "okta_v2_domain"},
		{"okta_aws_federation", &IntegrationOktaAwsFederationResource{}, "okta_aws_federation_domain"},
		{"okta_ciam", &IntegrationOktaCiamResource{}, "okta_ciam_domain"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := resource.SchemaRequest{}
			resp := &resource.SchemaResponse{}
			tc.resource.Schema(ctx, req, resp)
			attr, ok := resp.Schema.Attributes[tc.attrName].(*schema.StringAttribute)
			if !ok {
				t.Fatalf("attribute %q is not a schema.StringAttribute", tc.attrName)
			}
			found := false
			for _, pm := range attr.PlanModifiers {
				if _, isNorm := pm.(normalizeDomainModifier); isNorm {
					found = true
				}
			}
			if !found {
				t.Errorf("attribute %q missing normalizeDomainModifier in PlanModifiers", tc.attrName)
			}
		})
	}
}

// TestPopulateConfigNormalizesDomain verifies the production path: each of the
// four Okta resources' populateConfig() sends the normalized domain to the API.
func TestPopulateConfigNormalizesDomain(t *testing.T) {
	const admin = "integrator-3535680-admin.okta.com"
	const normalized = "integrator-3535680.okta.com"

	// okta v1
	cfg1 := (&IntegrationOktaResourceModel{OktaDomain: types.StringValue(admin)}).populateConfig()
	if got := cfg1["okta_domain"].(*string); got == nil || *got != normalized {
		t.Errorf("okta_domain populateConfig = %v, want %q", got, normalized)
	}

	// okta v2
	cfg2 := (&IntegrationOktaV2ResourceModel{OktaV2Domain: types.StringValue(admin)}).populateConfig()
	if got := cfg2["okta_v2_domain"].(*string); got == nil || *got != normalized {
		t.Errorf("okta_v2_domain populateConfig = %v, want %q", got, normalized)
	}

	// okta aws federation
	cfg3 := (&IntegrationOktaAwsFederationResourceModel{OktaAwsFederationDomain: types.StringValue(admin)}).populateConfig()
	if got := cfg3["okta_aws_federation_domain"].(*string); got == nil || *got != normalized {
		t.Errorf("okta_aws_federation_domain populateConfig = %v, want %q", got, normalized)
	}

	// okta ciam
	cfg4 := (&IntegrationOktaCiamResourceModel{OktaCiamDomain: types.StringValue(admin)}).populateConfig()
	if got := cfg4["okta_ciam_domain"].(*string); got == nil || *got != normalized {
		t.Errorf("okta_ciam_domain populateConfig = %v, want %q", got, normalized)
	}
}
