package provider

import "testing"

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
		{"admin in non-first component unchanged", "tenant.okta-admin.com", "tenant.okta-admin.com"},
		{"uppercase admin suffix unchanged", "Tenant-ADMIN.okta.com", "Tenant-ADMIN.okta.com"},
		{"leading whitespace preserved", " tenant-admin.okta.com", " tenant.okta.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeOktaDomain(tt.input)
			if got != tt.want {
				t.Errorf("normalizeOktaDomain(%q) = %q, want %q", tt.input, got, tt.want)
			}
			// The function must be idempotent for the Okta admin-domain case:
			// normalizing the result again must not change it.
			if again := normalizeOktaDomain(got); again != got {
				t.Errorf("normalizeOktaDomain(%q) not idempotent: second pass = %q, want %q", tt.input, again, got)
			}
		})
	}
}
