package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mcpApprovalAPI(t *testing.T, initial string, ignoreUpdate bool) *httptest.Server {
	t.Helper()
	var mu sync.Mutex
	approval := initial
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if r.URL.Path != "/api/v1/apps/app/mcp_servers/connector" && r.URL.Path != "/api/v1/apps/app/mcp_servers/connector/credentials" {
			http.NotFound(w, r)
			return
		}
		if r.Method == http.MethodPost && !strings.HasSuffix(r.URL.Path, "/credentials") {
			var update shared.MCPServerServiceUpdateRequest
			if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
				t.Errorf("decode metadata update: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !ignoreUpdate && update.UpdateMask != nil && strings.Contains(*update.UpdateMask, "requireToolApproval") {
				approval = "OPTIONAL_BOOL_UNSPECIFIED"
				if update.McpServer != nil && update.McpServer.RequireToolApproval != nil {
					approval = string(*update.McpServer.RequireToolApproval)
				}
			}
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{"mcpServer": map[string]any{
			"appId": "app", "connectorId": "connector", "requireToolApproval": approval,
		}}); err != nil {
			t.Errorf("encode server response: %v", err)
		}
	}))
	t.Cleanup(server.Close)
	return server
}

func mcpApprovalModel(hosted bool, approval types.String) MCPServerResourceModel {
	model := MCPServerResourceModel{
		AppID: types.StringValue("app"), ConnectorID: types.StringValue("connector"),
		RequireToolApproval: types.StringValue("OPTIONAL_BOOL_TRUE"),
	}
	if hosted {
		model.HostedConfig = &tfTypes.MCPServerHostedConfig{RequireToolApproval: approval}
	} else {
		model.ExternalConfig = &tfTypes.MCPServerExternalConfig{
			RequireToolApproval: approval,
			BearerToken:         &tfTypes.MCPServerAuthBearerToken{Token: types.StringValue("test-only-secret")},
		}
	}
	return model
}

func TestMCPApprovalUpdate(t *testing.T) {
	for _, tc := range []struct {
		name         string
		hosted       bool
		desired      types.String
		ignoreUpdate bool
		want         types.String
	}{
		{"external disables approval", false, types.StringValue("OPTIONAL_BOOL_FALSE"), false, types.StringValue("OPTIONAL_BOOL_FALSE")},
		{"hosted clears override", true, types.StringValue("OPTIONAL_BOOL_UNSPECIFIED"), false, types.StringValue("OPTIONAL_BOOL_UNSPECIFIED")},
		{"omitted remains unmanaged", false, types.StringNull(), false, types.StringNull()},
		{"server disagreement survives plan restoration", false, types.StringValue("OPTIONAL_BOOL_FALSE"), true, types.StringValue("OPTIONAL_BOOL_TRUE")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			server := mcpApprovalAPI(t, "OPTIONAL_BOOL_TRUE", tc.ignoreUpdate)
			r := &MCPServerResource{client: sdk.New(sdk.WithServerURL(server.URL))}
			var schemaResp resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
			state := tfsdk.State{Schema: schemaResp.Schema}
			before := mcpApprovalModel(tc.hosted, types.StringValue("OPTIONAL_BOOL_TRUE"))
			if diags := state.Set(ctx, &before); diags.HasError() {
				t.Fatal(diags)
			}
			plan := tfsdk.Plan{Schema: schemaResp.Schema}
			desired := mcpApprovalModel(tc.hosted, tc.desired)
			desired.RequireToolApproval = types.StringUnknown()
			if diags := plan.Set(ctx, &desired); diags.HasError() {
				t.Fatal(diags)
			}
			response := resource.UpdateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}
			r.Update(ctx, resource.UpdateRequest{State: state, Plan: plan}, &response)
			if response.Diagnostics.HasError() {
				t.Fatal(response.Diagnostics)
			}
			var actual MCPServerResourceModel
			if diags := response.State.Get(ctx, &actual); diags.HasError() {
				t.Fatal(diags)
			}
			if !actual.configuredToolApproval().Equal(tc.want) {
				t.Fatalf("approval state = %s, want %s", actual.configuredToolApproval(), tc.want)
			}
			if tc.desired.IsNull() && actual.RequireToolApproval.ValueString() != "OPTIONAL_BOOL_TRUE" {
				t.Fatal("omitted approval changed the stored override")
			}
			readback, err := server.Client().Get(server.URL + "/api/v1/apps/app/mcp_servers/connector")
			if err != nil {
				t.Fatal(err)
			}
			defer readback.Body.Close()
			var observed shared.MCPServerServiceGetResponse
			if err := json.NewDecoder(readback.Body).Decode(&observed); err != nil {
				t.Fatal(err)
			}
			wantRemote := tc.want.ValueString()
			if tc.desired.IsNull() {
				wantRemote = "OPTIONAL_BOOL_TRUE"
			}
			if observed.McpServer == nil || observed.McpServer.RequireToolApproval == nil || string(*observed.McpServer.RequireToolApproval) != wantRemote {
				t.Fatalf("API approval = %+v, want %s", observed.McpServer, wantRemote)
			}
		})
	}
}

func TestMCPApprovalReadExposesDriftWithoutLosingSecret(t *testing.T) {
	ctx := context.Background()
	server := mcpApprovalAPI(t, "OPTIONAL_BOOL_TRUE", false)
	r := &MCPServerResource{client: sdk.New(sdk.WithServerURL(server.URL))}
	var schemaResp resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
	state := tfsdk.State{Schema: schemaResp.Schema}
	before := mcpApprovalModel(false, types.StringValue("OPTIONAL_BOOL_FALSE"))
	if diags := state.Set(ctx, &before); diags.HasError() {
		t.Fatal(diags)
	}
	response := resource.ReadResponse{State: tfsdk.State{Schema: schemaResp.Schema}}
	r.Read(ctx, resource.ReadRequest{State: state}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	var actual MCPServerResourceModel
	if diags := response.State.Get(ctx, &actual); diags.HasError() {
		t.Fatal(diags)
	}
	if actual.ExternalConfig.RequireToolApproval.ValueString() != "OPTIONAL_BOOL_TRUE" {
		t.Fatal("read concealed the server's out-of-band approval change")
	}
	if !actual.ExternalConfig.BearerToken.Token.Equal(before.ExternalConfig.BearerToken.Token) {
		t.Fatal("approval refresh lost the write-only bearer credential")
	}
}
