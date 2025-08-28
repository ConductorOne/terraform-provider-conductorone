# AppAccessRequestsDefaults
(*AppAccessRequestsDefaults*)

## Overview

### Available Operations

* [GetAppAccessRequestsDefaults](#getappaccessrequestsdefaults) - Get App Access Requests Defaults
* [CreateAppAccessRequestsDefaults](#createappaccessrequestsdefaults) - Create App Access Requests Defaults
* [CancelAppAccessRequestsDefaults](#cancelappaccessrequestsdefaults) - Cancel App Access Requests Defaults

## GetAppAccessRequestsDefaults

Invokes the c1.api.app.v1.AppAccessRequestsDefaultsService.GetAppAccessRequestsDefaults method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppAccessRequestsDefaultsService.GetAppAccessRequestsDefaults" method="get" path="/api/v1/apps/{app_id}/access_request_defaults" -->
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

    res, err := s.AppAccessRequestsDefaults.GetAppAccessRequestsDefaults(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AppAccessRequestDefaults != nil {
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

**[*operations.C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsResponse](../../models/operations/c1apiappv1appaccessrequestsdefaultsservicegetappaccessrequestsdefaultsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateAppAccessRequestsDefaults

Invokes the c1.api.app.v1.AppAccessRequestsDefaultsService.CreateAppAccessRequestsDefaults method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppAccessRequestsDefaultsService.CreateAppAccessRequestsDefaults" method="post" path="/api/v1/apps/{app_id}/access_request_defaults" -->
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

    res, err := s.AppAccessRequestsDefaults.CreateAppAccessRequestsDefaults(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppAccessRequestDefaults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `appID`                                                                                       | *string*                                                                                      | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `appAccessRequestDefaults`                                                                    | [*components.AppAccessRequestDefaults1](../../models/components/appaccessrequestdefaults1.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*operations.C1APIAppV1AppAccessRequestsDefaultsServiceCreateAppAccessRequestsDefaultsResponse](../../models/operations/c1apiappv1appaccessrequestsdefaultsservicecreateappaccessrequestsdefaultsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CancelAppAccessRequestsDefaults

Invokes the c1.api.app.v1.AppAccessRequestsDefaultsService.CancelAppAccessRequestsDefaults method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppAccessRequestsDefaultsService.CancelAppAccessRequestsDefaults" method="post" path="/api/v1/apps/{app_id}/access_request_defaults/cancel" -->
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

    res, err := s.AppAccessRequestsDefaults.CancelAppAccessRequestsDefaults(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppAccessRequestDefaults != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |
| `appID`                                                                                                         | *string*                                                                                                        | :heavy_check_mark:                                                                                              | N/A                                                                                                             |
| `cancelAccessRequestDefaultsRequest`                                                                            | [*components.CancelAccessRequestDefaultsRequest](../../models/components/cancelaccessrequestdefaultsrequest.md) | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `opts`                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                        | :heavy_minus_sign:                                                                                              | The options for this request.                                                                                   |

### Response

**[*operations.C1APIAppV1AppAccessRequestsDefaultsServiceCancelAppAccessRequestsDefaultsResponse](../../models/operations/c1apiappv1appaccessrequestsdefaultsservicecancelappaccessrequestsdefaultsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |