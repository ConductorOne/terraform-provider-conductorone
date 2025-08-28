# AppEntitlements
(*AppEntitlements*)

## Overview

### Available Operations

* [List](#list) - List
* [Create](#create) - Create
* [AddManuallyManagedMembers](#addmanuallymanagedmembers) - Add Manually Managed Members
* [DeleteAutomation](#deleteautomation) - Delete Automation
* [GetAutomation](#getautomation) - Get Automation
* [CreateAutomation](#createautomation) - Create Automation
* [RemoveAutomationExclusion](#removeautomationexclusion) - Remove Automation Exclusion
* [ListAutomationExclusions](#listautomationexclusions) - List Automation Exclusions
* [AddAutomationExclusion](#addautomationexclusion) - Add Automation Exclusion
* [UpdateAutomation](#updateautomation) - Update Automation
* [RemoveEntitlementMembership](#removeentitlementmembership) - Remove Entitlement Membership
* [~~ListUsers~~](#listusers) - List Users :warning: **Deprecated**
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [Update](#update) - Update
* [ListForAppResource](#listforappresource) - List For App Resource
* [ListForAppUser](#listforappuser) - List For App User

## List

List app entitlements associated with an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.List" method="get" path="/api/v1/apps/{app_id}/entitlements" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.List(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementsListResponse](../../models/operations/c1apiappv1appentitlementslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Invokes the c1.api.app.v1.AppEntitlements.Create method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.Create" method="post" path="/api/v1/apps/{app_id}/entitlements" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.Create(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `appID`                                                                                           | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `createAppEntitlementRequest`                                                                     | [*components.CreateAppEntitlementRequest](../../models/components/createappentitlementrequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.C1APIAppV1AppEntitlementsCreateResponse](../../models/operations/c1apiappv1appentitlementscreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddManuallyManagedMembers

Invokes the c1.api.app.v1.AppEntitlements.AddManuallyManagedMembers method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.AddManuallyManagedMembers" method="post" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/add-manual-user" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.AddManuallyManagedMembers(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ManuallyManagedUsersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |
| `appID`                                                                                                 | *string*                                                                                                | :heavy_check_mark:                                                                                      | N/A                                                                                                     |
| `appEntitlementID`                                                                                      | *string*                                                                                                | :heavy_check_mark:                                                                                      | N/A                                                                                                     |
| `addManuallyManagedUsersRequest`                                                                        | [*components.AddManuallyManagedUsersRequest](../../models/components/addmanuallymanagedusersrequest.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |

### Response

**[*operations.C1APIAppV1AppEntitlementsAddManuallyManagedMembersResponse](../../models/operations/c1apiappv1appentitlementsaddmanuallymanagedmembersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAutomation

Invokes the c1.api.app.v1.AppEntitlements.DeleteAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.DeleteAutomation" method="delete" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/automation" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.DeleteAutomation(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `appID`                                                                                   | *string*                                                                                  | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `appEntitlementID`                                                                        | *string*                                                                                  | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `deleteAutomationRequest`                                                                 | [*components.DeleteAutomationRequest](../../models/components/deleteautomationrequest.md) | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.C1APIAppV1AppEntitlementsDeleteAutomationResponse](../../models/operations/c1apiappv1appentitlementsdeleteautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAutomation

Invokes the c1.api.app.v1.AppEntitlements.GetAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.GetAutomation" method="get" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/automation" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.GetAutomation(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementServiceGetAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `appEntitlementID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementsGetAutomationResponse](../../models/operations/c1apiappv1appentitlementsgetautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateAutomation

Invokes the c1.api.app.v1.AppEntitlements.CreateAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.CreateAutomation" method="post" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/automation/create" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.CreateAutomation(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |
| `appID`                                                                                             | *string*                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 |
| `appEntitlementID`                                                                                  | *string*                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 |
| `createAutomationRequest`                                                                           | [*components.CreateAutomationRequestInput](../../models/components/createautomationrequestinput.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |

### Response

**[*operations.C1APIAppV1AppEntitlementsCreateAutomationResponse](../../models/operations/c1apiappv1appentitlementscreateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveAutomationExclusion

Invokes the c1.api.app.v1.AppEntitlements.RemoveAutomationExclusion method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.RemoveAutomationExclusion" method="delete" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/automation/exclusions" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.RemoveAutomationExclusion(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveAutomationExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `appID`                                                                                                     | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `appEntitlementID`                                                                                          | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `removeAutomationExclusionRequest`                                                                          | [*components.RemoveAutomationExclusionRequest](../../models/components/removeautomationexclusionrequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.C1APIAppV1AppEntitlementsRemoveAutomationExclusionResponse](../../models/operations/c1apiappv1appentitlementsremoveautomationexclusionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAutomationExclusions

Invokes the c1.api.app.v1.AppEntitlements.ListAutomationExclusions method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.ListAutomationExclusions" method="get" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/automation/exclusions" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.ListAutomationExclusions(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAutomationExclusionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `appEntitlementID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementsListAutomationExclusionsResponse](../../models/operations/c1apiappv1appentitlementslistautomationexclusionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddAutomationExclusion

Invokes the c1.api.app.v1.AppEntitlements.AddAutomationExclusion method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.AddAutomationExclusion" method="post" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/automation/exclusions" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.AddAutomationExclusion(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AddAutomationExclusionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `appID`                                                                                               | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `appEntitlementID`                                                                                    | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `addAutomationExclusionRequest`                                                                       | [*components.AddAutomationExclusionRequest](../../models/components/addautomationexclusionrequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APIAppV1AppEntitlementsAddAutomationExclusionResponse](../../models/operations/c1apiappv1appentitlementsaddautomationexclusionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAutomation

Invokes the c1.api.app.v1.AppEntitlements.UpdateAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.UpdateAutomation" method="post" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/automation/update" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.UpdateAutomation(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementServiceUpdateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |
| `appID`                                                                                                                             | *string*                                                                                                                            | :heavy_check_mark:                                                                                                                  | N/A                                                                                                                                 |
| `appEntitlementID`                                                                                                                  | *string*                                                                                                                            | :heavy_check_mark:                                                                                                                  | N/A                                                                                                                                 |
| `appEntitlementServiceUpdateAutomationRequest`                                                                                      | [*components.AppEntitlementServiceUpdateAutomationRequest](../../models/components/appentitlementserviceupdateautomationrequest.md) | :heavy_minus_sign:                                                                                                                  | N/A                                                                                                                                 |
| `opts`                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                            | :heavy_minus_sign:                                                                                                                  | The options for this request.                                                                                                       |

### Response

**[*operations.C1APIAppV1AppEntitlementsUpdateAutomationResponse](../../models/operations/c1apiappv1appentitlementsupdateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveEntitlementMembership

Invokes the c1.api.app.v1.AppEntitlements.RemoveEntitlementMembership method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.RemoveEntitlementMembership" method="delete" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/remove-membership" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.RemoveEntitlementMembership(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveEntitlementMembershipResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |
| `appID`                                                                                                         | *string*                                                                                                        | :heavy_check_mark:                                                                                              | N/A                                                                                                             |
| `appEntitlementID`                                                                                              | *string*                                                                                                        | :heavy_check_mark:                                                                                              | N/A                                                                                                             |
| `removeEntitlementMembershipRequest`                                                                            | [*components.RemoveEntitlementMembershipRequest](../../models/components/removeentitlementmembershiprequest.md) | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `opts`                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                        | :heavy_minus_sign:                                                                                              | The options for this request.                                                                                   |

### Response

**[*operations.C1APIAppV1AppEntitlementsRemoveEntitlementMembershipResponse](../../models/operations/c1apiappv1appentitlementsremoveentitlementmembershipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~ListUsers~~

List the users, as AppEntitlementUsers objects, of an app entitlement.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.ListUsers" method="get" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/users" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.ListUsers(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementUsersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `appEntitlementID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementsListUsersResponse](../../models/operations/c1apiappv1appentitlementslistusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Invokes the c1.api.app.v1.AppEntitlements.Delete method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.Delete" method="delete" path="/api/v1/apps/{app_id}/entitlements/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.Delete(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `appID`                                                                                           | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `id`                                                                                              | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `deleteAppEntitlementRequest`                                                                     | [*components.DeleteAppEntitlementRequest](../../models/components/deleteappentitlementrequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.C1APIAppV1AppEntitlementsDeleteResponse](../../models/operations/c1apiappv1appentitlementsdeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get an app entitlement by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.Get" method="get" path="/api/v1/apps/{app_id}/entitlements/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.Get(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementsGetResponse](../../models/operations/c1apiappv1appentitlementsgetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an app entitlement by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.Update" method="post" path="/api/v1/apps/{app_id}/entitlements/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.Update(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `appID`                                                                                           | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `id`                                                                                              | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `updateAppEntitlementRequest`                                                                     | [*components.UpdateAppEntitlementRequest](../../models/components/updateappentitlementrequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.C1APIAppV1AppEntitlementsUpdateResponse](../../models/operations/c1apiappv1appentitlementsupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListForAppResource

List app entitlements associated with an app resource.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.ListForAppResource" method="get" path="/api/v1/apps/{app_id}/entitlements/resource_types/{app_resource_type_id}/resources/{app_resource_id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"undefined/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.ListForAppResource(ctx, operations.C1APIAppV1AppEntitlementsListForAppResourceRequest{
        AppID: "<id>",
        AppResourceTypeID: "<id>",
        AppResourceID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `request`                                                                                                                                      | [operations.C1APIAppV1AppEntitlementsListForAppResourceRequest](../../models/operations/c1apiappv1appentitlementslistforappresourcerequest.md) | :heavy_check_mark:                                                                                                                             | The request object to use for the request.                                                                                                     |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.C1APIAppV1AppEntitlementsListForAppResourceResponse](../../models/operations/c1apiappv1appentitlementslistforappresourceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListForAppUser

List app entitlements associated with an app user.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlements.ListForAppUser" method="get" path="/api/v1/apps/{app_id}/entitlements/users/{app_user_id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.AppEntitlements.ListForAppUser(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `appUserID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementsListForAppUserResponse](../../models/operations/c1apiappv1appentitlementslistforappuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |