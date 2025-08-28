# AppUsageControls
(*AppUsageControls*)

## Overview

### Available Operations

* [Get](#get) - Get
* [Update](#update) - Update

## Get

Get usage controls, as an AppUsageControls object which describes some peripheral configuration, for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUsageControlsService.Get" method="get" path="/api/v1/apps/{app_id}/usage_controls" -->
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

    res, err := s.AppUsageControls.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAppUsageControlsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppUsageControlsServiceGetResponse](../../models/operations/c1apiappv1appusagecontrolsservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update usage controls for an app.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUsageControlsService.Update" method="post" path="/api/v1/apps/{app_id}/usage_controls" -->
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

    res, err := s.AppUsageControls.Update(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateAppUsageControlsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `appID`                                                                                               | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `updateAppUsageControlsRequest`                                                                       | [*components.UpdateAppUsageControlsRequest](../../models/components/updateappusagecontrolsrequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APIAppV1AppUsageControlsServiceUpdateResponse](../../models/operations/c1apiappv1appusagecontrolsserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |