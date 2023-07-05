# Apps

### Available Operations

* [Get](#get) - Invokes the c1.api.app.v1.Apps.Get method.
* [List](#list) - Invokes the c1.api.app.v1.Apps.List method.

## Get

Invokes the c1.api.app.v1.Apps.Get method.

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
    res, err := s.Apps.Get(ctx, operations.C1APIAppV1AppsGetRequest{
        ID: "074f1547-1b5e-46e1-bb99-d488e1e91e45",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.C1APIAppV1AppsGetRequest](../../models/operations/c1apiappv1appsgetrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APIAppV1AppsGetResponse](../../models/operations/c1apiappv1appsgetresponse.md), error**


## List

Invokes the c1.api.app.v1.Apps.List method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Apps.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIAppV1AppsListResponse](../../models/operations/c1apiappv1appslistresponse.md), error**

