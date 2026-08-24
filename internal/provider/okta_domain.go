package provider

import "strings"

// normalizeOktaDomain strips a trailing "-admin" from the hostname portion
// (the first '.'-delimited component) of an Okta domain, matching the C1 API's
// server-side normalization. It is idempotent: an already-normalized domain is
// returned unchanged.
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
