package provider

import (
	"context"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The v1 Tailscale catalog entry is retired: creating it returns
// `400 connector is retired` (IGA-3835). The v2 resource must point at the v2
// catalog entry instead.
func TestTailscaleV2CatalogID(t *testing.T) {
	if got, want := tailscaleV2CatalogID, "2zvVTe8SQT4qEcTTMXJTrmG9OwV"; got != want {
		t.Errorf("tailscaleV2CatalogID = %q, want %q", got, want)
	}
	if tailscaleV2CatalogID == tailscaleCatalogID {
		t.Errorf("tailscale v2 resource is using the retired v1 catalog ID %q", tailscaleCatalogID)
	}
}

func tailscaleV2Schema(t *testing.T) schema.Schema {
	t.Helper()

	resp := &resource.SchemaResponse{}
	r, ok := NewIntegrationTailscaleV2Resource().(*IntegrationTailscaleV2Resource)
	if !ok {
		t.Fatal("NewIntegrationTailscaleV2Resource did not return an *IntegrationTailscaleV2Resource")
	}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected schema diagnostics: %v", resp.Diagnostics)
	}
	return resp.Schema
}

func TestTailscaleV2SchemaFieldGroups(t *testing.T) {
	s := tailscaleV2Schema(t)

	for _, group := range []struct {
		name       string
		attributes []string
	}{
		{
			name:       "api_key_flow_group",
			attributes: []string{"api_key", "tailnet", "ignore_ephemeral_devices"},
		},
		{
			name:       "oauth_client_credentials_flow_group",
			attributes: []string{"tailscale_client_id", "tailscale_client_secret", "tailnet", "ignore_ephemeral_devices"},
		},
	} {
		attribute, ok := s.Attributes[group.name]
		if !ok {
			t.Errorf("schema is missing the %q attribute", group.name)
			continue
		}
		nested, ok := attribute.(*schema.SingleNestedAttribute)
		if !ok {
			t.Errorf("%q = %T, want *schema.SingleNestedAttribute", group.name, attribute)
			continue
		}
		if len(nested.Attributes) != len(group.attributes) {
			t.Errorf("%q has %d attributes, want %d", group.name, len(nested.Attributes), len(group.attributes))
		}
		for _, name := range group.attributes {
			if _, ok := nested.Attributes[name]; !ok {
				t.Errorf("%q is missing the %q attribute", group.name, name)
			}
		}
	}
}

func TestTailscaleV2SecretsAreSensitive(t *testing.T) {
	s := tailscaleV2Schema(t)

	for _, tt := range []struct{ group, attribute string }{
		{"api_key_flow_group", "api_key"},
		{"oauth_client_credentials_flow_group", "tailscale_client_secret"},
	} {
		nested, ok := s.Attributes[tt.group].(*schema.SingleNestedAttribute)
		if !ok {
			t.Fatalf("%q is not a nested attribute", tt.group)
		}
		attribute, ok := nested.Attributes[tt.attribute].(*schema.StringAttribute)
		if !ok {
			t.Fatalf("%q.%q is not a string attribute", tt.group, tt.attribute)
		}
		if !attribute.Sensitive {
			t.Errorf("%q.%q is not marked sensitive", tt.group, tt.attribute)
		}
	}
}

