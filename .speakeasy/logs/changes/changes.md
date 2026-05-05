## Terraform SDK Changes:
* `conductoroneAPI.AccessReview.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `errorState` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReview.Get()`: `response.AccessReviewView.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `errorState` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReview.Update()`: 
  * `request.AccessReviewServiceUpdateRequest.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `errorState` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReview.List()`: `response.list[].AccessReview` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `CampaignHealthSnapshot` **Added**
    - `CampaignInsights` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `errorState` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReviewTemplate.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `RecurrenceRule` **Added**
    - `ReviewSignatureConfig` **Added**
    - `accuracyIssueAction` **Added**
    - `autoCloseCampaign` **Added**
    - `autoCloseDecision` **Added**
    - `autoGenerateReport` **Added**
    - `autoStartCampaign` **Added**
    - `defaultView` **Added**
    - `exemptCertifiedAccessConflicts` **Added**
    - `isCampaignScheduleEnabled` **Added**
    - `reviewInstructions` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
    - `usePolicyOverride` **Added**
  * `response.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.TaskActions.UpdateRequestData()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.UpdateGrantDuration()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.SkipStep()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.Restart()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.HardReset()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.Reassign()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.ProcessNow()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.Deny()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.Comment()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.Close()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.ApproveWithStepUp()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskActions.Approve()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.Task.Get()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.Task.CreateRevokeTask()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.Task.CreateOffboardingTask()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.Task.CreateGrantTask()`: `response.TaskView` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.TaskAudit.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `TaskAuditActionInstanceCreated.ActionInstance` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceCreated.PolicyActionInstance` **Added**
    - `TaskAuditActionInstanceFailed.ActionInstance` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceFailed.PolicyActionInstance` **Added**
    - `TaskAuditActionInstanceSucceeded.ActionInstance` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceSucceeded.PolicyActionInstance` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `TaskAuditConnectorActionResult.TaskAuditCancelledResult.cancelReason` **Added**
    - `TaskAuditConnectorActionResult.TaskAuditPendingResult` **Added**
    - `TaskAuditConnectorActionResult.TaskAuditSuccessResult.successReason` **Added**
    - `TaskAuditNewTaskCreatedFrom.originalTaskType` **Added**
    - `TaskAuditPolicyApprovalReassigned.users[].origin` **Added**
* `conductoroneAPI.TaskSearch.Search()`: 
  * `request` **Changed**
    - `excludeApplicationIds` **Added**
    - `requireApprovalReason` **Added**
    - `requireDenialReason` **Added**
    - `taskTypes[].TaskTypeAction.ScopeRole` **Added**
    - `taskTypes[].TaskTypeAction.TaskActionInstance` **Added**
    - `taskTypes[].TaskTypeFinding` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `Task.Form.fields[].StringField.TextField.suffix` **Added**
    - `Task.Form.fields[].StringMapField` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ActionInstance` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.CancelTicket` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceCancelTicket` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReassignToApprovers` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceReplacePolicy` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.EscalationInstanceSkipStep` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReassignToApprovers` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.ReplacePolicy` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.ApprovalInstance.EscalationInstance.SkipStep` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.PolicyStepInstance.PolicyActionInstance` **Added**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.PolicyInstance.next[].Action.ActionTargetClientIdApproval` **Added**
    - `Task.PolicyInstance.next[].Form` **Removed** (Breaking ⚠️)
    - `Task.PolicyInstance.next[].PolicyForm` **Added**
    - `Task.TaskType.TaskTypeAction.ScopeRole` **Added**
    - `Task.TaskType.TaskTypeAction.TaskActionInstance` **Added**
    - `Task.TaskType.TaskTypeAction.displayName` **Added**
    - `Task.TaskType.TaskTypeAction.type` **Added**
    - `Task.TaskType.TaskTypeFinding` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
    - `resourceBindingsPath` **Added**
    - `roleResourcePath` **Added**
    - `scopeResourcePath` **Added**
