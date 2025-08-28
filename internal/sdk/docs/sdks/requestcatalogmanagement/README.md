# RequestCatalogManagement
(*RequestCatalogManagement*)

## Overview

### Available Operations

* [List](#list) - List
* [Create](#create) - Create
* [ListAllEntitlementIdsPerApp](#listallentitlementidsperapp) - List All Entitlement Ids Per App
* [ListEntitlementsPerCatalog](#listentitlementspercatalog) - List Entitlements Per Catalog
* [UpdateAppEntitlements](#updateappentitlements) - Update App Entitlements
* [RemoveAppEntitlements](#removeappentitlements) - Remove App Entitlements
* [AddAppEntitlements](#addappentitlements) - Add App Entitlements
* [RemoveAccessEntitlements](#removeaccessentitlements) - Remove Access Entitlements
* [AddAccessEntitlements](#addaccessentitlements) - Add Access Entitlements
* [ListEntitlementsForAccess](#listentitlementsforaccess) - List Entitlements For Access
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [Update](#update) - Update
* [DeleteBundleAutomation](#deletebundleautomation) - Delete Bundle Automation
* [GetBundleAutomation](#getbundleautomation) - Get Bundle Automation
* [SetBundleAutomation](#setbundleautomation) - Set Bundle Automation
* [CreateBundleAutomation](#createbundleautomation) - Create Bundle Automation
* [ForceRunBundleAutomation](#forcerunbundleautomation) - Force Run Bundle Automation

## List

Get a list of request catalogs.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.List" method="get" path="/api/v1/catalogs" -->
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

    res, err := s.RequestCatalogManagement.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Creates a new request catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Create" method="post" path="/api/v1/catalogs" -->
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

    res, err := s.RequestCatalogManagement.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [components.RequestCatalogManagementServiceCreateRequest](../../models/components/requestcatalogmanagementservicecreaterequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAllEntitlementIdsPerApp

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.ListAllEntitlementIdsPerApp method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ListAllEntitlementIdsPerApp" method="get" path="/api/v1/catalogs/{catalog_id}/requestable_entitlementIDs" -->
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

    res, err := s.RequestCatalogManagement.ListAllEntitlementIdsPerApp(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListAllEntitlementIdsPerCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `catalogID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListAllEntitlementIdsPerAppResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistallentitlementidsperappresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListEntitlementsPerCatalog

List entitlements in a catalog that are requestable.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ListEntitlementsPerCatalog" method="get" path="/api/v1/catalogs/{catalog_id}/requestable_entitlements" -->
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

    res, err := s.RequestCatalogManagement.ListEntitlementsPerCatalog(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `catalogID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementspercatalogresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAppEntitlements

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.UpdateAppEntitlements method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.UpdateAppEntitlements" method="post" path="/api/v1/catalogs/{catalog_id}/requestable_entitlements/update" -->
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

    res, err := s.RequestCatalogManagement.UpdateAppEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceUpdateAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |
| `catalogID`                                                                                                                                                       | *string*                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |
| `requestCatalogManagementServiceUpdateAppEntitlementsRequest`                                                                                                     | [*components.RequestCatalogManagementServiceUpdateAppEntitlementsRequest](../../models/components/requestcatalogmanagementserviceupdateappentitlementsrequest.md) | :heavy_minus_sign:                                                                                                                                                | N/A                                                                                                                                                               |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateAppEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdateappentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveAppEntitlements

Remove requestable entitlements from a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.RemoveAppEntitlements" method="delete" path="/api/v1/catalogs/{catalog_id}/requestable_entries" -->
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

    res, err := s.RequestCatalogManagement.RemoveAppEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceRemoveAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |
| `catalogID`                                                                                                                                                       | *string*                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |
| `requestCatalogManagementServiceRemoveAppEntitlementsRequest`                                                                                                     | [*components.RequestCatalogManagementServiceRemoveAppEntitlementsRequest](../../models/components/requestcatalogmanagementserviceremoveappentitlementsrequest.md) | :heavy_minus_sign:                                                                                                                                                | N/A                                                                                                                                                               |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveappentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddAppEntitlements

Add requestable entitlements to a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.AddAppEntitlements" method="post" path="/api/v1/catalogs/{catalog_id}/requestable_entries" -->
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

    res, err := s.RequestCatalogManagement.AddAppEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceAddAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                   | Type                                                                                                                                                        | Required                                                                                                                                                    | Description                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                       | :heavy_check_mark:                                                                                                                                          | The context to use for the request.                                                                                                                         |
| `catalogID`                                                                                                                                                 | *string*                                                                                                                                                    | :heavy_check_mark:                                                                                                                                          | N/A                                                                                                                                                         |
| `requestCatalogManagementServiceAddAppEntitlementsRequest`                                                                                                  | [*components.RequestCatalogManagementServiceAddAppEntitlementsRequest](../../models/components/requestcatalogmanagementserviceaddappentitlementsrequest.md) | :heavy_minus_sign:                                                                                                                                          | N/A                                                                                                                                                         |
| `opts`                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                    | :heavy_minus_sign:                                                                                                                                          | The options for this request.                                                                                                                               |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddappentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RemoveAccessEntitlements

Remove visibility bindings (access entitlements) to a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.RemoveAccessEntitlements" method="delete" path="/api/v1/catalogs/{catalog_id}/visibility_bindings" -->
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

    res, err := s.RequestCatalogManagement.RemoveAccessEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                               | Type                                                                                                                                                                    | Required                                                                                                                                                                | Description                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                   | :heavy_check_mark:                                                                                                                                                      | The context to use for the request.                                                                                                                                     |
| `catalogID`                                                                                                                                                             | *string*                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                      | N/A                                                                                                                                                                     |
| `requestCatalogManagementServiceRemoveAccessEntitlementsRequest`                                                                                                        | [*components.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest](../../models/components/requestcatalogmanagementserviceremoveaccessentitlementsrequest.md) | :heavy_minus_sign:                                                                                                                                                      | N/A                                                                                                                                                                     |
| `opts`                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | The options for this request.                                                                                                                                           |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveaccessentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddAccessEntitlements

Add visibility bindings (access entitlements) to a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.AddAccessEntitlements" method="post" path="/api/v1/catalogs/{catalog_id}/visibility_bindings" -->
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

    res, err := s.RequestCatalogManagement.AddAccessEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceAddAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |
| `catalogID`                                                                                                                                                       | *string*                                                                                                                                                          | :heavy_check_mark:                                                                                                                                                | N/A                                                                                                                                                               |
| `requestCatalogManagementServiceAddAccessEntitlementsRequest`                                                                                                     | [*components.RequestCatalogManagementServiceAddAccessEntitlementsRequest](../../models/components/requestcatalogmanagementserviceaddaccessentitlementsrequest.md) | :heavy_minus_sign:                                                                                                                                                | N/A                                                                                                                                                               |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddaccessentitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListEntitlementsForAccess

List visibility bindings (access entitlements) for a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ListEntitlementsForAccess" method="get" path="/api/v1/catalogs/{catalog_id}/visibility_entitlements" -->
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

    res, err := s.RequestCatalogManagement.ListEntitlementsForAccess(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceListEntitlementsForAccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `catalogID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementsforaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Delete" method="delete" path="/api/v1/catalogs/{id}" -->
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

    res, err := s.RequestCatalogManagement.Delete(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |
| `id`                                                                                                                                | *string*                                                                                                                            | :heavy_check_mark:                                                                                                                  | N/A                                                                                                                                 |
| `requestCatalogManagementServiceDeleteRequest`                                                                                      | [*components.RequestCatalogManagementServiceDeleteRequest](../../models/components/requestcatalogmanagementservicedeleterequest.md) | :heavy_minus_sign:                                                                                                                  | N/A                                                                                                                                 |
| `opts`                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                            | :heavy_minus_sign:                                                                                                                  | The options for this request.                                                                                                       |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Get" method="get" path="/api/v1/catalogs/{id}" -->
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

    res, err := s.RequestCatalogManagement.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a catalog.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.Update" method="post" path="/api/v1/catalogs/{id}" -->
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

    res, err := s.RequestCatalogManagement.Update(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |
| `id`                                                                                                                                | *string*                                                                                                                            | :heavy_check_mark:                                                                                                                  | N/A                                                                                                                                 |
| `requestCatalogManagementServiceUpdateRequest`                                                                                      | [*components.RequestCatalogManagementServiceUpdateRequest](../../models/components/requestcatalogmanagementserviceupdaterequest.md) | :heavy_minus_sign:                                                                                                                  | N/A                                                                                                                                 |
| `opts`                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                            | :heavy_minus_sign:                                                                                                                  | The options for this request.                                                                                                       |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.DeleteBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.DeleteBundleAutomation" method="delete" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation" -->
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

    res, err := s.RequestCatalogManagement.DeleteBundleAutomation(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteBundleAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `requestCatalogID`                                                                                    | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `deleteBundleAutomationRequest`                                                                       | [*components.DeleteBundleAutomationRequest](../../models/components/deletebundleautomationrequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteBundleAutomationResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeletebundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetBundleAutomation

Get bundle automation

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.GetBundleAutomation" method="get" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation" -->
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

    res, err := s.RequestCatalogManagement.GetBundleAutomation(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `requestCatalogID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetBundleAutomationResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetbundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.SetBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.SetBundleAutomation" method="post" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation" -->
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

    res, err := s.RequestCatalogManagement.SetBundleAutomation(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |
| `requestCatalogID`                                                                              | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `setBundleAutomationRequest`                                                                    | [*components.SetBundleAutomationRequest](../../models/components/setbundleautomationrequest.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicesetbundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.CreateBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.CreateBundleAutomation" method="post" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation/create" -->
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

    res, err := s.RequestCatalogManagement.CreateBundleAutomation(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `requestCatalogID`                                                                                    | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `createBundleAutomationRequest`                                                                       | [*components.CreateBundleAutomationRequest](../../models/components/createbundleautomationrequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateBundleAutomationResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreatebundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ForceRunBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.ForceRunBundleAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.requestcatalog.v1.RequestCatalogManagementService.ForceRunBundleAutomation" method="post" path="/api/v1/catalogs/{request_catalog_id}/bundle_automation/run" -->
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

    res, err := s.RequestCatalogManagement.ForceRunBundleAutomation(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ForceRunBundleAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `requestCatalogID`                                                                                        | *string*                                                                                                  | :heavy_check_mark:                                                                                        | N/A                                                                                                       |
| `forceRunBundleAutomationRequest`                                                                         | [*components.ForceRunBundleAutomationRequest](../../models/components/forcerunbundleautomationrequest.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |

### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceforcerunbundleautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |