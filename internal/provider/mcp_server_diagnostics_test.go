package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	sdkerrors "github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/errors"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestMCPServerDiagnosticsProtectCredentialsAcrossOperations(t *testing.T) {
	const secret = "synthetic-mcp-response-secret"
	for _, operation := range []string{"register", "read", "update metadata for", "update credentials for", "delete"} {
		t.Run(operation, func(t *testing.T) {
			ctx := context.Background()
			var failedRequest atomic.Bool
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if operation == "update credentials for" && r.URL.Path == "/api/v1/apps/app/mcp_servers/connector" {
					w.Header().Set("Content-Type", "application/json")
					if err := json.NewEncoder(w).Encode(map[string]any{"mcpServer": map[string]any{
						"appId": "app", "connectorId": "connector", "description": "after",
					}}); err != nil {
						t.Errorf("encode metadata response: %v", err)
					}
					return
				}
				failedRequest.Store(true)
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-Error-Detail", secret)
				w.WriteHeader(http.StatusBadRequest)
				if err := json.NewEncoder(w).Encode(map[string]any{"code": 3, "message": "invalid credential: " + secret}); err != nil {
					t.Errorf("encode error response: %v", err)
				}
			}))
			t.Cleanup(server.Close)
			r, schemaResp := mcpRecoveryResource(t, server)
			before := mcpRecoveryModel("before", "old-secret")
			state := mcpRecoveryState(t, ctx, schemaResp, &before)
			after := mcpRecoveryModel("after", secret)
			plan := tfsdk.Plan{Schema: schemaResp.Schema}
			if diags := plan.Set(ctx, &after); diags.HasError() {
				t.Fatal(diags)
			}

			var diagnostics diag.Diagnostics
			switch operation {
			case "register":
				response := resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema}}
				r.Create(ctx, resource.CreateRequest{Plan: plan}, &response)
				diagnostics = response.Diagnostics
			case "read":
				response := resource.ReadResponse{State: state}
				r.Read(ctx, resource.ReadRequest{State: state}, &response)
				diagnostics = response.Diagnostics
			case "update metadata for", "update credentials for":
				response := resource.UpdateResponse{State: state}
				r.Update(ctx, resource.UpdateRequest{State: state, Plan: plan}, &response)
				diagnostics = response.Diagnostics
			case "delete":
				response := resource.DeleteResponse{State: state}
				r.Delete(ctx, resource.DeleteRequest{State: state}, &response)
				diagnostics = response.Diagnostics
			}

			message := fmt.Sprint(diagnostics)
			if !failedRequest.Load() || !diagnostics.HasError() {
				t.Fatalf("operation did not report the API failure: %v", diagnostics)
			}
			if strings.Contains(message, secret) {
				t.Fatal("MCP error diagnostics disclosed the response credential")
			}
			if !strings.Contains(message, "HTTP 400") || !strings.Contains(message, operation) {
				t.Fatalf("diagnostics lost operation/status context: %v", diagnostics)
			}
		})
	}
}

func TestMCPServerDiagnosticErrorProtectsWrappedAndTransportDetails(t *testing.T) {
	const secret = "synthetic-mcp-error-secret"
	apiError := sdkerrors.NewSDKError(secret, http.StatusBadRequest, secret, &http.Response{StatusCode: http.StatusBadRequest})
	for _, tc := range []struct {
		name string
		err  error
		want string
	}{
		{"API", apiError, "HTTP 400"},
		{"wrapped API", fmt.Errorf("%s: %w", secret, apiError), "HTTP 400"},
		{"transport", errors.New("request URL or response contains " + secret), "request failed"},
		{"canceled", fmt.Errorf("%s: %w", secret, context.Canceled), "canceled"},
		{"deadline", fmt.Errorf("%s: %w", secret, context.DeadlineExceeded), "deadline"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			message := mcpServerDiagnosticError(tc.err)
			if strings.Contains(message, secret) || !strings.Contains(message, tc.want) {
				t.Fatalf("unsafe or unhelpful MCP diagnostic: %q", message)
			}
		})
	}
}

func TestMCPServerMalformedImportDiagnosticOmitsInput(t *testing.T) {
	const secret = "synthetic-mcp-import-secret"
	ctx := context.Background()
	server := httptest.NewServer(http.NotFoundHandler())
	t.Cleanup(server.Close)
	r, schemaResp := mcpRecoveryResource(t, server)
	response := resource.ImportStateResponse{State: tfsdk.State{
		Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), tftypes.UnknownValue),
		Schema: schemaResp.Schema,
	}}
	r.ImportState(ctx, resource.ImportStateRequest{ID: `{"` + secret + `":"value"}`}, &response)
	if !response.Diagnostics.HasError() || strings.Contains(fmt.Sprint(response.Diagnostics), secret) {
		t.Fatalf("unsafe malformed import diagnostic: %v", response.Diagnostics)
	}
}
