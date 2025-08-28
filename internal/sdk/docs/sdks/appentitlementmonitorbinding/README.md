# AppEntitlementMonitorBinding
(*AppEntitlementMonitorBinding*)

## Overview

### Available Operations

* [DeleteAppEntitlementMonitorBinding](#deleteappentitlementmonitorbinding) - Delete App Entitlement Monitor Binding
* [CreateAppEntitlementMonitorBinding](#createappentitlementmonitorbinding) - Create App Entitlement Monitor Binding
* [GetAppEntitlementMonitorBinding](#getappentitlementmonitorbinding) - Get App Entitlement Monitor Binding

## DeleteAppEntitlementMonitorBinding

Invokes the c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.DeleteAppEntitlementMonitorBinding method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.DeleteAppEntitlementMonitorBinding" method="delete" path="/api/v1/appentitlementmonitorbinding" -->
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

    res, err := s.AppEntitlementMonitorBinding.DeleteAppEntitlementMonitorBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAppEntitlementMonitorBindingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [components.DeleteAppEntitlementMonitorBindingRequest](../../models/components/deleteappentitlementmonitorbindingrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.C1APIAccessconflictV1AppEntitlementMonitorBindingServiceDeleteAppEntitlementMonitorBindingResponse](../../models/operations/c1apiaccessconflictv1appentitlementmonitorbindingservicedeleteappentitlementmonitorbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateAppEntitlementMonitorBinding

Invokes the c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.CreateAppEntitlementMonitorBinding method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.CreateAppEntitlementMonitorBinding" method="post" path="/api/v1/appentitlementmonitorbinding" -->
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

    res, err := s.AppEntitlementMonitorBinding.CreateAppEntitlementMonitorBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementMonitorBinding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [components.CreateAppEntitlementMonitorBindingRequest](../../models/components/createappentitlementmonitorbindingrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.C1APIAccessconflictV1AppEntitlementMonitorBindingServiceCreateAppEntitlementMonitorBindingResponse](../../models/operations/c1apiaccessconflictv1appentitlementmonitorbindingservicecreateappentitlementmonitorbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAppEntitlementMonitorBinding

Invokes the c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.GetAppEntitlementMonitorBinding method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AppEntitlementMonitorBindingService.GetAppEntitlementMonitorBinding" method="post" path="/api/v1/appentitlementmonitorbinding/get" -->
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

    res, err := s.AppEntitlementMonitorBinding.GetAppEntitlementMonitorBinding(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementMonitorBinding != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [components.GetAppEntitlementMonitorBindingRequest](../../models/components/getappentitlementmonitorbindingrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.C1APIAccessconflictV1AppEntitlementMonitorBindingServiceGetAppEntitlementMonitorBindingResponse](../../models/operations/c1apiaccessconflictv1appentitlementmonitorbindingservicegetappentitlementmonitorbindingresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |