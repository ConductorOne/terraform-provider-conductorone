package provider

import (
	"context"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestCustomEntitlementPolicyAliasSelection(t *testing.T) {
	ctx := context.Background()
	for _, tt := range []struct {
		name    string
		a, b    *tfTypes.ProvisionPolicy
		want    string
		invalid bool
	}{
		{name: "first", a: delegatedCustomPolicy(), want: "provision_policy"},
		{name: "second", b: delegatedCustomPolicy(), want: "provisioner_policy"},
		{name: "neither"},
		{name: "both", a: delegatedCustomPolicy(), b: delegatedCustomPolicy(), invalid: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cfg := testCustomConfig(t, &CustomAppEntitlementResourceModel{ProvisionPolicy: tt.a, ProvisionerPolicy: tt.b})
			got, diags := selectedCustomPolicyAlias(ctx, cfg)
			if diags.HasError() != tt.invalid || got != tt.want {
				t.Fatalf("selected %q, diagnostics %v", got, diags)
			}
		})
	}
}

func TestCustomEntitlementPolicyAliasRefreshAndWrite(t *testing.T) {
	ctx := context.Background()
	m := &CustomAppEntitlementResourceModel{AppID: types.StringValue("app"), ProvisionPolicy: &tfTypes.ProvisionPolicy{Manual: &tfTypes.ManualProvision{Instructions: types.StringValue("old")}}}
	newValue := "new"
	diags := m.RefreshFromSharedAppEntitlement(ctx, &shared.AppEntitlement{ProvisionerPolicy: &shared.ProvisionPolicy{Manual: &shared.ManualProvision{Instructions: &newValue}}})
	if diags.HasError() {
		t.Fatal(diags)
	}
	syncCustomEntitlementPolicyAliases(m)
	if m.ProvisionPolicy.Manual.Instructions.ValueString() != "new" {
		t.Fatal("configured alias did not refresh")
	}
	m.ProvisionerPolicy.Manual.Instructions = types.StringValue("selected")
	cfg := testCustomConfig(t, &CustomAppEntitlementResourceModel{ProvisionerPolicy: delegatedCustomPolicy()})
	diags = selectCustomEntitlementPolicyForWrite(ctx, cfg, m)
	if diags.HasError() {
		t.Fatal(diags)
	}
	req, diags := m.ToSharedUpdateAppEntitlementRequest(ctx)
	if diags.HasError() {
		t.Fatal(diags)
	}
	if got := *req.Entitlement.ProvisionerPolicy.Manual.Instructions; got != "selected" {
		t.Fatalf("sent %q", got)
	}
}

func TestCustomEntitlementPolicyWritePreservesOrSelectsAlias(t *testing.T) {
	ctx := context.Background()
	manual := func() *tfTypes.ProvisionPolicy {
		return &tfTypes.ProvisionPolicy{Manual: &tfTypes.ManualProvision{Instructions: types.StringValue("old")}}
	}
	reset := func() *tfTypes.ProvisionPolicy {
		return &tfTypes.ProvisionPolicy{Unconfigured: &tfTypes.UnconfiguredProvision{}}
	}
	for _, tt := range []struct {
		name              string
		configured        *CustomAppEntitlementResourceModel
		canonical, mirror *tfTypes.ProvisionPolicy
		want              string
	}{
		{"imported mirror", &CustomAppEntitlementResourceModel{}, nil, manual(), "manual"},
		{"refreshed mirror takes precedence", &CustomAppEntitlementResourceModel{}, manual(), delegatedCustomPolicy(), "delegated"},
		{"legacy canonical only", &CustomAppEntitlementResourceModel{}, manual(), nil, "manual"},
		{"canonical switch", &CustomAppEntitlementResourceModel{ProvisionPolicy: delegatedCustomPolicy()}, delegatedCustomPolicy(), manual(), "delegated"},
		{"mirror switch", &CustomAppEntitlementResourceModel{ProvisionerPolicy: delegatedCustomPolicy()}, manual(), delegatedCustomPolicy(), "delegated"},
		{"canonical reset", &CustomAppEntitlementResourceModel{ProvisionPolicy: reset()}, reset(), manual(), "unconfigured"},
		{"mirror reset", &CustomAppEntitlementResourceModel{ProvisionerPolicy: reset()}, manual(), reset(), "unconfigured"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &CustomAppEntitlementResourceModel{ProvisionPolicy: tt.canonical, ProvisionerPolicy: tt.mirror}
			if d := selectCustomEntitlementPolicyForWrite(ctx, testCustomConfig(t, tt.configured), m); d.HasError() {
				t.Fatal(d)
			}
			got, d := selectedOneofArm("provision_policy", customAppEntitlementPolicyArms(m.ProvisionPolicy))
			if d.HasError() || got != tt.want {
				t.Fatalf("got %q, want %q: %v", got, tt.want, d)
			}
		})
	}
}
