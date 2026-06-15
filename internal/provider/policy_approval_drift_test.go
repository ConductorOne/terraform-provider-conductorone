package provider

import (
	"context"
	"testing"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// These tests lock the IGA-1898 / IGA-1899 fix: the `conductorone_policy`
// approval block must not perpetually diff `false -> null`.
//
// Root cause (plan-side): object-level plan-only (UseConfigValue) on the parent
// c1.api.policy.v1.Approval schema stamped the *config* object onto the plan.
// Config-unset Computed+Optional scalar leaves are null, so the plan wanted
// null while refresh (correctly) wrote the server's literal `false` into state.
//
// Fix (overlay.yaml, regenerated into policy_resource.go): drop the object-level
// plan-only on Approval and instead disable computed on each approver oneof
// member (x-speakeasy-param-computed: false), mirroring the ProvisionPolicy
// union. With no object plan modifier clobbering the subtree, Terraform's normal
// Computed+Optional behavior keeps each leaf's prior state value, so they no
// longer re-null.
//
// These are deterministic schema-introspection / refresh assertions that run
// without a live API. The end-to-end apply->empty-plan invariant is locked by
// the (credential-gated) acceptance test TestAccPolicyResource.

// approvalAttr navigates the generated PolicyResource schema down to the
// approval SingleNestedAttribute: policy_steps > steps > approval.
func approvalAttr(t *testing.T) schema.SingleNestedAttribute {
	t.Helper()

	resp := &resource.SchemaResponse{}
	NewPolicyResource().Schema(context.Background(), resource.SchemaRequest{}, resp)

	policySteps, ok := resp.Schema.Attributes["policy_steps"].(schema.MapNestedAttribute)
	if !ok {
		t.Fatalf("policy_steps is not a MapNestedAttribute, got %T", resp.Schema.Attributes["policy_steps"])
	}
	steps, ok := policySteps.NestedObject.Attributes["steps"].(schema.ListNestedAttribute)
	if !ok {
		t.Fatalf("policy_steps.steps is not a ListNestedAttribute, got %T", policySteps.NestedObject.Attributes["steps"])
	}
	approval, ok := steps.NestedObject.Attributes["approval"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("steps.approval is not a SingleNestedAttribute, got %T", steps.NestedObject.Attributes["approval"])
	}
	return approval
}

// TestApprovalParentHasNoObjectPlanModifier is the core regression guard: the
// parent approval object must carry no object plan modifier (the UseConfigValue
// that caused the drift). If a regen re-introduces
// x-speakeasy-terraform-plan-only on c1.api.policy.v1.Approval, this fails.
//
// The parent IS Computed+Optional, mirroring the ProvisionPolicy union (see
// overlay.yaml): keeping Computed: true on the parent while disabling computed
// on each oneof member is the documented fix. Computed alone does not drift —
// only the object-level plan modifier did. The guard is the absence of that
// modifier, not the absence of Computed.
func TestApprovalParentHasNoObjectPlanModifier(t *testing.T) {
	approval := approvalAttr(t)

	if mods := approval.ObjectPlanModifiers(); len(mods) != 0 {
		t.Errorf("approval object must have no plan modifiers (UseConfigValue re-introduces the false->null drift); got %d: %#v", len(mods), mods)
	}
	if !approval.IsOptional() {
		t.Errorf("approval object must be Optional")
	}
}

// TestApprovalMembersAreNotComputed verifies the ten approver oneof members are
// Optional-only (x-speakeasy-param-computed: false). Computed members would let
// Terraform copy a stale member's prior state into the plan when config is null,
// which both breaks oneof switching and reintroduces the need for the parent
// object plan modifier.
func TestApprovalMembersAreNotComputed(t *testing.T) {
	approval := approvalAttr(t)

	members := []string{
		"agent_approval",
		"app_group_approval",
		"app_owner_approval",
		"entitlement_owner_approval",
		"expression_approval",
		"manager_approval",
		"resource_owner_approval",
		"self_approval",
		"user_approval",
		"webhook_approval",
	}

	for _, name := range members {
		attr, ok := approval.Attributes[name].(schema.SingleNestedAttribute)
		if !ok {
			t.Errorf("approval.%s is not a SingleNestedAttribute, got %T", name, approval.Attributes[name])
			continue
		}
		if attr.IsComputed() {
			t.Errorf("approval.%s must NOT be Computed (param-computed:false) so unselected members plan as null and oneof switching works", name)
		}
		if !attr.IsOptional() {
			t.Errorf("approval.%s must be Optional", name)
		}
	}
}

// TestApprovalScalarLeavesAreComputedOptional verifies the previously-drifting
// scalar leaves remain Computed+Optional. Computed is what lets Terraform retain
// the server's `false` across plans (now that no object modifier clobbers them).
func TestApprovalScalarLeavesAreComputedOptional(t *testing.T) {
	approval := approvalAttr(t)

	computedOptionalBools := []string{
		"allow_delegation",
		"allow_reassignment",
		"require_approval_reason",
		"require_denial_reason",
		"require_reassignment_reason",
		"escalation_enabled",
	}
	for _, name := range computedOptionalBools {
		attr, ok := approval.Attributes[name].(schema.BoolAttribute)
		if !ok {
			t.Errorf("approval.%s is not a BoolAttribute, got %T", name, approval.Attributes[name])
			continue
		}
		if !attr.IsComputed() || !attr.IsOptional() {
			t.Errorf("approval.%s must be Computed+Optional (computed=%v optional=%v)", name, attr.IsComputed(), attr.IsOptional())
		}
	}

	// allowed_reassignees is the list-typed leaf that also drifted ([] -> null).
	if l, ok := approval.Attributes["allowed_reassignees"].(schema.ListAttribute); !ok {
		t.Errorf("approval.allowed_reassignees is not a ListAttribute, got %T", approval.Attributes["allowed_reassignees"])
	} else if !l.IsComputed() || !l.IsOptional() {
		t.Errorf("approval.allowed_reassignees must be Computed+Optional (computed=%v optional=%v)", l.IsComputed(), l.IsOptional())
	}

	// Per-member scalar bools called out in the issue.
	userApproval, ok := approval.Attributes["user_approval"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("approval.user_approval is not a SingleNestedAttribute, got %T", approval.Attributes["user_approval"])
	}
	for _, name := range []string{"allow_self_approval", "require_distinct_approvers"} {
		attr, ok := userApproval.Attributes[name].(schema.BoolAttribute)
		if !ok {
			t.Errorf("approval.user_approval.%s is not a BoolAttribute, got %T", name, userApproval.Attributes[name])
			continue
		}
		if !attr.IsComputed() || !attr.IsOptional() {
			t.Errorf("approval.user_approval.%s must be Computed+Optional (computed=%v optional=%v)", name, attr.IsComputed(), attr.IsOptional())
		}
	}
}

// TestApprovalAssignedIsComputedOnly covers the EXTRA WRINKLE from IGA-1898:
// `assigned` is read-only (cannot be set in config) yet still drifted
// false -> null under the parent object plan modifier. It must be Computed-only;
// once no object modifier clobbers it, a Computed-only attribute keeps its prior
// state value and no longer drifts.
func TestApprovalAssignedIsComputedOnly(t *testing.T) {
	approval := approvalAttr(t)

	assigned, ok := approval.Attributes["assigned"].(schema.BoolAttribute)
	if !ok {
		t.Fatalf("approval.assigned is not a BoolAttribute, got %T", approval.Attributes["assigned"])
	}
	if !assigned.IsComputed() {
		t.Errorf("approval.assigned must be Computed")
	}
	if assigned.IsOptional() {
		t.Errorf("approval.assigned must be read-only (Computed, not Optional)")
	}
}

// boolPtr / strPtrLocal are tiny helpers for building pointer fields in the
// refresh test.
func boolPtr(b bool) *bool         { return &b }
func strPtrLocal(s string) *string { return &s }

// TestRefreshFromSharedPolicyKeepsServerFalse documents the read side: the C1
// API marshals proto3 with EmitUnpopulated=true, so these scalars come back as
// literal `false` (non-nil pointer). RefreshFromSharedPolicy must write `false`
// into state via types.BoolPointerValue — NOT null. This is why the fix is
// plan-side: a read-side false->null normalization (issue option 1) would make
// state lie about the server and is unnecessary. It also confirms an absent
// approver member refreshes to nil (null), the basis for oneof switching.
func TestRefreshFromSharedPolicyKeepsServerFalse(t *testing.T) {
	policy := &shared.Policy{
		ID: strPtrLocal("policy-123"),
		PolicySteps: map[string]shared.PolicySteps{
			"certify": {
				Steps: []shared.PolicyStep{
					{
						Approval: &shared.Approval{
							// EmitUnpopulated=true => literal false, not omitted.
							AllowDelegation:           boolPtr(false),
							AllowReassignment:         boolPtr(false),
							Assigned:                  boolPtr(false),
							EscalationEnabled:         boolPtr(false),
							RequireApprovalReason:     boolPtr(false),
							RequireDenialReason:       boolPtr(false),
							RequireReassignmentReason: boolPtr(false),
							UserApproval: &shared.UserApproval{
								AllowSelfApproval:        boolPtr(false),
								RequireDistinctApprovers: boolPtr(false),
							},
						},
					},
				},
			},
		},
	}

	var model PolicyResourceModel
	if diags := model.RefreshFromSharedPolicy(context.Background(), policy); diags.HasError() {
		t.Fatalf("RefreshFromSharedPolicy returned errors: %v", diags)
	}

	steps := model.PolicySteps["certify"].Steps
	if len(steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(steps))
	}
	approval := steps[0].Approval
	if approval == nil {
		t.Fatal("expected non-nil approval")
	}

	// Every server `false` must land in state as a known false, never null.
	leaves := map[string]types.Bool{
		"allow_delegation":            approval.AllowDelegation,
		"allow_reassignment":          approval.AllowReassignment,
		"assigned":                    approval.Assigned,
		"escalation_enabled":          approval.EscalationEnabled,
		"require_approval_reason":     approval.RequireApprovalReason,
		"require_denial_reason":       approval.RequireDenialReason,
		"require_reassignment_reason": approval.RequireReassignmentReason,
	}
	for name, v := range leaves {
		if v.IsNull() || v.IsUnknown() {
			t.Errorf("approval.%s refreshed to null/unknown; expected known false", name)
			continue
		}
		if v.ValueBool() != false {
			t.Errorf("approval.%s = %v, expected false", name, v.ValueBool())
		}
	}

	if approval.UserApproval == nil {
		t.Fatal("expected non-nil user_approval")
	}
	if approval.UserApproval.AllowSelfApproval.IsNull() || approval.UserApproval.AllowSelfApproval.ValueBool() != false {
		t.Errorf("user_approval.allow_self_approval = %v, expected known false", approval.UserApproval.AllowSelfApproval)
	}

	// An approver member the server did not return must be nil (null in state),
	// which is what lets unselected members stay null without drift.
	if approval.ManagerApproval != nil {
		t.Errorf("manager_approval should refresh to nil when absent from the API response, got %#v", approval.ManagerApproval)
	}
}
