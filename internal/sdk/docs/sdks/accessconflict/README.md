# AccessConflict
(*AccessConflict*)

## Overview

### Available Operations

* [CreateMonitor](#createmonitor) - Create Monitor
* [DeleteMonitor](#deletemonitor) - Delete Monitor
* [GetMonitor](#getmonitor) - Get Monitor
* [UpdateMonitor](#updatemonitor) - Update Monitor

## CreateMonitor

Invokes the c1.api.accessconflict.v1.AccessConflictService.CreateMonitor method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AccessConflictService.CreateMonitor" method="post" path="/api/v1/accessconflict" -->
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

    res, err := s.AccessConflict.CreateMonitor(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConflictMonitor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.ConflictMonitorCreateRequest](../../models/components/conflictmonitorcreaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIAccessconflictV1AccessConflictServiceCreateMonitorResponse](../../models/operations/c1apiaccessconflictv1accessconflictservicecreatemonitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteMonitor

Invokes the c1.api.accessconflict.v1.AccessConflictService.DeleteMonitor method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AccessConflictService.DeleteMonitor" method="delete" path="/api/v1/accessconflict/{id}" -->
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

    res, err := s.AccessConflict.DeleteMonitor(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConflictMonitorDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |
| `id`                                                                                                | *string*                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 |
| `conflictMonitorDeleteRequest`                                                                      | [*components.ConflictMonitorDeleteRequest](../../models/components/conflictmonitordeleterequest.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |

### Response

**[*operations.C1APIAccessconflictV1AccessConflictServiceDeleteMonitorResponse](../../models/operations/c1apiaccessconflictv1accessconflictservicedeletemonitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetMonitor

Invokes the c1.api.accessconflict.v1.AccessConflictService.GetMonitor method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AccessConflictService.GetMonitor" method="get" path="/api/v1/accessconflict/{id}" -->
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

    res, err := s.AccessConflict.GetMonitor(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConflictMonitor != nil {
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

**[*operations.C1APIAccessconflictV1AccessConflictServiceGetMonitorResponse](../../models/operations/c1apiaccessconflictv1accessconflictservicegetmonitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateMonitor

Invokes the c1.api.accessconflict.v1.AccessConflictService.UpdateMonitor method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.accessconflict.v1.AccessConflictService.UpdateMonitor" method="post" path="/api/v1/accessconflict/{id}" -->
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

    res, err := s.AccessConflict.UpdateMonitor(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ConflictMonitor != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |
| `id`                                                                                                | *string*                                                                                            | :heavy_check_mark:                                                                                  | N/A                                                                                                 |
| `conflictMonitorUpdateRequest`                                                                      | [*components.ConflictMonitorUpdateRequest](../../models/components/conflictmonitorupdaterequest.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |

### Response

**[*operations.C1APIAccessconflictV1AccessConflictServiceUpdateMonitorResponse](../../models/operations/c1apiaccessconflictv1accessconflictserviceupdatemonitorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |