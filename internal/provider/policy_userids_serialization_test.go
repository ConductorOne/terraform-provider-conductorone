package provider

import (
	"context"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// These tests anchor the IGA-1898 investigation of the live-C1 acceptance
// failure:
//
//	TestAccPolicyResource -> 400 invalid Approval.Users:
//	invalid UserApproval.UserIds: value must contain between 1 and 32 items
//
// The reported hypothesis was that the approval bool drift fix
// (x-speakeasy-param-computed:false on each approver oneof member, object-level
// UseConfigValue removed) OVER-CORRECTED and stripped a *configured*
// UserApproval.UserIds on CREATE. These tests exercise the TF->API create
// builder (ToSharedCreatePolicyRequest) directly and show that is NOT the case:
// configured user_ids serialize verbatim. The create path is correct.
//
// The actual cause of the 400 is that the acceptance config selects a
// user_approval approver but never sets user_ids, so the create request carries
// an empty list -- which the live API (proto buf.validate, min 1 item) rejects.
// TestToSharedCreatePolicyRequestUserApprovalEmptyUserIds documents that shape.

// buildUserApprovalCreateModel returns a minimal PolicyResourceModel with a
// single certify step whose approval selects user_approval with the supplied
// user_ids. A nil userIds slice models "user_approval selected, user_ids unset".
func buildUserApprovalCreateModel(userIds []string) PolicyResourceModel {
	var tfUserIds []types.String
	if userIds != nil {
		tfUserIds = make([]types.String, 0, len(userIds))
		for _, id := range userIds {
			tfUserIds = append(tfUserIds, types.StringValue(id))
		}
	}

	return PolicyResourceModel{
		DisplayName: types.StringValue("acc user_approval policy"),
		PolicyType:  types.StringValue("POLICY_TYPE_CERTIFY"),
		PolicySteps: map[string]tfTypes.PolicySteps{
			"certify": {
				Steps: []tfTypes.PolicyStep{
					{
						Approval: &tfTypes.Approval{
							UserApproval: &tfTypes.UserApproval{
								AllowSelfApproval:        types.BoolValue(false),
								RequireDistinctApprovers: types.BoolValue(false),
								UserIds:                  tfUserIds,
							},
						},
					},
				},
			},
		},
	}
}

// userApprovalFromCreate navigates the create request down to the serialized
// shared.UserApproval for the certify baseline step.
func userApprovalFromCreate(t *testing.T, model PolicyResourceModel) (userIds []string) {
	t.Helper()

	req, diags := model.ToSharedCreatePolicyRequest(context.Background())
	if diags.HasError() {
		t.Fatalf("ToSharedCreatePolicyRequest returned errors: %v", diags)
	}
	steps, ok := req.PolicySteps["certify"]
	if !ok {
		t.Fatalf("create request missing certify policy steps; got keys %v", keysOf(req.PolicySteps))
	}
	if len(steps.Steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(steps.Steps))
	}
	approval := steps.Steps[0].Approval
	if approval == nil {
		t.Fatal("create request step has nil approval")
	}
	if approval.UserApproval == nil {
		t.Fatal("create request approval has nil user_approval; the selected member was dropped")
	}
	return approval.UserApproval.UserIds
}

func keysOf[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// TestToSharedCreatePolicyRequestSerializesUserApprovalUserIds is the core
// finding: a configured UserApproval.UserIds is serialized verbatim onto the
// CREATE request. If the drift fix had stripped the configured list (the
// reported hypothesis), this would observe an empty/nil slice. It does not --
// proving the create path is correct and no overlay/spec change is required.
func TestToSharedCreatePolicyRequestSerializesUserApprovalUserIds(t *testing.T) {
	want := []string{"user-aaa", "user-bbb"}
	got := userApprovalFromCreate(t, buildUserApprovalCreateModel(want))

	if len(got) != len(want) {
		t.Fatalf("serialized user_ids = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("serialized user_ids[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

// TestToSharedCreatePolicyRequestUserApprovalEmptyUserIds documents the actual
// cause of the live-C1 400: when user_approval is selected but user_ids is unset
// in config, the create request carries an empty user_ids list. The live API
// requires 1-32 items, so this shape is rejected. The fix lives in the
// acceptance config (supply user_ids), not in the provider serialization.
func TestToSharedCreatePolicyRequestUserApprovalEmptyUserIds(t *testing.T) {
	got := userApprovalFromCreate(t, buildUserApprovalCreateModel(nil))

	if len(got) != 0 {
		t.Fatalf("expected empty user_ids when unset in config, got %v", got)
	}
}
