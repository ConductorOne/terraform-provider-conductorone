package provider

import (
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// These tests verify that normalizeOktaDomain is actually wired into the
// write path (populateConfig) and the read path (RefreshFromGetResponse) of
// all four Okta integration resources. A missed call site would let the
// perpetual-diff bug (write normalizes, read does not) slip through even
// though the helper itself is unit-tested.

func connectorWithDomain(key, domain string) *shared.Connector {
	return &shared.Connector{
		Config: &shared.Config{
			AtType: sdk.String(envConfigType),
			AdditionalProperties: map[string]interface{}{
				"configuration": map[string]interface{}{
					key: map[string]interface{}{"stringValue": domain},
				},
			},
		},
	}
}

func TestPopulateConfigNormalizesOktaDomain(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		config map[string]interface{}
	}{
		{"okta", "okta_domain", (&IntegrationOktaResourceModel{OktaDomain: types.StringValue("tenant-admin.okta.com")}).populateConfig()},
		{"okta_v2", "okta_v2_domain", (&IntegrationOktaV2ResourceModel{OktaV2Domain: types.StringValue("tenant-admin.okta.com")}).populateConfig()},
		{"okta_ciam", "okta_ciam_domain", (&IntegrationOktaCiamResourceModel{OktaCiamDomain: types.StringValue("tenant-admin.okta.com")}).populateConfig()},
		{"okta_aws_federation", "okta_aws_federation_domain", (&IntegrationOktaAwsFederationResourceModel{OktaAwsFederationDomain: types.StringValue("tenant-admin.okta.com")}).populateConfig()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.config[tt.key].(*string)
			if !ok || got == nil {
				t.Fatalf("populateConfig() missing key %q", tt.key)
			}
			if *got != "tenant.okta.com" {
				t.Errorf("populateConfig()[%q] = %q, want %q", tt.key, *got, "tenant.okta.com")
			}
		})
	}
}

func TestRefreshFromCreateResponseNormalizesOktaDomain(t *testing.T) {
	tests := []struct {
		name  string
		check func() string
	}{
		{"okta", func() string {
			m := IntegrationOktaResourceModel{}
			m.RefreshFromCreateResponse(connectorWithDomain("okta_domain", "tenant-admin.okta.com"))
			return m.OktaDomain.ValueString()
		}},
		{"okta_v2", func() string {
			m := IntegrationOktaV2ResourceModel{}
			m.RefreshFromCreateResponse(connectorWithDomain("okta_v2_domain", "tenant-admin.okta.com"))
			return m.OktaV2Domain.ValueString()
		}},
		{"okta_ciam", func() string {
			m := IntegrationOktaCiamResourceModel{}
			m.RefreshFromCreateResponse(connectorWithDomain("okta_ciam_domain", "tenant-admin.okta.com"))
			return m.OktaCiamDomain.ValueString()
		}},
		{"okta_aws_federation", func() string {
			m := IntegrationOktaAwsFederationResourceModel{}
			m.RefreshFromCreateResponse(connectorWithDomain("okta_aws_federation_domain", "tenant-admin.okta.com"))
			return m.OktaAwsFederationDomain.ValueString()
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.check(); got != "tenant.okta.com" {
				t.Errorf("RefreshFromCreateResponse normalized domain = %q, want %q", got, "tenant.okta.com")
			}
		})
	}
}

func TestRefreshFromGetResponseNormalizesOktaDomain(t *testing.T) {
	tests := []struct {
		name  string
		check func() string
	}{
		{"okta", func() string {
			m := IntegrationOktaResourceModel{}
			m.RefreshFromGetResponse(connectorWithDomain("okta_domain", "tenant-admin.okta.com"))
			return m.OktaDomain.ValueString()
		}},
		{"okta_v2", func() string {
			m := IntegrationOktaV2ResourceModel{}
			m.RefreshFromGetResponse(connectorWithDomain("okta_v2_domain", "tenant-admin.okta.com"))
			return m.OktaV2Domain.ValueString()
		}},
		{"okta_ciam", func() string {
			m := IntegrationOktaCiamResourceModel{}
			m.RefreshFromGetResponse(connectorWithDomain("okta_ciam_domain", "tenant-admin.okta.com"))
			return m.OktaCiamDomain.ValueString()
		}},
		{"okta_aws_federation", func() string {
			m := IntegrationOktaAwsFederationResourceModel{}
			m.RefreshFromGetResponse(connectorWithDomain("okta_aws_federation_domain", "tenant-admin.okta.com"))
			return m.OktaAwsFederationDomain.ValueString()
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.check(); got != "tenant.okta.com" {
				t.Errorf("RefreshFromGetResponse normalized domain = %q, want %q", got, "tenant.okta.com")
			}
		})
	}
}
