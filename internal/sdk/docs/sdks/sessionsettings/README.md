# SessionSettings
(*SessionSettings*)

## Overview

### Available Operations

* [Get](#get) - Get
* [Update](#update) - Update
* [TestSourceIP](#testsourceip) - Test Source Ip

## Get

Invokes the c1.api.settings.v1.SessionSettingsService.Get method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.SessionSettingsService.Get" method="get" path="/api/v1/settings/session" -->
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

    res, err := s.SessionSettings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSessionSettingsResponse != nil {
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

**[*operations.C1APISettingsV1SessionSettingsServiceGetResponse](../../models/operations/c1apisettingsv1sessionsettingsservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Invokes the c1.api.settings.v1.SessionSettingsService.Update method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.SessionSettingsService.Update" method="post" path="/api/v1/settings/session" -->
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

    res, err := s.SessionSettings.Update(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateSessionSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.UpdateSessionSettingsRequest](../../models/components/updatesessionsettingsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APISettingsV1SessionSettingsServiceUpdateResponse](../../models/operations/c1apisettingsv1sessionsettingsserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## TestSourceIP

Invokes the c1.api.settings.v1.SessionSettingsService.TestSourceIP method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.SessionSettingsService.TestSourceIP" method="post" path="/api/v1/settings/session/test-source-ip" -->
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

    res, err := s.SessionSettings.TestSourceIP(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TestSourceIPResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.TestSourceIPRequest](../../models/components/testsourceiprequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.C1APISettingsV1SessionSettingsServiceTestSourceIPResponse](../../models/operations/c1apisettingsv1sessionsettingsservicetestsourceipresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |