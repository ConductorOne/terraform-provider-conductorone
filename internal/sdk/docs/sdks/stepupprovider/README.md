# StepUpProvider
(*StepUpProvider*)

## Overview

### Available Operations

* [Search](#search) - Search
* [List](#list) - List
* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [Update](#update) - Update
* [UpdateSecret](#updatesecret) - Update Secret
* [Test](#test) - Test

## Search

Search allows searching for step-up providers with various filters

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.Search" method="post" path="/api/v1/search/step-up/providers" -->
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

    res, err := s.StepUpProvider.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchStepUpProvidersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [components.SearchStepUpProvidersRequest](../../models/components/searchstepupprovidersrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.C1APIStepupV1StepUpProviderServiceSearchResponse](../../models/operations/c1apistepupv1stepupproviderservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Invokes the c1.api.stepup.v1.StepUpProviderService.List method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.List" method="get" path="/api/v1/step-up/providers" -->
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

    res, err := s.StepUpProvider.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListStepUpProvidersResponse != nil {
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

**[*operations.C1APIStepupV1StepUpProviderServiceListResponse](../../models/operations/c1apistepupv1stepupproviderservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Invokes the c1.api.stepup.v1.StepUpProviderService.Create method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.Create" method="post" path="/api/v1/step-up/providers" -->
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

    res, err := s.StepUpProvider.Create(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateStepUpProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateStepUpProviderRequest](../../models/components/createstepupproviderrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.C1APIStepupV1StepUpProviderServiceCreateResponse](../../models/operations/c1apistepupv1stepupproviderservicecreateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Invokes the c1.api.stepup.v1.StepUpProviderService.Delete method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.Delete" method="delete" path="/api/v1/step-up/providers/{id}" -->
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

    res, err := s.StepUpProvider.Delete(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteStepUpProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `id`                                                                                              | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `deleteStepUpProviderRequest`                                                                     | [*components.DeleteStepUpProviderRequest](../../models/components/deletestepupproviderrequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.C1APIStepupV1StepUpProviderServiceDeleteResponse](../../models/operations/c1apistepupv1stepupproviderservicedeleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Invokes the c1.api.stepup.v1.StepUpProviderService.Get method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.Get" method="get" path="/api/v1/step-up/providers/{id}" -->
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

    res, err := s.StepUpProvider.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetStepUpProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIStepupV1StepUpProviderServiceGetResponse](../../models/operations/c1apistepupv1stepupproviderservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Invokes the c1.api.stepup.v1.StepUpProviderService.Update method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.Update" method="post" path="/api/v1/step-up/providers/{id}" -->
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

    res, err := s.StepUpProvider.Update(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateStepUpProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `id`                                                                                              | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `updateStepUpProviderRequest`                                                                     | [*components.UpdateStepUpProviderRequest](../../models/components/updatestepupproviderrequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.C1APIStepupV1StepUpProviderServiceUpdateResponse](../../models/operations/c1apistepupv1stepupproviderserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateSecret

Invokes the c1.api.stepup.v1.StepUpProviderService.UpdateSecret method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.UpdateSecret" method="post" path="/api/v1/step-up/providers/{id}/secret" -->
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

    res, err := s.StepUpProvider.UpdateSecret(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateStepUpProviderSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |
| `id`                                                                                                          | *string*                                                                                                      | :heavy_check_mark:                                                                                            | N/A                                                                                                           |
| `updateStepUpProviderSecretRequest`                                                                           | [*components.UpdateStepUpProviderSecretRequest](../../models/components/updatestepupprovidersecretrequest.md) | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `opts`                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                      | :heavy_minus_sign:                                                                                            | The options for this request.                                                                                 |

### Response

**[*operations.C1APIStepupV1StepUpProviderServiceUpdateSecretResponse](../../models/operations/c1apistepupv1stepupproviderserviceupdatesecretresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Test

Invokes the c1.api.stepup.v1.StepUpProviderService.Test method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.stepup.v1.StepUpProviderService.Test" method="post" path="/api/v1/step-up/providers/{id}/test" -->
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

    res, err := s.StepUpProvider.Test(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TestStepUpProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `id`                                                                                          | *string*                                                                                      | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `testStepUpProviderRequest`                                                                   | [*components.TestStepUpProviderRequest](../../models/components/teststepupproviderrequest.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*operations.C1APIStepupV1StepUpProviderServiceTestResponse](../../models/operations/c1apistepupv1stepupproviderservicetestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |