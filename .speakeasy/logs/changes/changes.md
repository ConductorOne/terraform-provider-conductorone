## Terraform SDK Changes:
* `conductoroneAPI.AutomationSearch.SearchAutomationTemplateVersions()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `triggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `triggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `conductoroneAPI.AutomationSearch.SearchAutomations()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `conductoroneAPI.Automation.ListAutomations()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `conductoroneAPI.Automation.CreateAutomation()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
  * `response.Automation` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `conductoroneAPI.Automation.GetAutomation()`: `response.Automation` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `conductoroneAPI.Automation.UpdateAutomation()`: 
  * `request.UpdateAutomationRequest.Automation` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
  * `response.Automation` **Changed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword.passwordPolicyId` **Added**
    - `automationSteps[].GrantEntitlements.appEntitlementRefsCel` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements.appEntitlementRefs` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Added**
    - `draftTriggers[].FormTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger.GrantTriggerFilter.EntitlementFilter` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ManualAutomationTrigger` **Removed** (Breaking ⚠️)
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: 
  * `request.AccessReviewTemplateSetScopeByResourceTypeRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.Apps.Get()`: `response.App` **Changed**
    - `AppUserMapper` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.A2UI.CreateSurfaceFeedback()`: **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: **Added**
* `conductoroneAPI.Functions.Test()`: **Added**
* `conductoroneAPI.FunctionsInvocationSearch.Search()`: **Added**
* `conductoroneAPI.RoleMiningManagement.GetRoleMiningConfig()`: **Added**
* `conductoroneAPI.RoleMiningManagement.UpdateRoleMiningConfig()`: **Added**
* `conductoroneAPI.RoleMiningManagement.ListRuns()`: **Added**
* `conductoroneAPI.RoleMiningManagement.GetLatestRun()`: **Added**
* `conductoroneAPI.RoleMiningManagement.ListSuggestions()`: **Added**
* `conductoroneAPI.RoleMiningManagement.GetSuggestion()`: **Added**
* `conductoroneAPI.RoleMiningManagement.UpdateSuggestionState()`: **Added**
* `conductoroneAPI.RoleMiningManagement.TriggerAnalysis()`: **Added**
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
* `conductoroneAPI.OrgNotificationSettings.Get()`: **Added**
* `conductoroneAPI.OrgNotificationSettings.Update()`: **Added**
* `conductoroneAPI.UserNotificationSettings.Get()`: **Added**
* `conductoroneAPI.UserNotificationSettings.Update()`: **Added**
* `conductoroneAPI.AccessReview.Create()`: 
  * `request` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReview.Get()`: `response.AccessReviewView.AccessReview` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReview.Update()`: 
  * `request.AccessReviewServiceUpdateRequest.AccessReview` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
  * `response.AccessReviewView.AccessReview` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReview.List()`: `response.list[].AccessReview` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: 
  * `request.AccessReviewSetupEntitlementAndScopeServiceSetRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
  * `response.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: 
  * `request.AccessReviewSetScopeByResourceTypeRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.AccessReviewTemplate.Create()`: 
  * `request` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
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
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
    - `usePolicyOverride` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.AccessReviewTemplate.Get()`: `response.AccessReviewTemplate` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.Apps.Create()`: 
  *  `request.identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
  * `response.App` **Changed**
    - `AppUserMapper` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: 
  * `request.AccessReviewTemplateSetupEntitlementServiceSetRequest.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
  * `response.AccessReviewScopeV2` **Changed**
    - `AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `ResourceSelectionScope` **Added**
* `conductoroneAPI.A2UI.SubmitAction()`: **Added**
* `conductoroneAPI.Apps.List()`: `response.list[]` **Changed**
    - `AppUserMapper` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.AccessReviewTemplate.Update()`: 
  * `request.AccessReviewTemplateServiceUpdateRequest.AccessReviewTemplate` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
  * `response.AccessReviewTemplate` **Changed**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.complianceFrameworkAttributeValueIds` **Added**
    - `AccessReviewScopeV2.AppSelectionCriteriaScope.riskLevelAttributeValueIds` **Added**
    - `AccessReviewScopeV2.ResourceSelectionScope` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_RESOURCE)` **Added**
* `conductoroneAPI.A2UI.ListSurfaceFeedback()`: **Added**
* `conductoroneAPI.Directory.Create()`: 
  *  `request.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Connector.List()`:  `response.list[].Connector.connectorApiVersion` **Added**
* `conductoroneAPI.Connector.CreateDelegated()`:  `response.ConnectorView.Connector.connectorApiVersion` **Added**
* `conductoroneAPI.Connector.Get()`:  `response.ConnectorView.Connector.connectorApiVersion` **Added**
* `conductoroneAPI.Connector.Update()`:  `response.ConnectorView.Connector.connectorApiVersion` **Added**
* `conductoroneAPI.Connector.Create()`:  `response.ConnectorView.Connector.connectorApiVersion` **Added**
* `conductoroneAPI.Connector.UpdateDelegated()`:  `response.ConnectorView.Connector.connectorApiVersion` **Added**
* `conductoroneAPI.AppEntitlements.List()`:  `response.list[].AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlements.Create()`: 
  *  `request.CreateAppEntitlementRequest.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  *  `response.AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlements.AddManuallyManagedMembers()`:  `response.bulkActionId` **Added**
