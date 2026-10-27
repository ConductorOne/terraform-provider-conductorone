package provider

import (
	"context"
	"strings"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	mcpPlanCatalogID = "2xCatalogEntry0000000000001"
	mcpPlanURL       = "https://mcp.example.com/mcp"
)

func mcpPlanSchema(t *testing.T) resource.SchemaResponse {
	t.Helper()
	var schemaResp resource.SchemaResponse
	NewMCPServerResource().Schema(context.Background(), resource.SchemaRequest{}, &schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatal(schemaResp.Diagnostics)
	}
	return schemaResp
}

func mcpPlanValue(t *testing.T, model *MCPServerResourceModel) tftypes.Value {
	t.Helper()
	schemaResp := mcpPlanSchema(t)
	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(context.Background(), model); diags.HasError() {
		t.Fatal(diags)
	}
	return state.Raw
}

func mcpPlanDynamic(t *testing.T, value tftypes.Value) *tfprotov6.DynamicValue {
	t.Helper()
	dv, err := tfprotov6.NewDynamicValue(value.Type(), value)
	if err != nil {
		t.Fatal(err)
	}
	return &dv
}

// mcpProposedNewState mirrors Terraform core: config values, with computed
// attributes that are null in config taken from prior state.
func mcpProposedNewState(t *testing.T, prior, config tftypes.Value) tftypes.Value {
	t.Helper()
	ctx := context.Background()
	schemaResp := mcpPlanSchema(t)
	proposed, err := tftypes.Transform(config, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
		if !v.IsNull() || len(p.Steps()) == 0 {
			return v, nil
		}
		if attribute, err := schemaResp.Schema.AttributeAtTerraformPath(ctx, p); err != nil || !attribute.IsComputed() {
			return v, nil //nolint:nilerr // Block paths have no attribute and keep their config value.
		}
		priorValue, _, err := tftypes.WalkAttributePath(prior, p)
		if err != nil {
			return v, nil //nolint:nilerr // Prior state has no value under a null parent.
		}
		if value, ok := priorValue.(tftypes.Value); ok {
			return value, nil
		}
		return v, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return proposed
}

func mcpPlanServer(t *testing.T) tfprotov6.ProviderServer {
	t.Helper()
	server := providerserver.NewProtocol6(New("test")())()
	if _, err := server.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{}); err != nil {
		t.Fatal(err)
	}
	return server
}

func mcpPlan(t *testing.T, prior, config *MCPServerResourceModel) *tfprotov6.PlanResourceChangeResponse {
	t.Helper()
	priorValue := mcpPlanValue(t, prior)
	configValue := mcpPlanValue(t, config)
	resp, err := mcpPlanServer(t).PlanResourceChange(context.Background(), &tfprotov6.PlanResourceChangeRequest{
		TypeName:         "conductorone_mcp_server",
		PriorState:       mcpPlanDynamic(t, priorValue),
		ProposedNewState: mcpPlanDynamic(t, mcpProposedNewState(t, priorValue, configValue)),
		Config:           mcpPlanDynamic(t, configValue),
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range resp.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("plan diagnostic: %s: %s", d.Summary, d.Detail)
		}
	}
	return resp
}

func mcpPlanReplacePaths(resp *tfprotov6.PlanResourceChangeResponse) []string {
	paths := make([]string, len(resp.RequiresReplace))
	for i, p := range resp.RequiresReplace {
		var names []string
		for _, step := range p.Steps() {
			if name, ok := step.(tftypes.AttributeName); ok {
				names = append(names, string(name))
			}
		}
		paths[i] = strings.Join(names, ".")
	}
	return paths
}

// mcpImportedState is what Read stores after import: computed view attributes
// only, with both config blocks and every registration-only input null.
func mcpImportedState(hosted bool) *MCPServerResourceModel {
	model := &MCPServerResourceModel{
		AppID: types.StringValue("app"), ConnectorID: types.StringValue("connector"),
		DisplayName: types.StringValue("server"),
	}
	if hosted {
		model.ServerType = types.StringValue("MCP_SERVER_TYPE_HOSTED")
		model.McpServerCatalogID = types.StringValue(mcpPlanCatalogID)
		model.SourceAppID = types.StringValue("")
	} else {
		model.ServerType = types.StringValue("MCP_SERVER_TYPE_EXTERNAL")
		model.EndpointURL = types.StringValue(mcpPlanURL)
	}
	return model
}

