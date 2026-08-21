## Terraform SDK Changes:
* `conductoroneAPI.PaperSecretAdmin.Revoke()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `secret` **Added**
* `conductoroneAPI.RequestSchema.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `fieldRelationships[].AtLeastOne` **Removed** (Breaking ⚠️)
    - `fieldRelationships[].DependentOn` **Removed** (Breaking ⚠️)
    - `fieldRelationships[].MutuallyExclusive` **Removed** (Breaking ⚠️)
    - `fieldRelationships[].RequiredTogether` **Removed** (Breaking ⚠️)
    - `fieldRelationships[].atLeastOne` **Added**
    - `fieldRelationships[].dependentOn` **Added**
    - `fieldRelationships[].mutuallyExclusive` **Added**
    - `fieldRelationships[].requiredTogether` **Added**
    - `fields[].AdminProviderConfig` **Removed** (Breaking ⚠️)
    - `fields[].BoolField` **Removed** (Breaking ⚠️)
    - `fields[].FileField` **Removed** (Breaking ⚠️)
    - `fields[].FormStringField` **Removed** (Breaking ⚠️)
    - `fields[].FormStringMapField` **Removed** (Breaking ⚠️)
    - `fields[].Int64Field` **Removed** (Breaking ⚠️)
    - `fields[].Oauth2Field` **Removed** (Breaking ⚠️)
    - `fields[].SharedProviderConfig` **Removed** (Breaking ⚠️)
    - `fields[].UserProviderConfig` **Removed** (Breaking ⚠️)
    - `fields[].adminConfig` **Added**
    - `fields[].boolField` **Added**
    - `fields[].fileField` **Added**
    - `fields[].int64Field` **Added**
    - `fields[].oauth2Field` **Added**
    - `fields[].sharedConfig` **Added**
    - `fields[].stringField` **Added**
    - `fields[].stringMapField` **Added**
    - `fields[].userConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Removed** (Breaking ⚠️)
    - `requestSchema` **Added**
* `conductoroneAPI.AppResourceOwnersV2.CreateUserOwner()`: 
  * `request.CreateAppResourceUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceOwnerUser` **Removed** (Breaking ⚠️)
    - `appResourceOwnerUser` **Added**
* `conductoroneAPI.AppResourceOwnersV2.DeleteUserOwner()`: 
  * `request.DeleteAppResourceUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
  *  `response.status[404]` **Added** (Breaking ⚠️)
* `conductoroneAPI.AppResourceOwnersV2.SearchUserOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `user` **Added**
* `conductoroneAPI.AppResourceOwnersV2.CreateEntitlementOwner()`: 
  * `request.CreateAppResourceEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceOwnerEntitlement` **Removed** (Breaking ⚠️)
    - `appResourceOwnerEntitlement` **Added**
* `conductoroneAPI.AppResourceOwnersV2.DeleteEntitlementOwner()`: 
  * `request.DeleteAppResourceEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
  *  `response.status[404]` **Added** (Breaking ⚠️)
* `conductoroneAPI.AppResourceOwnersV2.SearchEntitlementOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppOwnersV2.CreateUserOwner()`: 
  * `request.CreateAppUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppOwnerUser` **Removed** (Breaking ⚠️)
    - `appOwnerUser` **Added**
* `conductoroneAPI.AppOwnersV2.GetUserOwner()`: `response` **Changed** (Breaking ⚠️)
    - `AppOwnerUser` **Removed** (Breaking ⚠️)
    - `appOwnerUser` **Added**
* `conductoroneAPI.AppOwnersV2.DeleteUserOwner()`: 
  * `request.DeleteAppUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
* `conductoroneAPI.AppOwnersV2.SearchUserOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `user` **Added**
* `conductoroneAPI.AppOwnersV2.CreateEntitlementOwner()`: 
  * `request.CreateAppEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppOwnerEntitlement` **Removed** (Breaking ⚠️)
    - `appOwnerEntitlement` **Added**
* `conductoroneAPI.AppOwnersV2.GetEntitlementOwner()`: `response` **Changed** (Breaking ⚠️)
    - `AppOwnerEntitlement` **Removed** (Breaking ⚠️)
    - `appOwnerEntitlement` **Added**
* `conductoroneAPI.AppOwnersV2.DeleteEntitlementOwner()`: 
  * `request.DeleteAppEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
* `conductoroneAPI.AppOwnersV2.SearchEntitlementOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppEntitlementOwnersV2.CreateUserOwner()`: 
  * `request.CreateAppEntitlementUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementOwnerUser` **Removed** (Breaking ⚠️)
    - `appEntitlementOwnerUser` **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.GetUserOwner()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementOwnerUser` **Removed** (Breaking ⚠️)
    - `appEntitlementOwnerUser` **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.DeleteUserOwner()`: 
  * `request.DeleteAppEntitlementUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.SearchUserOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `user` **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.CreateEntitlementOwner()`: 
  * `request.CreateAppEntitlementEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementOwnerEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlementOwnerEntitlement` **Added**
* `conductoroneAPI.AppOwners.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `departmentSources[].priority` **Added**
    - `profile` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppEntitlementOwnersV2.GetEntitlementOwner()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementOwnerEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlementOwnerEntitlement` **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.DeleteEntitlementOwner()`: 
  * `request.DeleteAppEntitlementEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
* `conductoroneAPI.AppEntitlementOwnersV2.SearchEntitlementOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.ConnectorOwnersV2.CreateUserOwner()`: 
  * `request.CreateConnectorUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorOwnerUser` **Removed** (Breaking ⚠️)
    - `connectorOwnerUser` **Added**
* `conductoroneAPI.ConnectorOwnersV2.GetUserOwner()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorOwnerUser` **Removed** (Breaking ⚠️)
    - `connectorOwnerUser` **Added**
* `conductoroneAPI.ConnectorOwnersV2.DeleteUserOwner()`: 
  * `request.DeleteConnectorUserOwnerRequest` **Changed** (Breaking ⚠️)
    - `UserRef` **Removed** (Breaking ⚠️)
    - `userRef` **Added**
* `conductoroneAPI.ConnectorOwnersV2.SearchUserOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `user` **Added**
* `conductoroneAPI.AppEntitlementOwners.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `departmentSources[].priority` **Added**
    - `profile` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.ConnectorOwnersV2.CreateEntitlementOwner()`: 
  * `request.CreateConnectorEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorOwnerEntitlement` **Removed** (Breaking ⚠️)
    - `connectorOwnerEntitlement` **Added**
* `conductoroneAPI.ConnectorOwnersV2.GetEntitlementOwner()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorOwnerEntitlement` **Removed** (Breaking ⚠️)
    - `connectorOwnerEntitlement` **Added**
* `conductoroneAPI.ConnectorOwnersV2.DeleteEntitlementOwner()`: 
  * `request.DeleteConnectorEntitlementOwnerRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `appEntitlementRef` **Added**
* `conductoroneAPI.ConnectorOwnersV2.SearchEntitlementOwners()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Webhooks.Test()`: `response` **Changed** (Breaking ⚠️)
    - `WebhookInstance` **Removed** (Breaking ⚠️)
    - `webhook` **Added**
* `conductoroneAPI.Webhooks.Update()`: 
  * `request.WebhooksServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `webhook` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `webhook` **Added**
* `conductoroneAPI.Webhooks.Get()`: `response` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `webhook` **Added**
* `conductoroneAPI.Webhooks.Create()`: 
  *  `request.callbackTimeout` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `WebhookEndpoint` **Removed** (Breaking ⚠️)
    - `webhook` **Added**
* `conductoroneAPI.Webhooks.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `callbackTimeout` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Vault.Update()`: 
  * `request.VaultServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Vault` **Removed** (Breaking ⚠️)
    - `vault` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Vault` **Removed** (Breaking ⚠️)
    - `vault` **Added**
* `conductoroneAPI.Vault.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Vault` **Removed** (Breaking ⚠️)
    - `vault` **Added**
* `conductoroneAPI.Vault.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `GroupAuthzVault` **Removed** (Breaking ⚠️)
    - `MagicVault` **Removed** (Breaking ⚠️)
    - `groupAuthzVault` **Added**
    - `magicVault` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Vault` **Removed** (Breaking ⚠️)
    - `vault` **Added**
* `conductoroneAPI.User.SetExpiringUserDelegationBindingByAdmin()`: 
  * `request.SetExpiringUserDelegationBindingByAdminRequest` **Changed**
    - `delegationExpireAt` **Changed**
    - `delegationStartAt` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `ExpiringUserDelegationBinding` **Removed** (Breaking ⚠️)
    - `item` **Added**
* `conductoroneAPI.User.Get()`: `response` **Changed** (Breaking ⚠️)
    - `UserView` **Removed** (Breaking ⚠️)
    - `userView` **Added**
* `conductoroneAPI.User.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `objectPermissions` **Added**
    - `userId` **Added**
    - `user` **Added**
* `conductoroneAPI.TerraformExport.GetSchema()`: `response` **Changed** (Breaking ⚠️)
    - `TFSchemaMapping` **Removed** (Breaking ⚠️)
    - `schema` **Added**
* `conductoroneAPI.TaskActions.UpdateRequestData()`: 
  * `request.TaskActionsServiceUpdateRequestDataRequest` **Changed**
    - `data` **Changed**
    - `expandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.UpdateGrantDuration()`: 
  * `request.TaskActionsServiceUpdateGrantDurationRequest` **Changed**
    - `duration` **Changed**
    - `expandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.SkipStep()`: 
  *  `request.TaskActionsServiceSkipStepRequest.expandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.Restart()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.HardReset()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.Reassign()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.ProcessNow()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.EscalateToEmergencyAccess()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.Deny()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.Comment()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.Close()`: 
  *  `request.TaskActionsServiceCloseRequest.expandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.ApproveWithStepUp()`: 
  *  `request.TaskActionsServiceApproveWithStepUpRequest.expandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskActions.Approve()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.Task.Get()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.Task.CreateRevokeTask()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.Task.CreateOffboardingTask()`: `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.Task.CreateGrantTask()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `TaskGrantSource` **Removed** (Breaking ⚠️)
    - `grantDuration` **Changed**
    - `requestData` **Changed**
    - `source` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TaskView` **Removed** (Breaking ⚠️)
    - `taskView` **Added**
* `conductoroneAPI.TaskAudit.List()`: 
  * `request` **Changed**
    - `commentsOnly` **Added**
    - `newestFirst` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `TaskAuditAccessRequestOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditAccountLifecycleActionCreated` **Removed** (Breaking ⚠️)
    - `TaskAuditAccountLifecycleActionFailed` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceCreated` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceFailed` **Removed** (Breaking ⚠️)
    - `TaskAuditActionInstanceSucceeded` **Removed** (Breaking ⚠️)
    - `TaskAuditActionSubmitted` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalAutoAcceptedByPolicy` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalAutoRejectedByPolicy` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalHappenedAutomatically` **Removed** (Breaking ⚠️)
    - `TaskAuditApprovalInstanceChange` **Removed** (Breaking ⚠️)
    - `TaskAuditBulkActionError` **Removed** (Breaking ⚠️)
    - `TaskAuditCertifyOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditComment` **Removed** (Breaking ⚠️)
    - `TaskAuditConditionalPolicyExecutionResult` **Removed** (Breaking ⚠️)
    - `TaskAuditConnectorActionResult` **Removed** (Breaking ⚠️)
    - `TaskAuditCreatedReplacementExtensionGrantTask` **Removed** (Breaking ⚠️)
    - `TaskAuditEscalateToEmergencyAccess` **Removed** (Breaking ⚠️)
    - `TaskAuditExpressionPolicyStepError` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketCreated` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketError` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketProvisionStepResolved` **Removed** (Breaking ⚠️)
    - `TaskAuditExternalTicketTriggered` **Removed** (Breaking ⚠️)
    - `TaskAuditFinishedConnectorActions` **Removed** (Breaking ⚠️)
    - `TaskAuditFormInstanceChange` **Removed** (Breaking ⚠️)
    - `TaskAuditGrantDurationUpdated` **Removed** (Breaking ⚠️)
    - `TaskAuditGrantOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditHardReset` **Removed** (Breaking ⚠️)
    - `TaskAuditMetaData` **Removed** (Breaking ⚠️)
    - `TaskAuditNewTaskCreatedFrom` **Removed** (Breaking ⚠️)
    - `TaskAuditNewTask` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyApprovalReassigned` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyChanged` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyEvaluationStep` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyProvisionCancelled` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyProvisionError` **Removed** (Breaking ⚠️)
    - `TaskAuditPolicyProvisionReassigned` **Removed** (Breaking ⚠️)
    - `TaskAuditReassignedToDelegate` **Removed** (Breaking ⚠️)
    - `TaskAuditReassignmentFallbackToAdmin` **Removed** (Breaking ⚠️)
    - `TaskAuditReassignmentListError` **Removed** (Breaking ⚠️)
    - `TaskAuditRestart` **Removed** (Breaking ⚠️)
    - `TaskAuditRevokeOutcome` **Removed** (Breaking ⚠️)
    - `TaskAuditSLAEscalation` **Removed** (Breaking ⚠️)
    - `TaskAuditStartedConnectorActions` **Removed** (Breaking ⚠️)
    - `TaskAuditStateChange` **Removed** (Breaking ⚠️)
    - `TaskAuditStepSkipped` **Removed** (Breaking ⚠️)
    - `TaskAuditStepUpApproval` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitForAnalysisStepSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitForAnalysisStepTimedOut` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitForAnalysisStepWaiting` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepTimedOut` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepUntilTime` **Removed** (Breaking ⚠️)
    - `TaskAuditWaitStepWaiting` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalAttempt` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalBadResponse` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalFatalError` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookApprovalTriggered` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookAttempt` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookSuccess` **Removed** (Breaking ⚠️)
    - `TaskAuditWebhookTriggered` **Removed** (Breaking ⚠️)
    - `accessRequestOutcome` **Added**
    - `accountDeleted` **Added**
    - `accountLifecycleActionCreated` **Added**
    - `accountLifecycleActionFailed` **Added**
    - `actionInstanceCreated` **Added**
    - `actionInstanceFailed` **Added**
    - `actionInstanceSucceeded` **Added**
    - `actionResult` **Added**
    - `actionSubmitted` **Added**
    - `approvalAutoAcceptedByPolicy` **Added**
    - `approvalAutoRejectedByPolicy` **Added**
    - `approvalInstanceChange` **Added**
    - `approvalReassigned` **Added**
    - `approvedAutomatically` **Added**
    - `automationTriggered` **Added**
    - `bulkActionError` **Added**
    - `certifyOutcome` **Added**
    - `comment` **Added**
    - `conditionalPolicyExecutionResult` **Added**
    - `connectorActionsEnd` **Added**
    - `connectorActionsStart` **Added**
    - `createdReplacementExtensionGrantTask` **Added**
    - `created` **Changed** (Breaking ⚠️)
    - `expressionPolicyStepError` **Added**
    - `externalTicketCreated` **Added**
    - `externalTicketError` **Added**
    - `externalTicketProvisionStepResolved` **Added**
    - `externalTicketTriggered` **Added**
    - `formInstanceChange` **Added**
    - `grantDurationUpdated` **Added**
    - `grantOutcome` **Added**
    - `hardReset` **Added**
    - `metadata` **Added**
    - `policyChanged` **Added**
    - `policyEvaluationStep` **Added**
    - `provisionCancelled` **Added**
    - `provisionEntitlementMergeCompleted` **Added**
    - `provisionEntitlementMergeTimedOut` **Added**
    - `provisionError` **Added**
    - `provisionReassigned` **Added**
    - `provisionWaitingForEntitlementMerge` **Added**
    - `reassignedToDelegate` **Added**
    - `reassignmentFallbackToAdmin` **Added**
    - `reassignmentListError` **Added**
    - `requestDefaultsApplied` **Added**
    - `revokeOutcome` **Added**
    - `slaEscalation` **Added**
    - `stateChange` **Added**
    - `stepSkipped` **Added**
    - `stepUpApproval` **Added**
    - `taskCreatedFrom` **Added**
    - `taskCreated` **Added**
    - `taskEscalated` **Added**
    - `taskRestarted` **Added**
    - `waitStepAnalysisSuccess` **Added**
    - `waitStepAnalysisTimedOut` **Added**
    - `waitStepAnalysisWaiting` **Added**
    - `waitStepSuccess` **Added**
    - `waitStepTimedOut` **Added**
    - `waitStepUntilTime` **Added**
    - `waitStepWaiting` **Added**
    - `webhookApprovalAttempt` **Added**
    - `webhookApprovalBadResponse` **Added**
    - `webhookApprovalFatalError` **Added**
    - `webhookApprovalSuccess` **Added**
    - `webhookApprovalTriggered` **Added**
    - `webhookAttempt` **Added**
    - `webhookSuccess` **Added**
    - `webhookTriggered` **Added**
* `conductoroneAPI.Export.Update()`: 
  * `request.ExportServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Exporter` **Removed** (Breaking ⚠️)
    - `exporter` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Exporter` **Removed** (Breaking ⚠️)
    - `exporter` **Added**
* `conductoroneAPI.Export.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Exporter` **Removed** (Breaking ⚠️)
    - `exporter` **Added**
* `conductoroneAPI.Export.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `ExportToDatasource` **Removed** (Breaking ⚠️)
    - `datasource` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Exporter` **Removed** (Breaking ⚠️)
    - `exporter` **Added**
* `conductoroneAPI.Export.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ExportToDatasource` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `datasource` **Added**
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.SSFReceiverEvent.List()`:  `response.list[].receivedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.SSFReceiverStream.GetStats()`: `response` **Changed** (Breaking ⚠️)
    - `SSFReceiverStreamStats` **Removed** (Breaking ⚠️)
    - `stats` **Added**
* `conductoroneAPI.SSFReceiverStream.Update()`: 
  * `request.SSFReceiverStreamServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `SSFReceiverStream` **Removed** (Breaking ⚠️)
    - `ssfReceiverStream` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `SSFReceiverStream` **Removed** (Breaking ⚠️)
    - `ssfReceiverStream` **Added**
* `conductoroneAPI.SSFReceiverStream.Get()`: `response` **Changed** (Breaking ⚠️)
    - `SSFReceiverStream` **Removed** (Breaking ⚠️)
    - `ssfReceiverStream` **Added**
* `conductoroneAPI.SSFReceiverStream.Create()`: 
  *  `request.pollInterval` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `SSFReceiverStream` **Removed** (Breaking ⚠️)
    - `ssfReceiverStream` **Added**
* `conductoroneAPI.SSFReceiverStream.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `SSFOutboundAuthBearer` **Removed** (Breaking ⚠️)
    - `SSFOutboundAuthOAuth2` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `lastErrorAt` **Changed** (Breaking ⚠️)
    - `lastVerifiedAt` **Changed** (Breaking ⚠️)
    - `outboundAuthBearer` **Added**
    - `outboundAuthOauth2` **Added**
    - `pollInterval` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.SessionSettings.TestSourceIP()`: `response` **Changed** (Breaking ⚠️)
    - `Status` **Removed** (Breaking ⚠️)
    - `details` **Added**
* `conductoroneAPI.SessionSettings.Update()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `SessionSettings` **Removed** (Breaking ⚠️)
    - `sessionSettings` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `SessionSettings` **Removed** (Breaking ⚠️)
    - `sessionSettings` **Added**
* `conductoroneAPI.SessionSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `SessionSettings` **Removed** (Breaking ⚠️)
    - `sessionSettings` **Added**
* `conductoroneAPI.RequestSettings.Update()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `RequestSettings` **Removed** (Breaking ⚠️)
    - `requestSettings` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `RequestSettings` **Removed** (Breaking ⚠️)
    - `requestSettings` **Added**
* `conductoroneAPI.RequestSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `RequestSettings` **Removed** (Breaking ⚠️)
    - `requestSettings` **Added**
* `conductoroneAPI.OnboardingSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `OnboardingOrgContext` **Removed** (Breaking ⚠️)
    - `mcpOnboardingGoal` **Added**
    - `mcpOnboardingStatus` **Added**
    - `mcpOnboardingTargets` **Added**
    - `orgContext` **Added**
* `conductoroneAPI.UserNotificationSettings.Update()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `ChannelSettings` **Removed** (Breaking ⚠️)
    - `channelSettings` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `UserNotificationSettings` **Removed** (Breaking ⚠️)
    - `userNotificationSettings` **Added**
* `conductoroneAPI.UserNotificationSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `UserNotificationSettings` **Removed** (Breaking ⚠️)
    - `userNotificationSettings` **Added**
* `conductoroneAPI.OrgNotificationSettings.Update()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `ChannelSettings` **Removed** (Breaking ⚠️)
    - `channelSettings` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `OrgNotificationSettings` **Removed** (Breaking ⚠️)
    - `orgNotificationSettings` **Added**
* `conductoroneAPI.OrgNotificationSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `OrgNotificationSettings` **Removed** (Breaking ⚠️)
    - `orgNotificationSettings` **Added**
* `conductoroneAPI.TenantEmailProvider.Update()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `TenantEmailProvider` **Removed** (Breaking ⚠️)
    - `emailProvider` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TenantEmailProvider` **Removed** (Breaking ⚠️)
    - `emailProvider` **Added**
* `conductoroneAPI.TenantEmailProvider.Get()`: `response` **Changed** (Breaking ⚠️)
    - `TenantEmailProvider` **Removed** (Breaking ⚠️)
    - `emailProvider` **Added**
* `conductoroneAPI.OrgDomain.Update()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.OrgDomain.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.UserDeveloperPreferences.Update()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `TerraformPreferences` **Removed** (Breaking ⚠️)
    - `terraform` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `UserDeveloperPreferences` **Removed** (Breaking ⚠️)
    - `userDeveloperPreferences` **Added**
* `conductoroneAPI.UserDeveloperPreferences.Get()`: `response` **Changed** (Breaking ⚠️)
    - `UserDeveloperPreferences` **Removed** (Breaking ⚠️)
    - `userDeveloperPreferences` **Added**
* `conductoroneAPI.Contacts.UpdateContacts()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `Contacts` **Removed** (Breaking ⚠️)
    - `contacts` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Contacts` **Removed** (Breaking ⚠️)
    - `contacts` **Added**
* `conductoroneAPI.Contacts.GetContacts()`: `response` **Changed** (Breaking ⚠️)
    - `Contacts` **Removed** (Breaking ⚠️)
    - `contacts` **Added**
* `conductoroneAPI.AWSExternalIDSettings.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AWSExternalID` **Removed** (Breaking ⚠️)
    - `awsExternalId` **Added**
