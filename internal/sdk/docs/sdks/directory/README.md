# Directory

### Available Operations

* [Create](#create) - Invokes the c1.api.directory.v1.DirectoryService.Create method.
* [Delete](#delete) - Invokes the c1.api.directory.v1.DirectoryService.Delete method.
* [Get](#get) - Invokes the c1.api.directory.v1.DirectoryService.Get method.
* [List](#list) - Invokes the c1.api.directory.v1.DirectoryService.List method.

## Create

Invokes the c1.api.directory.v1.DirectoryService.Create method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Directory.Create(ctx, shared.DirectoryServiceCreateRequest{
        AppID: conductoroneapi.String("consequatur"),
        ExpandMask: &shared.DirectoryExpandMask{
            Paths: []string{
                "quibusdam",
                "explicabo",
                "deserunt",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.DirectoryServiceCreateRequest](../../models/shared/directoryservicecreaterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceCreateResponse](../../models/operations/c1apidirectoryv1directoryservicecreateresponse.md), error**


## Delete

Invokes the c1.api.directory.v1.DirectoryService.Delete method.

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
    res, err := s.Directory.Delete(ctx, operations.C1APIDirectoryV1DirectoryServiceDeleteRequest{
        DirectoryServiceDeleteRequest: &shared.DirectoryServiceDeleteRequest{},
        AppID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `request`                                                                                                                            | [operations.C1APIDirectoryV1DirectoryServiceDeleteRequest](../../models/operations/c1apidirectoryv1directoryservicedeleterequest.md) | :heavy_check_mark:                                                                                                                   | The request object to use for the request.                                                                                           |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceDeleteResponse](../../models/operations/c1apidirectoryv1directoryservicedeleteresponse.md), error**


## Get

Invokes the c1.api.directory.v1.DirectoryService.Get method.

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
    res, err := s.Directory.Get(ctx, operations.C1APIDirectoryV1DirectoryServiceGetRequest{
        AppID: "quibusdam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APIDirectoryV1DirectoryServiceGetRequest](../../models/operations/c1apidirectoryv1directoryservicegetrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceGetResponse](../../models/operations/c1apidirectoryv1directoryservicegetresponse.md), error**


## List

Invokes the c1.api.directory.v1.DirectoryService.List method.

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
    res, err := s.Directory.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.DirectoryServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIDirectoryV1DirectoryServiceListResponse](../../models/operations/c1apidirectoryv1directoryservicelistresponse.md), error**

