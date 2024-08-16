# Export

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [ListEvents](#listevents) - List Events
* [Update](#update) - Update

## Create

Create a system log export.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Export.Create(ctx, shared.ExportServiceCreateRequest{
        ExportToDatasource: &shared.ExportToDatasource{
            DatasourceID: conductoronesdkterraform.String("quidem"),
            Format: shared.ExportToDatasourceFormatExportFormatUnspecified.ToPointer(),
            Prefix: conductoronesdkterraform.String("voluptas"),
        },
        DisplayName: conductoronesdkterraform.String("unde"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [shared.ExportServiceCreateRequest](../../models/shared/exportservicecreaterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.C1APISystemlogV1ExportServiceCreateResponse](../../models/operations/c1apisystemlogv1exportservicecreateresponse.md), error**


## Delete

Delete a policy by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Export.Delete(ctx, operations.C1APISystemlogV1ExportServiceDeleteRequest{
        ExportServiceDeleteRequest: &shared.ExportServiceDeleteRequest{},
        ExportID: "architecto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APISystemlogV1ExportServiceDeleteRequest](../../models/operations/c1apisystemlogv1exportservicedeleterequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APISystemlogV1ExportServiceDeleteResponse](../../models/operations/c1apisystemlogv1exportservicedeleteresponse.md), error**


## Get

Get a system log export by ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Export.Get(ctx, operations.C1APISystemlogV1ExportServiceGetRequest{
        ExportID: "suscipit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.C1APISystemlogV1ExportServiceGetRequest](../../models/operations/c1apisystemlogv1exportservicegetrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.C1APISystemlogV1ExportServiceGetResponse](../../models/operations/c1apisystemlogv1exportservicegetresponse.md), error**


## List

List Exports.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Export.List(ctx, operations.C1APISystemlogV1ExportServiceListRequest{
        PageSize: conductoronesdkterraform.Int(960257),
        PageToken: conductoronesdkterraform.String("debitis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APISystemlogV1ExportServiceListRequest](../../models/operations/c1apisystemlogv1exportservicelistrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.C1APISystemlogV1ExportServiceListResponse](../../models/operations/c1apisystemlogv1exportservicelistresponse.md), error**


## ListEvents

Invokes the c1.api.systemlog.v1.ExportService.ListEvents method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Export.ListEvents(ctx, operations.C1APISystemlogV1ExportServiceListEventsRequest{
        ExportServiceListEventsRequest: &shared.ExportServiceListEventsRequest{
            PageSize: conductoronesdkterraform.Int(72434),
            PageToken: conductoronesdkterraform.String("reiciendis"),
        },
        ExportID: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportServiceListEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.C1APISystemlogV1ExportServiceListEventsRequest](../../models/operations/c1apisystemlogv1exportservicelisteventsrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.C1APISystemlogV1ExportServiceListEventsResponse](../../models/operations/c1apisystemlogv1exportservicelisteventsresponse.md), error**


## Update

Update a system log export by providing a policy object and an update mask.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Export.Update(ctx, operations.C1APISystemlogV1ExportServiceUpdateRequest{
        ExportServiceUpdateRequestInput: &shared.ExportServiceUpdateRequestInput{
            Exporter: &shared.ExporterInput{
                ExportToDatasource: &shared.ExportToDatasource{
                    DatasourceID: conductoronesdkterraform.String("corrupti"),
                    Format: shared.ExportToDatasourceFormatExportFormatOcsfJSONGzip.ToPointer(),
                    Prefix: conductoronesdkterraform.String("incidunt"),
                },
                DisplayName: conductoronesdkterraform.String("sed"),
            },
            UpdateMask: conductoronesdkterraform.String("provident"),
        },
        ExportID: "eius",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.C1APISystemlogV1ExportServiceUpdateRequest](../../models/operations/c1apisystemlogv1exportserviceupdaterequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.C1APISystemlogV1ExportServiceUpdateResponse](../../models/operations/c1apisystemlogv1exportserviceupdateresponse.md), error**

