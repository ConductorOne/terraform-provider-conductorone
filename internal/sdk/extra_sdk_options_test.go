package sdk

import (
	"testing"
)

type ClientConfigTest struct {
	name               string
	input              string
	wantedURL          string
	wantedTenantDomain string
	tenantOnly         bool
	wantErr            bool
}

func TestParseClientID(t *testing.T) {
	tests := []ClientConfigTest{
		{name: "base", input: "foobar@insulator.conductor.one/pcc", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "dev", input: "foobar@c1dev.anthony.dev.ductone.com/pcc", wantedURL: "https://c1dev.anthony.dev.ductone.com", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseClientID(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseClientID returned err = '%v', but expected error: '%v'", err, tt.wantErr)
			}
			testClientConfig(t, tt, got)
		})
	}
}

func TestCustomTenantOptions(t *testing.T) {
	tests := []ClientConfigTest{
		{name: "base", input: "insulator.conductor.one", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "tenantOnly", input: "insulator", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "baseWithHTTP", input: "https://insulator.conductor.one", wantedURL: "https://insulator.conductor.one", wantedTenantDomain: "insulator", wantErr: false, tenantOnly: true},
		{name: "differentURL", input: "https://insulator.a.different.com", wantedURL: "https://insulator.a.different.com", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
		{name: "differentBase", input: "insulator.a.different.com", wantedURL: "https://insulator.a.different.com", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
		{name: "devURL", input: "https://c1dev.anthony.dev.ductone.com:2443", wantedURL: "https://c1dev.anthony.dev.ductone.com:2443", wantedTenantDomain: "", wantErr: false, tenantOnly: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WithTenantCustom(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WithTenant returned err = '%v', but expected error: '%v'", err, tt.wantErr)
			}
			customOptions := &CustomOptions{}
			got(customOptions)
			testClientConfig(t, tt, customOptions.ClientConfig)
		})
	}
}

func testClientConfig(t *testing.T, tt ClientConfigTest, got *ClientConfig) {
	if got.UseWithServer() == got.UseWithTenant() {
		t.Error("useWithServer and useWithTenant should not be the same")
	}
	if got.UseWithTenant() != tt.tenantOnly {
		t.Errorf("useWithTenant Got = '%v', want '%v'", got.UseWithTenant(), tt.tenantOnly)
	}
	if got.UseWithServer() {
		if got.ServerURL() != tt.wantedURL {
			t.Errorf("ServerURL Got = '%v', want '%v'", got.ServerURL(), tt.wantedURL)
		}
		if got.Tenant() != tt.wantedTenantDomain {
			t.Errorf("Tenant Got = '%v', want '%v'", got.Tenant(), "")
		}
	}
	if got.UseWithTenant() {
		if got.Tenant() != tt.wantedTenantDomain {
			t.Errorf("Tenant Got = '%v', want '%v'", got.Tenant(), tt.wantedTenantDomain)
		}
		if got.ServerURL() != "" {
			t.Errorf("ServerURL Got = '%v', want '%v'", got.ServerURL(), "")
		}
		if got.GetServerURL() != tt.wantedURL {
			t.Errorf("GetServerURL Got = '%v', want '%v'", got.GetServerURL(), tt.wantedURL)
		}
	}
}