* `conductoroneAPI.PolicySearch.Search()`: `response.list[].policySteps.Map<PolicySteps>.steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetClientIdApproval` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Policies.Update()`: 
  * `request.UpdatePolicyRequest.Policy.policySteps.Map<PolicySteps>.steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetClientIdApproval` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.policySteps.Map<PolicySteps>.steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetClientIdApproval` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Policies.Get()`: `response.Policy.policySteps.Map<PolicySteps>.steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetClientIdApproval` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Policies.Create()`: 
  * `request.policySteps.Map<PolicySteps>.steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetClientIdApproval` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.Policy.policySteps.Map<PolicySteps>.steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetClientIdApproval` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Policies.List()`: `response.list[].policySteps.Map<PolicySteps>.steps[]` **Changed** (Breaking ⚠️)
    - `Action.ActionTargetClientIdApproval` **Added**
    - `Form` **Removed** (Breaking ⚠️)
    - `PolicyForm` **Added**
    - `Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.ConnectorCatalog.ConfigurationSchema()`: `response` **Changed** (Breaking ⚠️)
    - `ConfigSchema.fieldGroups[].fieldNames` **Added**
    - `ConfigSchema.fieldGroups[].fields` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].AdminProviderConfig` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].BoolField` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].ConnectorCheckboxField` **Added**
    - `ConfigSchema.fields[].ConnectorSelectField` **Added**
    - `ConfigSchema.fields[].ConnectorStringField` **Added**
    - `ConfigSchema.fields[].ConnectorStringMapField` **Added**
    - `ConfigSchema.fields[].ConnectorTextField` **Added**
    - `ConfigSchema.fields[].FileField` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].ImportField` **Added**
    - `ConfigSchema.fields[].Int64Field` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].KeyValueField` **Added**
    - `ConfigSchema.fields[].OAuth2Field` **Added**
    - `ConfigSchema.fields[].Oauth2Field` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].OptionsField` **Added**
    - `ConfigSchema.fields[].RandomStringField` **Added**
    - `ConfigSchema.fields[].ReadOnlyField` **Added**
    - `ConfigSchema.fields[].RotatableSecretField` **Added**
    - `ConfigSchema.fields[].SharedProviderConfig` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].StringField` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].StringListField` **Added**
    - `ConfigSchema.fields[].UserProviderConfig` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].additionalPlaceholder` **Added**
    - `ConfigSchema.fields[].dependsOnFields` **Added**
    - `ConfigSchema.fields[].description` **Removed** (Breaking ⚠️)
    - `ConfigSchema.fields[].helpUrl` **Added**
    - `ConfigSchema.fields[].placeholder` **Added**
    - `ConfigSchema.fields[].postCreate` **Added**
    - `ConfigSchema.fields[].required` **Removed** (Breaking ⚠️)
    - `Form.fields[].StringField.TextField.suffix` **Added**
    - `Form.fields[].StringMapField` **Added**
* `conductoroneAPI.Automation.UpdateAutomation()`: 
  * `request.UpdateAutomationRequest.Automation` **Changed** (Breaking ⚠️)
    - `automationSteps[].AutomationsTaskAction` **Added**
    - `automationSteps[].AutomationsWebhook` **Added**
    - `automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail.emailCel` **Added**
    - `automationSteps[].SendEmail.email` **Added**
    - `automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `automationSteps[].SendSlackMessage.userRefs` **Added**
    - `automationSteps[].SetCredential` **Added**
    - `automationSteps[].StoreCredential` **Added**
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Added**
    - `draftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
    - `primaryTriggerType.enum(TRIGGER_TYPE_SCHEDULE_NO_USER)` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation.automationSteps[].AutomationsTaskAction` **Added**
    - `Automation.automationSteps[].AutomationsWebhook` **Added**
    - `Automation.automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `Automation.automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `Automation.automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `Automation.automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `Automation.automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `Automation.automationSteps[].SendEmail.emailCel` **Added**
    - `Automation.automationSteps[].SendEmail.email` **Added**
    - `Automation.automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `Automation.automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `Automation.automationSteps[].SendSlackMessage.userRefs` **Added**
    - `Automation.automationSteps[].SetCredential` **Added**
    - `Automation.automationSteps[].StoreCredential` **Added**
    - `Automation.automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `Automation.automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].ScheduleTriggerNoUser` **Added**
    - `Automation.draftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
    - `Automation.primaryTriggerType.enum(TRIGGER_TYPE_SCHEDULE_NO_USER)` **Added**
    - `webhookCapabilityUrl` **Added**
* `conductoroneAPI.Automation.GetAutomation()`: `response.Automation` **Changed** (Breaking ⚠️)
    - `automationSteps[].AutomationsTaskAction` **Added**
    - `automationSteps[].AutomationsWebhook` **Added**
    - `automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail.emailCel` **Added**
    - `automationSteps[].SendEmail.email` **Added**
    - `automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `automationSteps[].SendSlackMessage.userRefs` **Added**
    - `automationSteps[].SetCredential` **Added**
    - `automationSteps[].StoreCredential` **Added**
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Added**
    - `draftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
    - `primaryTriggerType.enum(TRIGGER_TYPE_SCHEDULE_NO_USER)` **Added**
* `conductoroneAPI.Automation.DeleteAutomation()`: `request` **Changed** (Breaking ⚠️)
    - `AutomationsDeleteAutomationRequest` **Added**
    - `DeleteAutomationRequest` **Removed** (Breaking ⚠️)
* `conductoroneAPI.Automation.CreateAutomation()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `automationSteps[].AutomationsTaskAction` **Added**
    - `automationSteps[].AutomationsWebhook` **Added**
    - `automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail.emailCel` **Added**
    - `automationSteps[].SendEmail.email` **Added**
    - `automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `automationSteps[].SendSlackMessage.userRefs` **Added**
    - `automationSteps[].SetCredential` **Added**
    - `automationSteps[].StoreCredential` **Added**
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Added**
    - `draftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation.automationSteps[].AutomationsTaskAction` **Added**
    - `Automation.automationSteps[].AutomationsWebhook` **Added**
    - `Automation.automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `Automation.automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `Automation.automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `Automation.automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `Automation.automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `Automation.automationSteps[].SendEmail.emailCel` **Added**
    - `Automation.automationSteps[].SendEmail.email` **Added**
    - `Automation.automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `Automation.automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `Automation.automationSteps[].SendSlackMessage.userRefs` **Added**
    - `Automation.automationSteps[].SetCredential` **Added**
    - `Automation.automationSteps[].StoreCredential` **Added**
    - `Automation.automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `Automation.automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `Automation.draftTriggers[].ScheduleTriggerNoUser` **Added**
    - `Automation.draftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
    - `Automation.primaryTriggerType.enum(TRIGGER_TYPE_SCHEDULE_NO_USER)` **Added**
    - `webhookCapabilityUrl` **Added**
* `conductoroneAPI.Automation.ListAutomations()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `automationSteps[].AutomationsTaskAction` **Added**
    - `automationSteps[].AutomationsWebhook` **Added**
    - `automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail.emailCel` **Added**
    - `automationSteps[].SendEmail.email` **Added**
    - `automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `automationSteps[].SendSlackMessage.userRefs` **Added**
    - `automationSteps[].SetCredential` **Added**
    - `automationSteps[].StoreCredential` **Added**
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Added**
    - `draftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
    - `primaryTriggerType.enum(TRIGGER_TYPE_SCHEDULE_NO_USER)` **Added**
