# AppEntitlementsProxy
(*AppEntitlementsProxy*)

## Overview

### Available Operations

* [Delete](#delete) - Delete
* [Get](#get) - Get
* [Create](#create) - Create

## Delete

Invokes the c1.api.app.v1.AppEntitlementsProxy.Delete method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementsProxy.Delete" method="delete" path="/api/v1/apps/{src_app_id}/{src_app_entitlement_id}/bindings/{dst_app_id}/{dst_app_entitlement_id}" -->
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

    res, err := s.AppEntitlementsProxy.Delete(ctx, operations.C1APIAppV1AppEntitlementsProxyDeleteRequest{
        SrcAppID: "<id>",
        SrcAppEntitlementID: "<id>",
        DstAppID: "<id>",
        DstAppEntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementProxyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV1AppEntitlementsProxyDeleteRequest](../../models/operations/c1apiappv1appentitlementsproxydeleterequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIAppV1AppEntitlementsProxyDeleteResponse](../../models/operations/c1apiappv1appentitlementsproxydeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Invokes the c1.api.app.v1.AppEntitlementsProxy.Get method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementsProxy.Get" method="get" path="/api/v1/apps/{src_app_id}/{src_app_entitlement_id}/bindings/{dst_app_id}/{dst_app_entitlement_id}" -->
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

    res, err := s.AppEntitlementsProxy.Get(ctx, "<id>", "<id>", "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppEntitlementProxyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `srcAppID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `srcAppEntitlementID`                                    | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `dstAppID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `dstAppEntitlementID`                                    | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementsProxyGetResponse](../../models/operations/c1apiappv1appentitlementsproxygetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Invokes the c1.api.app.v1.AppEntitlementsProxy.Create method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementsProxy.Create" method="post" path="/api/v1/apps/{src_app_id}/{src_app_entitlement_id}/bindings/{dst_app_id}/{dst_app_entitlement_id}" -->
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

    res, err := s.AppEntitlementsProxy.Create(ctx, operations.C1APIAppV1AppEntitlementsProxyCreateRequest{
        SrcAppID: "<id>",
        SrcAppEntitlementID: "<id>",
        DstAppID: "<id>",
        DstAppEntitlementID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAppEntitlementProxyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIAppV1AppEntitlementsProxyCreateRequest](../../models/operations/c1apiappv1appentitlementsproxycreaterequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.C1APIAppV1AppEntitlementsProxyCreateResponse](../../models/operations/c1apiappv1appentitlementsproxycreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |