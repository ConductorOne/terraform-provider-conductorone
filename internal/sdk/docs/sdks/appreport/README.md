# AppReport

### Available Operations

* [List](#list) - Invokes the c1.api.app.v1.AppReportService.List method.

## List

Invokes the c1.api.app.v1.AppReportService.List method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.AppReport.List(ctx, operations.C1APIAppV1AppReportServiceListRequest{
        AppID: "quod",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppReportServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.C1APIAppV1AppReportServiceListRequest](../../models/operations/c1apiappv1appreportservicelistrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.C1APIAppV1AppReportServiceListResponse](../../models/operations/c1apiappv1appreportservicelistresponse.md), error**

