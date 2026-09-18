package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &MCPSourceResource{}
var _ resource.ResourceWithImportState = &MCPSourceResource{}

func NewMCPSourceResource() resource.Resource {
	return &MCPSourceResource{}
}

type MCPSourceResource struct {
	client *sdk.ConductoroneAPI
}

type MCPSourceResourceModel struct {
	AppID               types.String         `tfsdk:"app_id"`
	ConnectorID         types.String         `tfsdk:"connector_id"`
	DataSensitivity     types.String         `tfsdk:"data_sensitivity"`
	Description         types.String         `tfsdk:"description"`
	DisplayName         types.String         `tfsdk:"display_name"`
	ExternalConfig      jsontypes.Normalized `tfsdk:"external_config"`
	ExternalURL         types.String         `tfsdk:"external_url"`
	HostedCatalogID     types.String         `tfsdk:"hosted_catalog_id"`
	HostedConfig        jsontypes.Normalized `tfsdk:"hosted_config"`
	RequireToolApproval types.Bool           `tfsdk:"require_tool_approval"`
	SourceType          types.String         `tfsdk:"source_type"`
	ToolPrefix          types.String         `tfsdk:"tool_prefix"`
}

func (r *MCPSourceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mcp_source"
}

func (r *MCPSourceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MCP source resource.",
		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ID of the existing app that owns the MCP source.",
			},
			"connector_id": schema.StringAttribute{
				Computed:    true,
				Description: "Connector ID of the registered MCP source.",
			},
			"data_sensitivity": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MCP_SERVER_DATA_SENSITIVITY_PUBLIC",
						"MCP_SERVER_DATA_SENSITIVITY_INTERNAL",
						"MCP_SERVER_DATA_SENSITIVITY_CONFIDENTIAL",
						"MCP_SERVER_DATA_SENSITIVITY_RESTRICTED",
					),
				},
				Description: "Data sensitivity classification.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Admin-provided description.",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "Admin-provided display name.",
			},
			"external_config": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Optional:    true,
				Sensitive:   true,
				Description: "JSON configuration for an external MCP source, excluding url. It contains exactly one auth configuration.",
			},
			"external_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "HTTPS endpoint for an external MCP source. Changing it replaces the source.",
			},
			"hosted_catalog_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Catalog entry ID for a hosted MCP source. Changing it replaces the source.",
			},
			"hosted_config": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Optional:    true,
				Sensitive:   true,
				Description: "JSON configuration for a hosted MCP source, excluding mcpServerCatalogId. It contains exactly one auth configuration.",
			},
			"require_tool_approval": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether users must enable source tools. When omitted, the source inherits the tenant setting.",
			},
			"source_type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("MCP_SERVER_TYPE_HOSTED", "MCP_SERVER_TYPE_EXTERNAL"),
				},
				Description: "Whether C1 hosts the source or the source uses an external endpoint.",
			},
			"tool_prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Optional prefix for tool names exposed by C1.",
			},
		},
	}
}

func (r *MCPSourceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.ConductoroneAPI)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *sdk.ConductoroneAPI, got: %T.", req.ProviderData))
		return
	}

	r.client = client
}

func (r *MCPSourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MCPSourceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	registration, err := buildMCPSourceRegistration(&data)
	if err != nil {
		resp.Diagnostics.AddError("Invalid MCP source configuration", err.Error())
		return
	}

	res, err := r.client.MCPServer.Register(ctx, operations.C1APIAiGovernanceV1MCPServerServiceRegisterRequest{
		AppID:                           data.AppID.ValueString(),
		MCPServerServiceRegisterRequest: registration,
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to register MCP source", err.Error())
		return
	}
	if res == nil || res.StatusCode != 200 || res.MCPServerServiceRegisterResponse == nil || res.MCPServerServiceRegisterResponse.McpServer == nil {
		resp.Diagnostics.AddError("Unexpected MCP source registration response", "The API did not return a registered MCP source.")
		return
	}

	applyMCPSourceView(&data, res.MCPServerServiceRegisterResponse.McpServer)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MCPSourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MCPSourceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	missing, err := r.refresh(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read MCP source", err.Error())
		return
	}
	if missing {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MCPSourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MCPSourceResourceModel
	var state MCPSourceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	metadata, err := buildMCPSourceMetadataUpdate(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Invalid MCP source metadata", err.Error())
		return
	}
	metadataResult, err := r.client.MCPServer.Update(ctx, operations.C1APIAiGovernanceV1MCPServerServiceUpdateRequest{
		AppID:                         plan.AppID.ValueString(),
		ConnectorID:                   state.ConnectorID.ValueString(),
		MCPServerServiceUpdateRequest: metadata,
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to update MCP source metadata", err.Error())
		return
	}
	if metadataResult == nil || metadataResult.StatusCode != 200 || metadataResult.MCPServerServiceUpdateResponse == nil || metadataResult.MCPServerServiceUpdateResponse.McpServer == nil {
		resp.Diagnostics.AddError("Unexpected MCP source metadata response", "The API did not return an updated MCP source.")
		return
	}

	applyMCPSourceView(&plan, metadataResult.MCPServerServiceUpdateResponse.McpServer)

	if mcpSourceConfigChanged(plan.HostedConfig, state.HostedConfig) || mcpSourceConfigChanged(plan.ExternalConfig, state.ExternalConfig) {
		credentials, err := buildMCPSourceCredentialsUpdate(&plan)
		if err != nil {
			resp.Diagnostics.AddError("Invalid MCP source credentials configuration", err.Error())
			return
		}
		credentialsResult, err := r.client.MCPServer.UpdateCredentials(ctx, operations.C1APIAiGovernanceV1MCPServerServiceUpdateCredentialsRequest{
			AppID:                                    plan.AppID.ValueString(),
			ConnectorID:                              state.ConnectorID.ValueString(),
			MCPServerServiceUpdateCredentialsRequest: credentials,
		})
		if err != nil {
			resp.Diagnostics.AddError("Failed to update MCP source credentials", err.Error())
			return
		}
		if credentialsResult == nil || credentialsResult.StatusCode != 200 || credentialsResult.MCPServerServiceUpdateCredentialsResponse == nil || credentialsResult.MCPServerServiceUpdateCredentialsResponse.McpServer == nil {
			resp.Diagnostics.AddError("Unexpected MCP source credentials response", "The API did not return an updated MCP source.")
			return
		}

		applyMCPSourceView(&plan, credentialsResult.MCPServerServiceUpdateCredentialsResponse.McpServer)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MCPSourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MCPSourceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.MCPServer.Delete(ctx, operations.C1APIAiGovernanceV1MCPServerServiceDeleteRequest{
		AppID:       data.AppID.ValueString(),
		ConnectorID: data.ConnectorID.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete MCP source", err.Error())
		return
	}
	if res == nil || (res.StatusCode != 200 && res.StatusCode != 404) {
		resp.Diagnostics.AddError("Unexpected MCP source deletion response", "The API did not confirm deletion of the MCP source.")
	}
}

func (r *MCPSourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var identity struct {
		AppID       string `json:"app_id"`
		ConnectorID string `json:"connector_id"`
	}
	dec := json.NewDecoder(bytes.NewBufferString(req.ID))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&identity); err != nil {
		resp.Diagnostics.AddError("Invalid MCP source import ID", `The import ID must be a JSON object: {"app_id":"...","connector_id":"..."}.`)
		return
	}
	if identity.AppID == "" || identity.ConnectorID == "" {
		resp.Diagnostics.AddError("Invalid MCP source import ID", `The import ID must include non-empty "app_id" and "connector_id" fields.`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("app_id"), identity.AppID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("connector_id"), identity.ConnectorID)...)
}

func (r *MCPSourceResource) refresh(ctx context.Context, data *MCPSourceResourceModel) (bool, error) {
	res, err := r.client.MCPServer.Get(ctx, operations.C1APIAiGovernanceV1MCPServerServiceGetRequest{
		AppID:       data.AppID.ValueString(),
		ConnectorID: data.ConnectorID.ValueString(),
	})
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, fmt.Errorf("API returned an empty response")
	}
	if res.StatusCode == 404 {
		return true, nil
	}
	if res.StatusCode != 200 {
		return false, fmt.Errorf("API returned status %d", res.StatusCode)
	}
	if res.MCPServerServiceGetResponse == nil || res.MCPServerServiceGetResponse.McpServer == nil {
		return false, fmt.Errorf("API returned a response without an MCP source")
	}

	applyMCPSourceView(data, res.MCPServerServiceGetResponse.McpServer)
	return false, nil
}

func buildMCPSourceRegistration(data *MCPSourceResourceModel) (*shared.MCPServerServiceRegisterRequest, error) {
	hosted, external, err := mcpSourceConfigs(data)
	if err != nil {
		return nil, err
	}

	serverType := shared.MCPServerServiceRegisterRequestServerType(data.SourceType.ValueString())
	request := &shared.MCPServerServiceRegisterRequest{
		DisplayName:    data.DisplayName.ValueStringPointer(),
		ExternalConfig: external,
		HostedConfig:   hosted,
		ServerType:     &serverType,
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		request.Description = data.Description.ValueStringPointer()
	}
	if !data.ToolPrefix.IsNull() && !data.ToolPrefix.IsUnknown() {
		request.ToolPrefix = data.ToolPrefix.ValueStringPointer()
	}
	if !data.DataSensitivity.IsNull() && !data.DataSensitivity.IsUnknown() {
		value := shared.MCPServerServiceRegisterRequestDataSensitivity(data.DataSensitivity.ValueString())
		request.DataSensitivity = &value
	}
	return request, nil
}

func buildMCPSourceMetadataUpdate(data *MCPSourceResourceModel) (*shared.MCPServerServiceUpdateRequest, error) {
	if data.DisplayName.IsNull() || data.DisplayName.IsUnknown() {
		return nil, fmt.Errorf("display_name is required")
	}

	mcpServer := &shared.MCPServerViewInput{DisplayName: data.DisplayName.ValueStringPointer()}
	mask := []string{"displayName"}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		mcpServer.Description = data.Description.ValueStringPointer()
		mask = append(mask, "description")
	}
	if !data.DataSensitivity.IsNull() && !data.DataSensitivity.IsUnknown() && data.DataSensitivity.ValueString() != "" {
		value := shared.DataSensitivity(data.DataSensitivity.ValueString())
		mcpServer.DataSensitivity = &value
		mask = append(mask, "dataSensitivity")
	}
	if !data.ToolPrefix.IsNull() && !data.ToolPrefix.IsUnknown() {
		mcpServer.ToolPrefix = data.ToolPrefix.ValueStringPointer()
		mask = append(mask, "toolPrefix")
	}
	if !data.RequireToolApproval.IsNull() && !data.RequireToolApproval.IsUnknown() {
		value := shared.RequireToolApprovalOptionalBoolFalse
		if data.RequireToolApproval.ValueBool() {
			value = shared.RequireToolApprovalOptionalBoolTrue
		}
		mcpServer.RequireToolApproval = &value
		mask = append(mask, "requireToolApproval")
	}
	updateMask := strings.Join(mask, ",")
	return &shared.MCPServerServiceUpdateRequest{McpServer: mcpServer, UpdateMask: &updateMask}, nil
}

func buildMCPSourceCredentialsUpdate(data *MCPSourceResourceModel) (*shared.MCPServerServiceUpdateCredentialsRequest, error) {
	hosted, external, err := mcpSourceConfigs(data)
	if err != nil {
		return nil, err
	}
	return &shared.MCPServerServiceUpdateCredentialsRequest{HostedConfig: hosted, ExternalConfig: external}, nil
}

func mcpSourceConfigs(data *MCPSourceResourceModel) (*shared.MCPServerHostedConfig, *shared.MCPServerExternalConfig, error) {
	switch data.SourceType.ValueString() {
	case "MCP_SERVER_TYPE_HOSTED":
		if data.HostedCatalogID.IsNull() || data.HostedCatalogID.IsUnknown() || data.HostedCatalogID.ValueString() == "" {
			return nil, nil, fmt.Errorf("hosted_catalog_id is required for a hosted MCP source")
		}
		if data.HostedConfig.IsNull() || data.HostedConfig.IsUnknown() {
			return nil, nil, fmt.Errorf("hosted_config is required for a hosted MCP source")
		}
		if (!data.ExternalConfig.IsNull() && !data.ExternalConfig.IsUnknown()) || (!data.ExternalURL.IsNull() && !data.ExternalURL.IsUnknown()) {
			return nil, nil, fmt.Errorf("external_config and external_url are only valid for an external MCP source")
		}
		var config shared.MCPServerHostedConfig
		if err := decodeMCPServerConfig(data.HostedConfig.ValueString(), &config); err != nil {
			return nil, nil, fmt.Errorf("invalid hosted_config: %w", err)
		}
		if config.McpServerCatalogID != nil {
			return nil, nil, fmt.Errorf("hosted_config must not set mcpServerCatalogId; use hosted_catalog_id")
		}
		if config.RequireToolApproval != nil {
			return nil, nil, fmt.Errorf("hosted_config must not set requireToolApproval; use require_tool_approval")
		}
		config.McpServerCatalogID = data.HostedCatalogID.ValueStringPointer()
		if !data.RequireToolApproval.IsNull() && !data.RequireToolApproval.IsUnknown() {
			value := shared.MCPServerHostedConfigRequireToolApprovalOptionalBoolFalse
			if data.RequireToolApproval.ValueBool() {
				value = shared.MCPServerHostedConfigRequireToolApprovalOptionalBoolTrue
			}
			config.RequireToolApproval = &value
		}
		return &config, nil, nil
	case "MCP_SERVER_TYPE_EXTERNAL":
		if data.ExternalURL.IsNull() || data.ExternalURL.IsUnknown() || data.ExternalURL.ValueString() == "" {
			return nil, nil, fmt.Errorf("external_url is required for an external MCP source")
		}
		if data.ExternalConfig.IsNull() || data.ExternalConfig.IsUnknown() {
			return nil, nil, fmt.Errorf("external_config is required for an external MCP source")
		}
		if (!data.HostedConfig.IsNull() && !data.HostedConfig.IsUnknown()) || (!data.HostedCatalogID.IsNull() && !data.HostedCatalogID.IsUnknown()) {
			return nil, nil, fmt.Errorf("hosted_config and hosted_catalog_id are only valid for a hosted MCP source")
		}
		var config shared.MCPServerExternalConfig
		if err := decodeMCPServerConfig(data.ExternalConfig.ValueString(), &config); err != nil {
			return nil, nil, fmt.Errorf("invalid external_config: %w", err)
		}
		if config.URL != nil {
			return nil, nil, fmt.Errorf("external_config must not set url; use external_url")
		}
		if config.RequireToolApproval != nil {
			return nil, nil, fmt.Errorf("external_config must not set requireToolApproval; use require_tool_approval")
		}
		config.URL = data.ExternalURL.ValueStringPointer()
		if !data.RequireToolApproval.IsNull() && !data.RequireToolApproval.IsUnknown() {
			value := shared.MCPServerExternalConfigRequireToolApprovalOptionalBoolFalse
			if data.RequireToolApproval.ValueBool() {
				value = shared.MCPServerExternalConfigRequireToolApprovalOptionalBoolTrue
			}
			config.RequireToolApproval = &value
		}
		return nil, &config, nil
	default:
		return nil, nil, fmt.Errorf("source_type must be MCP_SERVER_TYPE_HOSTED or MCP_SERVER_TYPE_EXTERNAL")
	}
}

func decodeMCPServerConfig(input string, destination any) error {
	if !strings.HasPrefix(strings.TrimSpace(input), "{") {
		return fmt.Errorf("must be a JSON object")
	}

	decoder := json.NewDecoder(strings.NewReader(input))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return fmt.Errorf("must contain a single JSON object")
		}
		return err
	}
	return nil
}

func mcpSourceConfigChanged(plan, state jsontypes.Normalized) bool {
	if plan.IsNull() != state.IsNull() {
		return true
	}
	if plan.IsNull() {
		return false
	}
	return plan.ValueString() != state.ValueString()
}

func applyMCPSourceView(data *MCPSourceResourceModel, view *shared.MCPServerView) {
	if view.AppID != nil {
		data.AppID = types.StringPointerValue(view.AppID)
	}
	if view.ConnectorID != nil {
		data.ConnectorID = types.StringPointerValue(view.ConnectorID)
	}
	if view.DataSensitivity != nil {
		data.DataSensitivity = types.StringValue(string(*view.DataSensitivity))
	}
	if view.Description != nil {
		data.Description = types.StringPointerValue(view.Description)
	}
	if view.DisplayName != nil {
		data.DisplayName = types.StringPointerValue(view.DisplayName)
	}
	if view.EndpointURL != nil {
		data.ExternalURL = types.StringPointerValue(view.EndpointURL)
	}
	if view.McpServerCatalogID != nil {
		data.HostedCatalogID = types.StringPointerValue(view.McpServerCatalogID)
	}
	if view.ServerType != nil {
		data.SourceType = types.StringValue(string(*view.ServerType))
	}
	if view.ToolPrefix != nil {
		data.ToolPrefix = types.StringPointerValue(view.ToolPrefix)
	}
	if view.RequireToolApproval != nil {
		switch *view.RequireToolApproval {
		case shared.RequireToolApprovalOptionalBoolTrue:
			data.RequireToolApproval = types.BoolValue(true)
		case shared.RequireToolApprovalOptionalBoolFalse:
			data.RequireToolApproval = types.BoolValue(false)
		case shared.RequireToolApprovalOptionalBoolUnspecified:
			data.RequireToolApproval = types.BoolNull()
		}
	}
}
