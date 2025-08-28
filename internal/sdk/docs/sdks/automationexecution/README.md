# AutomationExecution
(*AutomationExecution*)

## Overview

### Available Operations

* [ListAutomationExecutions](#listautomationexecutions) - List Automation Executions
* [GetAutomationExecution](#getautomationexecution) - Get Automation Execution

## ListAutomationExecutions

Invokes the c1.api.automations.v1.AutomationExecutionService.ListAutomationExecutions method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionService.ListAutomationExecutions" method="get" path="/api/v1/automation_executions" -->
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

    res, err := s.AutomationExecution.ListAutomationExecutions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAutomationExecutionsResponse != nil {
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

**[*operations.C1APIAutomationsV1AutomationExecutionServiceListAutomationExecutionsResponse](../../models/operations/c1apiautomationsv1automationexecutionservicelistautomationexecutionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAutomationExecution

Invokes the c1.api.automations.v1.AutomationExecutionService.GetAutomationExecution method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionService.GetAutomationExecution" method="get" path="/api/v1/automation_executions/{id}" -->
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

    res, err := s.AutomationExecution.GetAutomationExecution(ctx, 728203)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAutomationExecutionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAutomationsV1AutomationExecutionServiceGetAutomationExecutionResponse](../../models/operations/c1apiautomationsv1automationexecutionservicegetautomationexecutionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |