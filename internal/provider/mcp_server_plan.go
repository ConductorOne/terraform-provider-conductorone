package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = &MCPServerResource{}

// Auth arms of each config's auth_config oneof. Add new proto arms here.
var (
	mcpServerHostedAuthArms   = []string{"none", "bearer_token", "oauth2", "custom_header", "google_service_account", "aws_sigv4", "basic_auth"}
	mcpServerExternalAuthArms = []string{"none", "bearer_token", "oauth2", "custom_header", "basic_auth"}
)

// mcpServerConfigBlocks maps each config block to the server_type it registers.
var mcpServerConfigBlocks = []struct {
	name       string
	serverType string
	authArms   []string
	// Register rejects external servers without auth; hosted servers may defer it.
	authRequired bool
}{
	{name: "hosted_config", serverType: "MCP_SERVER_TYPE_HOSTED", authArms: mcpServerHostedAuthArms},
	{name: "external_config", serverType: "MCP_SERVER_TYPE_EXTERNAL", authArms: mcpServerExternalAuthArms, authRequired: true},
}

func (r *MCPServerResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{mcpServerConfigValidator{}}
}

type mcpServerConfigValidator struct{}

func (v mcpServerConfigValidator) Description(ctx context.Context) string {
	return "Exactly one of hosted_config or external_config is set, it matches server_type, and it sets at most one authentication method (exactly one for external_config)."
}

func (v mcpServerConfigValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v mcpServerConfigValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var serverType types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("server_type"), &serverType)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var configured []string
	unknown := false
	for _, block := range mcpServerConfigBlocks {
		var config types.Object
		resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root(block.name), &config)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.IsUnknown() {
			unknown = true
			continue
		}
		if config.IsNull() {
			continue
		}
		configured = append(configured, block.name)

		if !serverType.IsNull() && !serverType.IsUnknown() && serverType.ValueString() != block.serverType {
			resp.Diagnostics.AddAttributeError(path.Root("server_type"), "Mismatched MCP server type",
				fmt.Sprintf("%s requires server_type %q, got %q.", block.name, block.serverType, serverType.ValueString()))
		}
		validateMCPServerAuthArms(path.Root(block.name), config.Attributes(), block.authArms, block.authRequired, &resp.Diagnostics)
	}

	if unknown || len(configured) == 1 {
		return
	}
	names := make([]string, len(mcpServerConfigBlocks))
	for i, block := range mcpServerConfigBlocks {
		names[i] = block.name
	}
	resp.Diagnostics.AddAttributeError(path.Root(mcpServerConfigBlocks[0].name), "Invalid MCP server configuration",
		fmt.Sprintf("Exactly one of %s must be set; found %d.", strings.Join(names, ", "), len(configured)))
}

func validateMCPServerAuthArms(block path.Path, attributes map[string]attr.Value, arms []string, required bool, diags *diag.Diagnostics) {
	var set []string
	for _, arm := range arms {
		value, ok := attributes[arm]
		if !ok || value.IsNull() {
			continue
		}
		if value.IsUnknown() {
			return
		}
		set = append(set, arm)
	}
	if len(set) > 1 {
		diags.AddAttributeError(block, "Conflicting MCP server authentication",
			fmt.Sprintf("At most one authentication method may be set; found %s.", strings.Join(set, ", ")))
	}
	if required && len(set) == 0 {
		diags.AddAttributeError(block, "Missing MCP server authentication",
			fmt.Sprintf("Exactly one authentication method must be set: %s.", strings.Join(arms, ", ")))
	}
}

// mcpServerStateImported reports state produced by import: Read cannot
// recover either config block, while every Create stores the configured one.
func mcpServerStateImported(ctx context.Context, state tfsdk.State) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	for _, block := range mcpServerConfigBlocks {
		var config types.Object
		diags.Append(state.GetAttribute(ctx, path.Root(block.name), &config)...)
		if diags.HasError() || !config.IsNull() {
			return false, diags
		}
	}
	return true, diags
}

const mcpServerRegistrationInputDescription = "Requires replacement if changed. An import does not trigger replacement because registration inputs are not returned by the API."

// Registration-only inputs are never read back, so imported state cannot show
// whether the configured value differs from the registered one.
func mcpServerRegistrationInputReplace(ctx context.Context, state tfsdk.State, config attr.Value) (bool, diag.Diagnostics) {
	if config.IsNull() {
		return false, nil
	}
	imported, diags := mcpServerStateImported(ctx, state)
	return !imported, diags
}

func mcpServerRegistrationInputString() planmodifier.String {
	return stringplanmodifier.RequiresReplaceIf(func(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
		resp.RequiresReplace, resp.Diagnostics = mcpServerRegistrationInputReplace(ctx, req.State, req.ConfigValue)
	}, mcpServerRegistrationInputDescription, mcpServerRegistrationInputDescription)
}

func mcpServerRegistrationInputList() planmodifier.List {
	return listplanmodifier.RequiresReplaceIf(func(ctx context.Context, req planmodifier.ListRequest, resp *listplanmodifier.RequiresReplaceIfFuncResponse) {
		resp.RequiresReplace, resp.Diagnostics = mcpServerRegistrationInputReplace(ctx, req.State, req.ConfigValue)
	}, mcpServerRegistrationInputDescription, mcpServerRegistrationInputDescription)
}

func mcpServerRegistrationInputObject() planmodifier.Object {
	return objectplanmodifier.RequiresReplaceIf(func(ctx context.Context, req planmodifier.ObjectRequest, resp *objectplanmodifier.RequiresReplaceIfFuncResponse) {
		resp.RequiresReplace, resp.Diagnostics = mcpServerRegistrationInputReplace(ctx, req.State, req.ConfigValue)
	}, mcpServerRegistrationInputDescription, mcpServerRegistrationInputDescription)
}

// mcpServerIdentityString replaces the server when a configured registration
// identity changes. Read leaves the containing config block null, so imported
// state is compared with the computed top-level attribute that mirrors it.
func mcpServerIdentityString(block, mirror string) planmodifier.String {
	description := fmt.Sprintf("Requires replacement if changed. When %s is absent from state, such as after import, the configured value is compared with %s.", block, mirror)
	return stringplanmodifier.RequiresReplaceIf(func(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
		if req.ConfigValue.IsNull() {
			return
		}
		var config types.Object
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root(block), &config)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !config.IsNull() {
			resp.RequiresReplace = true
			return
		}
		var registered types.String
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root(mirror), &registered)...)
		if resp.Diagnostics.HasError() || registered.IsNull() {
			return
		}
		resp.RequiresReplace = !req.PlanValue.Equal(registered)
	}, description, description)
}
