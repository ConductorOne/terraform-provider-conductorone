package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	sdkerrors "github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/errors"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type mcpRecoveryAPI struct {
	t *testing.T

	mu                     sync.Mutex
	bearerToken            string
	connectorID            string
	description            string
	deleted                bool
	forbidCredentials      bool
	failedCredentialWrites int
	createCalls            int
	created                bool
	credentialCalls        int
	metadataCalls          int
	responseLostOnCreate   bool
}

func (api *mcpRecoveryAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.mu.Lock()
	defer api.mu.Unlock()

	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/api/v1/apps/app/mcp_servers":
		api.createCalls++
		api.created = true
		if api.responseLostOnCreate {
			return
		}
		api.writeView(w)
	case r.Method == http.MethodGet && r.URL.Path == "/api/v1/apps/app/mcp_servers/connector":
		if api.deleted {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		api.writeView(w)
	case r.Method == http.MethodPost && r.URL.Path == "/api/v1/apps/app/mcp_servers/connector":
		api.metadataCalls++

		var update shared.MCPServerServiceUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			api.t.Errorf("decode metadata update: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if update.McpServer != nil && update.McpServer.Description != nil {
			api.description = *update.McpServer.Description
		}
		api.writeView(w)
	case r.Method == http.MethodPost && r.URL.Path == "/api/v1/apps/app/mcp_servers/connector/credentials":
		api.credentialCalls++
		if api.forbidCredentials || api.credentialCalls <= api.failedCredentialWrites {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var update shared.MCPServerServiceUpdateCredentialsRequest
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			api.t.Errorf("decode credential update: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if update.ExternalConfig != nil && update.ExternalConfig.BearerToken != nil && update.ExternalConfig.BearerToken.Token != nil {
			api.bearerToken = *update.ExternalConfig.BearerToken.Token
		}
		api.writeView(w)
	default:
		http.NotFound(w, r)
	}
}

func (api *mcpRecoveryAPI) writeView(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]any{"mcpServer": map[string]any{
		"appId": "app", "connectorId": api.connectorID, "description": api.description,
	}}); err != nil {
		api.t.Errorf("encode MCP server response: %v", err)
	}
}

func mcpRecoveryResource(t *testing.T, server *httptest.Server) (*MCPServerResource, resource.SchemaResponse) {
	t.Helper()

	mcpResource := &MCPServerResource{client: sdk.New(sdk.WithServerURL(server.URL))}
	var schemaResp resource.SchemaResponse
	mcpResource.Schema(context.Background(), resource.SchemaRequest{}, &schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatalf("building MCP server resource schema: %v", schemaResp.Diagnostics)
	}

	return mcpResource, schemaResp
}

func mcpRecoveryState(t *testing.T, ctx context.Context, schemaResp resource.SchemaResponse, model *MCPServerResourceModel) tfsdk.State {
	t.Helper()

	state := tfsdk.State{Schema: schemaResp.Schema}
	if diags := state.Set(ctx, model); diags.HasError() {
		t.Fatalf("setting MCP server state: %v", diags)
	}

	return state
}

func mcpRecoveryModel(description, bearerToken string) MCPServerResourceModel {
	model := MCPServerResourceModel{
		AppID:       types.StringValue("app"),
		ConnectorID: types.StringValue("connector"),
		Description: types.StringValue(description),
	}
	if bearerToken != "" {
		model.ExternalConfig = &tfTypes.MCPServerExternalConfig{
			BearerToken: &tfTypes.MCPServerAuthBearerToken{Token: types.StringValue(bearerToken)},
		}
	}

	return model
}

