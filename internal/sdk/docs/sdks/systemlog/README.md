# SystemLog

### Available Operations

* [ListEvents](#listevents) - List Events

## ListEvents

ListEvents pulls Events from the ConductorOne system.

 This endpoint should be used to synchorize the
 system log events to external systems.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/types"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.SystemLog.ListEvents(ctx, shared.SystemLogServiceListEventsRequest{
        PageSize: conductoronesdkterraform.Int(855484),
        PageToken: conductoronesdkterraform.String("impedit"),
        Since: types.MustTimeFromString("2021-11-29T22:29:42.962Z"),
        SinceEventUID: conductoronesdkterraform.String("voluptates"),
        SortDirection: shared.SystemLogServiceListEventsRequestSortDirectionSortDirectionDesc.ToPointer(),
        Until: types.MustTimeFromString("2022-12-05T12:01:23.554Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SystemLogServiceListEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.SystemLogServiceListEventsRequest](../../models/shared/systemlogservicelisteventsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.C1APISystemlogV1SystemLogServiceListEventsResponse](../../models/operations/c1apisystemlogv1systemlogservicelisteventsresponse.md), error**