* `conductoroneAPI.Principal.ListBindings()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `ServicePrincipalBindingSubject` **Removed** (Breaking ⚠️)
    - `subject` **Added**
  * `response.bindings[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Principal.DeleteBinding()`: `request` **Changed** (Breaking ⚠️)
    - `ServicePrincipalBindingSubject` **Removed** (Breaking ⚠️)
    - `subject` **Added**
* `conductoroneAPI.Principal.AddBinding()`: `request` **Changed** (Breaking ⚠️)
    - `ServicePrincipalBindingSubject` **Removed** (Breaking ⚠️)
    - `subject` **Added**
* `conductoroneAPI.Principal.UpdateCredential()`: 
  * `request.ServicePrincipalServiceUpdateCredentialRequest` **Changed** (Breaking ⚠️)
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
    - `credential` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
    - `credential` **Added**
* `conductoroneAPI.Principal.GetCredential()`: `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
    - `credential` **Added**
* `conductoroneAPI.Principal.CreateCredential()`: 
  *  `request.ServicePrincipalServiceCreateCredentialRequest.expires` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipalCredential` **Removed** (Breaking ⚠️)
    - `credential` **Added**
* `conductoroneAPI.Principal.ListCredentials()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `expiresAt` **Changed** (Breaking ⚠️)
    - `lastUsedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Principal.Update()`: 
  * `request.ServicePrincipalServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
    - `servicePrincipal` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
    - `servicePrincipal` **Added**
* `conductoroneAPI.Principal.Get()`: `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
    - `servicePrincipal` **Added**
* `conductoroneAPI.Principal.Create()`: `response` **Changed** (Breaking ⚠️)
    - `ServicePrincipal` **Removed** (Breaking ⚠️)
    - `servicePrincipal` **Added**
* `conductoroneAPI.Principal.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `objectPermissions` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
    - `user` **Added**
* `conductoroneAPI.WorkloadFederation.UpdateProvider()`: 
  * `request.WorkloadFederationServiceUpdateProviderRequest` **Changed** (Breaking ⚠️)
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
    - `provider` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
    - `provider` **Added**
* `conductoroneAPI.WorkloadFederation.GetProvider()`: `response` **Changed** (Breaking ⚠️)
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
    - `provider` **Added**
* `conductoroneAPI.WorkloadFederation.CreateProvider()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `OIDCSettings` **Removed** (Breaking ⚠️)
    - `SPIFFESettings` **Removed** (Breaking ⚠️)
    - `oidc` **Added**
    - `spiffe` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `WorkloadFederationProvider` **Removed** (Breaking ⚠️)
    - `provider` **Added**
* `conductoroneAPI.WorkloadFederation.ListProviders()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `OIDCSettings` **Removed** (Breaking ⚠️)
    - `SPIFFESettings` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `oidc` **Added**
    - `spiffe` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.WorkloadFederation.TestToken()`: `response` **Changed** (Breaking ⚠️)
    - `TestTokenStepResult1` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult2` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult3` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult4` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult5` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult6` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult7` **Removed** (Breaking ⚠️)
    - `TestTokenStepResult` **Removed** (Breaking ⚠️)
    - `audienceValidation` **Added**
    - `celEvaluation` **Added**
    - `cidrCheck` **Added**
    - `issuerMatch` **Added**
    - `jwtDecode` **Added**
    - `signatureValidation` **Added**
    - `subjectValidation` **Added**
    - `tokenFreshness` **Added**
* `conductoroneAPI.WorkloadFederation.UpdateTrust()`: 
  * `request.WorkloadFederationServiceUpdateTrustRequest` **Changed** (Breaking ⚠️)
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
    - `trust` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
    - `trust` **Added**
* `conductoroneAPI.WorkloadFederation.GetTrust()`: `response` **Changed** (Breaking ⚠️)
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
    - `trust` **Added**
* `conductoroneAPI.WorkloadFederation.CreateTrust()`: `response` **Changed** (Breaking ⚠️)
    - `WorkloadFederationTrust` **Removed** (Breaking ⚠️)
    - `trust` **Added**
* `conductoroneAPI.WorkloadFederation.ListTrusts()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.WorkloadFederation.SearchTrusts()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.WebhooksSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `callbackTimeout` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.UserSearch.Search()`: 
  *  `request.sourceAppIds` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `objectPermissions` **Added**
    - `userId` **Added**
    - `user` **Added**
* `conductoroneAPI.TaskSearch.Search()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `accountStatuses` **Added**
    - `createdAfter` **Changed**
    - `createdBefore` **Changed**
    - `includeActedAfter` **Changed**
    - `olderThanDuration` **Changed**
    - `outcomeAfter` **Changed**
    - `outcomeBefore` **Changed**
    - `taskTypes[].TaskTypeAction` **Removed** (Breaking ⚠️)
    - `taskTypes[].TaskTypeCertify` **Removed** (Breaking ⚠️)
    - `taskTypes[].TaskTypeFinding` **Removed** (Breaking ⚠️)
    - `taskTypes[].TaskTypeGrant` **Removed** (Breaking ⚠️)
    - `taskTypes[].TaskTypeOffboarding` **Removed** (Breaking ⚠️)
    - `taskTypes[].TaskTypeRevoke` **Removed** (Breaking ⚠️)
    - `taskTypes[].action` **Added**
    - `taskTypes[].certify` **Added**
    - `taskTypes[].finding` **Added**
    - `taskTypes[].grant` **Added**
    - `taskTypes[].offboarding` **Added**
    - `taskTypes[].revoke` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `Task` **Removed** (Breaking ⚠️)
    - `objectPermissions` **Added**
    - `principalResourcePath` **Added**
    - `task` **Added**
* `conductoroneAPI.ExportsSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ExportToDatasource` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `datasource` **Added**
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.StepUpTransaction.Get()`: `response` **Changed** (Breaking ⚠️)
    - `StepUpTransaction` **Removed** (Breaking ⚠️)
    - `transaction` **Added**
* `conductoroneAPI.StepUpTransaction.Search()`: 
  * `request` **Changed**
    - `createdAfter` **Changed**
    - `createdBefore` **Changed**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `TargetTask` **Removed** (Breaking ⚠️)
    - `TargetTest` **Removed** (Breaking ⚠️)
    - `approveTask` **Added**
    - `claims` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `expiresAt` **Changed** (Breaking ⚠️)
    - `test` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.StepUpProvider.UpdateSecret()`: `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Removed** (Breaking ⚠️)
    - `stepUpProvider` **Added**
* `conductoroneAPI.StepUpProvider.Update()`: 
  * `request.UpdateStepUpProviderRequest` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Removed** (Breaking ⚠️)
    - `stepUpProvider` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Removed** (Breaking ⚠️)
    - `stepUpProvider` **Added**
* `conductoroneAPI.StepUpProvider.Get()`: `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Removed** (Breaking ⚠️)
    - `stepUpProvider` **Added**
* `conductoroneAPI.StepUpProvider.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `StepUpMicrosoftSettings` **Removed** (Breaking ⚠️)
    - `StepUpOAuth2Settings` **Removed** (Breaking ⚠️)
    - `microsoft` **Added**
    - `oauth2` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `StepUpProvider` **Removed** (Breaking ⚠️)
    - `stepUpProvider` **Added**
* `conductoroneAPI.StepUpProvider.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `StepUpMicrosoftSettings` **Removed** (Breaking ⚠️)
    - `StepUpOAuth2Settings` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `lastTestedAt` **Changed** (Breaking ⚠️)
    - `microsoft` **Added**
    - `oauth2` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.StepUpProvider.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `StepUpMicrosoftSettings` **Removed** (Breaking ⚠️)
    - `StepUpOAuth2Settings` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `lastTestedAt` **Changed** (Breaking ⚠️)
    - `microsoft` **Added**
    - `oauth2` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.SSFReceiverEventSearch.Search()`:  `response.list[].receivedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.PaperSecret.CreateInternal()`: 
  * `request` **Changed**
    - `expiresIn` **Changed**
    - `requiredAgeSuite` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `ageSuite` **Added**
    - `secret` **Added**
* `conductoroneAPI.PaperSecret.CreateExternal()`: 
  * `request` **Changed**
    - `expiresIn` **Changed**
    - `requiredAgeSuite` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `ageSuite` **Added**
    - `secret` **Added**
* `conductoroneAPI.PaperSecret.GetByShareCode()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `secret` **Added**
* `conductoroneAPI.PaperSecret.GetContent()`:  `response.createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.PaperSecret.SetTextContent()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `secret` **Added**
* `conductoroneAPI.PaperSecret.Get()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `secret` **Added**
* `conductoroneAPI.PaperSecret.Revoke()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `secret` **Added**
* `conductoroneAPI.PaperSecret.SearchMySecrets()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ageSuite` **Added**
    - `contentExpiresAt` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.PaperSecretAdmin.Get()`: `response` **Changed** (Breaking ⚠️)
    - `PaperSecret` **Removed** (Breaking ⚠️)
    - `secret` **Added**
* `conductoroneAPI.PaperSecretAdmin.Search()`: 
  * `request` **Changed**
    - `createdAfter` **Changed**
    - `createdBefore` **Changed**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `ageSuite` **Added**
    - `contentExpiresAt` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.RoleMiningManagementSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `lastGeneratedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.RequestCatalogSearch.SearchEntitlements()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
    - `appEntitlementUserBindings[].createdAt` **Changed** (Breaking ⚠️)
    - `appEntitlementUserBindings[].deletedAt` **Changed** (Breaking ⚠️)
    - `appEntitlementUserBindings[].deprovisionAt` **Changed** (Breaking ⚠️)
    - `entitlement` **Added**
* `conductoroneAPI.PolicySearch.Search()`: 
  * `request` **Changed**
    - `scopeAppEntitlementId` **Added**
    - `scopeAppId` **Added**
    - `scopeObjectType` **Added**
    - `scopeSlot` **Added**
    - `scopeView` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `baselinePolicyId` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Accept` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Action` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Approval` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Form` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Provision` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Reject` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Wait` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].accept` **Added**
    - `policySteps.Map<PolicySteps>.steps[].action` **Added**
    - `policySteps.Map<PolicySteps>.steps[].approval` **Added**
    - `policySteps.Map<PolicySteps>.steps[].form` **Added**
    - `policySteps.Map<PolicySteps>.steps[].provision` **Added**
    - `policySteps.Map<PolicySteps>.steps[].reject` **Added**
    - `policySteps.Map<PolicySteps>.steps[].wait` **Added**
    - `rules[].policyId` **Added**
    - `rules[].stepKey` **Added**
    - `scope` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.PersonalClientSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `expiresTime` **Changed** (Breaking ⚠️)
    - `lastUsedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.ExternalClientSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `clientIdType.enum(CLIENT_ID_TYPE_APP)` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `lastUsedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.HooksSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `BuiltInPattern` **Removed** (Breaking ⚠️)
    - `HookFilter` **Removed** (Breaking ⚠️)
    - `HookFunctionRef` **Removed** (Breaking ⚠️)
    - `builtinPattern` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `event.enum(HOOK_EVENT_TYPE_PRE_OUTPUT)` **Added**
    - `filter` **Added**
    - `function` **Added**
    - `jsonPatch` **Added**
    - `managedByGuardrails` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.FunctionsSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `hookRefs` **Added**
    - `provisionedConcurrency` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
    - `workflowTemplateRefs` **Added**
* `conductoroneAPI.AutomationSearch.SearchAutomations()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `DisabledReasonCircuitBreaker` **Removed** (Breaking ⚠️)
    - `automationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `automationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `automationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `automationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `automationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `automationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `automationSteps[].accountLifecycleAction` **Added**
    - `automationSteps[].callFunction` **Added**
    - `automationSteps[].connectorAction` **Added**
    - `automationSteps[].connectorCreateAccount` **Added**
    - `automationSteps[].createAccessReview` **Added**
    - `automationSteps[].createRevokeTasksV2` **Added**
    - `automationSteps[].createRevokeTasks` **Added**
    - `automationSteps[].evaluateExpressions` **Added**
    - `automationSteps[].generatePassword` **Added**
    - `automationSteps[].grantEntitlements` **Added**
    - `automationSteps[].removeFromDelegation` **Added**
    - `automationSteps[].runAutomation` **Added**
    - `automationSteps[].sendEmail` **Added**
    - `automationSteps[].sendSlackMessage` **Added**
    - `automationSteps[].setCredential` **Added**
    - `automationSteps[].storeCredential` **Added**
    - `automationSteps[].taskAction` **Added**
    - `automationSteps[].unenrollFromAllAccessProfiles` **Added**
    - `automationSteps[].updateUser` **Added**
    - `automationSteps[].waitForDuration` **Added**
    - `automationSteps[].webhook` **Added**
    - `circuitBreaker` **Added**
    - `context` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `draftTriggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].accessConflict` **Added**
    - `draftTriggers[].appUserCreated` **Added**
    - `draftTriggers[].appUserUpdated` **Added**
    - `draftTriggers[].grantDeleted` **Added**
    - `draftTriggers[].grantFound` **Added**
    - `draftTriggers[].scheduleAppUser` **Added**
    - `draftTriggers[].scheduleNoUser` **Added**
    - `draftTriggers[].schedule` **Added**
    - `draftTriggers[].usageBasedRevocation` **Added**
    - `draftTriggers[].userCreated` **Added**
    - `draftTriggers[].userProfileChange` **Added**
    - `draftTriggers[].webhook` **Added**
    - `lastExecutedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AutomationSearch.SearchAutomationTemplateVersions()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `automationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `automationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `automationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `automationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `automationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `automationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `automationSteps[].accountLifecycleAction` **Added**
    - `automationSteps[].callFunction` **Added**
    - `automationSteps[].connectorAction` **Added**
    - `automationSteps[].connectorCreateAccount` **Added**
    - `automationSteps[].createAccessReview` **Added**
    - `automationSteps[].createRevokeTasksV2` **Added**
    - `automationSteps[].createRevokeTasks` **Added**
    - `automationSteps[].evaluateExpressions` **Added**
    - `automationSteps[].generatePassword` **Added**
    - `automationSteps[].grantEntitlements` **Added**
    - `automationSteps[].removeFromDelegation` **Added**
    - `automationSteps[].runAutomation` **Added**
    - `automationSteps[].sendEmail` **Added**
    - `automationSteps[].sendSlackMessage` **Added**
    - `automationSteps[].setCredential` **Added**
    - `automationSteps[].storeCredential` **Added**
    - `automationSteps[].taskAction` **Added**
    - `automationSteps[].unenrollFromAllAccessProfiles` **Added**
    - `automationSteps[].updateUser` **Added**
    - `automationSteps[].waitForDuration` **Added**
    - `automationSteps[].webhook` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `triggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `triggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `triggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `triggers[].accessConflict` **Added**
    - `triggers[].appUserCreated` **Added**
    - `triggers[].appUserUpdated` **Added**
    - `triggers[].grantDeleted` **Added**
    - `triggers[].grantFound` **Added**
    - `triggers[].scheduleAppUser` **Added**
    - `triggers[].scheduleNoUser` **Added**
    - `triggers[].schedule` **Added**
    - `triggers[].usageBasedRevocation` **Added**
    - `triggers[].userCreated` **Added**
    - `triggers[].userProfileChange` **Added**
    - `triggers[].webhook` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AttributeSearch.SearchAttributeValues()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUserMapper` **Removed** (Breaking ⚠️)
    - `appOwners[].createdAt` **Changed** (Breaking ⚠️)
    - `appOwners[].deletedAt` **Changed** (Breaking ⚠️)
    - `appOwners[].departmentSources[].priority` **Added**
    - `appOwners[].profile` **Changed** (Breaking ⚠️)
    - `appOwners[].updatedAt` **Changed** (Breaking ⚠️)
    - `appUserMapper` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `revokeGrantSources` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppResourceSearch.SearchAppResources()`: 
  * `request` **Changed**
    - `agentStatuses` **Added**
    - `appIds` **Added**
    - `credentialTypes` **Added**
    - `direction` **Added**
    - `excludeDeletedApps` **Added**
    - `nhiTypes` **Added**
    - `secretAging` **Added**
    - `sortField` **Added**
    - `unownedOnly` **Added**
    - `withOpenFindings` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppResource` **Removed** (Breaking ⚠️)
    - `appResource` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.AppResourceSearch.SearchAppResourceTypes()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AutomationExecutionSearch.SearchAutomationExecutions()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AutomationExecutionExpandMask` **Removed** (Breaking ⚠️)
    - `expandMask` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `AutomationExecution` **Removed** (Breaking ⚠️)
    - `automationExecution` **Added**
* `conductoroneAPI.AutomationExecutionSearch.SearchAllAutomationExecutions()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AutomationExecutionExpandMask` **Removed** (Breaking ⚠️)
    - `expandMask` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `AutomationExecution` **Removed** (Breaking ⚠️)
    - `automationExecution` **Added**
* `conductoroneAPI.RoleMiningManagement.SearchCohortUsers()`: `response` **Changed** (Breaking ⚠️)
    - `list[].createdAt` **Changed** (Breaking ⚠️)
    - `list[].deletedAt` **Changed** (Breaking ⚠️)
    - `list[].departmentSources[].priority` **Added**
    - `list[].profile` **Changed** (Breaking ⚠️)
    - `list[].updatedAt` **Changed** (Breaking ⚠️)
    - `usersWithCoverage[].User` **Removed** (Breaking ⚠️)
    - `usersWithCoverage[].user` **Added**
* `conductoroneAPI.RoleMiningManagement.UpdateSuggestionState()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementSuggestion` **Removed** (Breaking ⚠️)
    - `suggestion` **Added**
* `conductoroneAPI.RoleMiningManagement.GetSuggestion()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementSuggestion` **Removed** (Breaking ⚠️)
    - `suggestion` **Added**
* `conductoroneAPI.RoleMiningManagement.ListSuggestions()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `lastGeneratedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.RoleMiningManagement.GetLatestRun()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementRun` **Removed** (Breaking ⚠️)
    - `run` **Added**
* `conductoroneAPI.RoleMiningManagement.ListRuns()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `completedAt` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.RoleMiningManagement.ListCustomAnalysisResults()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `completedAt` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.RoleMiningManagement.UpdateRoleMiningConfig()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementConfig` **Removed** (Breaking ⚠️)
    - `config` **Added**
* `conductoroneAPI.RoleMiningManagement.GetRoleMiningConfig()`: `response` **Changed** (Breaking ⚠️)
    - `RoleMiningManagementConfig` **Removed** (Breaking ⚠️)
    - `config` **Added**
* `conductoroneAPI.RequestSchema.Update()`: 
  * `request.RequestSchemaServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Removed** (Breaking ⚠️)
    - `requestSchema` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Removed** (Breaking ⚠️)
    - `requestSchema` **Added**
* `conductoroneAPI.RequestSchema.Get()`: `response` **Changed** (Breaking ⚠️)
    - `RequestSchema` **Removed** (Breaking ⚠️)
    - `requestSchema` **Added**
* `conductoroneAPI.Finding.GetFinding()`: `response` **Changed** (Breaking ⚠️)
    - `Finding` **Removed** (Breaking ⚠️)
    - `finding` **Added**
* `conductoroneAPI.RequestSchema.FindBindingForAppEntitlement()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `entitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `entitlementRef` **Added**
* `conductoroneAPI.RequestSchema.CreateEntitlementBinding()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `entitlementRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `entitlementRef` **Added**
* `conductoroneAPI.RequestSchema.RemoveEntitlementBinding()`: `request` **Changed** (Breaking ⚠️)
    - `AppEntitlementRef` **Removed** (Breaking ⚠️)
    - `entitlementRef` **Added**
* `conductoroneAPI.Policies.Update()`: 
  * `request.UpdatePolicyRequest` **Changed** (Breaking ⚠️)
    - `Policy` **Removed** (Breaking ⚠️)
    - `policy` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Policy` **Removed** (Breaking ⚠️)
    - `policy` **Added**
* `conductoroneAPI.Policies.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Policy` **Removed** (Breaking ⚠️)
    - `policy` **Added**
* `conductoroneAPI.Policies.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `baselinePolicyId` **Added**
    - `policySteps.Map<PolicySteps>.steps[].Accept` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Action` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Approval` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Form` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Provision` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Reject` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Wait` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].accept` **Added**
    - `policySteps.Map<PolicySteps>.steps[].action` **Added**
    - `policySteps.Map<PolicySteps>.steps[].approval` **Added**
    - `policySteps.Map<PolicySteps>.steps[].form` **Added**
    - `policySteps.Map<PolicySteps>.steps[].provision` **Added**
    - `policySteps.Map<PolicySteps>.steps[].reject` **Added**
    - `policySteps.Map<PolicySteps>.steps[].wait` **Added**
    - `rules[].policyId` **Added**
    - `rules[].stepKey` **Added**
    - `scope` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Policy` **Removed** (Breaking ⚠️)
    - `policy` **Added**
* `conductoroneAPI.Policies.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `baselinePolicyId` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Accept` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Action` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Approval` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Form` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Provision` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Reject` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].Wait` **Removed** (Breaking ⚠️)
    - `policySteps.Map<PolicySteps>.steps[].accept` **Added**
    - `policySteps.Map<PolicySteps>.steps[].action` **Added**
    - `policySteps.Map<PolicySteps>.steps[].approval` **Added**
    - `policySteps.Map<PolicySteps>.steps[].form` **Added**
    - `policySteps.Map<PolicySteps>.steps[].provision` **Added**
    - `policySteps.Map<PolicySteps>.steps[].reject` **Added**
    - `policySteps.Map<PolicySteps>.steps[].wait` **Added**
    - `rules[].policyId` **Added**
    - `rules[].stepKey` **Added**
    - `scope` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.LocalUserInvitation.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `acceptedAt` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `expiresAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.LocalUserInvitation.Revoke()`: `response` **Changed** (Breaking ⚠️)
    - `LocalUserInvitation` **Removed** (Breaking ⚠️)
    - `invitation` **Added**
* `conductoroneAPI.LocalUserInvitation.Get()`: `response` **Changed** (Breaking ⚠️)
    - `LocalUserInvitation` **Removed** (Breaking ⚠️)
    - `invitation` **Added**
* `conductoroneAPI.LocalUserInvitation.Create()`: `response` **Changed** (Breaking ⚠️)
    - `LocalUserInvitation` **Removed** (Breaking ⚠️)
    - `invitation` **Added**
* `conductoroneAPI.LocalDirectoryConfig.Update()`: 
  * `request.LocalDirectoryConfigServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
    - `localDirectoryConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
    - `localDirectoryConfig` **Added**
* `conductoroneAPI.LocalDirectoryConfig.Get()`: `response` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
    - `localDirectoryConfig` **Added**
* `conductoroneAPI.LocalDirectoryConfig.Create()`: 
  *  `request.invitationTtl` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `LocalDirectoryConfig` **Removed** (Breaking ⚠️)
    - `localDirectoryConfig` **Added**
* `conductoroneAPI.LocalDirectoryConfig.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `invitationTtl` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Roles.Update()`: 
  * `request.UpdateRoleRequest` **Changed** (Breaking ⚠️)
    - `Role` **Removed** (Breaking ⚠️)
    - `role` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Role` **Removed** (Breaking ⚠️)
    - `role` **Added**
* `conductoroneAPI.Roles.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Role` **Removed** (Breaking ⚠️)
    - `role` **Added**
* `conductoroneAPI.Roles.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.PersonalClient.Update()`: 
  * `request.PersonalClientServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `PersonalClient` **Removed** (Breaking ⚠️)
    - `client` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `PersonalClient` **Removed** (Breaking ⚠️)
    - `client` **Added**
* `conductoroneAPI.PersonalClient.Get()`: `response` **Changed** (Breaking ⚠️)
    - `PersonalClient` **Removed** (Breaking ⚠️)
    - `client` **Added**
* `conductoroneAPI.PersonalClient.Create()`: 
  *  `request.expires` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `PersonalClient` **Removed** (Breaking ⚠️)
    - `client` **Added**
* `conductoroneAPI.PersonalClient.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `expiresTime` **Changed** (Breaking ⚠️)
    - `lastUsedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Hooks.Update()`: 
  * `request.HooksServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Hook` **Removed** (Breaking ⚠️)
    - `hook` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Hook` **Removed** (Breaking ⚠️)
    - `hook` **Added**
* `conductoroneAPI.Hooks.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Hook` **Removed** (Breaking ⚠️)
    - `hook` **Added**
* `conductoroneAPI.Hooks.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `BuiltInPattern` **Removed** (Breaking ⚠️)
    - `HookFilter` **Removed** (Breaking ⚠️)
    - `HookFunctionRef` **Removed** (Breaking ⚠️)
    - `builtinPattern` **Added**
    - `event.enum(HOOK_EVENT_TYPE_PRE_OUTPUT)` **Added**
    - `filter` **Added**
    - `function` **Added**
    - `jsonPatch` **Added**
    - `managedByGuardrails` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Hook` **Removed** (Breaking ⚠️)
    - `hook` **Added**
* `conductoroneAPI.Hooks.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `BuiltInPattern` **Removed** (Breaking ⚠️)
    - `HookFilter` **Removed** (Breaking ⚠️)
    - `HookFunctionRef` **Removed** (Breaking ⚠️)
    - `builtinPattern` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `event.enum(HOOK_EVENT_TYPE_PRE_OUTPUT)` **Added**
    - `filter` **Added**
    - `function` **Added**
    - `jsonPatch` **Added**
    - `managedByGuardrails` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.FunctionsInvocationSearch.Search()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `input` **Changed** (Breaking ⚠️)
    - `output` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.FunctionsInvocation.Get()`: `response` **Changed** (Breaking ⚠️)
    - `FunctionInvocation` **Removed** (Breaking ⚠️)
    - `invocation` **Added**
* `conductoroneAPI.FunctionsInvocation.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `input` **Changed** (Breaking ⚠️)
    - `output` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Functions.UpdateFunction()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `Function` **Removed** (Breaking ⚠️)
    - `commitMessage` **Added**
    - `content` **Added**
    - `function` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Function` **Removed** (Breaking ⚠️)
    - `commit` **Added**
    - `function` **Added**
* `conductoroneAPI.Functions.GetFunction()`: `response` **Changed** (Breaking ⚠️)
    - `Function` **Removed** (Breaking ⚠️)
    - `function` **Added**
* `conductoroneAPI.Functions.Test()`: `response` **Changed** (Breaking ⚠️)
    - `FunctionTestResult` **Removed** (Breaking ⚠️)
    - `result` **Added**
* `conductoroneAPI.Functions.ListTags()`:  `response.tags.Map<FunctionCommit>.createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Functions.GetCommitContent()`: `response` **Changed** (Breaking ⚠️)
    - `FunctionCommit` **Removed** (Breaking ⚠️)
    - `commit` **Added**
* `conductoroneAPI.Functions.CreateFinalCommit()`: `response` **Changed** (Breaking ⚠️)
    - `FunctionCommit` **Removed** (Breaking ⚠️)
    - `commit` **Added**
* `conductoroneAPI.Functions.ListCommits()`:  `response.list[].createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Functions.CreateFunction()`: `response` **Changed** (Breaking ⚠️)
    - `FunctionCommit` **Removed** (Breaking ⚠️)
    - `Function` **Removed** (Breaking ⚠️)
    - `commit` **Added**
    - `function` **Added**
* `conductoroneAPI.Functions.ListFunctions()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `hookRefs` **Added**
    - `provisionedConcurrency` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
    - `workflowTemplateRefs` **Added**
* `conductoroneAPI.FindingSearch.Search()`: 
  * `request` **Changed**
    - `appResourceIds` **Added**
    - `appResourceTraitIds` **Added**
    - `appResourceTypeIds` **Added**
    - `appUserTypes` **Added**
    - `connectorIds` **Added**
    - `customSubTypes` **Added**
    - `decoyIds` **Added**
    - `findingTypes[].enum(FINDING_TYPE_CONNECTOR_ANOMALY_DETECTION_DISABLED)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_CREDENTIAL_EXPIRING)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_CREDENTIAL_PUBLICLY_EXPOSED)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_CUSTOM)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_DEACTIVATED_OWNER)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_DECOY_PUBLICLY_EXPOSED)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_NHI_UNOWNED)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_SERVICE_ACCOUNT_UNOWNED)` **Added**
    - `findingTypes[].enum(FINDING_TYPE_UNUSED_SECRET)` **Added**
    - `includeUnassigned` **Added**
    - `nhiTypes` **Added**
    - `ownerIdentityUserIds` **Added**
    - `refs` **Added**
    - `scopeToAppOwner` **Added**
    - `sourceKinds` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUserTarget` **Removed** (Breaking ⚠️)
    - `DecoyCredentialUsedType` **Removed** (Breaking ⚠️)
    - `FindingOwnerRef1` **Removed** (Breaking ⚠️)
    - `FindingOwnerRef` **Removed** (Breaking ⚠️)
    - `FindingRiskScore` **Removed** (Breaking ⚠️)
    - `IdentityUserTarget` **Removed** (Breaking ⚠️)
    - `ServiceAccountMisclassificationEvidence` **Removed** (Breaking ⚠️)
    - `ServiceAccountMisclassificationType` **Removed** (Breaking ⚠️)
    - `SimilarUsernameMatchEvidence` **Removed** (Breaking ⚠️)
    - `SimilarUsernameMatchType` **Removed** (Breaking ⚠️)
    - `annotations` **Added**
    - `appResourceTarget` **Added**
    - `appUserTarget` **Added**
    - `assignedOwner` **Added**
    - `computedOwner` **Added**
    - `connectorAnomalyDetectionDisabled` **Added**
    - `connectorTarget` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `credentialExpiringEvidence` **Added**
    - `credentialExpiring` **Added**
    - `credentialPubliclyExposedEvidence` **Added**
    - `credentialPubliclyExposed` **Added**
    - `customSubType` **Added**
    - `custom` **Added**
    - `deactivatedOwnerEvidence` **Added**
    - `deactivatedOwner` **Added**
    - `decoyCredentialUsed` **Added**
    - `decoyPubliclyExposedEvidence` **Added**
    - `decoyPubliclyExposed` **Added**
    - `decoyTarget` **Added**
    - `dedupKeyParts` **Added**
    - `description` **Added**
    - `firstObservedAt` **Changed** (Breaking ⚠️)
    - `identityUserTarget` **Added**
    - `lastAppearedAt` **Added**
    - `lastObservedAt` **Changed** (Breaking ⚠️)
    - `nhiUnowned` **Added**
    - `resolvedAt` **Changed** (Breaking ⚠️)
    - `riskAcceptanceExpiresAt` **Changed** (Breaking ⚠️)
    - `riskScore` **Added**
    - `serviceAccountMisclassificationEvidence` **Added**
    - `serviceAccountMisclassification` **Added**
    - `serviceAccountUnowned` **Added**
    - `similarUsernameMatchEvidence` **Added**
    - `similarUsernameMatch` **Added**
    - `snoozeUntil` **Changed** (Breaking ⚠️)
    - `sourceKind` **Added**
    - `tenantTarget` **Added**
    - `unusedSecretEvidence` **Added**
    - `unusedSecret` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.FindingRoutingRule.UpdateFindingRoutingRule()`: 
  * `request.UpdateFindingRoutingRuleRequest` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `routingRule` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `routingRule` **Added**
