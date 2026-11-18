package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/retry"
)

type releaseFakeServer struct {
	mu        sync.Mutex
	getStatus int
	getBody   string
	gets      int
	deletes   int
	srv       *httptest.Server
}

func newReleaseFakeServer(t *testing.T, getStatus int, getBody string) *releaseFakeServer {
	t.Helper()
	f := &releaseFakeServer{getStatus: getStatus, getBody: getBody}
	f.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		defer f.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet:
			f.gets++
			w.WriteHeader(f.getStatus)
			_, _ = w.Write([]byte(f.getBody))
		case r.Method == http.MethodDelete || (r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/delete")):
			f.deletes++
			_, _ = w.Write([]byte(`{}`))
		default:
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusTeapot)
		}
	}))
	t.Cleanup(f.srv.Close)
	return f
}

func (f *releaseFakeServer) client() *sdk.ConductoroneAPI {
	return sdk.New(
		sdk.WithServerURL(f.srv.URL),
		sdk.WithSecurity(shared.Security{BearerAuth: "test", Oauth: "test"}),
		sdk.WithRetryConfig(retry.Config{Strategy: "none"}),
	)
}

func releaseTestState(t *testing.T, ctx context.Context, r resource.Resource, attrs map[string]string) tfsdk.State {
	t.Helper()
	var sr resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &sr)
	if sr.Diagnostics.HasError() {
		t.Fatalf("schema: %v", sr.Diagnostics)
	}
	st := tfsdk.State{Schema: sr.Schema, Raw: tftypes.NewValue(sr.Schema.Type().TerraformType(ctx), nil)}
	for k, v := range attrs {
		if d := st.SetAttribute(ctx, path.Root(k), v); d.HasError() {
			t.Fatalf("set %s: %v", k, d)
		}
	}
	return st
}

func warningSummaries(d diag.Diagnostics) []string {
	var out []string
	for _, w := range d.Warnings() {
		out = append(out, w.Summary())
	}
	return out
}

type releaseCase struct {
	name        string
	getStatus   int
	getBody     string
	wantDeletes int
	wantWarning bool
	wantError   bool
}

func releaseCases(viewKey, rowKey string) []releaseCase {
	wrap := func(row string) string { return `{"` + viewKey + `":{"` + rowKey + `":` + row + `}}` }
	return []releaseCase{
		{name: "manually managed is deleted", getStatus: 200, getBody: wrap(`{"id":"x","isManuallyManaged":true}`), wantDeletes: 1},
		{name: "connector managed is released", getStatus: 200, getBody: wrap(`{"id":"x","isManuallyManaged":false}`), wantWarning: true},
		{name: "already gone 404", getStatus: 404, getBody: `{}`},
		{name: "deleted_at set", getStatus: 200, getBody: wrap(`{"id":"x","isManuallyManaged":false,"deletedAt":"2026-01-01T00:00:00Z"}`)},
		{name: "get 500 surfaces error", getStatus: 500, getBody: `{"message":"boom"}`, wantError: true},
		{name: "is_manually_managed absent falls through to delete", getStatus: 200, getBody: wrap(`{"id":"x"}`), wantDeletes: 1},
	}
}

func runReleaseCase(t *testing.T, tc releaseCase, newResource func(*sdk.ConductoroneAPI) resource.Resource, attrs map[string]string) {
	ctx := context.Background()
	f := newReleaseFakeServer(t, tc.getStatus, tc.getBody)
	r := newResource(f.client())
	resp := &resource.DeleteResponse{State: releaseTestState(t, ctx, r, attrs)}
	r.Delete(ctx, resource.DeleteRequest{State: releaseTestState(t, ctx, r, attrs)}, resp)

	if got := resp.Diagnostics.HasError(); got != tc.wantError {
		t.Fatalf("HasError=%v want %v: %v", got, tc.wantError, resp.Diagnostics)
	}
	if f.gets != 1 {
		t.Errorf("GET calls=%d want 1", f.gets)
	}
	if f.deletes != tc.wantDeletes {
		t.Errorf("DELETE calls=%d want %d", f.deletes, tc.wantDeletes)
	}
	ws := warningSummaries(resp.Diagnostics)
	if tc.wantWarning {
		if len(ws) != 1 || ws[0] != connectorManagedReleaseSummary {
			t.Errorf("warnings=%v want [%q]", ws, connectorManagedReleaseSummary)
		}
	} else if len(ws) != 0 {
		t.Errorf("unexpected warnings %v", ws)
	}
}

func TestAppResourceDeleteReleasesConnectorManaged(t *testing.T) {
	attrs := map[string]string{"app_id": "app1", "app_resource_type_id": "rt1", "id": "res1"}
	for _, tc := range releaseCases("appResourceView", "appResource") {
		t.Run(tc.name, func(t *testing.T) {
			runReleaseCase(t, tc, func(c *sdk.ConductoroneAPI) resource.Resource { return &AppResourceResource{client: c} }, attrs)
		})
	}
}

func TestCustomAppEntitlementDeleteReleasesConnectorManaged(t *testing.T) {
	attrs := map[string]string{"app_id": "app1", "id": "ent1"}
	for _, tc := range releaseCases("appEntitlementView", "appEntitlement") {
		t.Run(tc.name, func(t *testing.T) {
			runReleaseCase(t, tc, func(c *sdk.ConductoroneAPI) resource.Resource { return &CustomAppEntitlementResource{client: c} }, attrs)
		})
	}
}