func TestMCPServerImportedEndpointAdoptsOriginalConfig(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.NotFoundHandler())
	t.Cleanup(server.Close)
	_, schemaResp := mcpRecoveryResource(t, server)
	const endpoint = "https://fixture.example/mcp"

	imported := mcpRecoveryModel("imported", "")
	imported.EndpointURL = types.StringValue(endpoint)
	imported.AuthMethod = types.StringValue("MCP_SERVER_AUTH_METHOD_BEARER_TOKEN")
	state := mcpRecoveryState(t, ctx, schemaResp, &imported)

	desired := mcpRecoveryModel("imported", "asserted-secret")
	desired.ExternalConfig.URL = types.StringValue(endpoint)
	plan := tfsdk.Plan{Schema: schemaResp.Schema}
	if diags := plan.Set(ctx, &desired); diags.HasError() {
		t.Fatal(diags)
	}

	modifier := mcpServerURLRequiresReplace()
	response := &planmodifier.StringResponse{}
	modifier.PlanModifyString(ctx, planmodifier.StringRequest{
		State: state, Plan: plan, ConfigValue: types.StringValue(endpoint),
		StateValue: types.StringNull(), PlanValue: types.StringValue(endpoint),
	}, response)
	if response.Diagnostics.HasError() || response.RequiresReplace {
		t.Fatalf("matching imported URL requested replacement: %+v", response)
	}

	changed, diags := mcpServerCredentialsChanged(ctx, resource.UpdateRequest{State: state, Plan: plan})
	if diags.HasError() || changed {
		t.Fatalf("import adoption would replay unreadable credentials: changed=%t, diagnostics=%v", changed, diags)
	}

	response = &planmodifier.StringResponse{}
	modifier.PlanModifyString(ctx, planmodifier.StringRequest{
		State: state, Plan: plan, ConfigValue: types.StringValue("https://fixture.example/other"),
		StateValue: types.StringNull(), PlanValue: types.StringValue("https://fixture.example/other"),
	}, response)
	if response.Diagnostics.HasError() || !response.RequiresReplace {
		t.Fatalf("changed imported URL failed to request replacement: %+v", response)
	}
}

func TestMCPServerOwningAppChangeRequiresReplacement(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.NotFoundHandler())
	t.Cleanup(server.Close)
	_, schemaResp := mcpRecoveryResource(t, server)
	attribute, ok := schemaResp.Schema.Attributes["app_id"].(schema.StringAttribute)
	if !ok || len(attribute.PlanModifiers) != 1 {
		t.Fatal("app_id requires one replacement plan modifier")
	}

	before := mcpRecoveryModel("existing", "")
	state := mcpRecoveryState(t, ctx, schemaResp, &before)
	updated := before
	updated.AppID = types.StringValue("different-app")
	plan := tfsdk.Plan{Schema: schemaResp.Schema}
	if diags := plan.Set(ctx, &updated); diags.HasError() {
		t.Fatal(diags)
	}

	response := &planmodifier.StringResponse{}
	attribute.PlanModifiers[0].PlanModifyString(ctx, planmodifier.StringRequest{
		State: state, Plan: plan, ConfigValue: updated.AppID,
		StateValue: before.AppID, PlanValue: updated.AppID,
	}, response)
	if response.Diagnostics.HasError() || !response.RequiresReplace {
		t.Fatalf("app_id change failed to request replacement: %+v", response)
	}
}

func TestMCPServerDiagnosticsOmitEchoedCredentials(t *testing.T) {
	const synthetic = "synthetic-error-echo-secret"
	apiError := sdkerrors.NewSDKError("validation failed", http.StatusBadRequest, synthetic, &http.Response{StatusCode: http.StatusBadRequest})
	for _, err := range []error{apiError, fmt.Errorf("wrapped API error: %w", apiError)} {
		message := mcpServerDiagnosticError(err)
		if strings.Contains(message, synthetic) || !strings.Contains(message, "HTTP 400") {
			t.Fatalf("unsafe MCP diagnostic %q", message)
		}
	}
	if message := mcpServerResponseStatus(apiError.RawResponse); message != "HTTP 400" {
		t.Fatalf("unsafe response status diagnostic %q", message)
	}

	ctx := context.Background()
	server := httptest.NewServer(http.NotFoundHandler())
	t.Cleanup(server.Close)
	r, schemaResp := mcpRecoveryResource(t, server)
	importResp := resource.ImportStateResponse{State: tfsdk.State{
		Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), tftypes.UnknownValue),
		Schema: schemaResp.Schema,
	}}
	r.ImportState(ctx, resource.ImportStateRequest{ID: synthetic}, &importResp)
	if !importResp.Diagnostics.HasError() || strings.Contains(fmt.Sprint(importResp.Diagnostics), synthetic) {
		t.Fatalf("unsafe import diagnostic: %v", importResp.Diagnostics)
	}
}

func TestMCPServerImportMetadataUpdateDoesNotRewriteCredentials(t *testing.T) {
	ctx := context.Background()
	api := &mcpRecoveryAPI{
		t:                 t,
		bearerToken:       "remote-secret",
		connectorID:       "connector",
		description:       "before import",
		forbidCredentials: true,
	}
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)

	r, schemaResp := mcpRecoveryResource(t, server)
	importResp := resource.ImportStateResponse{State: tfsdk.State{
		Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), tftypes.UnknownValue),
		Schema: schemaResp.Schema,
	}}
	r.ImportState(ctx, resource.ImportStateRequest{ID: `{"app_id":"app","connector_id":"connector"}`}, &importResp)
	if importResp.Diagnostics.HasError() {
		t.Fatalf("importing MCP server: %v", importResp.Diagnostics)
	}

	readResp := resource.ReadResponse{State: tfsdk.State{Schema: schemaResp.Schema}}
	r.Read(ctx, resource.ReadRequest{State: importResp.State}, &readResp)
	if readResp.Diagnostics.HasError() {
		t.Fatalf("reading imported MCP server: %v", readResp.Diagnostics)
	}

	desired := mcpRecoveryModel("metadata only", "")
	plan := tfsdk.Plan{Schema: schemaResp.Schema}
	if diags := plan.Set(ctx, &desired); diags.HasError() {
		t.Fatalf("setting metadata-only MCP server plan: %v", diags)
	}
	if diags := plan.SetAttribute(ctx, path.Root("oauth2_extra_token_params"), types.MapUnknown(types.StringType)); diags.HasError() {
		t.Fatalf("setting computed OAuth map to unknown: %v", diags)
	}

	updateResp := resource.UpdateResponse{State: readResp.State}
	r.Update(ctx, resource.UpdateRequest{State: readResp.State, Plan: plan}, &updateResp)
	if updateResp.Diagnostics.HasError() {
		t.Fatalf("updating imported MCP server metadata: %v", updateResp.Diagnostics)
	}

	if api.credentialCalls != 0 {
		t.Fatalf("credential endpoint calls = %d, want 0", api.credentialCalls)
	}
	if api.bearerToken != "remote-secret" {
		t.Fatal("metadata-only update rewrote the unreadable bearer credential")
	}

	var actual MCPServerResourceModel
	if diags := updateResp.State.Get(ctx, &actual); diags.HasError() {
		t.Fatalf("reading updated MCP server state: %v", diags)
	}
	if actual.ConnectorID.ValueString() != "connector" {
		t.Fatalf("connector ID = %q, want connector", actual.ConnectorID.ValueString())
	}
}

func TestMCPServerReadRemovesStateAfterExternalDeletion(t *testing.T) {
	ctx := context.Background()
	api := &mcpRecoveryAPI{t: t, connectorID: "connector", deleted: true}
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)

	r, schemaResp := mcpRecoveryResource(t, server)
	model := mcpRecoveryModel("before deletion", "state-only-secret")
	state := mcpRecoveryState(t, ctx, schemaResp, &model)
	readResp := resource.ReadResponse{State: state}

	r.Read(ctx, resource.ReadRequest{State: state}, &readResp)
	if readResp.Diagnostics.HasError() {
		t.Fatalf("reading externally deleted MCP server: %v", readResp.Diagnostics)
	}
	if !readResp.State.Raw.IsNull() {
		t.Fatal("read retained state after C1 reported the MCP server missing")
	}
}