* `conductoroneAPI.AutomationSearch.SearchAutomations()`: 
  * `request` **Changed**
    - `direction` **Added**
    - `sortField` **Added**
    - `triggerTypes[].enum(TRIGGER_TYPE_SCHEDULE_NO_USER)` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `automationSteps[].AutomationsTaskAction` **Added**
    - `automationSteps[].AutomationsWebhook` **Added**
    - `automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail.emailCel` **Added**
    - `automationSteps[].SendEmail.email` **Added**
    - `automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `automationSteps[].SendSlackMessage.userRefs` **Added**
    - `automationSteps[].SetCredential` **Added**
    - `automationSteps[].StoreCredential` **Added**
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Added**
    - `draftTriggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
    - `primaryTriggerType.enum(TRIGGER_TYPE_SCHEDULE_NO_USER)` **Added**
* `conductoroneAPI.AutomationSearch.SearchAutomationTemplateVersions()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `automationSteps[].AutomationsTaskAction` **Added**
    - `automationSteps[].AutomationsWebhook` **Added**
    - `automationSteps[].ConnectorCreateAccount.passwordCel` **Added**
    - `automationSteps[].GeneratePassword.GeneratePasswordPolicy` **Added**
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail.emailCel` **Added**
    - `automationSteps[].SendEmail.email` **Added**
    - `automationSteps[].SendSlackMessage.useSubjectUser` **Added**
    - `automationSteps[].SendSlackMessage.userIdsCel` **Added**
    - `automationSteps[].SendSlackMessage.userRefs` **Added**
    - `automationSteps[].SetCredential` **Added**
    - `automationSteps[].StoreCredential` **Added**
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `triggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `triggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].ScheduleTriggerNoUser` **Added**
    - `triggers[].WebhookAutomationTrigger.WebhookListenerAuthCapabilityURL` **Added**
* `conductoroneAPI.AppEntitlements.CreateAutomation()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AppCreateAutomationRequest` **Added**
    - `CreateAutomationRequest` **Removed** (Breaking ⚠️)
  *  `response.AppEntitlementAutomation.managedByRequestCatalogId` **Added**
* `conductoroneAPI.AppEntitlements.DeleteAutomation()`: `request` **Changed** (Breaking ⚠️)
    - `AppDeleteAutomationRequest` **Added**
    - `DeleteAutomationRequest` **Removed** (Breaking ⚠️)
* `conductoroneAPI.AccessConflict.UpdateMonitor()`: 
  * `request.ConflictMonitorUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
* `conductoroneAPI.AccessConflict.GetMonitor()`: `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
* `conductoroneAPI.AccessConflict.CreateMonitor()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
* `conductoroneAPI.AccessReviewTemplate.Update()`: 
  * `request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
  * `response.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Added**
    - `AccessReviewNotificationConfig` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `defaultView.enum(ACCESS_REVIEW_VIEW_TYPE_BY_RESOURCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_INHERITANCE)` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.SSFReceiverStream.Update()`: **Added**
* `conductoroneAPI.Connector.List()`: `response.list[].Connector` **Changed**
    - `configUpdatedAt` **Added**
    - `connectorApiVersion` **Added**
    - `parallelSyncWorkerCount` **Added**
* `conductoroneAPI.LocalUserInvitation.Search()`: **Added**
* `conductoroneAPI.RoleMiningManagement.CreateAccessProfileFromCohort()`: **Added**
* `conductoroneAPI.RoleMiningManagement.GetRoleMiningConfig()`: **Added**
* `conductoroneAPI.RoleMiningManagement.UpdateRoleMiningConfig()`: **Added**
* `conductoroneAPI.RoleMiningManagement.GetCustomAnalysisResult()`: **Added**
* `conductoroneAPI.RoleMiningManagement.TriggerCustomAnalysis()`: **Added**
* `conductoroneAPI.RoleMiningManagement.ListRuns()`: **Added**
* `conductoroneAPI.RoleMiningManagement.GetLatestRun()`: **Added**
* `conductoroneAPI.RoleMiningManagement.ListSuggestions()`: **Added**
* `conductoroneAPI.RoleMiningManagement.GetSuggestion()`: **Added**
* `conductoroneAPI.RoleMiningManagement.UpdateSuggestionState()`: **Added**
* `conductoroneAPI.RoleMiningManagement.SearchCohortUsers()`: **Added**
* `conductoroneAPI.RoleMiningManagement.TriggerAnalysis()`: **Added**
* `conductoroneAPI.AutomationExecutionSearch.SearchAllAutomationExecutions()`: **Added**
* `conductoroneAPI.AppSearch.SearchUserOwnership()`: **Added**
* `conductoroneAPI.HooksSearch.Search()`: **Added**
* `conductoroneAPI.RoleMiningManagementSearch.Search()`: **Added**
* `conductoroneAPI.PaperSecretAdmin.Search()`: **Added**
* `conductoroneAPI.PaperSecretAdmin.SearchAuditEvents()`: **Added**
* `conductoroneAPI.PaperSecretAdmin.Revoke()`: **Added**
* `conductoroneAPI.PaperSecretAdmin.Get()`: **Added**
* `conductoroneAPI.PaperSecret.SearchAuditEvents()`: **Added**
* `conductoroneAPI.PaperSecret.SearchMySecrets()`: **Added**
* `conductoroneAPI.PaperSecret.Revoke()`: **Added**
* `conductoroneAPI.PaperSecret.Get()`: **Added**
* `conductoroneAPI.PaperSecret.SetTextContent()`: **Added**
* `conductoroneAPI.PaperSecret.GetContent()`: **Added**
* `conductoroneAPI.PaperSecret.GetByShareCode()`: **Added**
* `conductoroneAPI.PaperSecret.CreateExternal()`: **Added**
* `conductoroneAPI.PaperSecret.CreateInternal()`: **Added**
* `conductoroneAPI.SSFReceiverEventSearch.Search()`: **Added**
* `conductoroneAPI.WorkloadFederation.SearchTrusts()`: **Added**
* `conductoroneAPI.WorkloadFederation.ListTrusts()`: **Added**
* `conductoroneAPI.WorkloadFederation.CreateTrust()`: **Added**
* `conductoroneAPI.WorkloadFederation.DeleteTrust()`: **Added**
* `conductoroneAPI.WorkloadFederation.GetTrust()`: **Added**
* `conductoroneAPI.WorkloadFederation.UpdateTrust()`: **Added**
* `conductoroneAPI.WorkloadFederation.TestToken()`: **Added**
* `conductoroneAPI.WorkloadFederation.ListProviders()`: **Added**
* `conductoroneAPI.WorkloadFederation.CreateProvider()`: **Added**
* `conductoroneAPI.WorkloadFederation.DeleteProvider()`: **Added**
* `conductoroneAPI.WorkloadFederation.GetProvider()`: **Added**
* `conductoroneAPI.WorkloadFederation.UpdateProvider()`: **Added**
* `conductoroneAPI.WorkloadFederation.TestCEL()`: **Added**
* `conductoroneAPI.Principal.List()`: **Added**
* `conductoroneAPI.Principal.Create()`: **Added**
* `conductoroneAPI.Principal.Delete()`: **Added**
* `conductoroneAPI.Principal.Get()`: **Added**
* `conductoroneAPI.Principal.Update()`: **Added**
* `conductoroneAPI.Principal.ListCredentials()`: **Added**
* `conductoroneAPI.Principal.CreateCredential()`: **Added**
* `conductoroneAPI.Principal.RevokeCredential()`: **Added**
* `conductoroneAPI.Principal.GetCredential()`: **Added**
* `conductoroneAPI.Principal.UpdateCredential()`: **Added**
* `conductoroneAPI.Principal.AddBinding()`: **Added**
* `conductoroneAPI.Principal.DeleteBinding()`: **Added**
* `conductoroneAPI.Principal.ListBindings()`: **Added**
* `conductoroneAPI.Contacts.GetContacts()`: **Added**
* `conductoroneAPI.Contacts.UpdateContacts()`: **Added**
* `conductoroneAPI.TenantEmailProvider.GetEmailCapabilities()`: **Added**
* `conductoroneAPI.TenantEmailProvider.Get()`: **Added**
* `conductoroneAPI.TenantEmailProvider.Update()`: **Added**
* `conductoroneAPI.TenantEmailProvider.SearchAuditEvents()`: **Added**
* `conductoroneAPI.TenantEmailProvider.Test()`: **Added**
* `conductoroneAPI.OrgNotificationSettings.Get()`: **Added**
* `conductoroneAPI.OrgNotificationSettings.Update()`: **Added**
* `conductoroneAPI.UserNotificationSettings.Get()`: **Added**
* `conductoroneAPI.UserNotificationSettings.Update()`: **Added**
* `conductoroneAPI.OnboardingSettings.Get()`: **Added**
* `conductoroneAPI.OnboardingSettings.Update()`: **Added**
* `conductoroneAPI.SSFReceiverStream.List()`: **Added**
* `conductoroneAPI.SSFReceiverStream.Create()`: **Added**
* `conductoroneAPI.SSFReceiverStream.Delete()`: **Added**
* `conductoroneAPI.SSFReceiverStream.Get()`: **Added**
* `conductoroneAPI.LocalUserInvitation.Get()`: **Added**
* `conductoroneAPI.SSFReceiverStream.Test()`: **Added**
* `conductoroneAPI.SSFReceiverStream.GetStats()`: **Added**
* `conductoroneAPI.SSFReceiverEvent.List()`: **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.Set()`: **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.SearchEntitlementOwners()`: **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.SearchUserOwners()`: **Added**
* `conductoroneAPI.AppOwnersV2.Set()`: **Added**
* `conductoroneAPI.AppOwnersV2.SearchEntitlementOwners()`: **Added**
* `conductoroneAPI.AppOwnersV2.SearchUserOwners()`: **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: 
  * `request.AccessReviewSetupEntitlementAndScopeServiceSetRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
  * `response.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.LocalUserInvitation.Create()`: **Added**
* `conductoroneAPI.LocalDirectoryConfig.Update()`: **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: 
  * `request.AccessReviewSetScopeByResourceTypeRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: 
  * `request.AccessReviewTemplateSetupEntitlementServiceSetRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
  * `response.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: 
  * `request.AccessReviewTemplateSetScopeByResourceTypeRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.LocalDirectoryConfig.Get()`: **Added**
* `conductoroneAPI.LocalDirectoryConfig.Delete()`: **Added**
* `conductoroneAPI.LocalDirectoryConfig.Create()`: **Added**
* `conductoroneAPI.Apps.List()`: `response.list[]` **Changed**
    - `AppUserMapper` **Added**
    - `accessModel` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.Apps.Create()`: 
  * `request` **Changed**
    - `appEntitlementOwnerRefs` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
  * `response.App` **Changed**
    - `AppUserMapper` **Added**
    - `accessModel` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.Apps.Get()`: `response.App` **Changed**
    - `AppUserMapper` **Added**
    - `accessModel` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.Apps.Update()`: 
  * `request.UpdateAppRequest.App` **Changed**
    - `AppUserMapper` **Added**
    - `accessModel` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
  * `response.App` **Changed**
    - `AppUserMapper` **Added**
    - `accessModel` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.LocalUserInvitation.Revoke()`: **Added**
* `conductoroneAPI.Connector.CreateDelegated()`: 
  *  `request.ConnectorServiceCreateDelegatedRequest.appEntitlementOwnerRefs` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `configUpdatedAt` **Added**
    - `connectorApiVersion` **Added**
    - `parallelSyncWorkerCount` **Added**
* `conductoroneAPI.Connector.Get()`: `response.ConnectorView.Connector` **Changed**
    - `configUpdatedAt` **Added**
    - `connectorApiVersion` **Added**
    - `parallelSyncWorkerCount` **Added**
* `conductoroneAPI.Connector.Update()`: 
  *  `request.ConnectorServiceUpdateRequest.Connector.parallelSyncWorkerCount` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `configUpdatedAt` **Added**
    - `connectorApiVersion` **Added**
    - `parallelSyncWorkerCount` **Added**
* `conductoroneAPI.Connector.Create()`: `response.ConnectorView.Connector` **Changed**
    - `configUpdatedAt` **Added**
    - `connectorApiVersion` **Added**
    - `parallelSyncWorkerCount` **Added**
* `conductoroneAPI.Connector.UpdateDelegated()`: 
  *  `request.ConnectorServiceUpdateDelegatedRequest.Connector.parallelSyncWorkerCount` **Added**
  * `response.ConnectorView.Connector` **Changed**
    - `configUpdatedAt` **Added**
    - `connectorApiVersion` **Added**
    - `parallelSyncWorkerCount` **Added**
* `conductoroneAPI.AppEntitlements.List()`: `response.list[].AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlements.Create()`: 
  *  `request.CreateAppEntitlementRequest.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.AppEntitlementView.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlements.AddManuallyManagedMembers()`:  `response.bulkActionId` **Added**
* `conductoroneAPI.LocalDirectoryConfig.List()`: **Added**
* `conductoroneAPI.AppEntitlements.GetAutomation()`:  `response.AppEntitlementAutomation.managedByRequestCatalogId` **Added**
* `conductoroneAPI.Hooks.Update()`: **Added**
* `conductoroneAPI.AppEntitlements.ListAutomationExclusions()`:  `response.list[].User.origin` **Added**
* `conductoroneAPI.AppEntitlements.UpdateAutomation()`:  `response.AppEntitlementAutomation.managedByRequestCatalogId` **Added**
* `conductoroneAPI.AppEntitlements.Get()`: `response.AppEntitlementView.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlements.Update()`: 
  *  `request.UpdateAppEntitlementRequest.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  * `response.AppEntitlementView.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlements.ListForAppResource()`: `response.list[].AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlements.ListForAppUser()`: `response.list[].AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`:  `response.list[].User.origin` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`: `response.list[].AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlementSearch.Search()`: `response.list[].AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchGrants()`: 
  *  `request.entitlementSlugs` **Added**
  * `response.list[].AppEntitlementView.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.AppEntitlementOwners.List()`:  `response.list[].origin` **Added**
* `conductoroneAPI.AppOwners.List()`:  `response.list[].origin` **Added**
* `conductoroneAPI.AppResource.List()`:  `response.list[].AppResource.externalId` **Added**
* `conductoroneAPI.AppResource.CreateManuallyManagedAppResource()`:  `response.AppResource.externalId` **Added**
* `conductoroneAPI.AppResource.Get()`:  `response.AppResourceView.AppResource.externalId` **Added**
* `conductoroneAPI.AppResource.Update()`:  `response.AppResourceView.AppResource.externalId` **Added**
* `conductoroneAPI.AppResourceOwners.List()`:  `response.list[].origin` **Added**
* `conductoroneAPI.Hooks.Get()`: **Added**
* `conductoroneAPI.Hooks.Delete()`: **Added**
* `conductoroneAPI.Hooks.Create()`: **Added**
* `conductoroneAPI.Hooks.List()`: **Added**
* `conductoroneAPI.FunctionsInvocationSearch.Search()`: **Added**
* `conductoroneAPI.Functions.Test()`: **Added**
* `conductoroneAPI.Functions.GetCommitContent()`: **Added**
* `conductoroneAPI.RequestCatalogManagement.List()`:  `response.list[].RequestCatalog.grantPolicyId` **Added**
* `conductoroneAPI.RequestCatalogManagement.Create()`: 
  *  `request.grantPolicyId` **Added**
  *  `response.RequestCatalogView.RequestCatalog.grantPolicyId` **Added**
