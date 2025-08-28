# AppReport
(*AppReport*)

## Overview

### Available Operations

* [List](#list) - List

## List

Get a list of reports for the given app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppReportService.List" method="get" path="/api/v1/apps/{app_id}/report" -->
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

    res, err := s.AppReport.List(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppReportServiceListResponse != nil {
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

**[*operations.C1APIAppV1AppReportServiceListResponse](../../models/operations/c1apiappv1appreportservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |