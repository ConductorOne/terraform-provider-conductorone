# AppResourceType
(*AppResourceType*)

## Overview

### Available Operations

* [List](#list) - List
* [CreateManuallyManagedResourceType](#createmanuallymanagedresourcetype) - Create Manually Managed Resource Type
* [DeleteManuallyManagedResourceType](#deletemanuallymanagedresourcetype) - Delete Manually Managed Resource Type
* [Get](#get) - Get
* [UpdateManuallyManagedResourceType](#updatemanuallymanagedresourcetype) - Update Manually Managed Resource Type

## List

List app resource types.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceTypeService.List" method="get" path="/api/v1/apps/{app_id}/resource_types" -->
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

    res, err := s.AppResourceType.List(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppResourceTypeServiceListResponse != nil {
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

**[*operations.C1APIAppV1AppResourceTypeServiceListResponse](../../models/operations/c1apiappv1appresourcetypeservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateManuallyManagedResourceType

Invokes the c1.api.app.v1.AppResourceTypeService.CreateManuallyManagedResourceType method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceTypeService.CreateManuallyManagedResourceType" method="post" path="/api/v1/apps/{app_id}/resource_types" -->
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

    res, err := s.AppResourceType.CreateManuallyManagedResourceType(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateManuallyManagedResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |
| `appID`                                                                                                                     | *string*                                                                                                                    | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         |
| `createManuallyManagedResourceTypeRequest`                                                                                  | [*components.CreateManuallyManagedResourceTypeRequest](../../models/components/createmanuallymanagedresourcetyperequest.md) | :heavy_minus_sign:                                                                                                          | N/A                                                                                                                         |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceCreateManuallyManagedResourceTypeResponse](../../models/operations/c1apiappv1appresourcetypeservicecreatemanuallymanagedresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteManuallyManagedResourceType

Invokes the c1.api.app.v1.AppResourceTypeService.DeleteManuallyManagedResourceType method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceTypeService.DeleteManuallyManagedResourceType" method="delete" path="/api/v1/apps/{app_id}/resource_types/{id}" -->
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

    res, err := s.AppResourceType.DeleteManuallyManagedResourceType(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteManuallyManagedResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |
| `appID`                                                                                                                     | *string*                                                                                                                    | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         |
| `id`                                                                                                                        | *string*                                                                                                                    | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         |
| `deleteManuallyManagedResourceTypeRequest`                                                                                  | [*components.DeleteManuallyManagedResourceTypeRequest](../../models/components/deletemanuallymanagedresourcetyperequest.md) | :heavy_minus_sign:                                                                                                          | N/A                                                                                                                         |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceDeleteManuallyManagedResourceTypeResponse](../../models/operations/c1apiappv1appresourcetypeservicedeletemanuallymanagedresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get an app resource type.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceTypeService.Get" method="get" path="/api/v1/apps/{app_id}/resource_types/{id}" -->
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

    res, err := s.AppResourceType.Get(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AppResourceTypeServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceGetResponse](../../models/operations/c1apiappv1appresourcetypeservicegetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateManuallyManagedResourceType

Invokes the c1.api.app.v1.AppResourceTypeService.UpdateManuallyManagedResourceType method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceTypeService.UpdateManuallyManagedResourceType" method="post" path="/api/v1/apps/{app_id}/resource_types/{id}" -->
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

    res, err := s.AppResourceType.UpdateManuallyManagedResourceType(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateManuallyManagedResourceTypeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                   | Type                                                                                                                        | Required                                                                                                                    | Description                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                       | :heavy_check_mark:                                                                                                          | The context to use for the request.                                                                                         |
| `appID`                                                                                                                     | *string*                                                                                                                    | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         |
| `id`                                                                                                                        | *string*                                                                                                                    | :heavy_check_mark:                                                                                                          | N/A                                                                                                                         |
| `updateManuallyManagedResourceTypeRequest`                                                                                  | [*components.UpdateManuallyManagedResourceTypeRequest](../../models/components/updatemanuallymanagedresourcetyperequest.md) | :heavy_minus_sign:                                                                                                          | N/A                                                                                                                         |
| `opts`                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                    | :heavy_minus_sign:                                                                                                          | The options for this request.                                                                                               |

### Response

**[*operations.C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeResponse](../../models/operations/c1apiappv1appresourcetypeserviceupdatemanuallymanagedresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |