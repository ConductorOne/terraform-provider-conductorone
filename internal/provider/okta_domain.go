package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeOktaDomain strips trailing "-admin" occurrences from the hostname
// portion (the first '.'-delimited component) of an Okta domain, matching the
// C1 API's server-side normalization. It is idempotent: an already-normalized
// domain is returned unchanged, and repeated application never changes the
// result.
//
// It is applied on the write path (populateConfig), the read path
// (RefreshFromGetResponse/RefreshFromCreateResponse), and the plan path
// (oktaDomainPlanModifier) so the planned value, the value sent to the API,
// and the stored state all hold the canonical form. An "-admin" domain is
// accepted on write and normalized automatically; the plan shows the
// normalized value, so there is no perpetual diff.
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

// oktaDomainPlanModifier normalizes the planned value of an Okta domain
// attribute so the plan matches the state the provider will store after apply.
// Without it, Terraform's "Provider produced inconsistent result after apply"
// check compares the raw config value (e.g. "tenant-admin.okta.com") against
// the normalized state value ("tenant.okta.com") and fails the apply.
//
// Null and unknown planned values are passed through unchanged: a null value
// (attribute omitted) must stay null, and an unknown value (e.g. a computed
// reference) must not be clobbered.
type oktaDomainPlanModifier struct{}

func (m oktaDomainPlanModifier) Description(_ context.Context) string {
	return "Normalizes the Okta domain to its canonical (non-admin) form in the plan."
}

func (m oktaDomainPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m oktaDomainPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		resp.PlanValue = req.PlanValue
		return
	}
	resp.PlanValue = types.StringValue(normalizeOktaDomain(req.PlanValue.ValueString()))
}
