# AutomationExecutionActions
(*AutomationExecutionActions*)

## Overview

### Available Operations

* [TerminateAutomation](#terminateautomation) - Terminate Automation

## TerminateAutomation

Invokes the c1.api.automations.v1.AutomationExecutionActionsService.TerminateAutomation method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationExecutionActionsService.TerminateAutomation" method="post" path="/api/v1/automation_executions/{id}/actions/terminate" -->
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

    res, err := s.AutomationExecutionActions.TerminateAutomation(ctx, 839265, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TerminateAutomationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |
| `id`                                                                                            | *int64*                                                                                         | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `terminateAutomationRequest`                                                                    | [*components.TerminateAutomationRequest](../../models/components/terminateautomationrequest.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |

### Response

**[*operations.C1APIAutomationsV1AutomationExecutionActionsServiceTerminateAutomationResponse](../../models/operations/c1apiautomationsv1automationexecutionactionsserviceterminateautomationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |