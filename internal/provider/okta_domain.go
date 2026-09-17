package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeOktaDomain strips a trailing "-admin" from the hostname portion
// (the first '.'-delimited component) of an Okta domain, matching the C1 API's
// server-side normalization. It is idempotent: an already-normalized domain is
// returned unchanged, and applying it twice is a no-op.
func normalizeOktaDomain(domain string) string {
	if domain == "" {
		return ""
	}
	hostname, suffix, hasSuffix := strings.Cut(domain, ".")
	for strings.HasSuffix(hostname, "-admin") {
		hostname = strings.TrimSuffix(hostname, "-admin")
	}
	if hasSuffix {
		return hostname + "." + suffix
	}
	return hostname
}

// normalizeOktaDomainPlanModifier returns a plan modifier that normalizes the
// planned value of an Okta domain attribute. The C1 API stores the normalized
// domain and returns it in the read response, so without this modifier a config
// value like "tenant-admin.okta.com" would produce a perpetual plan diff
// against the normalized value ("tenant.okta.com") held in state.
func normalizeOktaDomainPlanModifier() planmodifier.String {
	return oktaDomainPlanModifier{}
}

type oktaDomainPlanModifier struct{}

func (m oktaDomainPlanModifier) Description(_ context.Context) string {
	return "Normalizes the Okta domain so a trailing -admin does not cause a perpetual plan diff."
}

func (m oktaDomainPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m oktaDomainPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		return
	}
	resp.PlanValue = types.StringValue(normalizeOktaDomain(req.PlanValue.ValueString()))
}
