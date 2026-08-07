package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

const replacementModifierDescription = "Terraform will destroy and recreate the resource"

func TestAccessReviewImmutableConfiguredAttributesRequireReplacement(t *testing.T) {
	ctx := context.Background()
	resp := &resource.SchemaResponse{}
	NewAccessReviewResource().Schema(ctx, resource.SchemaRequest{}, resp)

	scope, ok := resp.Schema.Attributes["access_review_scope_v2"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("access_review_scope_v2 is %T, want schema.SingleNestedAttribute", resp.Schema.Attributes["access_review_scope_v2"])
	}
	assertObjectAttributeRequiresReplacement(t, ctx, "access_review_scope_v2", scope.PlanModifiers)

	policyID, ok := resp.Schema.Attributes["policy_id"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("policy_id is %T, want schema.StringAttribute", resp.Schema.Attributes["policy_id"])
	}
	assertStringAttributeRequiresReplacement(t, ctx, "policy_id", policyID.PlanModifiers)
}

func assertObjectAttributeRequiresReplacement(t *testing.T, ctx context.Context, name string, modifiers []planmodifier.Object) {
	t.Helper()
	for _, modifier := range modifiers {
		if strings.Contains(modifier.Description(ctx), replacementModifierDescription) {
			return
		}
	}
	t.Errorf("attribute %q has no replacement plan modifier", name)
}

func assertStringAttributeRequiresReplacement(t *testing.T, ctx context.Context, name string, modifiers []planmodifier.String) {
	t.Helper()
	for _, modifier := range modifiers {
		if strings.Contains(modifier.Description(ctx), replacementModifierDescription) {
			return
		}
	}
	t.Errorf("attribute %q has no replacement plan modifier", name)
}