// Terraform attribute names cannot contain dashes, but the connector's config
// fields do. populateConfig must translate back, otherwise the connector never
// sees its credentials.
func TestTailscaleV2PopulateConfigUsesConnectorFieldNames(t *testing.T) {
	for _, tt := range []struct {
		name  string
		model IntegrationTailscaleV2ResourceModel
		want  map[string]string
	}{
		{
			name: "api key flow",
			model: IntegrationTailscaleV2ResourceModel{
				ApiKeyFlowGroup: types.ObjectValueMust(
					map[string]attr.Type{
						"api_key":                  types.StringType,
						"tailnet":                  types.StringType,
						"ignore_ephemeral_devices": types.BoolType,
					},
					map[string]attr.Value{
						"api_key":                  types.StringValue("tskey-api-secret"),
						"tailnet":                  types.StringValue("example.com"),
						"ignore_ephemeral_devices": types.BoolValue(true),
					},
				),
				OauthClientCredentialsFlowGroup: types.ObjectNull(map[string]attr.Type{}),
			},
			want: map[string]string{
				"C1_selected_field_group_name": "api-key-flow-group",
				"api-key":                      "tskey-api-secret",
				"tailnet":                      "example.com",
				"ignore-ephemeral-devices":     "true",
			},
		},
		{
			name: "oauth client credentials flow",
			model: IntegrationTailscaleV2ResourceModel{
				ApiKeyFlowGroup: types.ObjectNull(map[string]attr.Type{}),
				OauthClientCredentialsFlowGroup: types.ObjectValueMust(
					map[string]attr.Type{
						"tailscale_client_id":      types.StringType,
						"tailscale_client_secret":  types.StringType,
						"tailnet":                  types.StringType,
						"ignore_ephemeral_devices": types.BoolType,
					},
					map[string]attr.Value{
						"tailscale_client_id":      types.StringValue("client-id"),
						"tailscale_client_secret":  types.StringValue("client-secret"),
						"tailnet":                  types.StringValue("example.com"),
						"ignore_ephemeral_devices": types.BoolValue(false),
					},
				),
			},
			want: map[string]string{
				"C1_selected_field_group_name": "oauth-client-credentials-flow-group",
				"tailscale-client-id":          "client-id",
				"tailscale-client-secret":      "client-secret",
				"tailnet":                      "example.com",
				"ignore-ephemeral-devices":     "false",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.model.populateConfig()
			if len(got) != len(tt.want) {
				t.Errorf("populateConfig() has %d keys (%v), want %d", len(got), got, len(tt.want))
			}
			for key, want := range tt.want {
				value, ok := got[key]
				if !ok {
					t.Errorf("populateConfig() is missing the %q key, got %v", key, got)
					continue
				}
				if value != want {
					t.Errorf("populateConfig()[%q] = %v, want %q", key, value, want)
				}
			}
		})
	}
}

// RefreshFromGetResponse builds the group object by hand; its attribute names
// have to match the ones declared in the schema or terraform rejects the state
// with a value conversion error.
func TestTailscaleV2RefreshUsesSchemaAttributeNames(t *testing.T) {
	data := &IntegrationTailscaleV2ResourceModel{
		ApiKeyFlowGroup: types.ObjectValueMust(
			map[string]attr.Type{
				"api_key":                  types.StringType,
				"tailnet":                  types.StringType,
				"ignore_ephemeral_devices": types.BoolType,
			},
			map[string]attr.Value{
				"api_key":                  types.StringValue("tskey-api-secret"),
				"tailnet":                  types.StringValue("example.com"),
				"ignore_ephemeral_devices": types.BoolValue(true),
			},
		),
		OauthClientCredentialsFlowGroup: types.ObjectNull(map[string]attr.Type{}),
	}

	// The API returns each config value as {"stringValue": "..."}.
	stringValue := func(v string) map[string]interface{} {
		return map[string]interface{}{"stringValue": v}
	}
	data.RefreshFromGetResponse(&shared.Connector{
		Config: makeConnectorConfig(map[string]interface{}{
			"C1_selected_field_group_name": stringValue("api-key-flow-group"),
			"api-key":                      stringValue("tskey-api-secret"),
			"tailnet":                      stringValue("example.com"),
			"ignore-ephemeral-devices":     stringValue("true"),
		}),
	})

	nested, ok := tailscaleV2Schema(t).Attributes["api_key_flow_group"].(*schema.SingleNestedAttribute)
	if !ok {
		t.Fatal("api_key_flow_group is not a nested attribute")
	}
	refreshed := data.ApiKeyFlowGroup.Attributes()
	if len(refreshed) != len(nested.Attributes) {
		t.Errorf("refreshed state has %d attributes (%v), want %d", len(refreshed), refreshed, len(nested.Attributes))
	}
	for name := range refreshed {
		if _, ok := nested.Attributes[name]; !ok {
			t.Errorf("refreshed state has attribute %q, which the schema does not declare", name)
		}
	}
	for name, want := range map[string]attr.Value{
		"api_key":                  types.StringValue("tskey-api-secret"),
		"tailnet":                  types.StringValue("example.com"),
		"ignore_ephemeral_devices": types.BoolValue(true),
	} {
		got, ok := refreshed[name]
		if !ok {
			t.Errorf("refreshed state is missing the %q attribute", name)
			continue
		}
		if !got.Equal(want) {
			t.Errorf("refreshed %q = %v, want %v", name, got, want)
		}
	}
}
