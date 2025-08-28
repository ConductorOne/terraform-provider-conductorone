# SystemLog
(*SystemLog*)

## Overview

### Available Operations

* [ListEvents](#listevents) - List Events

## ListEvents

ListEvents pulls Events from the ConductorOne system.

 This endpoint should be used to synchorize the
 system log events to external systems.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.systemlog.v1.SystemLogService.ListEvents" method="post" path="/api/v1/systemlog/events" -->
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

    res, err := s.SystemLog.ListEvents(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SystemLogServiceListEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [components.SystemLogServiceListEventsRequest](../../models/components/systemlogservicelisteventsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.C1APISystemlogV1SystemLogServiceListEventsResponse](../../models/operations/c1apisystemlogv1systemlogservicelisteventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |