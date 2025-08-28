# Automation
(*Automation*)

## Overview

### Available Operations

* [ListAutomations](#listautomations) - List Automations
* [CreateAutomation](#createautomation) - Create Automation
* [DeleteAutomation](#deleteautomation) - Delete Automation
* [GetAutomation](#getautomation) - Get Automation
* [UpdateAutomation](#updateautomation) - Update Automation
* [ExecuteAutomation](#executeautomation) - Execute Automation

## ListAutomations

Invokes the c1.api.automations.v1.AutomationService.ListAutomations method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ListAutomations" method="get" path="/api/v1/automations" -->
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

    res, err := s.Automation.ListAutomations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAutomationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceListAutomationsResponse](../../models/operations/c1apiautomationsv1automationservicelistautomationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateAutomation

Invokes the c1.api.automations.v1.AutomationService.CreateAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.CreateAutomation" method="post" path="/api/v1/automations" -->
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

    res, err := s.Automation.CreateAutomation(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.CreateAutomationRequest](../../models/components/createautomationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceCreateAutomationResponse](../../models/operations/c1apiautomationsv1automationservicecreateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAutomation

Invokes the c1.api.automations.v1.AutomationService.DeleteAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.DeleteAutomation" method="delete" path="/api/v1/automations/{id}" -->
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

    res, err := s.Automation.DeleteAutomation(ctx, "<id>", nil)
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
| `id`                                                                                      | *string*                                                                                  | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `deleteAutomationRequest`                                                                 | [*components.DeleteAutomationRequest](../../models/components/deleteautomationrequest.md) | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceDeleteAutomationResponse](../../models/operations/c1apiautomationsv1automationservicedeleteautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAutomation

Invokes the c1.api.automations.v1.AutomationService.GetAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.GetAutomation" method="get" path="/api/v1/automations/{id}" -->
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

    res, err := s.Automation.GetAutomation(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAutomationResponse != nil {
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

**[*operations.C1APIAutomationsV1AutomationServiceGetAutomationResponse](../../models/operations/c1apiautomationsv1automationservicegetautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAutomation

Invokes the c1.api.automations.v1.AutomationService.UpdateAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.UpdateAutomation" method="post" path="/api/v1/automations/{id}" -->
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

    res, err := s.Automation.UpdateAutomation(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `id`                                                                                      | *string*                                                                                  | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `updateAutomationRequest`                                                                 | [*components.UpdateAutomationRequest](../../models/components/updateautomationrequest.md) | :heavy_minus_sign:                                                                        | N/A                                                                                       |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceUpdateAutomationResponse](../../models/operations/c1apiautomationsv1automationserviceupdateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ExecuteAutomation

Invokes the c1.api.automations.v1.AutomationService.ExecuteAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationService.ExecuteAutomation" method="post" path="/api/v1/automations/{id}/execute" -->
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

    res, err := s.Automation.ExecuteAutomation(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExecuteAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `executeAutomationRequest`                                                                  | [*components.ExecuteAutomationRequest](../../models/components/executeautomationrequest.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.C1APIAutomationsV1AutomationServiceExecuteAutomationResponse](../../models/operations/c1apiautomationsv1automationserviceexecuteautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |