# AppResource

### Available Operations

* [Get](#get) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [List](#list) - Invokes the c1.api.app.v1.AppResourceService.List method.

## Get

Invokes the c1.api.app.v1.AppResourceService.Get method.

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
    res, err := s.AppResource.Get(ctx, operations.C1APIAppV1AppResourceServiceGetRequest{
        AppID: "totam",
        AppResourceTypeID: "porro",
        ID: "a1ba928f-c816-4742-8b73-9205929396fe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppResourceServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.C1APIAppV1AppResourceServiceGetRequest](../../models/operations/c1apiappv1appresourceservicegetrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[*operations.C1APIAppV1AppResourceServiceGetResponse](../../models/operations/c1apiappv1appresourceservicegetresponse.md), error**


## List

Invokes the c1.api.app.v1.AppResourceService.List method.

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
    res, err := s.AppResource.List(ctx, operations.C1APIAppV1AppResourceServiceListRequest{
        AppID: "fuga",
        AppResourceTypeID: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppResourceServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.C1APIAppV1AppResourceServiceListRequest](../../models/operations/c1apiappv1appresourceservicelistrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.C1APIAppV1AppResourceServiceListResponse](../../models/operations/c1apiappv1appresourceservicelistresponse.md), error**

