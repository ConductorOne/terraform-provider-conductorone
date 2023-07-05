# User

### Available Operations

* [Get](#get) - Invokes the c1.api.user.v1.UserService.Get method.
* [List](#list) - Invokes the c1.api.user.v1.UserService.List method.

## Get

Invokes the c1.api.user.v1.UserService.Get method.

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
    res, err := s.User.Get(ctx, operations.C1APIUserV1UserServiceGetRequest{
        ID: "9817ee17-cbe6-41e6-b7b9-5bc0ab3c20c4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.C1APIUserV1UserServiceGetRequest](../../models/operations/c1apiuserv1userservicegetrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.C1APIUserV1UserServiceGetResponse](../../models/operations/c1apiuserv1userservicegetresponse.md), error**


## List

Invokes the c1.api.user.v1.UserService.List method.

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
    res, err := s.User.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.UserServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIUserV1UserServiceListResponse](../../models/operations/c1apiuserv1userservicelistresponse.md), error**