func TestMCPServerPartialCredentialFailureRetainsStateAndConverges(t *testing.T) {
	ctx := context.Background()
	api := &mcpRecoveryAPI{
		t:                      t,
		bearerToken:            "old-secret",
		connectorID:            "connector",
		description:            "before update",
		failedCredentialWrites: 1,
	}
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)

	r, schemaResp := mcpRecoveryResource(t, server)
	before := mcpRecoveryModel("before update", "old-secret")
	state := mcpRecoveryState(t, ctx, schemaResp, &before)
	desired := mcpRecoveryModel("after update", "new-secret")
	plan := tfsdk.Plan{Schema: schemaResp.Schema}
	if diags := plan.Set(ctx, &desired); diags.HasError() {
		t.Fatalf("setting MCP server credential plan: %v", diags)
	}

	firstResp := resource.UpdateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}
	r.Update(ctx, resource.UpdateRequest{State: state, Plan: plan}, &firstResp)
	if !firstResp.Diagnostics.HasError() {
		t.Fatal("first credential update unexpectedly succeeded")
	}
	if api.description != "after update" {
		t.Fatal("metadata endpoint did not complete before the credential failure")
	}

	var retained MCPServerResourceModel
	if diags := firstResp.State.Get(ctx, &retained); diags.HasError() {
		t.Fatalf("reading retained state after partial update: %v", diags)
	}
	if retained.AppID.ValueString() != "app" || retained.ConnectorID.ValueString() != "connector" {
		t.Fatalf("partial update lost resource identity: app=%q connector=%q", retained.AppID.ValueString(), retained.ConnectorID.ValueString())
	}
	if retained.ExternalConfig == nil || retained.ExternalConfig.BearerToken == nil || retained.ExternalConfig.BearerToken.Token.ValueString() != "old-secret" {
		t.Fatal("partial update rewrote state before the credential endpoint succeeded")
	}

	secondResp := resource.UpdateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}
	r.Update(ctx, resource.UpdateRequest{State: firstResp.State, Plan: plan}, &secondResp)
	if secondResp.Diagnostics.HasError() {
		t.Fatalf("retrying partial MCP server update: %v", secondResp.Diagnostics)
	}
	if api.metadataCalls != 2 || api.credentialCalls != 2 {
		t.Fatalf("retry calls metadata=%d credentials=%d, want 2 and 2", api.metadataCalls, api.credentialCalls)
	}
	if api.bearerToken != "new-secret" || api.description != "after update" {
		t.Fatal("retry did not converge the MCP server update")
	}
}

func TestMCPServerCreateResponseLossRequiresExplicitRecovery(t *testing.T) {
	ctx := context.Background()
	api := &mcpRecoveryAPI{t: t, connectorID: "created-connector", responseLostOnCreate: true}
	server := httptest.NewServer(api)
	t.Cleanup(server.Close)

	r, schemaResp := mcpRecoveryResource(t, server)
	model := mcpRecoveryModel("created server", "create-secret")
	plan := tfsdk.Plan{Schema: schemaResp.Schema}
	if diags := plan.Set(ctx, &model); diags.HasError() {
		t.Fatalf("setting MCP server create plan: %v", diags)
	}

	createResp := resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}
	r.Create(ctx, resource.CreateRequest{Plan: plan}, &createResp)
	if !createResp.Diagnostics.HasError() {
		t.Fatal("response-loss create unexpectedly succeeded")
	}
	if api.createCalls != 1 {
		t.Fatalf("register calls = %d, want 1", api.createCalls)
	}
	if !api.created {
		t.Fatal("response-loss fixture did not simulate a committed registration")
	}

	foundRecoveryDiagnostic := false
	for _, diagnostic := range createResp.Diagnostics.Errors() {
		if diagnostic.Summary() == "MCP server registration outcome is unknown" && strings.Contains(diagnostic.Detail(), "cannot safely locate or retry") {
			foundRecoveryDiagnostic = true
		}
	}
	if !foundRecoveryDiagnostic {
		t.Fatalf("missing explicit response-loss recovery diagnostic: %v", createResp.Diagnostics)
	}
}

func TestMCPServerCreateRecoveryDiagnosticUsesKnownImportIdentity(t *testing.T) {
	model := mcpRecoveryModel("created server", "")
	var createResp resource.CreateResponse

	addMCPServerCreatedStateRecoveryDiagnostic(&createResp, &model)
	if !createResp.Diagnostics.HasError() {
		t.Fatal("known MCP server identity did not produce a recovery diagnostic")
	}

	for _, diagnostic := range createResp.Diagnostics.Errors() {
		if diagnostic.Summary() == "MCP server was created but Terraform state was not saved" && strings.Contains(diagnostic.Detail(), `{"app_id":"app","connector_id":"connector"}`) {
			return
		}
	}
	t.Fatalf("known-identity recovery diagnostic omitted the exact import ID: %v", createResp.Diagnostics)
}