* `conductoroneAPI.AppEntitlements.ListAutomationExclusions()`:  `response.list[].User.origin` **Added**
* `conductoroneAPI.AppEntitlements.Get()`:  `response.AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlements.Update()`: 
  *  `request.UpdateAppEntitlementRequest.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  *  `response.AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlements.ListForAppResource()`:  `response.list[].AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlements.ListForAppUser()`:  `response.list[].AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`:  `response.list[].User.origin` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`:  `response.list[].AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.A2UI.ListSurfaces()`: **Added**
* `conductoroneAPI.AppEntitlementSearch.Search()`:  `response.list[].AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchGrants()`:  `response.list[].AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.AppEntitlementOwners.List()`:  `response.list[].origin` **Added**
* `conductoroneAPI.AppOwners.List()`:  `response.list[].origin` **Added**
* `conductoroneAPI.AppResourceOwners.List()`:  `response.list[].origin` **Added**
* `conductoroneAPI.RequestCatalogManagement.ListEntitlementsPerCatalog()`:  `response.list[].AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.RequestCatalogManagement.ListEntitlementsForAccess()`:  `response.list[].AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Directory.List()`:  `response.list[].Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Apps.Update()`: 
  * `request.UpdateAppRequest.App` **Changed**
    - `AppUserMapper` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
  * `response.App` **Changed**
    - `AppUserMapper` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.Directory.Get()`:  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Directory.Update()`: 
  *  `request.DirectoryServiceUpdateRequest.DirectoryMergeConfig` **Added**
  *  `response.DirectoryView.Directory.DirectoryMergeConfig` **Added**
* `conductoroneAPI.Policies.List()`:  `response.list[].policySteps.Map<PolicySteps>.steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Policies.Create()`: 
  *  `request.policySteps.Map<PolicySteps>.steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  *  `response.Policy.policySteps.Map<PolicySteps>.steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Policies.Get()`:  `response.Policy.policySteps.Map<PolicySteps>.steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.Policies.Update()`: 
  *  `request.UpdatePolicyRequest.Policy.policySteps.Map<PolicySteps>.steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
  *  `response.Policy.policySteps.Map<PolicySteps>.steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.RequestSchema.Create()`: `request` **Changed**
    - `fieldGroups` **Added**
    - `fieldRelationships` **Added**
    - `justificationVisibility` **Added**
* `conductoroneAPI.AppSearch.Search()`: `response.list[]` **Changed**
    - `AppUserMapper` **Added**
    - `appOwners[].origin` **Added**
    - `identityMatching.enum(APP_USER_IDENTITY_MATCHING_CUSTOM)` **Added**
* `conductoroneAPI.ExternalClientSearch.Search()`: 
  *  `request.clientIdUrls` **Added**
  * `response.list[]` **Changed**
    - `clientIdType` **Added**
    - `verifiedDomain` **Added**
* `conductoroneAPI.PolicySearch.Search()`:  `response.list[].policySteps.Map<PolicySteps>.steps[].Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.RequestCatalogSearch.SearchEntitlements()`:  `response.list[].AppEntitlementView.AppEntitlement.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
* `conductoroneAPI.TaskSearch.Search()`: 
  * `request` **Changed**
    - `excludeApplicationIds` **Added**
    - `requireApprovalReason` **Added**
    - `requireDenialReason` **Added**
  * `response.list[]` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
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
* `conductoroneAPI.TaskAudit.List()`: `response.list[]` **Changed**
    - `TaskAuditNewTaskCreatedFrom.originalTaskType` **Added**
    - `TaskAuditPolicyApprovalReassigned.users[].origin` **Added**
* `conductoroneAPI.Task.CreateGrantTask()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.Task.CreateOffboardingTask()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.Task.CreateRevokeTask()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.Task.Get()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.Approve()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.ApproveWithStepUp()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.Close()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.Comment()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.Deny()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.EscalateToEmergencyAccess()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.ProcessNow()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.Reassign()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.HardReset()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.Restart()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.SkipStep()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.UpdateGrantDuration()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
* `conductoroneAPI.TaskActions.UpdateRequestData()`: `response.TaskView` **Changed**
    - `Task.PolicyInstance.PolicyStepInstance.ProvisionInstance.Provision.ProvisionPolicy.ManualProvision.ProvisionerAssignment` **Added**
    - `Task.approverIds` **Added**
    - `Task.revocationTargets` **Added**
    - `approversPath` **Added**
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
