package provider

import "strings"

// normalizeOktaDomain strips a trailing "-admin" from the hostname portion
// (the first '.'-delimited component) of an Okta domain, matching the C1 API's
// server-side normalization. It is idempotent for the Okta admin-domain case:
// an already-normalized domain is returned unchanged.
//
// It is applied on both the write path (populateConfig) and the read path
// (RefreshFromGetResponse/RefreshFromCreateResponse) so state always holds the
// canonical form. Users should supply the normalized (non-admin) domain; an
// "-admin" domain is accepted on write but is normalized, so the stored state
// will differ from the config and Terraform will report a diff on the next
// plan.
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
