package provider

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestBuildMCPSourceRegistrationHosted(t *testing.T) {
	data := MCPSourceResourceModel{
		AppID:               types.StringValue("app-id"),
		DataSensitivity:     types.StringUnknown(),
		Description:         types.StringUnknown(),
		DisplayName:         types.StringValue("GitHub MCP"),
		HostedCatalogID:     types.StringValue("catalog-id"),
		HostedConfig:        jsontypes.NewNormalizedValue(`{"none":{}}`),
		ExternalConfig:      jsontypes.NewNormalizedNull(),
		ExternalURL:         types.StringUnknown(),
		RequireToolApproval: types.BoolValue(true),
		SourceType:          types.StringValue("MCP_SERVER_TYPE_HOSTED"),
		ToolPrefix:          types.StringUnknown(),
	}

	request, err := buildMCPSourceRegistration(&data)
	if err != nil {
		t.Fatalf("buildMCPSourceRegistration() error = %v", err)
	}
	if request.ServerType == nil || string(*request.ServerType) != "MCP_SERVER_TYPE_HOSTED" {
		t.Fatalf("source type = %v, want hosted", request.ServerType)
	}
	if request.HostedConfig == nil || request.HostedConfig.McpServerCatalogID == nil || *request.HostedConfig.McpServerCatalogID != "catalog-id" {
		t.Fatalf("hosted catalog ID = %#v, want catalog-id", request.HostedConfig)
	}
	if request.HostedConfig.None == nil {
		t.Fatal("hosted auth config was not decoded")
	}
	if request.HostedConfig.RequireToolApproval == nil || string(*request.HostedConfig.RequireToolApproval) != "OPTIONAL_BOOL_TRUE" {
		t.Fatalf("require tool approval = %#v, want OPTIONAL_BOOL_TRUE", request.HostedConfig.RequireToolApproval)
	}
	if request.ExternalConfig != nil {
		t.Fatalf("external config = %#v, want nil", request.ExternalConfig)
	}
	if request.DataSensitivity != nil || request.Description != nil || request.ToolPrefix != nil {
		t.Fatalf("unconfigured computed attributes = %#v, want omitted", request)
	}
}

func TestBuildMCPSourceRegistrationExternal(t *testing.T) {
	data := MCPSourceResourceModel{
		AppID:               types.StringValue("app-id"),
		DisplayName:         types.StringValue("Partner MCP"),
		HostedCatalogID:     types.StringUnknown(),
		HostedConfig:        jsontypes.NewNormalizedNull(),
		ExternalConfig:      jsontypes.NewNormalizedValue(`{"none":{},"transportType":"MCP_SERVER_TRANSPORT_TYPE_SSE"}`),
		ExternalURL:         types.StringValue("https://mcp.example.com/sse"),
		RequireToolApproval: types.BoolValue(true),
		SourceType:          types.StringValue("MCP_SERVER_TYPE_EXTERNAL"),
	}

	request, err := buildMCPSourceRegistration(&data)
	if err != nil {
		t.Fatalf("buildMCPSourceRegistration() error = %v", err)
	}
	if request.ExternalConfig == nil || request.ExternalConfig.URL == nil || *request.ExternalConfig.URL != "https://mcp.example.com/sse" {
		t.Fatalf("external URL = %#v, want configured endpoint", request.ExternalConfig)
	}
	if request.ExternalConfig.None == nil {
		t.Fatal("external auth config was not decoded")
	}
	if request.HostedConfig != nil {
		t.Fatalf("hosted config = %#v, want nil", request.HostedConfig)
	}
	if request.ExternalConfig.RequireToolApproval == nil || string(*request.ExternalConfig.RequireToolApproval) != "OPTIONAL_BOOL_TRUE" {
		t.Fatalf("require tool approval = %#v, want OPTIONAL_BOOL_TRUE", request.ExternalConfig.RequireToolApproval)
	}
}

func TestBuildMCPSourceRegistrationRejectsConfigurationOwnedByRootAttribute(t *testing.T) {
	data := MCPSourceResourceModel{
		AppID:           types.StringValue("app-id"),
		DisplayName:     types.StringValue("Partner MCP"),
		HostedCatalogID: types.StringNull(),
		HostedConfig:    jsontypes.NewNormalizedNull(),
		ExternalConfig:  jsontypes.NewNormalizedValue(`{"none":{},"url":"https://mcp.example.com/sse"}`),
		ExternalURL:     types.StringValue("https://mcp.example.com/sse"),
		SourceType:      types.StringValue("MCP_SERVER_TYPE_EXTERNAL"),
	}

	_, err := buildMCPSourceRegistration(&data)
	if err == nil || !strings.Contains(err.Error(), "must not set url") {
		t.Fatalf("buildMCPSourceRegistration() error = %v, want external URL ownership error", err)
	}
}

func TestBuildMCPSourceRegistrationRejectsToolEnablementInConfiguration(t *testing.T) {
	data := MCPSourceResourceModel{
		AppID:           types.StringValue("app-id"),
		DisplayName:     types.StringValue("Partner MCP"),
		HostedCatalogID: types.StringNull(),
		HostedConfig:    jsontypes.NewNormalizedNull(),
		ExternalConfig:  jsontypes.NewNormalizedValue(`{"none":{},"requireToolApproval":"OPTIONAL_BOOL_TRUE"}`),
		ExternalURL:     types.StringValue("https://mcp.example.com/sse"),
		SourceType:      types.StringValue("MCP_SERVER_TYPE_EXTERNAL"),
	}

	_, err := buildMCPSourceRegistration(&data)
	if err == nil || !strings.Contains(err.Error(), "must not set requireToolApproval") {
		t.Fatalf("buildMCPSourceRegistration() error = %v, want tool enablement ownership error", err)
	}
}

func TestBuildMCPSourceRegistrationOmitsInheritedToolEnablement(t *testing.T) {
	data := MCPSourceResourceModel{
		AppID:               types.StringValue("app-id"),
		DisplayName:         types.StringValue("Partner MCP"),
		HostedCatalogID:     types.StringUnknown(),
		HostedConfig:        jsontypes.NewNormalizedNull(),
		ExternalConfig:      jsontypes.NewNormalizedValue(`{"none":{}}`),
		ExternalURL:         types.StringValue("https://mcp.example.com/sse"),
		RequireToolApproval: types.BoolNull(),
		SourceType:          types.StringValue("MCP_SERVER_TYPE_EXTERNAL"),
	}

	request, err := buildMCPSourceRegistration(&data)
	if err != nil {
		t.Fatalf("buildMCPSourceRegistration() error = %v", err)
	}
	if request.ExternalConfig.RequireToolApproval != nil {
		t.Fatalf("require tool approval = %#v, want inherited setting", request.ExternalConfig.RequireToolApproval)
	}
}

func TestBuildMCPSourceMetadataUpdateUsesAcceptedFieldMaskPaths(t *testing.T) {
	data := MCPSourceResourceModel{
		DisplayName:         types.StringValue("Updated MCP"),
		Description:         types.StringValue("Updated description"),
		DataSensitivity:     types.StringValue("MCP_SERVER_DATA_SENSITIVITY_CONFIDENTIAL"),
		RequireToolApproval: types.BoolValue(false),
		ToolPrefix:          types.StringValue("updated"),
	}

	request, err := buildMCPSourceMetadataUpdate(&data)
	if err != nil {
		t.Fatalf("buildMCPSourceMetadataUpdate() error = %v", err)
	}
	if request.UpdateMask == nil || *request.UpdateMask != "displayName,description,dataSensitivity,toolPrefix,requireToolApproval" {
		t.Fatalf("update mask = %v, want all MCP metadata fields", request.UpdateMask)
	}
	if request.McpServer == nil || request.McpServer.DisplayName == nil || *request.McpServer.DisplayName != "Updated MCP" {
		t.Fatalf("metadata request = %#v, want display name", request.McpServer)
	}
	if request.McpServer.RequireToolApproval == nil || string(*request.McpServer.RequireToolApproval) != "OPTIONAL_BOOL_FALSE" {
		t.Fatalf("require tool approval = %#v, want OPTIONAL_BOOL_FALSE", request.McpServer.RequireToolApproval)
	}
}

func TestBuildMCPSourceMetadataUpdateOmitsEmptyComputedSensitivity(t *testing.T) {
	data := MCPSourceResourceModel{
		DataSensitivity: types.StringValue(""),
		DisplayName:     types.StringValue("Updated MCP"),
	}

	request, err := buildMCPSourceMetadataUpdate(&data)
	if err != nil {
		t.Fatalf("buildMCPSourceMetadataUpdate() error = %v", err)
	}
	if request.UpdateMask == nil || *request.UpdateMask != "displayName" {
		t.Fatalf("update mask = %v, want displayName", request.UpdateMask)
	}
	if request.McpServer.DataSensitivity != nil {
		t.Fatalf("data sensitivity = %v, want nil", request.McpServer.DataSensitivity)
	}
}

func TestAccMCPSourceResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "conductorone_app" "mcp" {
  display_name = "terraform-mcp-source-test"
  description  = "Terraform MCP source acceptance test"
}

resource "conductorone_mcp_source" "test" {
  app_id                = conductorone_app.mcp.id
  display_name          = "Terraform MCP source"
  description           = "Terraform-managed external MCP source"
  external_config       = jsonencode({ none = {}, transportType = "MCP_SERVER_TRANSPORT_TYPE_SSE" })
  source_type           = "MCP_SERVER_TYPE_EXTERNAL"
  external_url          = "https://example.com/mcp"
  tool_prefix           = "terraform"
  require_tool_approval = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "display_name", "Terraform MCP source"),
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "description", "Terraform-managed external MCP source"),
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "external_url", "https://example.com/mcp"),
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "tool_prefix", "terraform"),
					resource.TestCheckResourceAttrSet("conductorone_mcp_source.test", "connector_id"),
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "require_tool_approval", "true"),
				),
			},
			{
				Config: providerConfig + `
resource "conductorone_app" "mcp" {
  display_name = "terraform-mcp-source-test"
  description  = "Terraform MCP source acceptance test"
}

resource "conductorone_mcp_source" "test" {
  app_id                = conductorone_app.mcp.id
  display_name          = "Terraform MCP source updated"
  description           = "Updated Terraform-managed external MCP source"
  external_config       = jsonencode({ none = {}, transportType = "MCP_SERVER_TRANSPORT_TYPE_SSE" })
  source_type           = "MCP_SERVER_TYPE_EXTERNAL"
  external_url          = "https://example.com/mcp"
  tool_prefix           = "terraform-updated"
  require_tool_approval = false
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "display_name", "Terraform MCP source updated"),
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "description", "Updated Terraform-managed external MCP source"),
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "tool_prefix", "terraform-updated"),
					resource.TestCheckResourceAttr("conductorone_mcp_source.test", "require_tool_approval", "false"),
				),
			},
		},
	})
}