* `conductoroneAPI.A2UI.ListSurfaces()`: `response.surfaces[]` **Changed** (Breaking ⚠️)
    - `components[].ButtonComponent` **Removed** (Breaking ⚠️)
    - `components[].C1CodeBlockComponent` **Removed** (Breaking ⚠️)
    - `components[].C1ConnectorConfigFormComponent` **Removed** (Breaking ⚠️)
    - `components[].C1ConnectorSyncDetailComponent` **Removed** (Breaking ⚠️)
    - `components[].C1ConnectorSyncProgressComponent` **Removed** (Breaking ⚠️)
    - `components[].C1DurationPickerComponent` **Removed** (Breaking ⚠️)
    - `components[].C1MSTeamsNotificationsComponent` **Removed** (Breaking ⚠️)
    - `components[].C1OnboardingPlanComponent` **Removed** (Breaking ⚠️)
    - `components[].C1OnboardingWelcomeComponent` **Removed** (Breaking ⚠️)
    - `components[].C1ResourcePickerComponent` **Removed** (Breaking ⚠️)
    - `components[].C1SlackNotificationsComponent` **Removed** (Breaking ⚠️)
    - `components[].C1StatusIndicatorComponent` **Removed** (Breaking ⚠️)
    - `components[].C1TodoListComponent` **Removed** (Breaking ⚠️)
    - `components[].CardComponent` **Removed** (Breaking ⚠️)
    - `components[].CheckBoxComponent` **Removed** (Breaking ⚠️)
    - `components[].ChoicePickerComponent` **Removed** (Breaking ⚠️)
    - `components[].ColumnComponent` **Removed** (Breaking ⚠️)
    - `components[].DateTimeInputComponent` **Removed** (Breaking ⚠️)
    - `components[].DividerComponent` **Removed** (Breaking ⚠️)
    - `components[].ProgressBarComponent` **Removed** (Breaking ⚠️)
    - `components[].RowComponent` **Removed** (Breaking ⚠️)
    - `components[].SliderComponent` **Removed** (Breaking ⚠️)
    - `components[].TextComponent` **Removed** (Breaking ⚠️)
    - `components[].TextFieldComponent` **Removed** (Breaking ⚠️)
    - `components[].button` **Added**
    - `components[].c1Chart` **Added**
    - `components[].c1CodeBlock` **Added**
    - `components[].c1ConnectorConfigForm` **Added**
    - `components[].c1ConnectorSyncDetail` **Added**
    - `components[].c1ConnectorSyncProgress` **Added**
    - `components[].c1DurationPicker` **Added**
    - `components[].c1MetricCards` **Added**
    - `components[].c1MsTeamsNotifications` **Added**
    - `components[].c1OnboardingPlan` **Added**
    - `components[].c1OnboardingWelcome` **Added**
    - `components[].c1ResourcePicker` **Added**
    - `components[].c1SlackNotifications` **Added**
    - `components[].c1StatusIndicator` **Added**
    - `components[].c1Table` **Added**
    - `components[].c1TodoList` **Added**
    - `components[].card` **Added**
    - `components[].checkBox` **Added**
    - `components[].choicePicker` **Added**
    - `components[].column` **Added**
    - `components[].dateTimeInput` **Added**
    - `components[].divider` **Added**
    - `components[].progressBar` **Added**
    - `components[].row` **Added**
    - `components[].slider` **Added**
    - `components[].textField` **Added**
    - `components[].text` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `role` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.FindingRoutingRule.GetFindingRoutingRule()`: `response` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `routingRule` **Added**
* `conductoroneAPI.A2UI.ListSurfaceFeedback()`:  `response.feedback[].createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.A2UI.CreateSurfaceFeedback()`: `response` **Changed** (Breaking ⚠️)
    - `A2UISurfaceFeedback` **Removed** (Breaking ⚠️)
    - `feedback` **Added**
* `conductoroneAPI.AccessReview.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `completionDate` **Changed**
    - `expandMask` **Added**
    - `notificationConfig` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_USERS)` **Added**
    - `scopeV2` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewView` **Removed** (Breaking ⚠️)
    - `accessReview` **Added**
* `conductoroneAPI.AccessReview.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewView` **Removed** (Breaking ⚠️)
    - `accessReview` **Added**
* `conductoroneAPI.AccessReview.Update()`: 
  * `request.AccessReviewServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessReview` **Removed** (Breaking ⚠️)
    - `accessReview` **Added**
    - `expandMask` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewView` **Removed** (Breaking ⚠️)
    - `accessReview` **Added**
* `conductoroneAPI.AccessReview.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AccessReview` **Removed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `accessReview` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.GetCampaignScopeAndEntitlements()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `list[].AccessReviewSetupEntitlement` **Removed** (Breaking ⚠️)
    - `list[].accessReviewEntitlement` **Added**
    - `scopeV2` **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.SetCampaignScopeAndEntitlements()`: 
  * `request.AccessReviewSetupEntitlementAndScopeServiceSetRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `AccessReviewSetupEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `expandMask` **Added**
    - `scopeV2` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `list[].AccessReviewSetupEntitlement` **Removed** (Breaking ⚠️)
    - `list[].accessReviewEntitlement` **Added**
    - `scopeV2` **Added**
* `conductoroneAPI.AccessReviewSetupEntitlement.SetCampaignScopeByResourceType()`: 
  * `request.AccessReviewSetScopeByResourceTypeRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `scopeV2` **Added**
* `conductoroneAPI.AccessReviewTemplate.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AccessReviewColumnConfig` **Removed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `NotificationConfig` **Removed** (Breaking ⚠️)
    - `RecurrenceRule` **Removed** (Breaking ⚠️)
    - `ReviewSignatureConfig` **Removed** (Breaking ⚠️)
    - `accessReviewDuration` **Changed**
    - `columnConfig` **Added**
    - `notificationConfig` **Added**
    - `recurrenceRule` **Added**
    - `reviewerAttributeConfig` **Added**
    - `scopeType.enum(ACCESS_REVIEW_SCOPE_TYPE_BY_USERS)` **Added**
    - `scope` **Added**
    - `signatureConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
    - `accessReviewTemplate` **Added**
* `conductoroneAPI.AccessReviewTemplate.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
    - `accessReviewTemplate` **Added**
* `conductoroneAPI.AccessReviewTemplate.Update()`: 
  * `request.AccessReviewTemplateServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
    - `accessReviewTemplate` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewTemplate` **Removed** (Breaking ⚠️)
    - `accessReviewTemplate` **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.GetScopeAndEntitlements()`: `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `list[].AccessReviewTemplateSetupEntitlement` **Removed** (Breaking ⚠️)
    - `list[].accessReviewTemplateEntitlement` **Added**
    - `scope` **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.SetScopeAndEntitlements()`: 
  * `request.AccessReviewTemplateSetupEntitlementServiceSetRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `AccessReviewTemplateSetupEntitlementExpandMask` **Removed** (Breaking ⚠️)
    - `expandMask` **Added**
    - `scope` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `list[].AccessReviewTemplateSetupEntitlement` **Removed** (Breaking ⚠️)
    - `list[].accessReviewTemplateEntitlement` **Added**
    - `scope` **Added**
* `conductoroneAPI.AccessReviewTemplateSetupEntitlement.SetScopeByResourceType()`: 
  * `request.AccessReviewTemplateSetScopeByResourceTypeRequest` **Changed** (Breaking ⚠️)
    - `AccessReviewScopeV2` **Removed** (Breaking ⚠️)
    - `scope` **Added**
* `conductoroneAPI.AccessConflict.CreateMonitor()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `notificationConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `negateGroupB` **Added**
    - `notificationConfig` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AccessConflict.GetMonitor()`: `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `negateGroupB` **Added**
    - `notificationConfig` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AccessConflict.UpdateMonitor()`: 
  * `request.ConflictMonitorUpdateRequest` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `negateGroupB` **Added**
    - `notificationConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AccessConflictNotificationConfig` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `negateGroupB` **Added**
    - `notificationConfig` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppEntitlementMonitorBinding.CreateAppEntitlementMonitorBinding()`: `response` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppEntitlementMonitorBinding.GetAppEntitlementMonitorBinding()`: `response` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Apps.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUserMapper` **Removed** (Breaking ⚠️)
    - `appOwners[].createdAt` **Changed** (Breaking ⚠️)
    - `appOwners[].deletedAt` **Changed** (Breaking ⚠️)
    - `appOwners[].departmentSources[].priority` **Added**
    - `appOwners[].profile` **Changed** (Breaking ⚠️)
    - `appOwners[].updatedAt` **Changed** (Breaking ⚠️)
    - `appUserMapper` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `revokeGrantSources` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Apps.Create()`: `response` **Changed** (Breaking ⚠️)
    - `App` **Removed** (Breaking ⚠️)
    - `app` **Added**
* `conductoroneAPI.Apps.Get()`: `response` **Changed** (Breaking ⚠️)
    - `App` **Removed** (Breaking ⚠️)
    - `app` **Added**
* `conductoroneAPI.AppReport.List()`:  `response.list[].createdAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults()`:  `response.durationGrant` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults()`: 
  *  `request.AppAccessRequestDefaults.durationGrant` **Changed**
  *  `response.durationGrant` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppAccessRequestsDefaults.CancelAppAccessRequestsDefaults()`:  `response.durationGrant` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppUser.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Removed** (Breaking ⚠️)
    - `appUser` **Added**
* `conductoroneAPI.AppUser.ListAppUserCredentials()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `EncryptedData` **Removed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `encryptedData` **Added**
    - `expiresAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppUser.ListAppUsersForUser()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Removed** (Breaking ⚠️)
    - `appUser` **Added**
* `conductoroneAPI.AppUser.Update()`: 
  * `request.AppUserServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AppUser` **Removed** (Breaking ⚠️)
    - `appUser` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppUserView` **Removed** (Breaking ⚠️)
    - `appUserView` **Added**
* `conductoroneAPI.AppUser.Search()`: 
  * `request` **Changed**
    - `agentStatuses` **Added**
    - `appIds` **Added**
    - `excludeDeletedApps` **Added**
    - `expandMask` **Added**
    - `nhiTypes` **Added**
    - `sortBy` **Added**
    - `withOpenFindings` **Added**
    - `withoutResponsibleParty` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Removed** (Breaking ⚠️)
    - `appUser` **Added**
* `conductoroneAPI.Connector.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `Connector` **Removed** (Breaking ⚠️)
    - `connector` **Added**
* `conductoroneAPI.Connector.CreateDelegated()`: 
  * `request.ConnectorServiceCreateDelegatedRequest` **Changed** (Breaking ⚠️)
    - `AppManagedStateBindingRef` **Removed** (Breaking ⚠️)
    - `appManagedStateBindingRef` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Removed** (Breaking ⚠️)
    - `connectorView` **Added**
* `conductoroneAPI.Connector.GetCredentials()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorCredential` **Removed** (Breaking ⚠️)
    - `credential` **Added**
* `conductoroneAPI.Connector.UpdateConnectorSchedule()`: 
  * `request.UpdateConnectorScheduleRequest` **Changed** (Breaking ⚠️)
    - `ConnectorScheduleCron` **Removed** (Breaking ⚠️)
    - `cron` **Added**
* `conductoroneAPI.Connector.Get()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Removed** (Breaking ⚠️)
    - `connectorView` **Added**
* `conductoroneAPI.Connector.Update()`: 
  * `request.ConnectorServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Connector` **Removed** (Breaking ⚠️)
    - `connector` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Removed** (Breaking ⚠️)
    - `connectorView` **Added**
* `conductoroneAPI.Connector.Create()`: 
  *  `request.ConnectorServiceCreateRequest.config` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Removed** (Breaking ⚠️)
    - `connectorView` **Added**
* `conductoroneAPI.Connector.UpdateDelegated()`: 
  * `request.ConnectorServiceUpdateDelegatedRequest` **Changed** (Breaking ⚠️)
    - `Connector` **Removed** (Breaking ⚠️)
    - `connector` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `ConnectorView` **Removed** (Breaking ⚠️)
    - `connectorView` **Added**
* `conductoroneAPI.Connector.RotateCredential()`: `response` **Changed** (Breaking ⚠️)
    - `ConnectorCredential` **Removed** (Breaking ⚠️)
    - `credential` **Added**
* `conductoroneAPI.AppEntitlements.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.AppEntitlements.Create()`: 
  * `request.CreateAppEntitlementRequest` **Changed** (Breaking ⚠️)
    - `ProvisionPolicy` **Removed** (Breaking ⚠️)
    - `durationGrant` **Changed**
    - `provisionPolicy` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
    - `appEntitlementView` **Added**
* `conductoroneAPI.AppEntitlements.GetAutomation()`: `response.AppEntitlementAutomation` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomationLastRunStatus` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleBasic` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleNone` **Removed** (Breaking ⚠️)
    - `basic` **Added**
    - `cel` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `entitlements` **Added**
    - `lastRunStatus` **Added**
    - `none` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppEntitlements.CreateAutomation()`: 
  * `request.CreateAutomationRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomation` **Removed** (Breaking ⚠️)
    - `automation` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomation` **Removed** (Breaking ⚠️)
    - `automation` **Added**
* `conductoroneAPI.AppEntitlements.ListAutomationExclusions()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `user` **Added**
* `conductoroneAPI.AppEntitlements.UpdateAutomation()`: 
  * `request.AppEntitlementServiceUpdateAutomationRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleBasic` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleNone` **Removed** (Breaking ⚠️)
    - `basic` **Added**
    - `cel` **Added**
    - `entitlements` **Added**
    - `none` **Added**
  * `response.AppEntitlementAutomation` **Changed** (Breaking ⚠️)
    - `AppEntitlementAutomationLastRunStatus` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleBasic` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `AppEntitlementAutomationRuleNone` **Removed** (Breaking ⚠️)
    - `basic` **Added**
    - `cel` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `entitlements` **Added**
    - `lastRunStatus` **Added**
    - `none` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppEntitlements.ListUsers()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUserView` **Removed** (Breaking ⚠️)
    - `appEntitlementId` **Added**
    - `appEntitlementUserBindingCreatedAt` **Changed** (Breaking ⚠️)
    - `appEntitlementUserBindingDeprovisionAt` **Changed** (Breaking ⚠️)
    - `appId` **Added**
    - `appUserId` **Added**
    - `appUser` **Added**
* `conductoroneAPI.AppEntitlements.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
    - `appEntitlementView` **Added**
* `conductoroneAPI.AppEntitlements.Update()`: 
  * `request.UpdateAppEntitlementRequest` **Changed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `entitlement` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
    - `appEntitlementView` **Added**
* `conductoroneAPI.AppEntitlements.ListForAppResource()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.AppEntitlements.ListForAppUser()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchAppEntitlementsWithExpired()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppUser` **Removed** (Breaking ⚠️)
    - `User` **Removed** (Breaking ⚠️)
    - `appUser` **Added**
    - `discovered` **Changed** (Breaking ⚠️)
    - `expired` **Changed** (Breaking ⚠️)
    - `grantReasons[].createdAt` **Changed** (Breaking ⚠️)
    - `grantReasons[].deletedAt` **Changed** (Breaking ⚠️)
    - `grantReasons[].reasonExpiresAt` **Changed** (Breaking ⚠️)
    - `grantReasons[].updatedAt` **Changed** (Breaking ⚠️)
    - `user` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchAppEntitlementsForAppUser()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.AppEntitlementSearch.Search()`: `response` **Changed** (Breaking ⚠️)
    - `facets` **Added**
    - `list[].ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `list[].AppEntitlement` **Removed** (Breaking ⚠️)
    - `list[].appEntitlement` **Added**
    - `list[].objectPermissions` **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchGrants()`: 
  *  `request.expandMask` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserView` **Removed** (Breaking ⚠️)
    - `AppEntitlementView` **Removed** (Breaking ⚠️)
    - `appEntitlementUserBinding` **Added**
    - `entitlement` **Added**
* `conductoroneAPI.FindingRoutingRule.CreateFindingRoutingRule()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `routingRule` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `FindingRoutingRule` **Removed** (Breaking ⚠️)
    - `routingRule` **Added**
* `conductoroneAPI.AppEntitlementUserBinding.RemoveGrantDuration()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBinding` **Removed** (Breaking ⚠️)
    - `binding` **Added**
* `conductoroneAPI.AppEntitlementUserBinding.UpdateGrantDuration()`: 
  *  `request.UpdateGrantDurationRequest.newDeprovisionAt` **Changed**
  * `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBinding` **Removed** (Breaking ⚠️)
    - `binding` **Added**
* `conductoroneAPI.AppEntitlementUserBinding.ListAppUsersForIdentityWithGrant()`: `response.bindings[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `deprovisionAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppEntitlementUserBinding.SearchGrantFeed()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingExpandHistoryMask` **Removed** (Breaking ⚠️)
    - `after` **Changed**
    - `before` **Changed**
    - `expandMask` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingFeed` **Removed** (Breaking ⚠️)
    - `feed` **Added**
* `conductoroneAPI.AppEntitlementUserBinding.SearchPastGrants()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingExpandHistoryMask` **Removed** (Breaking ⚠️)
    - `expandMask` **Added**
  * `response.list[]` **Changed** (Breaking ⚠️)
    - `AppEntitlementUserBindingHistory` **Removed** (Breaking ⚠️)
    - `history` **Added**
* `conductoroneAPI.FindingRoutingRule.ListFindingRoutingRules()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `FindingRoutingRuleAction` **Removed** (Breaking ⚠️)
    - `action` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `dispatchers` **Added**
    - `findingType` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Finding.BulkCreateFindingTasks()`: `request` **Changed** (Breaking ⚠️)
    - `FindingSearchRequest` **Removed** (Breaking ⚠️)
    - `searchRequest` **Added**
* `conductoroneAPI.Apps.Update()`: 
  * `request.UpdateAppRequest` **Changed** (Breaking ⚠️)
    - `App` **Removed** (Breaking ⚠️)
    - `app` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `App` **Removed** (Breaking ⚠️)
    - `app` **Added**
* `conductoroneAPI.AppResourceType.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Removed** (Breaking ⚠️)
    - `appResourceType` **Added**
* `conductoroneAPI.AppResourceType.CreateManuallyManagedResourceType()`: 
  * `request.CreateManuallyManagedResourceTypeRequest.resourceType` **Changed**
    - `enum(CLAW_AGENT)` **Added**
    - `enum(SESSION_POLICY)` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Removed** (Breaking ⚠️)
    - `appResourceType` **Added**
* `conductoroneAPI.AppResourceType.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppResourceTypeView` **Removed** (Breaking ⚠️)
    - `appResourceTypeView` **Added**
* `conductoroneAPI.AppResourceType.UpdateManuallyManagedResourceType()`: 
  * `request.UpdateManuallyManagedResourceTypeRequest` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Removed** (Breaking ⚠️)
    - `appResourceType` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceType` **Removed** (Breaking ⚠️)
    - `appResourceType` **Added**
* `conductoroneAPI.AppResource.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppResource` **Removed** (Breaking ⚠️)
    - `appResource` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.AppResource.CreateManuallyManagedAppResource()`: `response` **Changed** (Breaking ⚠️)
    - `AppResource` **Removed** (Breaking ⚠️)
    - `appResource` **Added**
* `conductoroneAPI.AppResource.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppResourceView` **Removed** (Breaking ⚠️)
    - `appResourceView` **Added**
* `conductoroneAPI.AppResource.Update()`: 
  * `request.AppResourceServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `AppResource` **Removed** (Breaking ⚠️)
    - `appResource` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppResourceView` **Removed** (Breaking ⚠️)
    - `appResourceView` **Added**
* `conductoroneAPI.AppResourceOwners.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `departmentSources[].priority` **Added**
    - `profile` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AppUsageControls.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppUsageControls` **Removed** (Breaking ⚠️)
    - `appUsageControls` **Added**
* `conductoroneAPI.AppUsageControls.Update()`: 
  * `request.UpdateAppUsageControlsRequest` **Changed** (Breaking ⚠️)
    - `AppUsageControls` **Removed** (Breaking ⚠️)
    - `appUsageControls` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `AppUsageControls` **Removed** (Breaking ⚠️)
    - `appUsageControls` **Added**
* `conductoroneAPI.AppEntitlementsProxy.Get()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementProxyView` **Removed** (Breaking ⚠️)
    - `appProxyEntitlementView` **Added**
* `conductoroneAPI.AppEntitlementsProxy.Create()`: `response` **Changed** (Breaking ⚠️)
    - `AppEntitlementProxyView` **Removed** (Breaking ⚠️)
    - `appProxyEntitlementView` **Added**
* `conductoroneAPI.Attributes.CreateAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `value` **Added**
* `conductoroneAPI.Attributes.GetAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `value` **Added**
* `conductoroneAPI.Attributes.ListComplianceFrameworks()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Attributes.CreateComplianceFrameworkAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `value` **Added**
* `conductoroneAPI.Attributes.GetComplianceFrameworkAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `value` **Added**
* `conductoroneAPI.Attributes.ListRiskLevels()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Attributes.CreateRiskLevelAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `value` **Added**
* `conductoroneAPI.Attributes.GetRiskLevelAttributeValue()`: `response` **Changed** (Breaking ⚠️)
    - `AttributeValue` **Removed** (Breaking ⚠️)
    - `value` **Added**
* `conductoroneAPI.Attributes.ListAttributeValues()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.TenantAuthConfig.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AuthConfigC1Local` **Removed** (Breaking ⚠️)
    - `AuthConfigGoogle` **Removed** (Breaking ⚠️)
    - `AuthConfigJumpCloud` **Removed** (Breaking ⚠️)
    - `AuthConfigMicrosoft` **Removed** (Breaking ⚠️)
    - `AuthConfigOIDC` **Removed** (Breaking ⚠️)
    - `AuthConfigOkta` **Removed** (Breaking ⚠️)
    - `AuthConfigOneLogin` **Removed** (Breaking ⚠️)
    - `AuthConfigPingOne` **Removed** (Breaking ⚠️)
    - `c1Local` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deprecationDeadline` **Changed** (Breaking ⚠️)
    - `google` **Added**
    - `jumpcloud` **Added**
    - `microsoft` **Added**
    - `oidc` **Added**
    - `okta` **Added**
    - `onelogin` **Added**
    - `pingone` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.TenantAuthConfig.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AuthConfigC1Local` **Removed** (Breaking ⚠️)
    - `AuthConfigGoogle` **Removed** (Breaking ⚠️)
    - `AuthConfigJumpCloud` **Removed** (Breaking ⚠️)
    - `AuthConfigMicrosoft` **Removed** (Breaking ⚠️)
    - `AuthConfigOIDC` **Removed** (Breaking ⚠️)
    - `AuthConfigOkta` **Removed** (Breaking ⚠️)
    - `AuthConfigOneLogin` **Removed** (Breaking ⚠️)
    - `AuthConfigPingOne` **Removed** (Breaking ⚠️)
    - `c1Local` **Added**
    - `deprecationDeadline` **Changed**
    - `google` **Added**
    - `jumpcloud` **Added**
    - `microsoft` **Added**
    - `oidc` **Added**
    - `okta` **Added**
    - `onelogin` **Added**
    - `pingone` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
    - `authConfig` **Added**
* `conductoroneAPI.TenantAuthConfig.Get()`: `response` **Changed** (Breaking ⚠️)
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
    - `authConfig` **Added**
* `conductoroneAPI.TenantAuthConfig.Update()`: 
  * `request.TenantAuthConfigServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
    - `authConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `TenantAuthConfig` **Removed** (Breaking ⚠️)
    - `authConfig` **Added**
* `conductoroneAPI.Finding.BulkUpdateFindingState()`: `request` **Changed** (Breaking ⚠️)
    - `BulkAcceptRiskAction` **Removed** (Breaking ⚠️)
    - `BulkAssignOwnerAction` **Removed** (Breaking ⚠️)
    - `BulkReopenAction` **Removed** (Breaking ⚠️)
    - `BulkSnoozeAction` **Removed** (Breaking ⚠️)
    - `BulkSuppressAction` **Removed** (Breaking ⚠️)
    - `BulkUnsuppressAction` **Removed** (Breaking ⚠️)
    - `FindingSearchRequest` **Removed** (Breaking ⚠️)
    - `acceptRisk` **Added**
    - `assignOwner` **Added**
    - `reopen` **Added**
    - `searchRequest` **Added**
    - `snooze` **Added**
    - `suppress` **Added**
    - `unsuppress` **Added**
* `conductoroneAPI.AutomationExecution.ListAutomationExecutions()`: `response.automationExecutions[]` **Changed** (Breaking ⚠️)
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `completedAt` **Changed** (Breaking ⚠️)
    - `context` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `duration` **Changed** (Breaking ⚠️)
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.AutomationExecution.GetAutomationExecution()`: `response` **Changed** (Breaking ⚠️)
    - `AutomationExecutionView` **Removed** (Breaking ⚠️)
    - `AutomationExecution` **Removed** (Breaking ⚠️)
    - `automationExecution` **Added**
    - `view` **Added**
* `conductoroneAPI.Automation.ListAutomations()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `DisabledReasonCircuitBreaker` **Removed** (Breaking ⚠️)
    - `automationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `automationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `automationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `automationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `automationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `automationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `automationSteps[].accountLifecycleAction` **Added**
    - `automationSteps[].callFunction` **Added**
    - `automationSteps[].connectorAction` **Added**
    - `automationSteps[].connectorCreateAccount` **Added**
    - `automationSteps[].createAccessReview` **Added**
    - `automationSteps[].createRevokeTasksV2` **Added**
    - `automationSteps[].createRevokeTasks` **Added**
    - `automationSteps[].evaluateExpressions` **Added**
    - `automationSteps[].generatePassword` **Added**
    - `automationSteps[].grantEntitlements` **Added**
    - `automationSteps[].removeFromDelegation` **Added**
    - `automationSteps[].runAutomation` **Added**
    - `automationSteps[].sendEmail` **Added**
    - `automationSteps[].sendSlackMessage` **Added**
    - `automationSteps[].setCredential` **Added**
    - `automationSteps[].storeCredential` **Added**
    - `automationSteps[].taskAction` **Added**
    - `automationSteps[].unenrollFromAllAccessProfiles` **Added**
    - `automationSteps[].updateUser` **Added**
    - `automationSteps[].waitForDuration` **Added**
    - `automationSteps[].webhook` **Added**
    - `circuitBreaker` **Added**
    - `context` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `draftTriggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].accessConflict` **Added**
    - `draftTriggers[].appUserCreated` **Added**
    - `draftTriggers[].appUserUpdated` **Added**
    - `draftTriggers[].grantDeleted` **Added**
    - `draftTriggers[].grantFound` **Added**
    - `draftTriggers[].scheduleAppUser` **Added**
    - `draftTriggers[].scheduleNoUser` **Added**
    - `draftTriggers[].schedule` **Added**
    - `draftTriggers[].usageBasedRevocation` **Added**
    - `draftTriggers[].userCreated` **Added**
    - `draftTriggers[].userProfileChange` **Added**
    - `draftTriggers[].webhook` **Added**
    - `lastExecutedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Automation.CreateAutomation()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `automationSteps[].AccountLifecycleAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].CallFunction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].ConnectorCreateAccount` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateAccessReview` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasksV2` **Removed** (Breaking ⚠️)
    - `automationSteps[].CreateRevokeTasks` **Removed** (Breaking ⚠️)
    - `automationSteps[].EvaluateExpressions` **Removed** (Breaking ⚠️)
    - `automationSteps[].GeneratePassword` **Removed** (Breaking ⚠️)
    - `automationSteps[].GrantEntitlements` **Removed** (Breaking ⚠️)
    - `automationSteps[].RemoveFromDelegation` **Removed** (Breaking ⚠️)
    - `automationSteps[].RunAutomation` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendEmail` **Removed** (Breaking ⚠️)
    - `automationSteps[].SendSlackMessage` **Removed** (Breaking ⚠️)
    - `automationSteps[].SetCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].StoreCredential` **Removed** (Breaking ⚠️)
    - `automationSteps[].TaskAction` **Removed** (Breaking ⚠️)
    - `automationSteps[].UnenrollFromAllAccessProfiles` **Removed** (Breaking ⚠️)
    - `automationSteps[].UpdateUser` **Removed** (Breaking ⚠️)
    - `automationSteps[].WaitForDuration` **Removed** (Breaking ⚠️)
    - `automationSteps[].Webhook` **Removed** (Breaking ⚠️)
    - `automationSteps[].accountLifecycleAction` **Added**
    - `automationSteps[].callFunction` **Added**
    - `automationSteps[].connectorAction` **Added**
    - `automationSteps[].connectorCreateAccount` **Added**
    - `automationSteps[].createAccessReview` **Added**
    - `automationSteps[].createRevokeTasksV2` **Added**
    - `automationSteps[].createRevokeTasks` **Added**
    - `automationSteps[].evaluateExpressions` **Added**
    - `automationSteps[].generatePassword` **Added**
    - `automationSteps[].grantEntitlements` **Added**
    - `automationSteps[].removeFromDelegation` **Added**
    - `automationSteps[].runAutomation` **Added**
    - `automationSteps[].sendEmail` **Added**
    - `automationSteps[].sendSlackMessage` **Added**
    - `automationSteps[].setCredential` **Added**
    - `automationSteps[].storeCredential` **Added**
    - `automationSteps[].taskAction` **Added**
    - `automationSteps[].unenrollFromAllAccessProfiles` **Added**
    - `automationSteps[].updateUser` **Added**
    - `automationSteps[].waitForDuration` **Added**
    - `automationSteps[].webhook` **Added**
    - `context` **Added**
    - `draftTriggers[].AccessConflictTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].AppUserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].AppUserUpdatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantDeletedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].GrantFoundTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerAppUser` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTriggerNoUser` **Removed** (Breaking ⚠️)
    - `draftTriggers[].ScheduleTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UsageBasedRevocationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UserCreatedTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].UserProfileChangeTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].WebhookAutomationTrigger` **Removed** (Breaking ⚠️)
    - `draftTriggers[].accessConflict` **Added**
    - `draftTriggers[].appUserCreated` **Added**
    - `draftTriggers[].appUserUpdated` **Added**
    - `draftTriggers[].grantDeleted` **Added**
    - `draftTriggers[].grantFound` **Added**
    - `draftTriggers[].scheduleAppUser` **Added**
    - `draftTriggers[].scheduleNoUser` **Added**
    - `draftTriggers[].schedule` **Added**
    - `draftTriggers[].usageBasedRevocation` **Added**
    - `draftTriggers[].userCreated` **Added**
    - `draftTriggers[].userProfileChange` **Added**
    - `draftTriggers[].webhook` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation` **Removed** (Breaking ⚠️)
    - `automation` **Added**
* `conductoroneAPI.Automation.GetAutomation()`: `response` **Changed** (Breaking ⚠️)
    - `Automation` **Removed** (Breaking ⚠️)
    - `automation` **Added**
* `conductoroneAPI.Automation.UpdateAutomation()`: 
  * `request.UpdateAutomationRequest` **Changed** (Breaking ⚠️)
    - `Automation` **Removed** (Breaking ⚠️)
    - `automation` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Automation` **Removed** (Breaking ⚠️)
    - `automation` **Added**
* `conductoroneAPI.Automation.ExecuteAutomation()`: 
  * `request.ExecuteAutomationRequest` **Changed** (Breaking ⚠️)
    - `AutomationContext` **Removed** (Breaking ⚠️)
    - `context` **Added**
* `conductoroneAPI.RequestCatalogManagement.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `RequestCatalog` **Removed** (Breaking ⚠️)
    - `requestCatalog` **Added**
* `conductoroneAPI.RequestCatalogManagement.Create()`: `response` **Changed** (Breaking ⚠️)
    - `RequestCatalogView` **Removed** (Breaking ⚠️)
    - `requestCatalogView` **Added**
* `conductoroneAPI.RequestCatalogManagement.ListEntitlementsPerCatalog()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.RequestCatalogManagement.GetRequestableEntry()`: `response` **Changed** (Breaking ⚠️)
    - `RequestableEntry` **Removed** (Breaking ⚠️)
    - `requestableEntry` **Added**
* `conductoroneAPI.RequestCatalogManagement.CreateRequestableEntry()`: `response` **Changed** (Breaking ⚠️)
    - `RequestableEntry` **Removed** (Breaking ⚠️)
    - `requestableEntry` **Added**
* `conductoroneAPI.RequestCatalogManagement.ListEntitlementsForAccess()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `ActorObjectPermissions` **Removed** (Breaking ⚠️)
    - `AppEntitlement` **Removed** (Breaking ⚠️)
    - `appEntitlement` **Added**
    - `objectPermissions` **Added**
* `conductoroneAPI.RequestCatalogManagement.Get()`: `response` **Changed** (Breaking ⚠️)
    - `RequestCatalogView` **Removed** (Breaking ⚠️)
    - `requestCatalogView` **Added**
* `conductoroneAPI.RequestCatalogManagement.Update()`: 
  * `request.RequestCatalogManagementServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `RequestCatalog` **Removed** (Breaking ⚠️)
    - `catalog` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `RequestCatalogView` **Removed** (Breaking ⚠️)
    - `requestCatalogView` **Added**
* `conductoroneAPI.RequestCatalogManagement.GetBundleAutomation()`: `response` **Changed** (Breaking ⚠️)
    - `BundleAutomationCircuitBreaker` **Removed** (Breaking ⚠️)
    - `BundleAutomationLastRunState` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `cel` **Added**
    - `circuitBreaker` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `entitlements` **Added**
    - `state` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.RequestCatalogManagement.SetBundleAutomation()`: 
  * `request.SetBundleAutomationRequest` **Changed** (Breaking ⚠️)
    - `BundleAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `cel` **Added**
    - `entitlements` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `BundleAutomationCircuitBreaker` **Removed** (Breaking ⚠️)
    - `BundleAutomationLastRunState` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `cel` **Added**
    - `circuitBreaker` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `entitlements` **Added**
    - `state` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.RequestCatalogManagement.CreateBundleAutomation()`: 
  * `request.CreateBundleAutomationRequest` **Changed** (Breaking ⚠️)
    - `BundleAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `cel` **Added**
    - `entitlements` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `BundleAutomationCircuitBreaker` **Removed** (Breaking ⚠️)
    - `BundleAutomationLastRunState` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleCEL` **Removed** (Breaking ⚠️)
    - `BundleAutomationRuleEntitlement` **Removed** (Breaking ⚠️)
    - `cel` **Added**
    - `circuitBreaker` **Added**
    - `createdAt` **Changed** (Breaking ⚠️)
    - `deletedAt` **Changed** (Breaking ⚠️)
    - `entitlements` **Added**
    - `state` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.ConnectorCatalog.ConfigurationSchema()`: `response` **Changed** (Breaking ⚠️)
    - `ConfigSchema` **Removed** (Breaking ⚠️)
    - `RequestSchemaForm` **Removed** (Breaking ⚠️)
    - `formSchema` **Added**
    - `schema` **Added**
* `conductoroneAPI.Decoy.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `createdAt` **Changed** (Breaking ⚠️)
    - `lastUsedAt` **Added**
    - `materialFingerprintSha256` **Added**
    - `updatedAt` **Changed** (Breaking ⚠️)
* `conductoroneAPI.Decoy.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `DecoyAccessTokenInput` **Removed** (Breaking ⚠️)
    - `DecoyConnectorClientInput` **Removed** (Breaking ⚠️)
    - `DecoyUserClientCredentialInput` **Removed** (Breaking ⚠️)
    - `DecoyWorkloadFederationInput` **Removed** (Breaking ⚠️)
    - `accessToken` **Added**
    - `connectorClient` **Added**
    - `userClientCredential` **Added**
    - `workloadFed` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `DecoyVendingMaterial` **Removed** (Breaking ⚠️)
    - `Decoy` **Removed** (Breaking ⚠️)
    - `decoy` **Added**
    - `material` **Added**
* `conductoroneAPI.Decoy.Get()`: `response` **Changed** (Breaking ⚠️)
    - `Decoy` **Removed** (Breaking ⚠️)
    - `decoy` **Added**
* `conductoroneAPI.Decoy.Update()`: 
  * `request.DecoyServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `Decoy` **Removed** (Breaking ⚠️)
    - `decoy` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Decoy` **Removed** (Breaking ⚠️)
    - `decoy` **Added**
* `conductoroneAPI.Decoy.Rotate()`: `response` **Changed** (Breaking ⚠️)
    - `DecoyVendingMaterial` **Removed** (Breaking ⚠️)
    - `Decoy` **Removed** (Breaking ⚠️)
    - `decoy` **Added**
    - `material` **Added**
* `conductoroneAPI.Directory.List()`: `response.list[]` **Changed** (Breaking ⚠️)
    - `Directory` **Removed** (Breaking ⚠️)
    - `directory` **Added**
* `conductoroneAPI.Directory.Create()`: 
  * `request` **Changed** (Breaking ⚠️)
    - `DirectoryAccountFilterAll` **Removed** (Breaking ⚠️)
    - `DirectoryAccountFilterCel` **Removed** (Breaking ⚠️)
    - `DirectoryMergeConfig` **Removed** (Breaking ⚠️)
    - `all` **Added**
    - `celExpression` **Added**
    - `mergeConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `DirectoryView` **Removed** (Breaking ⚠️)
    - `directoryView` **Added**
* `conductoroneAPI.Directory.Get()`: `response` **Changed** (Breaking ⚠️)
    - `DirectoryView` **Removed** (Breaking ⚠️)
    - `directoryView` **Added**
* `conductoroneAPI.Directory.Update()`: 
  * `request.DirectoryServiceUpdateRequest` **Changed** (Breaking ⚠️)
    - `DirectoryAccountFilterAll` **Removed** (Breaking ⚠️)
    - `DirectoryAccountFilterCel` **Removed** (Breaking ⚠️)
    - `DirectoryMergeConfig` **Removed** (Breaking ⚠️)
    - `all` **Added**
    - `celExpression` **Added**
    - `expandMask` **Added**
    - `mergeConfig` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `DirectoryView` **Removed** (Breaking ⚠️)
    - `directoryView` **Added**
* `conductoroneAPI.Finding.UpdateFindingState()`: 
  * `request.UpdateFindingStateRequest` **Changed** (Breaking ⚠️)
    - `AcceptRiskAction` **Removed** (Breaking ⚠️)
    - `ReopenAction` **Removed** (Breaking ⚠️)
    - `ResolveAction` **Removed** (Breaking ⚠️)
    - `SnoozeAction` **Removed** (Breaking ⚠️)
    - `SuppressStateAction` **Removed** (Breaking ⚠️)
    - `UnsuppressAction` **Removed** (Breaking ⚠️)
    - `acceptRisk` **Added**
    - `reopen` **Added**
    - `resolve` **Added**
    - `snooze` **Added**
    - `suppress` **Added**
    - `unsuppress` **Added**
  * `response` **Changed** (Breaking ⚠️)
    - `Finding` **Removed** (Breaking ⚠️)
    - `finding` **Added**
* `conductoroneAPI.Finding.CreateFindingTask()`: `response` **Changed** (Breaking ⚠️)
    - `Finding` **Removed** (Breaking ⚠️)
    - `finding` **Added**
* `conductoroneAPI.RecoveryPolicy.Create()`: **Added**
* `conductoroneAPI.Auth.Introspect()`: `response` **Changed**
    - `deviceClientId` **Added**
    - `disabledModules` **Added**
    - `tenantId` **Added**
* `conductoroneAPI.MCPAccessProfile.SearchAccessProfiles()`: **Added**
* `conductoroneAPI.MCPAccessProfileToolBinding.GetAccessProfilesForTools()`: **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchGraph()`: 
  *  `request.maxNodes` **Added**
  * `response` **Changed**
    - `nodeCeilingHit` **Added**
    - `nodes[].secondaryText` **Added**
* `conductoroneAPI.A2UI.SubmitAction()`: 
  *  `request.A2UIServiceSubmitActionRequest.clientTimestamp` **Changed**
* `conductoroneAPI.UserOwnersV2.CreateUserOwner()`: **Added**
* `conductoroneAPI.UserOwnersV2.DeleteUserOwner()`: **Added**
* `conductoroneAPI.UserOwnersV2.SearchUserOwners()`: **Added**
* `conductoroneAPI.UserOwnersV2.CreateEntitlementOwner()`: **Added**
* `conductoroneAPI.UserOwnersV2.DeleteEntitlementOwner()`: **Added**
* `conductoroneAPI.UserOwnersV2.SearchEntitlementOwners()`: **Added**
* `conductoroneAPI.UserOwnersV2.Set()`: **Added**
* `conductoroneAPI.AppUserOwnersV2.CreateUserOwner()`: **Added**
* `conductoroneAPI.AppUserOwnersV2.DeleteUserOwner()`: **Added**
* `conductoroneAPI.AppUserOwnersV2.SearchUserOwners()`: **Added**
* `conductoroneAPI.AppUserOwnersV2.CreateEntitlementOwner()`: **Added**
* `conductoroneAPI.AppUserOwnersV2.DeleteEntitlementOwner()`: **Added**
* `conductoroneAPI.AppUserOwnersV2.SearchEntitlementOwners()`: **Added**
* `conductoroneAPI.AppUserOwnersV2.Set()`: **Added**
* `conductoroneAPI.AppResourceOwnersV2.GetUserOwner()`: **Added**
* `conductoroneAPI.AppResourceOwnersV2.GetEntitlementOwner()`: **Added**
* `conductoroneAPI.User.Introspect()`: **Added**
* `conductoroneAPI.TaskActions.RetryProvisioning()`: **Added**
* `conductoroneAPI.Task.CreateResourceActionTask()`: **Added**
* `conductoroneAPI.Task.CreateActionTask()`: **Added**
* `conductoroneAPI.SSOSettings.ListHistory()`: **Added**
* `conductoroneAPI.SSOSettings.Update()`: **Added**
* `conductoroneAPI.SSOSettings.Get()`: **Added**
* `conductoroneAPI.IdentityPolicyTenantDefaults.Update()`: **Added**
* `conductoroneAPI.IdentityPolicyTenantDefaults.Get()`: **Added**
* `conductoroneAPI.XAASettings.ListHistory()`: **Added**
* `conductoroneAPI.XAASettings.Update()`: **Added**
* `conductoroneAPI.XAASettings.Get()`: **Added**
* `conductoroneAPI.AIGovernanceSettings.GetTenantDefaults()`: **Added**
* `conductoroneAPI.AIGovernanceSettings.ListHistory()`: **Added**
* `conductoroneAPI.AIGovernanceSettings.Update()`: **Added**
* `conductoroneAPI.AIGovernanceSettings.Get()`: **Added**
* `conductoroneAPI.XAAClientAudienceMapping.ListHistory()`: **Added**
* `conductoroneAPI.XAAClientAudienceMapping.Delete()`: **Added**
* `conductoroneAPI.XAAClientAudienceMapping.Get()`: **Added**
* `conductoroneAPI.XAAClientAudienceMapping.Create()`: **Added**
* `conductoroneAPI.XAAClientAudienceMapping.List()`: **Added**
* `conductoroneAPI.XAAClientAudienceMapping.Update()`: **Added**
* `conductoroneAPI.XAAClientAudienceMapping.Search()`: **Added**
* `conductoroneAPI.SignInPolicy.Update()`: **Added**
* `conductoroneAPI.SignInPolicy.Get()`: **Added**
* `conductoroneAPI.SignInPolicy.Delete()`: **Added**
* `conductoroneAPI.SignInPolicy.Create()`: **Added**
* `conductoroneAPI.SignInPolicy.List()`: **Added**
* `conductoroneAPI.SignInPolicy.Search()`: **Added**
* `conductoroneAPI.SessionPolicy.UnassignUser()`: **Added**
* `conductoroneAPI.SessionPolicy.AssignUser()`: **Added**
* `conductoroneAPI.SessionPolicy.UnassignGroup()`: **Added**
* `conductoroneAPI.SessionPolicy.AssignGroup()`: **Added**
* `conductoroneAPI.SessionPolicy.ListAssignments()`: **Added**
* `conductoroneAPI.SessionPolicy.Update()`: **Added**
* `conductoroneAPI.SessionPolicy.Get()`: **Added**
* `conductoroneAPI.SessionPolicy.Delete()`: **Added**
* `conductoroneAPI.SessionPolicy.Create()`: **Added**
* `conductoroneAPI.SessionPolicy.List()`: **Added**
* `conductoroneAPI.SessionPolicy.Search()`: **Added**
* `conductoroneAPI.FindingAudit.Search()`: **Added**
* `conductoroneAPI.Reporting.GetRunProvenance()`: **Added**
* `conductoroneAPI.Reporting.Run()`: **Added**
* `conductoroneAPI.Reporting.Update()`: **Added**
* `conductoroneAPI.Reporting.Get()`: **Added**
* `conductoroneAPI.Reporting.Delete()`: **Added**
* `conductoroneAPI.Reporting.Save()`: **Added**
* `conductoroneAPI.Reporting.List()`: **Added**
* `conductoroneAPI.RecoveryPolicy.Search()`: **Added**
* `conductoroneAPI.RecoveryPolicy.Update()`: **Added**
* `conductoroneAPI.RecoveryPolicy.Get()`: **Added**
* `conductoroneAPI.RecoveryPolicy.Delete()`: **Added**
* `conductoroneAPI.A2UI.GetSurfaceProvenance()`: **Added**
* `conductoroneAPI.AccessReviewReport.List()`: **Added**
* `conductoroneAPI.RecoveryPolicy.List()`: **Added**
* `conductoroneAPI.ProviderCredential.Set()`: **Added**
* `conductoroneAPI.ProviderCredential.Get()`: **Added**
* `conductoroneAPI.ProviderCredential.Clear()`: **Added**
* `conductoroneAPI.GatewayKey.Revoke()`: **Added**
* `conductoroneAPI.GatewayKey.Mint()`: **Added**
* `conductoroneAPI.GatewayKey.List()`: **Added**
* `conductoroneAPI.TunnelCredentials.RevokeBridgeCredential()`: **Added**
* `conductoroneAPI.TunnelCredentials.UpdateBridge()`: **Added**
* `conductoroneAPI.TunnelCredentials.GetBridge()`: **Added**
* `conductoroneAPI.TunnelCredentials.DeleteBridge()`: **Added**
* `conductoroneAPI.TunnelCredentials.CreateBridgeCredential()`: **Added**
* `conductoroneAPI.TunnelCredentials.ListBridgeCredentials()`: **Added**
* `conductoroneAPI.TunnelCredentials.ListBridgeAnnouncedServices()`: **Added**
* `conductoroneAPI.TunnelCredentials.CreateBridge()`: **Added**
* `conductoroneAPI.TunnelCredentials.ListBridges()`: **Added**
* `conductoroneAPI.PersonalDevice.Search()`: **Added**
* `conductoroneAPI.PersonalDevice.RevokeDeviceClient()`: **Added**
* `conductoroneAPI.PersonalDevice.ListDeviceClients()`: **Added**
* `conductoroneAPI.PersonalDevice.UpdateDevice()`: **Added**
* `conductoroneAPI.PersonalDevice.GetDevice()`: **Added**
* `conductoroneAPI.PersonalDevice.RevokeDevice()`: **Added**
* `conductoroneAPI.FindingTransformationRule.UpdateFindingTransformationRule()`: **Added**
* `conductoroneAPI.FindingTransformationRule.GetFindingTransformationRule()`: **Added**
* `conductoroneAPI.FindingTransformationRule.DeleteFindingTransformationRule()`: **Added**
* `conductoroneAPI.FindingTransformationRule.CreateFindingTransformationRule()`: **Added**
* `conductoroneAPI.FindingTransformationRule.ListFindingTransformationRules()`: **Added**
* `conductoroneAPI.FindingSettings.UpdateFindingSettings()`: **Added**
* `conductoroneAPI.FindingSettings.ListFindingSettings()`: **Added**
* `conductoroneAPI.Finding.CreateFinding()`: **Added**
* `conductoroneAPI.DecoySearch.Search()`: **Added**
* `conductoroneAPI.CredentialInventoryPolicy.Search()`: **Added**
* `conductoroneAPI.CredentialInventoryPolicy.Update()`: **Added**
* `conductoroneAPI.CredentialInventoryPolicy.Get()`: **Added**
* `conductoroneAPI.CredentialInventoryPolicy.Delete()`: **Added**
* `conductoroneAPI.CredentialInventoryPolicy.Create()`: **Added**
* `conductoroneAPI.CredentialInventoryPolicy.List()`: **Added**
* `conductoroneAPI.UIConversations.EnsureOnboardingSession()`: **Added**
* `conductoroneAPI.ConnectorAuthoringActivation.RollbackRevision()`: **Added**
* `conductoroneAPI.ConnectorAuthoringActivation.ActivateRevision()`: **Added**
* `conductoroneAPI.XAAScope.Search()`: **Added**
* `conductoroneAPI.XAAScope.ListHistory()`: **Added**
* `conductoroneAPI.XAAScope.Update()`: **Added**
* `conductoroneAPI.XAAScope.Get()`: **Added**
* `conductoroneAPI.XAAScope.Delete()`: **Added**
* `conductoroneAPI.XAAScope.Create()`: **Added**
* `conductoroneAPI.XAAScope.List()`: **Added**
* `conductoroneAPI.XAAResourceServer.Search()`: **Added**
* `conductoroneAPI.XAAResourceServer.ListHistory()`: **Added**
* `conductoroneAPI.XAAResourceServer.Update()`: **Added**
* `conductoroneAPI.XAAResourceServer.Get()`: **Added**
* `conductoroneAPI.XAAResourceServer.Delete()`: **Added**
* `conductoroneAPI.XAAResourceServer.Create()`: **Added**
* `conductoroneAPI.XAAResourceServer.List()`: **Added**
* `conductoroneAPI.XAAAccessProfileScopeBinding.Search()`: **Added**
* `conductoroneAPI.XAAAccessProfileScopeBinding.DeleteBindings()`: **Added**
* `conductoroneAPI.OnboardingSettings.Update()`: 
  * `request` **Changed**
    - `mcpOnboardingGoal` **Added**
    - `mcpOnboardingStatus` **Added**
    - `mcpOnboardingTargets` **Added**
  * `response` **Changed**
    - `mcpOnboardingGoal` **Added**
    - `mcpOnboardingStatus` **Added**
    - `mcpOnboardingTargets` **Added**
* `conductoroneAPI.XAAAccessProfileScopeBinding.CreateBindings()`: **Added**
* `conductoroneAPI.XAAAccessProfileScopeBinding.List()`: **Added**
* `conductoroneAPI.XAAAccessProfile.Search()`: **Added**
* `conductoroneAPI.XAAAccessProfile.GetByAppEntitlementId()`: **Added**
* `conductoroneAPI.XAAAccessProfile.ListHistory()`: **Added**
* `conductoroneAPI.XAAAccessProfile.Update()`: **Added**
* `conductoroneAPI.XAAAccessProfile.Get()`: **Added**
* `conductoroneAPI.XAAAccessProfile.Delete()`: **Added**
* `conductoroneAPI.XAAAccessProfile.Create()`: **Added**
* `conductoroneAPI.XAAAccessProfile.List()`: **Added**
* `conductoroneAPI.SSOApplication.ParseSAMLServiceProviderMetadata()`: **Added**
* `conductoroneAPI.SystemLog.ListEvents()`: `request` **Changed**
    - `since` **Changed**
    - `until` **Changed**
* `conductoroneAPI.SSOApplication.Search()`: **Added**
* `conductoroneAPI.SSOApplication.BatchImportSubjectCompatibility()`: **Added**
* `conductoroneAPI.SSOApplication.BatchDeleteSubjectCompatibility()`: **Added**
* `conductoroneAPI.SSOApplication.ListHistory()`: **Added**
* `conductoroneAPI.SSOApplication.UpdateClient()`: **Added**
* `conductoroneAPI.SSOApplication.RotateClientSecret()`: **Added**
* `conductoroneAPI.SSOApplication.DeleteClient()`: **Added**
* `conductoroneAPI.SSOApplication.CreateClient()`: **Added**
* `conductoroneAPI.SSOApplication.ListClients()`: **Added**
* `conductoroneAPI.SSOApplication.Update()`: **Added**
* `conductoroneAPI.SSOApplication.Get()`: **Added**
* `conductoroneAPI.SSOApplication.Delete()`: **Added**
* `conductoroneAPI.SSOApplication.Create()`: **Added**
* `conductoroneAPI.SSOApplication.List()`: **Added**
* `conductoroneAPI.AppManagedState.Promote()`: **Added**
* `conductoroneAPI.AppManagedState.Get()`: **Added**
* `conductoroneAPI.AppManagedState.List()`: **Added**
* `conductoroneAPI.MCPServer.TestConnection()`: **Added**
* `conductoroneAPI.MCPServer.DiscoverOIDCEndpoints()`: **Added**
* `conductoroneAPI.MCPServer.ListConnections()`: **Added**
* `conductoroneAPI.MCPServer.GetCatalog()`: **Added**
* `conductoroneAPI.MCPServer.ListCatalog()`: **Added**
* `conductoroneAPI.MCPServer.SearchWithToolCount()`: **Added**
* `conductoroneAPI.MCPServer.ResyncTools()`: **Added**
* `conductoroneAPI.MCPServer.UpdateCredentials()`: **Added**
* `conductoroneAPI.MCPServer.Update()`: **Added**
* `conductoroneAPI.MCPServer.Get()`: **Added**
* `conductoroneAPI.MCPServer.Delete()`: **Added**
* `conductoroneAPI.MCPServer.Register()`: **Added**
* `conductoroneAPI.MCPServer.List()`: **Added**
* `conductoroneAPI.AppEntitlementSearch.SearchReachableResourcesForUser()`: **Added**
* `conductoroneAPI.AppEntitlementSearch.CountGrantsForUserByApp()`: **Added**
* `conductoroneAPI.AppEntitlementRoutingRule.ReorderAppEntitlementRoutingRules()`: **Added**
* `conductoroneAPI.AppEntitlementRoutingRule.UpdateAppEntitlementRoutingRule()`: **Added**
* `conductoroneAPI.AppEntitlementRoutingRule.GetAppEntitlementRoutingRule()`: **Added**
* `conductoroneAPI.AppEntitlementRoutingRule.DeleteAppEntitlementRoutingRule()`: **Added**
* `conductoroneAPI.AppEntitlementRoutingRule.CreateAppEntitlementRoutingRule()`: **Added**
* `conductoroneAPI.AppEntitlementRoutingRule.ListAppEntitlementRoutingRules()`: **Added**
* `conductoroneAPI.MCPAccessProfileToolBinding.ListProfilesByToolHistory()`: **Added**
* `conductoroneAPI.MCPAccessProfileToolBinding.ListToolsByProfileHistory()`: **Added**
* `conductoroneAPI.MCPAccessProfileToolBinding.DeleteBindings()`: **Added**
* `conductoroneAPI.MCPAccessProfileToolBinding.CreateBindings()`: **Added**
* `conductoroneAPI.MCPAccessProfileToolBinding.List()`: **Added**
* `conductoroneAPI.MCPAccessProfile.SearchRequestableConnectors()`: **Added**
* `conductoroneAPI.MCPAccessProfile.ListRequestableConnectors()`: **Added**
* `conductoroneAPI.MCPAccessProfile.GetByAppEntitlementId()`: **Added**
* `conductoroneAPI.MCPAccessProfile.Update()`: **Added**
* `conductoroneAPI.MCPAccessProfile.Get()`: **Added**
* `conductoroneAPI.MCPAccessProfile.Delete()`: **Added**
* `conductoroneAPI.MCPAccessProfile.Create()`: **Added**
* `conductoroneAPI.MCPAccessProfile.List()`: **Added**
* `conductoroneAPI.MCPTool.Search()`: **Added**
* `conductoroneAPI.MCPTool.ListHistory()`: **Added**
* `conductoroneAPI.MCPTool.Update()`: **Added**
* `conductoroneAPI.MCPTool.Get()`: **Added**
* `conductoroneAPI.MCPTool.Delete()`: **Added**
* `conductoroneAPI.MCPTool.List()`: **Added**
* `conductoroneAPI.MCPResource.Search()`: **Added**
* `conductoroneAPI.MCPResource.ListHistory()`: **Added**
* `conductoroneAPI.MCPResource.Update()`: **Added**
* `conductoroneAPI.MCPResource.Get()`: **Added**
* `conductoroneAPI.MCPResource.List()`: **Added**
* `conductoroneAPI.AppUser.ListOwnedServiceAccounts()`: **Added**
* `conductoroneAPI.AccessReviewActions.GenerateReport()`: **Added**
