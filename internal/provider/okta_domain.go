package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeOktaDomain strips a single trailing "-admin" from the hostname
// portion (the first '.'-delimited component) of an Okta domain, matching the
// C1 API's server-side normalization. The strip is gated on the known Okta
// TLDs (okta.com, oktapreview.com, okta-emea.com) or a suffix-less domain, so
// a legitimate custom domain whose hostname happens to end in "-admin" (e.g.
// login-admin.mycompany.com) is left unchanged. An already-normalized domain
// (hostname not ending in "-admin") is returned unchanged.
func normalizeOktaDomain(domain string) string {
	if domain == "" {
		return ""
	}
	hostname, suffix, hasSuffix := strings.Cut(domain, ".")
	if strings.HasSuffix(hostname, "-admin") && (!hasSuffix || suffix == "okta.com" || suffix == "oktapreview.com" || suffix == "okta-emea.com") {
		hostname = strings.TrimSuffix(hostname, "-admin")
	}
	if hasSuffix {
		return hostname + "." + suffix
	}
	return hostname
}

// normalizeDomainModifier is a plan modifier that normalizes the Okta domain
// during planning, so the planned value matches the value the C1 API returns
// (which is normalized server-side). Without it, Terraform's post-apply
// consistency check compares the user's un-normalized config value against the
// API-normalized state and fails with "Provider produced inconsistent result
// after apply".
type normalizeDomainModifier struct{}

func (m normalizeDomainModifier) Description(_ context.Context) string {
	return "Normalizes the Okta domain by stripping a trailing -admin from the hostname."
}

func (m normalizeDomainModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m normalizeDomainModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = types.StringValue(normalizeOktaDomain(req.ConfigValue.ValueString()))
}
