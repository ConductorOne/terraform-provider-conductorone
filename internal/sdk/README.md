# conductorone-api

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get terraform/internal/sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.AppEntitlementUserBinding.ListAppUsersForIdentityWithGrant(ctx, operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantRequest{
        AppEntitlementID: "corrupti",
        AppID: "provident",
        IdentityUserID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppUsersForIdentityWithGrantResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AppEntitlementUserBinding](docs/sdks/appentitlementuserbinding/README.md)

* [ListAppUsersForIdentityWithGrant](docs/sdks/appentitlementuserbinding/README.md#listappusersforidentitywithgrant) - Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.

### [AppEntitlements](docs/sdks/appentitlements/README.md)

* [Get](docs/sdks/appentitlements/README.md#get) - Invokes the c1.api.app.v1.AppEntitlements.Get method.

### [AppOwners](docs/sdks/appowners/README.md)

* [Add](docs/sdks/appowners/README.md#add) - Invokes the c1.api.app.v1.AppOwners.Add method.
* [List](docs/sdks/appowners/README.md#list) - Invokes the c1.api.app.v1.AppOwners.List method.
* [Remove](docs/sdks/appowners/README.md#remove) - Invokes the c1.api.app.v1.AppOwners.Remove method.

### [AppReport](docs/sdks/appreport/README.md)

* [List](docs/sdks/appreport/README.md#list) - Invokes the c1.api.app.v1.AppReportService.List method.

### [AppReportAction](docs/sdks/appreportaction/README.md)

* [GenerateReport](docs/sdks/appreportaction/README.md#generatereport) - Invokes the c1.api.app.v1.AppReportActionService.GenerateReport method.

### [AppResource](docs/sdks/appresource/README.md)

* [Get](docs/sdks/appresource/README.md#get) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [List](docs/sdks/appresource/README.md#list) - Invokes the c1.api.app.v1.AppResourceService.List method.

### [AppResourceOwners](docs/sdks/appresourceowners/README.md)

* [List](docs/sdks/appresourceowners/README.md#list) - Invokes the c1.api.app.v1.AppResourceOwners.List method.

### [AppResourceSearch](docs/sdks/appresourcesearch/README.md)

* [SearchAppResourceTypes](docs/sdks/appresourcesearch/README.md#searchappresourcetypes) - Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes method.

### [AppResourceType](docs/sdks/appresourcetype/README.md)

* [Get](docs/sdks/appresourcetype/README.md#get) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.
* [List](docs/sdks/appresourcetype/README.md#list) - Invokes the c1.api.app.v1.AppResourceTypeService.List method.

### [AppSearch](docs/sdks/appsearch/README.md)

* [Search](docs/sdks/appsearch/README.md#search) - Invokes the c1.api.app.v1.AppSearch.Search method.

### [AppUsageControls](docs/sdks/appusagecontrols/README.md)

* [Get](docs/sdks/appusagecontrols/README.md#get) - Invokes the c1.api.app.v1.AppUsageControlsService.Get method.

### [Apps](docs/sdks/apps/README.md)

* [Get](docs/sdks/apps/README.md#get) - Invokes the c1.api.app.v1.Apps.Get method.
* [List](docs/sdks/apps/README.md#list) - Invokes the c1.api.app.v1.Apps.List method.

### [Auth](docs/sdks/auth/README.md)

* [Introspect](docs/sdks/auth/README.md#introspect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

### [Directory](docs/sdks/directory/README.md)

* [Create](docs/sdks/directory/README.md#create) - Invokes the c1.api.directory.v1.DirectoryService.Create method.
* [Delete](docs/sdks/directory/README.md#delete) - Invokes the c1.api.directory.v1.DirectoryService.Delete method.
* [Get](docs/sdks/directory/README.md#get) - Invokes the c1.api.directory.v1.DirectoryService.Get method.
* [List](docs/sdks/directory/README.md#list) - Invokes the c1.api.directory.v1.DirectoryService.List method.

### [PersonalClient](docs/sdks/personalclient/README.md)

* [Create](docs/sdks/personalclient/README.md#create) - Invokes the c1.api.iam.v1.PersonalClientService.Create method.

### [Policies](docs/sdks/policies/README.md)

* [Create](docs/sdks/policies/README.md#create) - Invokes the c1.api.policy.v1.Policies.Create method.
* [Delete](docs/sdks/policies/README.md#delete) - Invokes the c1.api.policy.v1.Policies.Delete method.
* [Get](docs/sdks/policies/README.md#get) - Invokes the c1.api.policy.v1.Policies.Get method.
* [List](docs/sdks/policies/README.md#list) - Invokes the c1.api.policy.v1.Policies.List method.
* [Update](docs/sdks/policies/README.md#update) - Invokes the c1.api.policy.v1.Policies.Update method.

### [RequestCatalogSearch](docs/sdks/requestcatalogsearch/README.md)

* [SearchEntitlements](docs/sdks/requestcatalogsearch/README.md#searchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

### [Roles](docs/sdks/roles/README.md)

* [Get](docs/sdks/roles/README.md#get) - Invokes the c1.api.iam.v1.Roles.Get method.
* [List](docs/sdks/roles/README.md#list) - Invokes the c1.api.iam.v1.Roles.List method.

### [Task](docs/sdks/task/README.md)

* [CreateGrantTask](docs/sdks/task/README.md#creategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [CreateRevokeTask](docs/sdks/task/README.md#createrevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [Get](docs/sdks/task/README.md#get) - Invokes the c1.api.task.v1.TaskService.Get method.

### [TaskActions](docs/sdks/taskactions/README.md)

* [Approve](docs/sdks/taskactions/README.md#approve) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [Comment](docs/sdks/taskactions/README.md#comment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [Deny](docs/sdks/taskactions/README.md#deny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.

### [TaskSearch](docs/sdks/tasksearch/README.md)

* [Search](docs/sdks/tasksearch/README.md#search) - Invokes the c1.api.task.v1.TaskSearchService.Search method.

### [User](docs/sdks/user/README.md)

* [Get](docs/sdks/user/README.md#get) - Invokes the c1.api.user.v1.UserService.Get method.
* [List](docs/sdks/user/README.md#list) - Invokes the c1.api.user.v1.UserService.List method.

### [UserSearch](docs/sdks/usersearch/README.md)

* [Search](docs/sdks/usersearch/README.md#search) - Invokes the c1.api.user.v1.UserSearch.Search method.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