func mcpPlanConfig(hosted bool, identity string) *MCPServerResourceModel {
	model := &MCPServerResourceModel{
		AppID: types.StringValue("app"), DisplayName: types.StringValue("server"),
		AccessProfileIds:       []types.String{types.StringValue("profile")},
		AcknowledgedFindingIds: []types.String{types.StringValue("finding")},
		AppManagedStateBindingRef: &tfTypes.AppManagedStateBindingRef{
			AppID: types.StringValue("app"), ResourceID: types.StringValue("resource"), ResourceTypeID: types.StringValue("type"),
		},
		OauthDiagnosticID: types.StringValue("diagnostic"),
		TunnelPath:        types.StringValue("/mcp"),
	}
	token := &tfTypes.MCPServerAuthBearerToken{Token: types.StringValue("test-only-secret")}
	if hosted {
		model.ServerType = types.StringValue("MCP_SERVER_TYPE_HOSTED")
		model.HostedConfig = &tfTypes.MCPServerHostedConfig{
			McpServerCatalogID: types.StringValue(identity), SourceAppID: types.StringValue(""), BearerToken: token,
		}
	} else {
		model.ServerType = types.StringValue("MCP_SERVER_TYPE_EXTERNAL")
		model.ExternalConfig = &tfTypes.MCPServerExternalConfig{URL: types.StringValue(identity), BearerToken: token}
	}
	return model
}

func TestMCPServerImportPlanDoesNotReplace(t *testing.T) {
	for _, tc := range []struct {
		name     string
		hosted   bool
		identity string
	}{
		{"hosted", true, mcpPlanCatalogID},
		{"external", false, mcpPlanURL},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resp := mcpPlan(t, mcpImportedState(tc.hosted), mcpPlanConfig(tc.hosted, tc.identity))
			if len(resp.RequiresReplace) != 0 {
				t.Fatalf("import plan requires replacement of %v", mcpPlanReplacePaths(resp))
			}
		})
	}
}

func TestMCPServerImportPlanReplacesOnIdentityMismatch(t *testing.T) {
	for _, tc := range []struct {
		name     string
		hosted   bool
		identity string
		path     string
	}{
		{"hosted catalog", true, "2xCatalogEntry0000000000002", "hosted_config.mcp_server_catalog_id"},
		{"external url", false, "https://other.example.com/mcp", "external_config.url"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			paths := mcpPlanReplacePaths(mcpPlan(t, mcpImportedState(tc.hosted), mcpPlanConfig(tc.hosted, tc.identity)))
			if strings.Join(paths, ",") != tc.path {
				t.Fatalf("replacement paths = %v, want [%s]", paths, tc.path)
			}
		})
	}
}

func TestMCPServerManagedPlanReplacesOnRegistrationChange(t *testing.T) {
	for _, tc := range []struct {
		name   string
		hosted bool
		change func(*MCPServerResourceModel)
		path   string
	}{
		{"hosted catalog", true, func(m *MCPServerResourceModel) {
			m.HostedConfig.McpServerCatalogID = types.StringValue("2xCatalogEntry0000000000002")
		}, "hosted_config.mcp_server_catalog_id"},
		{"external url", false, func(m *MCPServerResourceModel) {
			m.ExternalConfig.URL = types.StringValue("https://other.example.com/mcp")
		}, "external_config.url"},
		{"tunnel path", false, func(m *MCPServerResourceModel) {
			m.TunnelPath = types.StringValue("/sse")
		}, "tunnel_path"},
		{"access profiles", true, func(m *MCPServerResourceModel) {
			m.AccessProfileIds = []types.String{types.StringValue("other")}
		}, "access_profile_ids"},
		{"binding ref", true, func(m *MCPServerResourceModel) {
			m.AppManagedStateBindingRef.ResourceID = types.StringValue("other")
		}, "app_managed_state_binding_ref,app_managed_state_binding_ref.resource_id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			identity := mcpPlanURL
			if tc.hosted {
				identity = mcpPlanCatalogID
			}
			prior := mcpPlanConfig(tc.hosted, identity)
			prior.ConnectorID = types.StringValue("connector")
			unchanged := mcpPlan(t, prior, mcpPlanConfig(tc.hosted, identity))
			if len(unchanged.RequiresReplace) != 0 {
				t.Fatalf("unchanged plan requires replacement of %v", mcpPlanReplacePaths(unchanged))
			}
			config := mcpPlanConfig(tc.hosted, identity)
			tc.change(config)
			paths := mcpPlanReplacePaths(mcpPlan(t, prior, config))
			if strings.Join(paths, ",") != tc.path {
				t.Fatalf("replacement paths = %v, want [%s]", paths, tc.path)
			}
		})
	}
}

