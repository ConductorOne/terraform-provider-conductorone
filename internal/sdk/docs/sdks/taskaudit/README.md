# TaskAudit
(*TaskAudit*)

## Overview

### Available Operations

* [List](#list) - List

## List

Invokes the c1.api.task.v1.TaskAudit.List method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskAudit.List" method="post" path="/api/v1/task/audits" -->
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

    res, err := s.TaskAudit.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskAuditListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.TaskAuditListRequest](../../models/components/taskauditlistrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.C1APITaskV1TaskAuditListResponse](../../models/operations/c1apitaskv1taskauditlistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |