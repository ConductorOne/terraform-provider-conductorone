# undefined

Developer-friendly & type-safe Go SDK specifically catered to leverage *undefined* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=undefined&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/conductorone-v37/testws). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

ConductorOne API: The ConductorOne API is a HTTP API for managing ConductorOne resources.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [undefined](#undefined)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get undefined
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/models/components"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ConflictMonitor != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type   | Scheme       |
| ------------ | ------ | ------------ |
| `BearerAuth` | http   | HTTP Bearer  |
| `Oauth`      | oauth2 | OAuth2 token |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/models/components"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ConflictMonitor != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccessConflict](docs/sdks/accessconflict/README.md)

* [CreateMonitor](docs/sdks/accessconflict/README.md#createmonitor) - Create Monitor
* [DeleteMonitor](docs/sdks/accessconflict/README.md#deletemonitor) - Delete Monitor
* [GetMonitor](docs/sdks/accessconflict/README.md#getmonitor) - Get Monitor
* [UpdateMonitor](docs/sdks/accessconflict/README.md#updatemonitor) - Update Monitor

### [AccountProvisionPolicyTest](docs/sdks/accountprovisionpolicytest/README.md)

* [Test](docs/sdks/accountprovisionpolicytest/README.md#test) - Test

### [AppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md)

* [GetAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#getappaccessrequestsdefaults) - Get App Access Requests Defaults
* [CreateAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#createappaccessrequestsdefaults) - Create App Access Requests Defaults
* [CancelAppAccessRequestsDefaults](docs/sdks/appaccessrequestsdefaults/README.md#cancelappaccessrequestsdefaults) - Cancel App Access Requests Defaults

### [AppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md)

* [DeleteAppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md#deleteappentitlementmonitorbinding) - Delete App Entitlement Monitor Binding
* [CreateAppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md#createappentitlementmonitorbinding) - Create App Entitlement Monitor Binding
* [GetAppEntitlementMonitorBinding](docs/sdks/appentitlementmonitorbinding/README.md#getappentitlementmonitorbinding) - Get App Entitlement Monitor Binding

### [AppEntitlementOwners](docs/sdks/appentitlementowners/README.md)

* [List](docs/sdks/appentitlementowners/README.md#list) - List
* [Add](docs/sdks/appentitlementowners/README.md#add) - Add
* [Set](docs/sdks/appentitlementowners/README.md#set) - Set
* [Remove](docs/sdks/appentitlementowners/README.md#remove) - Remove

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [List](docs/sdks/appentitlements/README.md#list) - List
* [Create](docs/sdks/appentitlements/README.md#create) - Create
* [AddManuallyManagedMembers](docs/sdks/appentitlements/README.md#addmanuallymanagedmembers) - Add Manually Managed Members
* [DeleteAutomation](docs/sdks/appentitlements/README.md#deleteautomation) - Delete Automation
* [GetAutomation](docs/sdks/appentitlements/README.md#getautomation) - Get Automation
* [CreateAutomation](docs/sdks/appentitlements/README.md#createautomation) - Create Automation
* [RemoveAutomationExclusion](docs/sdks/appentitlements/README.md#removeautomationexclusion) - Remove Automation Exclusion
* [ListAutomationExclusions](docs/sdks/appentitlements/README.md#listautomationexclusions) - List Automation Exclusions
* [AddAutomationExclusion](docs/sdks/appentitlements/README.md#addautomationexclusion) - Add Automation Exclusion
* [UpdateAutomation](docs/sdks/appentitlements/README.md#updateautomation) - Update Automation
* [RemoveEntitlementMembership](docs/sdks/appentitlements/README.md#removeentitlementmembership) - Remove Entitlement Membership
* [~~ListUsers~~](docs/sdks/appentitlements/README.md#listusers) - List Users :warning: **Deprecated**
* [Delete](docs/sdks/appentitlements/README.md#delete) - Delete
* [Get](docs/sdks/appentitlements/README.md#get) - Get
* [Update](docs/sdks/appentitlements/README.md#update) - Update
* [ListForAppResource](docs/sdks/appentitlements/README.md#listforappresource) - List For App Resource
* [ListForAppUser](docs/sdks/appentitlements/README.md#listforappuser) - List For App User

### [AppEntitlementSearch](docs/sdks/appentitlementsearch/README.md)

* [SearchAppEntitlementsWithExpired](docs/sdks/appentitlementsearch/README.md#searchappentitlementswithexpired) - Search App Entitlements With Expired
* [SearchAppEntitlementsForAppUser](docs/sdks/appentitlementsearch/README.md#searchappentitlementsforappuser) - Search App Entitlements For App User
* [Search](docs/sdks/appentitlementsearch/README.md#search) - Search
* [SearchGrants](docs/sdks/appentitlementsearch/README.md#searchgrants) - Search Grants

### [AppEntitlementsProxy](docs/sdks/appentitlementsproxy/README.md)

* [Delete](docs/sdks/appentitlementsproxy/README.md#delete) - Delete
* [Get](docs/sdks/appentitlementsproxy/README.md#get) - Get
* [Create](docs/sdks/appentitlementsproxy/README.md#create) - Create

### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [RemoveGrantDuration](docs/sdks/appentitlementuserbinding/README.md#removegrantduration) - Remove Grant Duration
* [UpdateGrantDuration](docs/sdks/appentitlementuserbinding/README.md#updategrantduration) - Update Grant Duration
* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - List App Users For Identity With Grant
* [SearchGrantFeed](docs/sdks/appentitlementuserbinding/README.md#searchgrantfeed) - Search Grant Feed
* [SearchPastGrants](docs/sdks/appentitlementuserbinding/README.md#searchpastgrants) - Search Past Grants

### [AppOwners](docs/sdks/appowners/README.md)

* [List](docs/sdks/appowners/README.md#list) - List
* [Set](docs/sdks/appowners/README.md#set) - Set
* [Remove](docs/sdks/appowners/README.md#remove) - Remove
* [Add](docs/sdks/appowners/README.md#add) - Add

### [AppReport](docs/sdks/appreport/README.md)

* [List](docs/sdks/appreport/README.md#list) - List

### [AppReportAction](docs/sdks/appreportaction/README.md)

* [GenerateReport](docs/sdks/appreportaction/README.md#generatereport) - Generate Report

### [AppResource](docs/sdks/appresource/README.md)

* [List](docs/sdks/appresource/README.md#list) - List
* [CreateManuallyManagedAppResource](docs/sdks/appresource/README.md#createmanuallymanagedappresource) - Create Manually Managed App Resource
* [DeleteManuallyManagedAppResource](docs/sdks/appresource/README.md#deletemanuallymanagedappresource) - Delete Manually Managed App Resource
* [Get](docs/sdks/appresource/README.md#get) - Get
* [Update](docs/sdks/appresource/README.md#update) - Update

### [AppResourceOwners](docs/sdks/appresourceowners/README.md)

* [Remove](docs/sdks/appresourceowners/README.md#remove) - Remove
* [List](docs/sdks/appresourceowners/README.md#list) - List
* [Add](docs/sdks/appresourceowners/README.md#add) - Add

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Search App Resource Types
* [SearchAppResources](docs/sdks/appresourcesearch/README.md#searchappresources) - Search App Resources

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [List](docs/sdks/appresourcetype/README.md#list) - List
* [CreateManuallyManagedResourceType](docs/sdks/appresourcetype/README.md#createmanuallymanagedresourcetype) - Create Manually Managed Resource Type
* [DeleteManuallyManagedResourceType](docs/sdks/appresourcetype/README.md#deletemanuallymanagedresourcetype) - Delete Manually Managed Resource Type
* [Get](docs/sdks/appresourcetype/README.md#get) - Get
* [UpdateManuallyManagedResourceType](docs/sdks/appresourcetype/README.md#updatemanuallymanagedresourcetype) - Update Manually Managed Resource Type

### [Apps](docs/sdks/apps/README.md)

* [List](docs/sdks/apps/README.md#list) - List
* [Create](docs/sdks/apps/README.md#create) - Create
* [Delete](docs/sdks/apps/README.md#delete) - Delete
* [Get](docs/sdks/apps/README.md#get) - Get
* [Update](docs/sdks/apps/README.md#update) - Update

### [AppSearch](docs/sdks/appsearch/README.md)

* [Search](docs/sdks/appsearch/README.md#search) - Search

### [AppUsageControls](docs/sdks/appusagecontrols/README.md)

* [Get](docs/sdks/appusagecontrols/README.md#get) - Get
* [Update](docs/sdks/appusagecontrols/README.md#update) - Update

### [AppUser](docs/sdks/appuser/README.md)

* [List](docs/sdks/appuser/README.md#list) - List
* [ListAppUserCredentials](docs/sdks/appuser/README.md#listappusercredentials) - List App User Credentials
* [ListAppUsersForUser](docs/sdks/appuser/README.md#listappusersforuser) - List App Users For User
* [Update](docs/sdks/appuser/README.md#update) - Update
* [Search](docs/sdks/appuser/README.md#search) - Search

### [Attributes](docs/sdks/attributes/README.md)

* [DeleteAttributeValue](docs/sdks/attributes/README.md#deleteattributevalue) - Delete Attribute Value
* [CreateAttributeValue](docs/sdks/attributes/README.md#createattributevalue) - Create Attribute Value
* [GetAttributeValue](docs/sdks/attributes/README.md#getattributevalue) - Get Attribute Value
* [CreateComplianceFrameworkAttributeValue](docs/sdks/attributes/README.md#createcomplianceframeworkattributevalue) - Create Compliance Framework Attribute Value
* [DeleteComplianceFrameworkAttributeValue](docs/sdks/attributes/README.md#deletecomplianceframeworkattributevalue) - Delete Compliance Framework Attribute Value
* [GetComplianceFrameworkAttributeValue](docs/sdks/attributes/README.md#getcomplianceframeworkattributevalue) - Get Compliance Framework Attribute Value
* [CreateRiskLevelAttributeValue](docs/sdks/attributes/README.md#createrisklevelattributevalue) - Create Risk Level Attribute Value
* [DeleteRiskLevelAttributeValue](docs/sdks/attributes/README.md#deleterisklevelattributevalue) - Delete Risk Level Attribute Value
* [GetRiskLevelAttributeValue](docs/sdks/attributes/README.md#getrisklevelattributevalue) - Get Risk Level Attribute Value
* [ListAttributeTypes](docs/sdks/attributes/README.md#listattributetypes) - List Attribute Types
* [ListAttributeValues](docs/sdks/attributes/README.md#listattributevalues) - List Attribute Values

### [AttributeSearch](docs/sdks/attributesearch/README.md)

* [SearchAttributeValues](docs/sdks/attributesearch/README.md#searchattributevalues) - Search Attribute Values

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Introspect

### [Automation](docs/sdks/automation/README.md)

* [ListAutomations](docs/sdks/automation/README.md#listautomations) - List Automations
* [CreateAutomation](docs/sdks/automation/README.md#createautomation) - Create Automation
* [DeleteAutomation](docs/sdks/automation/README.md#deleteautomation) - Delete Automation
* [GetAutomation](docs/sdks/automation/README.md#getautomation) - Get Automation
* [UpdateAutomation](docs/sdks/automation/README.md#updateautomation) - Update Automation
* [ExecuteAutomation](docs/sdks/automation/README.md#executeautomation) - Execute Automation

### [AutomationExecution](docs/sdks/automationexecution/README.md)

* [ListAutomationExecutions](docs/sdks/automationexecution/README.md#listautomationexecutions) - List Automation Executions
* [GetAutomationExecution](docs/sdks/automationexecution/README.md#getautomationexecution) - Get Automation Execution

### [AutomationExecutionActions](docs/sdks/automationexecutionactions/README.md)

* [TerminateAutomation](docs/sdks/automationexecutionactions/README.md#terminateautomation) - Terminate Automation

### [AutomationExecutionSearch](docs/sdks/automationexecutionsearch/README.md)

* [SearchAutomationExecutions](docs/sdks/automationexecutionsearch/README.md#searchautomationexecutions) - Search Automation Executions

### [AutomationSearch](docs/sdks/automationsearch/README.md)

* [SearchAutomationTemplateVersions](docs/sdks/automationsearch/README.md#searchautomationtemplateversions) - Search Automation Template Versions
* [SearchAutomations](docs/sdks/automationsearch/README.md#searchautomations) - Search Automations

### [AWSExternalIDSettings](docs/sdks/awsexternalidsettings/README.md)

* [Get](docs/sdks/awsexternalidsettings/README.md#get) - Get


### [Connector](docs/sdks/connector/README.md)

* [List](docs/sdks/connector/README.md#list) - List
* [CreateDelegated](docs/sdks/connector/README.md#createdelegated) - Create Delegated
* [RotateCredential](docs/sdks/connector/README.md#rotatecredential) - Rotate Credential
* [GetCredentials](docs/sdks/connector/README.md#getcredentials) - Get Credentials
* [RevokeCredential](docs/sdks/connector/README.md#revokecredential) - Revoke Credential
* [ForceSync](docs/sdks/connector/README.md#forcesync) - Force Sync
* [Delete](docs/sdks/connector/README.md#delete) - Delete
* [Get](docs/sdks/connector/README.md#get) - Get
* [Update](docs/sdks/connector/README.md#update) - Update
* [Create](docs/sdks/connector/README.md#create) - Create
* [UpdateDelegated](docs/sdks/connector/README.md#updatedelegated) - Update Delegated
* [ValidateHTTPConnectorConfig](docs/sdks/connector/README.md#validatehttpconnectorconfig) - Validate Http Connector Config

### [Directory](docs/sdks/directory/README.md)

* [List](docs/sdks/directory/README.md#list) - List
* [Create](docs/sdks/directory/README.md#create) - Create
* [Delete](docs/sdks/directory/README.md#delete) - Delete
* [Get](docs/sdks/directory/README.md#get) - Get

### [Export](docs/sdks/export/README.md)

* [List](docs/sdks/export/README.md#list) - List
* [Create](docs/sdks/export/README.md#create) - Create
* [Delete](docs/sdks/export/README.md#delete) - Delete
* [Get](docs/sdks/export/README.md#get) - Get
* [Update](docs/sdks/export/README.md#update) - Update
* [ListEvents](docs/sdks/export/README.md#listevents) - List Events

### [ExportsSearch](docs/sdks/exportssearch/README.md)

* [Search](docs/sdks/exportssearch/README.md#search) - Search

### [Functions](docs/sdks/functions/README.md)

* [ListFunctions](docs/sdks/functions/README.md#listfunctions) - List Functions
* [CreateFunction](docs/sdks/functions/README.md#createfunction) - Create Function
* [ListCommits](docs/sdks/functions/README.md#listcommits) - List Commits
* [Commit](docs/sdks/functions/README.md#commit) - Commit
* [GetCommit](docs/sdks/functions/README.md#getcommit) - Get Commit
* [Invoke](docs/sdks/functions/README.md#invoke) - Invoke
* [ListTags](docs/sdks/functions/README.md#listtags) - List Tags
* [CreateTag](docs/sdks/functions/README.md#createtag) - Create Tag
* [DeleteFunction](docs/sdks/functions/README.md#deletefunction) - Delete Function
* [GetFunction](docs/sdks/functions/README.md#getfunction) - Get Function
* [UpdateFunction](docs/sdks/functions/README.md#updatefunction) - Update Function

### [FunctionsSearch](docs/sdks/functionssearch/README.md)

* [Search](docs/sdks/functionssearch/README.md#search) - Search

### [OrgDomain](docs/sdks/orgdomain/README.md)

* [List](docs/sdks/orgdomain/README.md#list) - List
* [Update](docs/sdks/orgdomain/README.md#update) - Update

### [PersonalClient](docs/sdks/personalclient/README.md)

* [List](docs/sdks/personalclient/README.md#list) - NOTE: Only shows personal clients for the current user.
* [Create](docs/sdks/personalclient/README.md#create) - Create
* [Delete](docs/sdks/personalclient/README.md#delete) - Delete
* [Get](docs/sdks/personalclient/README.md#get) - Get
* [Update](docs/sdks/personalclient/README.md#update) - Update

### [PersonalClientSearch](docs/sdks/personalclientsearch/README.md)

* [Search](docs/sdks/personalclientsearch/README.md#search) - NOTE: Searches personal clients for all users

### [Policies](docs/sdks/policies/README.md)

* [List](docs/sdks/policies/README.md#list) - List
* [Create](docs/sdks/policies/README.md#create) - Create
* [Delete](docs/sdks/policies/README.md#delete) - Delete
* [Get](docs/sdks/policies/README.md#get) - Get
* [Update](docs/sdks/policies/README.md#update) - Update

### [PolicySearch](docs/sdks/policysearch/README.md)

* [Search](docs/sdks/policysearch/README.md#search) - Search

### [PolicyValidate](docs/sdks/policyvalidate/README.md)

* [ValidateCEL](docs/sdks/policyvalidate/README.md#validatecel) - Validate Cel

### [RequestCatalogManagement](docs/sdks/requestcatalogmanagement/README.md)

* [List](docs/sdks/requestcatalogmanagement/README.md#list) - List
* [Create](docs/sdks/requestcatalogmanagement/README.md#create) - Create
* [ListAllEntitlementIdsPerApp](docs/sdks/requestcatalogmanagement/README.md#listallentitlementidsperapp) - List All Entitlement Ids Per App
* [ListEntitlementsPerCatalog](docs/sdks/requestcatalogmanagement/README.md#listentitlementspercatalog) - List Entitlements Per Catalog
* [UpdateAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#updateappentitlements) - Update App Entitlements
* [RemoveAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeappentitlements) - Remove App Entitlements
* [AddAppEntitlements](docs/sdks/requestcatalogmanagement/README.md#addappentitlements) - Add App Entitlements
* [RemoveAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#removeaccessentitlements) - Remove Access Entitlements
* [AddAccessEntitlements](docs/sdks/requestcatalogmanagement/README.md#addaccessentitlements) - Add Access Entitlements
* [ListEntitlementsForAccess](docs/sdks/requestcatalogmanagement/README.md#listentitlementsforaccess) - List Entitlements For Access
* [Delete](docs/sdks/requestcatalogmanagement/README.md#delete) - Delete
* [Get](docs/sdks/requestcatalogmanagement/README.md#get) - Get
* [Update](docs/sdks/requestcatalogmanagement/README.md#update) - Update
* [DeleteBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#deletebundleautomation) - Delete Bundle Automation
* [GetBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#getbundleautomation) - Get Bundle Automation
* [SetBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#setbundleautomation) - Set Bundle Automation
* [CreateBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#createbundleautomation) - Create Bundle Automation
* [ForceRunBundleAutomation](docs/sdks/requestcatalogmanagement/README.md#forcerunbundleautomation) - Force Run Bundle Automation

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Search Entitlements

### [Roles](docs/sdks/roles/README.md)

* [List](docs/sdks/roles/README.md#list) - List
* [Get](docs/sdks/roles/README.md#get) - Get
* [Update](docs/sdks/roles/README.md#update) - Update

### [SessionSettings](docs/sdks/sessionsettings/README.md)

* [Get](docs/sdks/sessionsettings/README.md#get) - Get
* [Update](docs/sdks/sessionsettings/README.md#update) - Update
* [TestSourceIP](docs/sdks/sessionsettings/README.md#testsourceip) - Test Source Ip

### [StepUpProvider](docs/sdks/stepupprovider/README.md)

* [Search](docs/sdks/stepupprovider/README.md#search) - Search
* [List](docs/sdks/stepupprovider/README.md#list) - List
* [Create](docs/sdks/stepupprovider/README.md#create) - Create
* [Delete](docs/sdks/stepupprovider/README.md#delete) - Delete
* [Get](docs/sdks/stepupprovider/README.md#get) - Get
* [Update](docs/sdks/stepupprovider/README.md#update) - Update
* [UpdateSecret](docs/sdks/stepupprovider/README.md#updatesecret) - Update Secret
* [Test](docs/sdks/stepupprovider/README.md#test) - Test

### [StepUpTransaction](docs/sdks/stepuptransaction/README.md)

* [Search](docs/sdks/stepuptransaction/README.md#search) - Search
* [Get](docs/sdks/stepuptransaction/README.md#get) - Get

### [SystemLog](docs/sdks/systemlog/README.md)

* [ListEvents](docs/sdks/systemlog/README.md#listevents) - List Events

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Create Grant Task
* [CreateOffboardingTask](docs/sdks/task/README.md#createoffboardingtask) - Create Offboarding Task
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Create Revoke Task
* [Get](docs/sdks/task/README.md#get) - Get

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Approve
* [ApproveWithStepUp](docs/sdks/taskactions/README.md#approvewithstepup) - Approve With Step Up
* [Close](docs/sdks/taskactions/README.md#close) - Close
* [Comment](docs/sdks/taskactions/README.md#comment) - Comment
* [Deny](docs/sdks/taskactions/README.md#deny) - Deny
* [EscalateToEmergencyAccess](docs/sdks/taskactions/README.md#escalatetoemergencyaccess) - Escalate To Emergency Access
* [ProcessNow](docs/sdks/taskactions/README.md#processnow) - Process Now
* [Reassign](docs/sdks/taskactions/README.md#reassign) - Reassign
* [HardReset](docs/sdks/taskactions/README.md#hardreset) - Hard Reset
* [Restart](docs/sdks/taskactions/README.md#restart) - Restart
* [SkipStep](docs/sdks/taskactions/README.md#skipstep) - Skip Step
* [UpdateGrantDuration](docs/sdks/taskactions/README.md#updategrantduration) - Update Grant Duration
* [UpdateRequestData](docs/sdks/taskactions/README.md#updaterequestdata) - Update Request Data

### [TaskAudit](docs/sdks/taskaudit/README.md)

* [List](docs/sdks/taskaudit/README.md#list) - List

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Search

### [User](docs/sdks/user/README.md)

* [List](docs/sdks/user/README.md#list) - List
* [Get](docs/sdks/user/README.md#get) - Get

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Search

### [Webhooks](docs/sdks/webhooks/README.md)

* [List](docs/sdks/webhooks/README.md#list) - List
* [Create](docs/sdks/webhooks/README.md#create) - Create
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete
* [Get](docs/sdks/webhooks/README.md#get) - Get
* [Update](docs/sdks/webhooks/README.md#update) - Update
* [Test](docs/sdks/webhooks/README.md#test) - Test

### [WebhooksSearch](docs/sdks/webhookssearch/README.md)

* [Search](docs/sdks/webhookssearch/README.md#search) - Search

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/models/components"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AppEntitlementSearch.Search(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.AppEntitlementSearchServiceSearchResponse != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"log"
	"models/operations"
	"undefined"
	"undefined/models/components"
	"undefined/retry"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ConflictMonitor != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/models/components"
	"undefined/retry"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ConflictMonitor != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `CreateMonitor` function may return the following errors:

| Error Type         | Status Code | Content Type |
| ------------------ | ----------- | ------------ |
| apierrors.APIError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
	"context"
	"errors"
	"log"
	"undefined"
	"undefined/models/apierrors"
	"undefined/models/components"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil)
	if err != nil {

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Server Variables

The default server `https://{tenantDomain}.conductor.one` contains variables and is set to `https://example.conductor.one` by default. To override default values, the following options are available when initializing the SDK client instance:

| Variable       | Option                                  | Default     | Description                                       |
| -------------- | --------------------------------------- | ----------- | ------------------------------------------------- |
| `tenantDomain` | `WithTenantDomain(tenantDomain string)` | `"example"` | The domain of the tenant to use for this request. |

#### Example

```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/models/components"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithTenantDomain("<value>"),
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ConflictMonitor != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/models/components"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithServerURL("https://example.conductor.one"),
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ConflictMonitor != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"undefined"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = undefined.New(undefined.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=undefined&utm_campaign=go)
