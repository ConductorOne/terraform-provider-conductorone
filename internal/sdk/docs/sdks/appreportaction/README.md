# AppReportAction

### Available Operations

* [GenerateReport](#generatereport) - Invokes the c1.api.app.v1.AppReportActionService.GenerateReport method.

## GenerateReport

Invokes the c1.api.app.v1.AppReportActionService.GenerateReport method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/operations"
	"terraform/internal/sdk/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.AppReportAction.GenerateReport(ctx, operations.C1APIAppV1AppReportActionServiceGenerateReportRequest{
        AppActionsServiceGenerateReportRequest: &shared.AppActionsServiceGenerateReportRequest{},
        AppID: "esse",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppActionsServiceGenerateReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                            | Type                                                                                                                                                 | Required                                                                                                                                             | Description                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                | :heavy_check_mark:                                                                                                                                   | The context to use for the request.                                                                                                                  |
| `request`                                                                                                                                            | [operations.C1APIAppV1AppReportActionServiceGenerateReportRequest](../../models/operations/c1apiappv1appreportactionservicegeneratereportrequest.md) | :heavy_check_mark:                                                                                                                                   | The request object to use for the request.                                                                                                           |


### Response

**[*operations.C1APIAppV1AppReportActionServiceGenerateReportResponse](../../models/operations/c1apiappv1appreportactionservicegeneratereportresponse.md), error**

