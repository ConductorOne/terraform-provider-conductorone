package provider

import (
	"strings"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type updateMaskField struct {
	terraformName string
	apiPath       string
	serialized    bool
}

// updateMaskForChanges builds a mask from fields which both changed in the
// Terraform plan and are present in the JSON payload. Update APIs reject an
// empty mask, while including an omitted field in a mask can overwrite it.
func updateMaskForChanges(state, plan types.Object, fields []updateMaskField) (*string, diag.Diagnostics) {
	var diags diag.Diagnostics
	stateAttributes := state.Attributes()
	planAttributes := plan.Attributes()
	paths := make([]string, 0, len(fields))

	for _, field := range fields {
		stateValue, stateOK := stateAttributes[field.terraformName]
		planValue, planOK := planAttributes[field.terraformName]
		if !stateOK || !planOK || planValue.Equal(stateValue) {
			continue
		}
		if field.serialized {
			paths = append(paths, field.apiPath)
		}
	}

	if len(paths) == 0 {
		diags.AddError(
			"Unable to build update mask",
			"The planned changes do not include a field that this provider can safely serialize for this update. No request was sent.",
		)
		return nil, diags
	}

	mask := strings.Join(paths, ",")
	return &mask, diags
}

func functionUpdateMask(input *shared.FunctionInput) (*string, diag.Diagnostics) {
	if input == nil {
		return updateMaskForSerializedFields(nil)
	}

	return updateMaskForSerializedFields([]updateMaskField{
		{apiPath: "displayName", serialized: input.DisplayName != nil},
		{apiPath: "description", serialized: input.Description != nil},
		{apiPath: "functionType", serialized: input.FunctionType != nil},
		{apiPath: "publishedCommitId", serialized: input.PublishedCommitID != nil},
		{apiPath: "isDraft", serialized: input.IsDraft != nil},
		{apiPath: "secret", serialized: len(input.Secret) > 0},
		{apiPath: "outboundNetworkAllowlist", serialized: len(input.OutboundNetworkAllowlist) > 0},
		{apiPath: "scopedRoleIds", serialized: len(input.ScopedRoleIds) > 0},
	})
}

func functionUpdateMaskForChanges(state, plan types.Object, input *shared.FunctionInput) (*string, diag.Diagnostics) {
	return updateMaskForChanges(state, plan, []updateMaskField{
		{terraformName: "display_name", apiPath: "displayName", serialized: input != nil && input.DisplayName != nil},
		{terraformName: "description", apiPath: "description", serialized: input != nil && input.Description != nil},
		{terraformName: "function_type", apiPath: "functionType", serialized: input != nil && input.FunctionType != nil},
		{terraformName: "published_commit_id", apiPath: "publishedCommitId", serialized: input != nil && input.PublishedCommitID != nil},
		{terraformName: "is_draft", apiPath: "isDraft", serialized: input != nil && input.IsDraft != nil},
		{terraformName: "secret", apiPath: "secret", serialized: input != nil && len(input.Secret) > 0},
		{terraformName: "outbound_network_allowlist", apiPath: "outboundNetworkAllowlist", serialized: input != nil && len(input.OutboundNetworkAllowlist) > 0},
		{terraformName: "scoped_role_ids", apiPath: "scopedRoleIds", serialized: input != nil && len(input.ScopedRoleIds) > 0},
	})
}

func accessReviewUpdateMask(input *shared.AccessReviewInput) (*string, diag.Diagnostics) {
	if input == nil {
		return updateMaskForSerializedFields(nil)
	}

	return updateMaskForSerializedFields([]updateMaskField{
		{apiPath: "displayName", serialized: input.DisplayName != nil},
		{apiPath: "description", serialized: input.Description != nil},
		{apiPath: "reviewInstructions", serialized: input.ReviewInstructions != nil},
		{apiPath: "defaultView", serialized: input.DefaultView != nil},
		{apiPath: "completionDate", serialized: input.CompletionDate != nil},
		{apiPath: "autoResolve", serialized: input.AutoResolve != nil},
		{apiPath: "usePolicyOverride", serialized: input.UsePolicyOverride != nil},
		{apiPath: "autoGenerateReport", serialized: input.AutoGenerateReport != nil},
		{apiPath: "scopeType", serialized: input.ScopeType != nil},
		{apiPath: "exemptCertifiedAccessConflicts", serialized: input.ExemptCertifiedAccessConflicts != nil},
		{apiPath: "autoStartCampaign", serialized: input.AutoStartCampaign != nil},
		{apiPath: "scheduledStartDate", serialized: input.ScheduledStartDate != nil},
		{apiPath: "accuracyIssueAction", serialized: input.AccuracyIssueAction != nil},
		{apiPath: "autoCloseCampaign", serialized: input.AutoCloseCampaign != nil},
		{apiPath: "autoCloseDecision", serialized: input.AutoCloseDecision != nil},
		{apiPath: "columnConfig", serialized: input.AccessReviewColumnConfig != nil},
		{apiPath: "notificationConfig", serialized: input.NotificationConfig != nil},
		{apiPath: "signatureConfig", serialized: input.ReviewSignatureConfig != nil},
		{apiPath: "inclusionScope", serialized: input.AccessReviewInclusionScope != nil},
	})
}

func accessReviewUpdateMaskForChanges(state, plan types.Object, input *shared.AccessReviewInput) (*string, diag.Diagnostics) {
	if input == nil {
		return updateMaskForChanges(state, plan, nil)
	}

	return updateMaskForChanges(state, plan, []updateMaskField{
		{terraformName: "display_name", apiPath: "displayName", serialized: input.DisplayName != nil},
		{terraformName: "description", apiPath: "description", serialized: input.Description != nil},
		{terraformName: "review_instructions", apiPath: "reviewInstructions", serialized: input.ReviewInstructions != nil},
		{terraformName: "default_view", apiPath: "defaultView", serialized: input.DefaultView != nil},
		{terraformName: "completion_date", apiPath: "completionDate", serialized: input.CompletionDate != nil},
		{terraformName: "auto_resolve", apiPath: "autoResolve", serialized: input.AutoResolve != nil},
		{terraformName: "use_policy_override", apiPath: "usePolicyOverride", serialized: input.UsePolicyOverride != nil},
		{terraformName: "auto_generate_report", apiPath: "autoGenerateReport", serialized: input.AutoGenerateReport != nil},
		{terraformName: "scope_type", apiPath: "scopeType", serialized: input.ScopeType != nil},
		{terraformName: "exempt_certified_access_conflicts", apiPath: "exemptCertifiedAccessConflicts", serialized: input.ExemptCertifiedAccessConflicts != nil},
		{terraformName: "auto_start_campaign", apiPath: "autoStartCampaign", serialized: input.AutoStartCampaign != nil},
		{terraformName: "scheduled_start_date", apiPath: "scheduledStartDate", serialized: input.ScheduledStartDate != nil},
		{terraformName: "accuracy_issue_action", apiPath: "accuracyIssueAction", serialized: input.AccuracyIssueAction != nil},
		{terraformName: "auto_close_campaign", apiPath: "autoCloseCampaign", serialized: input.AutoCloseCampaign != nil},
		{terraformName: "auto_close_decision", apiPath: "autoCloseDecision", serialized: input.AutoCloseDecision != nil},
		{terraformName: "access_review_column_config", apiPath: "columnConfig", serialized: input.AccessReviewColumnConfig != nil},
		{terraformName: "notification_config", apiPath: "notificationConfig", serialized: input.NotificationConfig != nil},
		{terraformName: "review_signature_config", apiPath: "signatureConfig", serialized: input.ReviewSignatureConfig != nil},
		{terraformName: "access_review_inclusion_scope", apiPath: "inclusionScope", serialized: input.AccessReviewInclusionScope != nil},
	})
}

func updateMaskForSerializedFields(fields []updateMaskField) (*string, diag.Diagnostics) {
	var diags diag.Diagnostics
	paths := make([]string, 0, len(fields))
	for _, field := range fields {
		if field.serialized {
			paths = append(paths, field.apiPath)
		}
	}
	if len(paths) == 0 {
		diags.AddError(
			"Unable to build update mask",
			"The update payload has no maskable fields. No request was sent.",
		)
		return nil, diags
	}
	mask := strings.Join(paths, ",")
	return &mask, diags
}
