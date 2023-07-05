package sdk

import (
	"testing"
)

func TestNormalizeTenantInput(t *testing.T) {

	tests := []struct {
		name               string
		input              string
		wantedURL          string
		wantedTenantDomain string
		wantErr            bool
	}{
		{
			name:               "base",
			input:              "insulator.conductor.one",
			wantedURL:          "",
			wantedTenantDomain: "insulator",
			wantErr:            false,
		},
		{
			name:               "tenantOnly",
			input:              "insulator",
			wantedURL:          "",
			wantedTenantDomain: "insulator",
			wantErr:            false,
		},
		{
			name:               "baseWithHTTP",
			input:              "https://insulator.conductor.one",
			wantedURL:          "",
			wantedTenantDomain: "insulator",
			wantErr:            false,
		},
		{
			name:               "differentUrlWithHTTP",
			input:              "https://insulator.a.different.com",
			wantedURL:          "https://insulator.a.different.com",
			wantedTenantDomain: "invalid-example", // Doesn't match {tenant}.conductor.one, so the go sdk default is not changed
			wantErr:            false,
		},
		{
			name:               "differentUrlWithHTTP",
			input:              "insulator.a.different.com",
			wantedURL:          "https://insulator.a.different.com",
			wantedTenantDomain: "invalid-example", // Doesn't match {tenant}.conductor.one, so the go sdk default is not changed
			wantErr:            false,
		},
		{
			name:               "differentURL",
			input:              "insulator.a.different.com",
			wantedURL:          "https://insulator.a.different.com",
			wantedTenantDomain: "invalid-example", // Doesn't match {tenant}.conductor.one, so the go sdk default is not changed
			wantErr:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WithTenant(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithTenant returned err = '%v', but expected error: '%v'", err, tt.wantErr)
			}
			client := New(got)
			if client.sdkConfiguration.ServerURL != tt.wantedURL {
				t.Errorf("ServerURL Got = '%v', want '%v'", client.sdkConfiguration.ServerURL, tt.wantedURL)
			}
			if client.sdkConfiguration.ServerDefaults[0]["tenantDomain"] != tt.wantedTenantDomain {
				t.Errorf("TenantDomain Got = '%v', want '%v'", client.sdkConfiguration.ServerDefaults[0]["tenantDomain"], tt.wantedTenantDomain)
			}
		})
	}
}