* `conductoroneAPI.RequestCatalogManagement.ListEntitlementsPerCatalog()`: `response.list[].AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.RequestCatalogManagement.ListEntitlementsForAccess()`: `response.list[].AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.RequestCatalogManagement.Get()`:  `response.RequestCatalogView.RequestCatalog.grantPolicyId` **Added**
* `conductoroneAPI.RequestCatalogManagement.Update()`: 
  *  `request.RequestCatalogManagementServiceUpdateRequest.RequestCatalog.grantPolicyId` **Added**
  *  `response.RequestCatalogView.RequestCatalog.grantPolicyId` **Added**
* `conductoroneAPI.RequestCatalogManagement.GetBundleAutomation()`: `response` **Changed**
    - `BundleAutomationCircuitBreaker.state.enum(CIRCUIT_BREAKER_STATE_SUPPORT_DISABLED)` **Added**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCEL` **Added**
* `conductoroneAPI.RequestCatalogManagement.SetBundleAutomation()`: 
  *  `request.SetBundleAutomationRequest.BundleAutomationRuleCEL` **Added**
  * `response` **Changed**
    - `BundleAutomationCircuitBreaker.state.enum(CIRCUIT_BREAKER_STATE_SUPPORT_DISABLED)` **Added**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCEL` **Added**
* `conductoroneAPI.RequestCatalogManagement.CreateBundleAutomation()`: 
  *  `request.CreateBundleAutomationRequest.BundleAutomationRuleCEL` **Added**
  * `response` **Changed**
    - `BundleAutomationCircuitBreaker.state.enum(CIRCUIT_BREAKER_STATE_SUPPORT_DISABLED)` **Added**
    - `BundleAutomationLastRunState.BundleAutomationCelEvaluationState` **Added**
    - `BundleAutomationRuleCEL` **Added**
* `conductoroneAPI.Functions.GetLockFile()`: **Added**
* `conductoroneAPI.Directory.List()`:  `response.list[].Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Directory.Create()`: 
  *  `request.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Directory.Get()`:  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Directory.Update()`: 
  *  `request.DirectoryServiceUpdateRequest.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Functions.ListFunctions()`: `response.list[]` **Changed**
    - `functionType.enum(FUNCTION_TYPE_CODE_MODE)` **Added**
    - `useSpn` **Added**
* `conductoroneAPI.Functions.CreateFunction()`: 
  *  `request.functionType.enum(FUNCTION_TYPE_CODE_MODE)` **Added**
  * `response.Function` **Changed**
    - `functionType.enum(FUNCTION_TYPE_CODE_MODE)` **Added**
    - `useSpn` **Added**
* `conductoroneAPI.Functions.Invoke()`: 
  *  `request.FunctionsServiceInvokeRequest.vfsId` **Added**
* `conductoroneAPI.Functions.GetFunction()`: `response.Function` **Changed**
    - `functionType.enum(FUNCTION_TYPE_CODE_MODE)` **Added**
    - `useSpn` **Added**
* `conductoroneAPI.Functions.UpdateFunction()`: 
  *  `request.Function.functionType.enum(FUNCTION_TYPE_CODE_MODE)` **Added**
  * `response.Function` **Changed**
    - `functionType.enum(FUNCTION_TYPE_CODE_MODE)` **Added**
    - `useSpn` **Added**
* `conductoroneAPI.Functions.CreateFinalCommit()`: **Added**
* `conductoroneAPI.Functions.CreateInitialCommit()`: **Added**
* `conductoroneAPI.FindingSearch.Search()`: **Added**
* `conductoroneAPI.FindingRoutingRule.UpdateFindingRoutingRule()`: **Added**
* `conductoroneAPI.RequestSchema.Create()`: 
  * `request` **Changed**
    - `fieldGroups` **Added**
    - `fieldRelationships` **Added**
    - `fields[].StringField.TextField.suffix` **Added**
    - `fields[].StringMapField` **Added**
    - `justificationVisibility` **Added**
  * `response.RequestSchema.Form.fields[]` **Changed**
    - `StringField.TextField.suffix` **Added**
    - `StringMapField` **Added**
* `conductoroneAPI.RequestSchema.Get()`: `response.RequestSchema.Form.fields[]` **Changed**
    - `StringField.TextField.suffix` **Added**
    - `StringMapField` **Added**
* `conductoroneAPI.RequestSchema.Update()`: 
  * `request.RequestSchemaServiceUpdateRequest.RequestSchema.Form.fields[]` **Changed**
    - `StringField.TextField.suffix` **Added**
    - `StringMapField` **Added**
  * `response.RequestSchema.Form.fields[]` **Changed**
    - `StringField.TextField.suffix` **Added**
    - `StringMapField` **Added**
* `conductoroneAPI.AppResourceSearch.SearchAppResources()`:  `response.list[].AppResource.externalId` **Added**
* `conductoroneAPI.AppSearch.Search()`: `response.list[]` **Changed**
    - `AppUserMapper` **Added**
    - `accessModel` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.FunctionsSearch.Search()`: 
  *  `request.functionTypes[].enum(FUNCTION_TYPE_CODE_MODE)` **Added**
  * `response.list[]` **Changed**
    - `functionType.enum(FUNCTION_TYPE_CODE_MODE)` **Added**
    - `useSpn` **Added**
* `conductoroneAPI.ExternalClientSearch.Search()`: 
  *  `request.clientIdUrls` **Added**
  * `response.list[]` **Changed**
    - `clientIdType` **Added**
    - `clientIdUrl` **Added**
    - `mcpClientId` **Added**
    - `verifiedDomain` **Added**
* `conductoroneAPI.FindingRoutingRule.GetFindingRoutingRule()`: **Added**
* `conductoroneAPI.RequestCatalogSearch.SearchEntitlements()`: `response.list[].AppEntitlementView.AppEntitlement` **Changed**
    - `ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `externalId` **Added**
* `conductoroneAPI.FindingRoutingRule.DeleteFindingRoutingRule()`: **Added**
* `conductoroneAPI.UserSearch.Search()`: 
  * `request` **Changed**
    - `delegateStatus` **Added**
    - `delegatedUserIds` **Added**
    - `excludeOrigins` **Added**
    - `isDelegate` **Added**
    - `origins` **Added**
  *  `response.list[].User.origin` **Added**
* `conductoroneAPI.WebhooksSearch.Search()`:  `response.list[].callbackTimeout` **Added**
* `conductoroneAPI.SessionSettings.Get()`: `response.SessionSettings` **Changed**
    - `clientIdApprovalRequestPolicyId` **Added**
    - `clientIdMetadataDocumentPolicy` **Added**
* `conductoroneAPI.SessionSettings.Update()`: 
  * `request.SessionSettings` **Changed**
    - `clientIdApprovalRequestPolicyId` **Added**
    - `clientIdMetadataDocumentPolicy` **Added**
  * `response.SessionSettings` **Changed**
    - `clientIdApprovalRequestPolicyId` **Added**
    - `clientIdMetadataDocumentPolicy` **Added**
* `conductoroneAPI.SystemLog.ListEvents()`: 
  *  `request.untilEventUid` **Added**
* `conductoroneAPI.FindingRoutingRule.CreateFindingRoutingRule()`: **Added**
* `conductoroneAPI.FindingRoutingRule.ListFindingRoutingRules()`: **Added**
* `conductoroneAPI.Finding.BulkCreateFindingTasks()`: **Added**
* `conductoroneAPI.Finding.BulkUpdateFindingState()`: **Added**
* `conductoroneAPI.Finding.GetFinding()`: **Added**
* `conductoroneAPI.Finding.CreateFindingTask()`: **Added**
* `conductoroneAPI.Finding.UpdateFindingState()`: **Added**
* `conductoroneAPI.TenantAuthConfig.Update()`: **Added**
* `conductoroneAPI.TenantAuthConfig.Get()`: **Added**
* `conductoroneAPI.TenantAuthConfig.Delete()`: **Added**
* `conductoroneAPI.TenantAuthConfig.Create()`: **Added**
* `conductoroneAPI.TenantAuthConfig.List()`: **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: **Added**
* `conductoroneAPI.A2UI.CreateSurfaceFeedback()`: **Added**
* `conductoroneAPI.A2UI.ListSurfaceFeedback()`: **Added**
* `conductoroneAPI.A2UI.SubmitAction()`: **Added**
* `conductoroneAPI.A2UI.ListSurfaces()`: **Added**
* `conductoroneAPI.User.List()`:  `response.list[].User.origin` **Added**
* `conductoroneAPI.User.Get()`:  `response.UserView.User.origin` **Added**
* `conductoroneAPI.Webhooks.List()`:  `response.list[].callbackTimeout` **Added**
* `conductoroneAPI.Webhooks.Create()`: 
  *  `request.callbackTimeout` **Added**
  *  `response.Webhook.callbackTimeout` **Added**
* `conductoroneAPI.Webhooks.Get()`:  `response.Webhook.callbackTimeout` **Added**
* `conductoroneAPI.Webhooks.Update()`: 
  *  `request.WebhooksServiceUpdateRequest.Webhook.callbackTimeout` **Added**
  *  `response.Webhook.callbackTimeout` **Added**
