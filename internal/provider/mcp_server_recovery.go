package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	sdkerrors "github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/errors"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// mcpServerCredentialsChanged excludes approval because the metadata endpoint
// owns that leaf. Sending an unchanged credential configuration can otherwise
// replace sealed values which C1 deliberately does not return after import.
func mcpServerCredentialsChanged(ctx context.Context, req resource.UpdateRequest) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	var stateExternal, planExternal *tfTypes.MCPServerExternalConfig
	var stateHosted, planHosted *tfTypes.MCPServerHostedConfig

	diags.Append(req.State.GetAttribute(ctx, path.Root("external_config"), &stateExternal)...)
	diags.Append(req.Plan.GetAttribute(ctx, path.Root("external_config"), &planExternal)...)
	diags.Append(req.State.GetAttribute(ctx, path.Root("hosted_config"), &stateHosted)...)
	diags.Append(req.Plan.GetAttribute(ctx, path.Root("hosted_config"), &planHosted)...)
	if diags.HasError() {
		return false, diags
	}

	if stateExternal == nil && planExternal != nil && stateHosted == nil && planHosted == nil {
		var endpointURL, authMethod types.String
		diags.Append(req.State.GetAttribute(ctx, path.Root("endpoint_url"), &endpointURL)...)
		diags.Append(req.State.GetAttribute(ctx, path.Root("auth_method"), &authMethod)...)
		if diags.HasError() {
			return false, diags
		}
		if !endpointURL.IsNull() && !endpointURL.IsUnknown() && !planExternal.URL.IsNull() && !planExternal.URL.IsUnknown() &&
			endpointURL.Equal(planExternal.URL) && mcpServerExternalAuthMatches(planExternal, authMethod) {
			return false, diags
		}
	}

	equal := reflect.DeepEqual(
		mcpServerExternalConfigWithoutApproval(stateExternal),
		mcpServerExternalConfigWithoutApproval(planExternal),
	) && reflect.DeepEqual(
		mcpServerHostedConfigWithoutApproval(stateHosted),
		mcpServerHostedConfigWithoutApproval(planHosted),
	)
	return !equal, diags
}

func mcpServerExternalAuthMatches(config *tfTypes.MCPServerExternalConfig, authMethod types.String) bool {
	if config == nil || authMethod.IsNull() || authMethod.IsUnknown() {
		return false
	}
	method := authMethod.ValueString()
	switch {
	case config.None != nil:
		return method == "MCP_SERVER_AUTH_METHOD_NONE"
	case config.BearerToken != nil:
		return method == "MCP_SERVER_AUTH_METHOD_BEARER_TOKEN"
	case config.BasicAuth != nil:
		return method == "MCP_SERVER_AUTH_METHOD_BASIC_AUTH"
	case config.CustomHeader != nil:
		return method == "MCP_SERVER_AUTH_METHOD_CUSTOM_HEADER"
	case config.Oauth2 != nil:
		return method == "MCP_SERVER_AUTH_METHOD_OAUTH2"
	default:
		return false
	}
}

func mcpServerDiagnosticError(err error) string {
	var apiError *sdkerrors.SDKError
	if errors.As(err, &apiError) {
		return fmt.Sprintf("%s: HTTP %d", apiError.Message, apiError.StatusCode)
	}
	return err.Error()
}

func mcpServerResponseStatus(response *http.Response) string {
	if response == nil {
		return "C1 returned no HTTP response"
	}
	return fmt.Sprintf("HTTP %d", response.StatusCode)
}

func mcpServerExternalConfigWithoutApproval(config *tfTypes.MCPServerExternalConfig) *tfTypes.MCPServerExternalConfig {
	if config == nil {
		return nil
	}

	withoutApproval := *config
	withoutApproval.RequireToolApproval = types.StringNull()

	return &withoutApproval
}

func mcpServerHostedConfigWithoutApproval(config *tfTypes.MCPServerHostedConfig) *tfTypes.MCPServerHostedConfig {
	if config == nil {
		return nil
	}

	withoutApproval := *config
	withoutApproval.RequireToolApproval = types.StringNull()

	return &withoutApproval
}

type mcpServerImportID struct {
	AppID       string `json:"app_id"`
	ConnectorID string `json:"connector_id"`
}

func mcpServerImportIdentity(data *MCPServerResourceModel) (string, bool) {
	if data == nil || data.AppID.IsNull() || data.AppID.IsUnknown() || data.ConnectorID.IsNull() || data.ConnectorID.IsUnknown() {
		return "", false
	}

	appID := data.AppID.ValueString()
	connectorID := data.ConnectorID.ValueString()
	if appID == "" || connectorID == "" {
		return "", false
	}

	encoded, err := json.Marshal(mcpServerImportID{
		AppID:       appID,
		ConnectorID: connectorID,
	})
	if err != nil {
		return "", false
	}

	return string(encoded), true
}

func addMCPServerCreatedStateRecoveryDiagnostic(resp *resource.CreateResponse, data *MCPServerResourceModel) {
	importID, ok := mcpServerImportIdentity(data)
	if !ok {
		return
	}

	resp.Diagnostics.AddError(
		"MCP server was created but Terraform state was not saved",
		fmt.Sprintf(
			"C1 identified the MCP server as %s, but the provider could not persist its Terraform state. Do not create another server. Recover the existing server with `terraform import <resource-address> '%s'`.",
			importID,
			importID,
		),
	)
}

func addMCPServerUnknownCreateOutcomeDiagnostic(resp *resource.CreateResponse, data *MCPServerResourceModel, cause error) {
	var apiError *sdkerrors.SDKError
	if errors.As(cause, &apiError) && apiError.StatusCode >= 400 && apiError.StatusCode < 500 {
		return
	}

	appID := "<app_id>"
	if data != nil && !data.AppID.IsNull() && !data.AppID.IsUnknown() && data.AppID.ValueString() != "" {
		appID = data.AppID.ValueString()
	}

	resp.Diagnostics.AddError(
		"MCP server registration outcome is unknown",
		fmt.Sprintf(
			"The Register request did not produce a usable connector ID. C1 may have created a server under app %q, but Register has no client-provided connector ID or idempotency key, so the provider cannot safely locate or retry it. Do not re-apply blindly. First locate the server in C1, then recover it with `terraform import <resource-address> '{\"app_id\":%q,\"connector_id\":\"<connector_id>\"}'`.",
			appID,
			appID,
		),
	)
}
