package provider

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Exercise import, refresh, protocol planning, and the update request.
func TestIndependentIGA5240ImportedPolicyUnrelatedUpdate(t *testing.T) {
	ctx := context.Background()
	var updateBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			updateBody, _ = io.ReadAll(r.Body)
			t.Logf("request %s %s body=%s", r.Method, r.URL, updateBody)
		}
		name := "original"
		if r.Method != "GET" {
			name = "renamed"
		}
		_, _ = io.WriteString(w, `{"appEntitlementView":{"appEntitlement":{"id":"ent1","appId":"app1","displayName":"`+name+`","appResourceId":"res1","appResourceTypeId":"type1","grantPolicyId":"grant1","revokePolicyId":"revoke1","deprovisionerPolicy":{"manual":{"instructions":"retain"}},"provisionerPolicy":{"delegated":{"appId":"app2","entitlementId":"ent2"}}}}}`)
	}))
	defer server.Close()
	r := &CustomAppEntitlementResource{client: sdk.New(sdk.WithServerURL(server.URL))}
	sr := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, sr)
	initial := tfsdk.State{Schema: sr.Schema}
	if d := initial.Set(ctx, &CustomAppEntitlementResourceModel{AppID: types.StringValue("app1"), ID: types.StringValue("ent1")}); d.HasError() {
		t.Fatal(d)
	}
	imp := &resource.ImportStateResponse{State: initial}
	r.ImportState(ctx, resource.ImportStateRequest{ID: `{"app_id":"app1","id":"ent1"}`}, imp)
	if imp.Diagnostics.HasError() {
		t.Fatal(imp.Diagnostics)
	}
	read := &resource.ReadResponse{State: imp.State}
	r.Read(ctx, resource.ReadRequest{State: imp.State}, read)
	if read.Diagnostics.HasError() {
		t.Fatal(read.Diagnostics)
	}
	var values map[string]tftypes.Value
	if err := read.State.Raw.As(&values); err != nil {
		t.Fatal(err)
	}
	if !values["provision_policy"].IsNull() || values["provisioner_policy"].IsNull() {
		t.Fatalf("unexpected read aliases: %v", values)
	}
	t.Log("Actual ImportState + Read: provision_policy null; provisioner_policy populated")
	cfg := testCustomConfig(t, &CustomAppEntitlementResourceModel{AppID: types.StringValue("app1"), DisplayName: types.StringValue("renamed"), AppResourceID: types.StringValue("res1"), AppResourceTypeID: types.StringValue("type1")})
	values["display_name"] = tftypes.NewValue(tftypes.String, "renamed")
	proposed := tftypes.NewValue(read.State.Raw.Type(), values)
	dynamic := func(v tftypes.Value) *tfprotov6.DynamicValue {
		d, e := tfprotov6.NewDynamicValue(v.Type(), v)
		if e != nil {
			t.Fatal(e)
		}
		return &d
	}
	ps := providerserver.NewProtocol6(New("test")())()
	validated, err := ps.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{TypeName: "conductorone_custom_app_entitlement", Config: dynamic(cfg.Raw)})
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range validated.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("config validation %s: %s", d.Summary, d.Detail)
		}
	}
	t.Log("Real ValidateResourceConfig accepts config with both policy aliases absent")
	planned, e := ps.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{TypeName: "conductorone_custom_app_entitlement", PriorState: dynamic(read.State.Raw), ProposedNewState: dynamic(proposed), Config: dynamic(cfg.Raw)})
	if e != nil {
		t.Fatal(e)
	}
	for _, d := range planned.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("plan %s: %s", d.Summary, d.Detail)
		}
	}
	if len(planned.RequiresReplace) != 0 {
		t.Fatalf("replacement required: %v", planned.RequiresReplace)
	}
	p, e := planned.PlannedState.Unmarshal(read.State.Raw.Type())
	if e != nil {
		t.Fatal(e)
	}
	var pv map[string]tftypes.Value
	if e := p.As(&pv); e != nil {
		t.Fatal(e)
	}
	t.Logf("RPC planned aliases: canonical null=%t mirror null=%t fully known=%t; replacements=%d", pv["provision_policy"].IsNull(), pv["provisioner_policy"].IsNull(), pv["provisioner_policy"].IsFullyKnown(), len(planned.RequiresReplace))
	out := &resource.UpdateResponse{State: read.State}
	r.Update(ctx, resource.UpdateRequest{State: read.State, Plan: tfsdk.Plan{Schema: sr.Schema, Raw: p}, Config: cfg}, out)
	if out.Diagnostics.HasError() {
		t.Fatal(out.Diagnostics)
	}
	if len(updateBody) == 0 {
		t.Fatal("no update request")
	}
	var body map[string]any
	if e := json.Unmarshal(updateBody, &body); e != nil {
		t.Fatal(e)
	}
	entitlement, ok := body["entitlement"].(map[string]any)
	if !ok {
		t.Fatal(string(updateBody))
	}
	policy, ok := entitlement["provisionerPolicy"].(map[string]any)
	if !ok {
		t.Fatalf("refreshed policy missing: %s", updateBody)
	}
	delegated, ok := policy["delegated"].(map[string]any)
	if !ok || delegated["appId"] != "app2" || delegated["entitlementId"] != "ent2" {
		t.Fatalf("refreshed delegated policy changed: %s", updateBody)
	}
	if entitlement["grantPolicyId"] != "grant1" || entitlement["revokePolicyId"] != "revoke1" {
		t.Fatalf("unrelated policy IDs changed: %s", updateBody)
	}
	deprovisioner, ok := entitlement["deprovisionerPolicy"].(map[string]any)
	if !ok {
		t.Fatalf("deprovisioner missing: %s", updateBody)
	}
	manual, ok := deprovisioner["manual"].(map[string]any)
	if !ok || manual["instructions"] != "retain" {
		t.Fatalf("deprovisioner changed: %s", updateBody)
	}
	if _, ok := body["updateMask"]; ok {
		t.Fatalf("mask unexpectedly included: %s", updateBody)
	}
	var result CustomAppEntitlementResourceModel
	if d := out.State.Get(ctx, &result); d.HasError() || result.ID.ValueString() != "ent1" {
		t.Fatalf("ID changed or state invalid: %v %v", result.ID, d)
	}
	second, err := ps.PlanResourceChange(ctx, &tfprotov6.PlanResourceChangeRequest{TypeName: "conductorone_custom_app_entitlement", PriorState: dynamic(out.State.Raw), ProposedNewState: dynamic(out.State.Raw), Config: dynamic(cfg.Raw)})
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range second.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("second plan %s: %s", d.Summary, d.Detail)
		}
	}
	secondState, err := second.PlannedState.Unmarshal(out.State.Raw.Type())
	if err != nil || !secondState.Equal(out.State.Raw) {
		t.Fatalf("subsequent plan changed state: %v", err)
	}
}
