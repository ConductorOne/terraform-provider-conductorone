package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// TestAccPolicyResource locks the IGA-1898 / IGA-1899 invariant: a
// conductorone_policy approval step must not perpetually diff false → null on
// the approval scalar bools (allow_delegation, allow_reassignment, assigned,
// escalation_enabled, require_approval_reason, require_denial_reason,
// require_reassignment_reason, and per-approver allow_self_approval /
// require_distinct_approvers).
//
// This is the exact repro from the issue, encoded as three steps:
//
//	Step 1 — explicit falses + app_owner_approval: every drift-prone scalar bool
//	         is spelled out as false (including escalation_enabled = false) and an
//	         app_owner_approval approver is selected (its own drift bools
//	         allow_self_approval / require_distinct_approvers spelled false too).
//	         After apply the plan must be empty (no false → null churn).
//
//	         NOTE: this step does NOT set an `escalation = { skip_step = {} }`
//	         object. With escalation_enabled = false, live C1 does not persist or
//	         return the escalation block, so a config-set escalation object reads
//	         back as null — a permanent null↔object [update] that is NOT the
//	         IGA-1898 bool drift and is not fixable via any plan modifier (a known
//	         config value can't be suppressed). It was a self-contradictory config;
//	         the escalation_enabled bool (which IS a drift-prone IGA-1898 scalar)
//	         stays covered. Drift attribution proven by the diagnostic PlanCheck in
//	         the separate PR (ExpectEmptyPlanWithDriftLog).
//	Step 2 — NOTHING specified: an approval step with a single approver and
//	         none of the optional scalar bools set. This is the original bug
//	         shape — config-unset leaves planned null while the API returned
//	         literal false. After apply the plan must be empty.
//	Step 3 — read-only `assigned`: re-applies step 1's config to confirm the
//	         read-only `assigned` field (which cannot be set in config and
//	         still drifted under the old object plan modifier) is stable.
//
// NOTE (IGA-1898): the selected approver is app_owner_approval rather
// than user_approval. A user_approval step requires a non-empty user_ids list
// (live API: 1–32 items), which has no stable CI-tenant fixture; app_owner_approval
// carries the same per-member drift bools with no required user list, so the
// false→null drift signature is still fully exercised under ExpectEmptyPlan.
// user_ids serialization on CREATE is locked separately by the unit tests in
// policy_userids_serialization_test.go.
//
// Each step asserts an empty post-apply / post-refresh plan via
// plancheck.ExpectEmptyPlan(); terraform-plugin-testing also fails any step on
// a non-empty implicit plan. A regen that re-introduces object-level plan-only
// (UseConfigValue) on c1.api.policy.v1.Approval, or marks the approver oneof
// members Computed again, fails here.
//
// Was skipped while the parent Approval schema carried object-level plan-only,
// which stamped config-null leaves back into the plan. The fix (overlay.yaml)
// drops that cascade and disables computed on the approver oneof members,
// mirroring the existing ProvisionPolicy union treatment.
//
// NOTE: acceptance tests require TF_ACC=1 and live ConductorOne tenant
// credentials (see provider_test.go / CONDUCTORONE_* env). They are skipped by
// the standard `make test` unit run.
func TestAccPolicyResource(t *testing.T) {
	emptyPlanAfterApply := resource.ConfigPlanChecks{
		PostApplyPreRefresh:  []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
		PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: explicit falses + app_owner_approval.
			{
				Config: providerConfig + `
				resource "conductorone_policy" "test" {
					display_name                = "Terraform Created Certify Policy"
					description                 = "explicit falses + app_owner_approval"
					policy_type                 = "POLICY_TYPE_CERTIFY"
					reassign_tasks_to_delegates = "false"
					policy_steps = {
						certify = {
							steps = [
								{
									approval = {
										allow_delegation            = "false"
										allow_reassignment          = "false"
										require_approval_reason     = "false"
										require_denial_reason       = "false"
										require_reassignment_reason = "false"
										escalation_enabled          = "false"
										app_owner_approval = {
											allow_self_approval        = "false"
											require_distinct_approvers = "false"
										}
									}
								}
							]
						}
					}
				}
				`,
				ConfigPlanChecks: emptyPlanAfterApply,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_type", "POLICY_TYPE_CERTIFY"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.allow_delegation", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.allow_reassignment", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_approval_reason", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_denial_reason", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_reassignment_reason", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.escalation_enabled", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.app_owner_approval.allow_self_approval", "false"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.app_owner_approval.require_distinct_approvers", "false"),
					// assigned is read-only; live C1 returns it as null on a fresh
					// policy-definition read (it is populated per-task at runtime, not
					// on the template). Drift is covered by ExpectEmptyPlan +
					// TestApprovalAssignedIsComputedOnly, so just assert it stays absent.
					resource.TestCheckNoResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.assigned"),
				),
			},
			// Step 2: NOTHING specified — a single approver, no optional scalar
			// bools. The original drift shape.
			{
				Config: providerConfig + `
				resource "conductorone_policy" "test" {
					display_name                = "Terraform Created Certify Policy"
					description                 = "nothing specified"
					policy_type                 = "POLICY_TYPE_CERTIFY"
					reassign_tasks_to_delegates = "false"
					policy_steps = {
						certify = {
							steps = [
								{
									approval = {
										app_owner_approval = {}
									}
								}
							]
						}
					}
				}
				`,
				ConfigPlanChecks: emptyPlanAfterApply,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("conductorone_policy.test", "description", "nothing specified"),
					// Never-set optional scalar bools read back absent (null) from a
					// fresh policy-definition read on live C1 — the same shape as the
					// read-only `assigned` field below. The IGA-1898 invariant locked
					// here is convergence (an empty post-apply plan via ExpectEmptyPlan),
					// not a literal `false`: these are stable at null (no false↔null
					// flapping), so assert absence rather than "false".
					resource.TestCheckNoResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.allow_delegation"),
					resource.TestCheckNoResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.require_approval_reason"),
					// assigned is read-only and null on a fresh policy-definition read (see Step 1).
					resource.TestCheckNoResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.assigned"),
				),
			},
			// Step 3: re-apply step 1 to confirm read-only `assigned` and the
			// scalar bools remain stable across an approver switch back.
			{
				Config: providerConfig + `
				resource "conductorone_policy" "test" {
					display_name                = "Terraform Created Certify Policy"
					description                 = "explicit falses + app_owner_approval"
					policy_type                 = "POLICY_TYPE_CERTIFY"
					reassign_tasks_to_delegates = "false"
					policy_steps = {
						certify = {
							steps = [
								{
									approval = {
										allow_delegation            = "false"
										allow_reassignment          = "false"
										require_approval_reason     = "false"
										require_denial_reason       = "false"
										require_reassignment_reason = "false"
										escalation_enabled          = "false"
										app_owner_approval = {
											allow_self_approval        = "false"
											require_distinct_approvers = "false"
										}
									}
								}
							]
						}
					}
				}
				`,
				ConfigPlanChecks: emptyPlanAfterApply,
				Check: resource.ComposeAggregateTestCheckFunc(
					// assigned is read-only and null on a fresh policy-definition read (see Step 1).
					resource.TestCheckNoResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.assigned"),
					resource.TestCheckResourceAttr("conductorone_policy.test", "policy_steps.certify.steps.0.approval.app_owner_approval.allow_self_approval", "false"),
				),
			},
		},
	})
}
