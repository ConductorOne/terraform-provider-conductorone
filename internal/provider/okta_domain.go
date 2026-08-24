package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeOktaDomain strips a trailing "-admin" from the hostname portion
// (the first '.'-delimited component) of an Okta domain, matching the C1 API's
// server-side normalization. For already-normalized domains it is idempotent:
// an already-normalized domain is returned unchanged.
func normalizeOktaDomain(domain string) string {
	if domain == "" {
		return ""
	}
	hostname, suffix, hasSuffix := strings.Cut(domain, ".")
	if strings.HasSuffix(hostname, "-admin") {
		hostname = strings.TrimSuffix(hostname, "-admin")
	}
	if hasSuffix {
		return hostname + "." + suffix
	}
	return hostname
}

// oktaDomainPlanModifier normalizes the planned value of an Okta domain
// attribute to match the C1 API's server-side normalization, so the planned
// value equals the applied state value. Without it, Terraform's strict
// plan-vs-state equality check fails with "Provider produced inconsistent
// result after apply" when a user configures an Okta admin domain (e.g.
// "integrator-3535680-admin.okta.com"), because the server stores the
// normalized form.
type oktaDomainPlanModifier struct{}

func (m oktaDomainPlanModifier) Description(_ context.Context) string {
	return "Normalizes the Okta domain (strips a trailing -admin from the hostname) so the planned value matches the server-normalized state."
}

func (m oktaDomainPlanModifier) MarkdownDescription(_ context.Context) string {
	return m.Description(context.Background())
}

func (m oktaDomainPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}
	resp.PlanValue = types.StringValue(normalizeOktaDomain(req.ConfigValue.ValueString()))
}