func mcpValidate(t *testing.T, config *MCPServerResourceModel) []*tfprotov6.Diagnostic {
	t.Helper()
	resp, err := mcpPlanServer(t).ValidateResourceConfig(context.Background(), &tfprotov6.ValidateResourceConfigRequest{
		TypeName: "conductorone_mcp_server",
		Config:   mcpPlanDynamic(t, mcpPlanValue(t, config)),
	})
	if err != nil {
		t.Fatal(err)
	}
	return resp.Diagnostics
}

func TestMCPServerConfigValidation(t *testing.T) {
	none := &tfTypes.MCPServerAuthNone{}
	token := &tfTypes.MCPServerAuthBearerToken{Token: types.StringValue("test-only-secret")}
	for _, tc := range []struct {
		name   string
		config func(*MCPServerResourceModel)
		want   string
	}{
		{"external with one auth", func(m *MCPServerResourceModel) {
			m.ExternalConfig = &tfTypes.MCPServerExternalConfig{URL: types.StringValue(mcpPlanURL), BearerToken: token}
		}, ""},
		{"external with none auth", func(m *MCPServerResourceModel) {
			m.ExternalConfig = &tfTypes.MCPServerExternalConfig{URL: types.StringValue(mcpPlanURL), None: none}
		}, ""},
		{"hosted with deferred auth", func(m *MCPServerResourceModel) {
			m.ServerType = types.StringValue("MCP_SERVER_TYPE_HOSTED")
			m.HostedConfig = &tfTypes.MCPServerHostedConfig{McpServerCatalogID: types.StringValue(mcpPlanCatalogID)}
		}, ""},
		{"server type omitted", func(m *MCPServerResourceModel) {
			m.ServerType = types.StringNull()
			m.HostedConfig = &tfTypes.MCPServerHostedConfig{McpServerCatalogID: types.StringValue(mcpPlanCatalogID), BearerToken: token}
		}, ""},
		{"no config", func(m *MCPServerResourceModel) {}, "Invalid MCP server configuration"},
		{"both configs", func(m *MCPServerResourceModel) {
			m.ServerType = types.StringNull()
			m.ExternalConfig = &tfTypes.MCPServerExternalConfig{URL: types.StringValue(mcpPlanURL), BearerToken: token}
			m.HostedConfig = &tfTypes.MCPServerHostedConfig{McpServerCatalogID: types.StringValue(mcpPlanCatalogID)}
		}, "Invalid MCP server configuration"},
		{"external without auth", func(m *MCPServerResourceModel) {
			m.ExternalConfig = &tfTypes.MCPServerExternalConfig{URL: types.StringValue(mcpPlanURL)}
		}, "Missing MCP server authentication"},
		{"external with two auths", func(m *MCPServerResourceModel) {
			m.ExternalConfig = &tfTypes.MCPServerExternalConfig{URL: types.StringValue(mcpPlanURL), None: none, BearerToken: token}
		}, "Conflicting MCP server authentication"},
		{"hosted with two auths", func(m *MCPServerResourceModel) {
			m.ServerType = types.StringValue("MCP_SERVER_TYPE_HOSTED")
			m.HostedConfig = &tfTypes.MCPServerHostedConfig{
				McpServerCatalogID: types.StringValue(mcpPlanCatalogID), BearerToken: token,
				AwsSigv4: &tfTypes.MCPServerAuthAWSSigV4{AccessKeyID: types.StringValue("key")},
			}
		}, "Conflicting MCP server authentication"},
		{"server type mismatch", func(m *MCPServerResourceModel) {
			m.HostedConfig = &tfTypes.MCPServerHostedConfig{McpServerCatalogID: types.StringValue(mcpPlanCatalogID)}
		}, "Mismatched MCP server type"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			config := &MCPServerResourceModel{
				AppID: types.StringValue("app"), DisplayName: types.StringValue("server"),
				ServerType: types.StringValue("MCP_SERVER_TYPE_EXTERNAL"),
			}
			tc.config(config)
			var errors []string
			for _, d := range mcpValidate(t, config) {
				if d.Severity == tfprotov6.DiagnosticSeverityError {
					errors = append(errors, d.Summary)
				}
			}
			if tc.want == "" && len(errors) != 0 {
				t.Fatalf("unexpected validation errors %v", errors)
			}
			if tc.want != "" && (len(errors) != 1 || errors[0] != tc.want) {
				t.Fatalf("validation errors = %v, want [%s]", errors, tc.want)
			}
		})
	}
}
