# RequestCatalogManagement

### Available Operations

* [AddAccessEntitlements](#addaccessentitlements) - Add Access Entitlements
* [AddAppEntitlements](#addappentitlements) - Add App Entitlements
* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [GetBundleAutomation](#getbundleautomation) - Get Bundle Automation
* [List](#list) - List
* [ListEntitlementsForAccess](#listentitlementsforaccess) - List Entitlements For Access
* [ListEntitlementsPerCatalog](#listentitlementspercatalog) - List Entitlements Per Catalog
* [RemoveAccessEntitlements](#removeaccessentitlements) - Remove Access Entitlements
* [RemoveAppEntitlements](#removeappentitlements) - Remove App Entitlements
* [SetBundleAutomation](#setbundleautomation) - Set Bundle Automation
* [Update](#update) - Update

## AddAccessEntitlements

Add visibility bindings (access entitlements) to a catalog.

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
    res, err := s.RequestCatalogManagement.AddAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest{
        RequestCatalogManagementServiceAddAccessEntitlementsRequest: &shared.RequestCatalogManagementServiceAddAccessEntitlementsRequest{
            AccessEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("doloremque"),
                    ID: conductoronesdkterraform.String("06d67887-8ba8-4581-a582-08c54fefa9c9"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("quaerat"),
                    ID: conductoronesdkterraform.String("f2eac556-5d30-47cf-ae81-206e2813fa4a"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("modi"),
                    ID: conductoronesdkterraform.String("1c480d3f-2132-4af0-b102-d514f4cc6f18"),
                },
            },
        },
        CatalogID: "expedita",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceAddAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddaccessentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAccessEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddaccessentitlementsresponse.md), error**


## AddAppEntitlements

Add requestable entitlements to a catalog.

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
    res, err := s.RequestCatalogManagement.AddAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest{
        RequestCatalogManagementServiceAddAppEntitlementsRequest: &shared.RequestCatalogManagementServiceAddAppEntitlementsRequest{
            AppEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("excepturi"),
                    ID: conductoronesdkterraform.String("621a6a4f-77a8-47ee-be4b-e752c65b3441"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("praesentium"),
                    ID: conductoronesdkterraform.String("e3bb91c8-d975-4e0e-8419-d8f84f144f3e"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("consequatur"),
                    ID: conductoronesdkterraform.String("7edcc4aa-5f3c-4abd-905a-972e05672822"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("iusto"),
                    ID: conductoronesdkterraform.String("b2d30947-0bf7-4a4f-a87c-f535a6fae54e"),
                },
            },
        },
        CatalogID: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceAddAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                            | Type                                                                                                                                                                                                 | Required                                                                                                                                                                                             | Description                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                  |
| `request`                                                                                                                                                                                            | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddappentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                           |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceAddAppEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceaddappentitlementsresponse.md), error**


## Create

Creates a new request catalog.

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
    res, err := s.RequestCatalogManagement.Create(ctx, shared.RequestCatalogManagementServiceCreateRequest{
        RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
            Paths: []string{
                "laboriosam",
                "voluptatem",
                "optio",
                "sequi",
            },
        },
        Description: conductoronesdkterraform.String("sunt"),
        DisplayName: conductoronesdkterraform.String("vitae"),
        Published: conductoronesdkterraform.Bool(false),
        RequestBundle: conductoronesdkterraform.Bool(false),
        VisibleToEveryone: conductoronesdkterraform.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [shared.RequestCatalogManagementServiceCreateRequest](../../models/shared/requestcatalogmanagementservicecreaterequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicecreateresponse.md), error**


## Delete

Delete a catalog.

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
    res, err := s.RequestCatalogManagement.Delete(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest{
        RequestCatalogManagementServiceDeleteRequest: &shared.RequestCatalogManagementServiceDeleteRequest{},
        ID: "f023b75d-2367-4fe1-a0cc-8df79f0a396d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleterequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicedeleteresponse.md), error**


## Get

Get a catalog.

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
    res, err := s.RequestCatalogManagement.Get(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest{
        ID: "90c364b7-c15d-4fba-8e18-8b1c4ee2c8c6",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                              | Type                                                                                                                                                                   | Required                                                                                                                                                               | Description                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                  | :heavy_check_mark:                                                                                                                                                     | The context to use for the request.                                                                                                                                    |
| `request`                                                                                                                                                              | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetrequest.md) | :heavy_check_mark:                                                                                                                                                     | The request object to use for the request.                                                                                                                             |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetresponse.md), error**


## GetBundleAutomation

Get bundle automation

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
    res, err := s.RequestCatalogManagement.GetBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetBundleAutomationRequest{
        RequestCatalogID: "eligendi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                              | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                    |
| `request`                                                                                                                                                                                              | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetBundleAutomationRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetbundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                             |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceGetBundleAutomationResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicegetbundleautomationresponse.md), error**


## List

Get a list of request catalogs.

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
    res, err := s.RequestCatalogManagement.List(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListRequest{
        PageSize: conductoronesdkterraform.Int(923159),
        PageToken: conductoronesdkterraform.String("ex"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `request`                                                                                                                                                                | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistrequest.md) | :heavy_check_mark:                                                                                                                                                       | The request object to use for the request.                                                                                                                               |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistresponse.md), error**


## ListEntitlementsForAccess

List visibility bindings (access entitlements) for a catalog.

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
    res, err := s.RequestCatalogManagement.ListEntitlementsForAccess(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest{
        CatalogID: "beatae",
        PageSize: conductoronesdkterraform.Int(86955),
        PageToken: conductoronesdkterraform.String("maiores"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceListEntitlementsForAccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                          | Type                                                                                                                                                                                                               | Required                                                                                                                                                                                                           | Description                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                |
| `request`                                                                                                                                                                                                          | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementsforaccessrequest.md) | :heavy_check_mark:                                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                                         |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementsforaccessresponse.md), error**


## ListEntitlementsPerCatalog

List entitlements in a catalog that are requestable.

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
    res, err := s.RequestCatalogManagement.ListEntitlementsPerCatalog(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest{
        CatalogID: "itaque",
        PageSize: conductoronesdkterraform.Int(875452),
        PageToken: conductoronesdkterraform.String("quidem"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |
| `request`                                                                                                                                                                                                            | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementspercatalogrequest.md) | :heavy_check_mark:                                                                                                                                                                                                   | The request object to use for the request.                                                                                                                                                                           |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicelistentitlementspercatalogresponse.md), error**


## RemoveAccessEntitlements

Remove visibility bindings (access entitlements) to a catalog.

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
    res, err := s.RequestCatalogManagement.RemoveAccessEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest{
        RequestCatalogManagementServiceRemoveAccessEntitlementsRequest: &shared.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest{
            AccessEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("quo"),
                    ID: conductoronesdkterraform.String("7cbdb6ee-c743-478b-a253-17747dc915ad"),
                },
            },
        },
        CatalogID: "fugit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                        | Type                                                                                                                                                                                                             | Required                                                                                                                                                                                                         | Description                                                                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                               | The context to use for the request.                                                                                                                                                                              |
| `request`                                                                                                                                                                                                        | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveaccessentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                               | The request object to use for the request.                                                                                                                                                                       |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveaccessentitlementsresponse.md), error**


## RemoveAppEntitlements

Remove requestable entitlements from a catalog.

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
    res, err := s.RequestCatalogManagement.RemoveAppEntitlements(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest{
        RequestCatalogManagementServiceRemoveAppEntitlementsRequest: &shared.RequestCatalogManagementServiceRemoveAppEntitlementsRequest{
            AppEntitlements: []shared.AppEntitlementRef{
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("dolorum"),
                    ID: conductoronesdkterraform.String("f5dd6723-dc0f-45ae-af3a-6b7008787561"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("quaerat"),
                    ID: conductoronesdkterraform.String("3f5a6c98-b555-4540-80d4-0bcacc6cbd6b"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("enim"),
                    ID: conductoronesdkterraform.String("f3ec9093-04f9-426b-ad25-53819b474b0e"),
                },
                shared.AppEntitlementRef{
                    AppID: conductoronesdkterraform.String("possimus"),
                    ID: conductoronesdkterraform.String("20e56248-fff6-439a-910a-bdcab6267669"),
                },
            },
        },
        CatalogID: "laboriosam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceRemoveAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                  | Type                                                                                                                                                                                                       | Required                                                                                                                                                                                                   | Description                                                                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                         | The context to use for the request.                                                                                                                                                                        |
| `request`                                                                                                                                                                                                  | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveappentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                                         | The request object to use for the request.                                                                                                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceremoveappentitlementsresponse.md), error**


## SetBundleAutomation

Invokes the c1.api.requestcatalog.v1.RequestCatalogManagementService.SetBundleAutomation method.

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
    res, err := s.RequestCatalogManagement.SetBundleAutomation(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationRequest{
        SetBundleAutomationRequest: &shared.SetBundleAutomationRequest{
            BundleAutomationRuleEntitlement: &shared.BundleAutomationRuleEntitlement{
                EntitlementRefs: []shared.AppEntitlementRef{
                    shared.AppEntitlementRef{
                        AppID: conductoronesdkterraform.String("ab"),
                        ID: conductoronesdkterraform.String("ec00221b-335d-489a-8b3e-cfda8d0c549e"),
                    },
                    shared.AppEntitlementRef{
                        AppID: conductoronesdkterraform.String("earum"),
                        ID: conductoronesdkterraform.String("03004978-a61f-4a1c-b206-88f77c1ffc71"),
                    },
                    shared.AppEntitlementRef{
                        AppID: conductoronesdkterraform.String("facere"),
                        ID: conductoronesdkterraform.String("ca163f2a-3c80-4a97-bf33-4cddf857a9e6"),
                    },
                    shared.AppEntitlementRef{
                        AppID: conductoronesdkterraform.String("quasi"),
                        ID: conductoronesdkterraform.String("876c6ab2-1d29-4dfc-94d6-fecd79939006"),
                    },
                },
            },
        },
        RequestCatalogID: "laboriosam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BundleAutomation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                              | Type                                                                                                                                                                                                   | Required                                                                                                                                                                                               | Description                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                    |
| `request`                                                                                                                                                                                              | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicesetbundleautomationrequest.md) | :heavy_check_mark:                                                                                                                                                                                     | The request object to use for the request.                                                                                                                                                             |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceSetBundleAutomationResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementservicesetbundleautomationresponse.md), error**


## Update

Update a catalog.

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
    res, err := s.RequestCatalogManagement.Update(ctx, operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest{
        RequestCatalogManagementServiceUpdateRequestInput: &shared.RequestCatalogManagementServiceUpdateRequestInput{
            RequestCatalog: &shared.RequestCatalogInput{
                AccessEntitlements: []shared.AppEntitlementInput{
                    shared.AppEntitlementInput{
                        ProvisionPolicy: &shared.ProvisionPolicy{
                            ConnectorProvision: &shared.ConnectorProvision{},
                            DelegatedProvision: &shared.DelegatedProvision{
                                AppID: conductoronesdkterraform.String("autem"),
                                EntitlementID: conductoronesdkterraform.String("assumenda"),
                                Implicit: conductoronesdkterraform.Bool(false),
                            },
                            ExternalTicketProvision: &shared.ExternalTicketProvision{
                                AppID: conductoronesdkterraform.String("explicabo"),
                                ConnectorID: conductoronesdkterraform.String("fugiat"),
                                ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("doloremque"),
                                Instructions: conductoronesdkterraform.String("voluptatem"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoronesdkterraform.String("alias"),
                                UserIds: []string{
                                    "ullam",
                                },
                            },
                            MultiStep: &shared.MultiStep{
                                ProvisionSteps: []shared.ProvisionPolicy{
                                    shared.ProvisionPolicy{},
                                    shared.ProvisionPolicy{},
                                },
                            },
                            WebhookProvision: &shared.WebhookProvision{
                                WebhookID: conductoronesdkterraform.String("velit"),
                            },
                        },
                        AppID: conductoronesdkterraform.String("ratione"),
                        AppResourceID: conductoronesdkterraform.String("quas"),
                        AppResourceTypeID: conductoronesdkterraform.String("maxime"),
                        CertifyPolicyID: conductoronesdkterraform.String("recusandae"),
                        ComplianceFrameworkValueIds: []string{
                            "doloremque",
                            "totam",
                            "iure",
                            "maiores",
                        },
                        DefaultValuesApplied: conductoronesdkterraform.Bool(false),
                        Description: conductoronesdkterraform.String("est"),
                        DisplayName: conductoronesdkterraform.String("fugit"),
                        DurationGrant: conductoronesdkterraform.String("veritatis"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoronesdkterraform.Bool(false),
                        EmergencyGrantPolicyID: conductoronesdkterraform.String("necessitatibus"),
                        GrantPolicyID: conductoronesdkterraform.String("iste"),
                        IsManuallyManaged: conductoronesdkterraform.Bool(false),
                        OverrideAccessRequestsDefaults: conductoronesdkterraform.Bool(false),
                        RevokePolicyID: conductoronesdkterraform.String("dicta"),
                        RiskLevelValueID: conductoronesdkterraform.String("ipsam"),
                        Slug: conductoronesdkterraform.String("consequuntur"),
                        SourceConnectorIds: map[string]string{
                            "quidem": "non",
                            "beatae": "sunt",
                            "molestias": "beatae",
                            "autem": "ducimus",
                        },
                        UserEditedMask: conductoronesdkterraform.String("libero"),
                    },
                    shared.AppEntitlementInput{
                        ProvisionPolicy: &shared.ProvisionPolicy{
                            ConnectorProvision: &shared.ConnectorProvision{},
                            DelegatedProvision: &shared.DelegatedProvision{
                                AppID: conductoronesdkterraform.String("molestias"),
                                EntitlementID: conductoronesdkterraform.String("necessitatibus"),
                                Implicit: conductoronesdkterraform.Bool(false),
                            },
                            ExternalTicketProvision: &shared.ExternalTicketProvision{
                                AppID: conductoronesdkterraform.String("ipsum"),
                                ConnectorID: conductoronesdkterraform.String("impedit"),
                                ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("quos"),
                                Instructions: conductoronesdkterraform.String("illum"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoronesdkterraform.String("distinctio"),
                                UserIds: []string{
                                    "non",
                                },
                            },
                            MultiStep: &shared.MultiStep{
                                ProvisionSteps: []shared.ProvisionPolicy{
                                    shared.ProvisionPolicy{},
                                    shared.ProvisionPolicy{},
                                },
                            },
                            WebhookProvision: &shared.WebhookProvision{
                                WebhookID: conductoronesdkterraform.String("consequatur"),
                            },
                        },
                        AppID: conductoronesdkterraform.String("laudantium"),
                        AppResourceID: conductoronesdkterraform.String("repellendus"),
                        AppResourceTypeID: conductoronesdkterraform.String("commodi"),
                        CertifyPolicyID: conductoronesdkterraform.String("quibusdam"),
                        ComplianceFrameworkValueIds: []string{
                            "voluptas",
                        },
                        DefaultValuesApplied: conductoronesdkterraform.Bool(false),
                        Description: conductoronesdkterraform.String("quaerat"),
                        DisplayName: conductoronesdkterraform.String("earum"),
                        DurationGrant: conductoronesdkterraform.String("tenetur"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoronesdkterraform.Bool(false),
                        EmergencyGrantPolicyID: conductoronesdkterraform.String("assumenda"),
                        GrantPolicyID: conductoronesdkterraform.String("dolore"),
                        IsManuallyManaged: conductoronesdkterraform.Bool(false),
                        OverrideAccessRequestsDefaults: conductoronesdkterraform.Bool(false),
                        RevokePolicyID: conductoronesdkterraform.String("enim"),
                        RiskLevelValueID: conductoronesdkterraform.String("ullam"),
                        Slug: conductoronesdkterraform.String("perspiciatis"),
                        SourceConnectorIds: map[string]string{
                            "ex": "quibusdam",
                        },
                        UserEditedMask: conductoronesdkterraform.String("dicta"),
                    },
                    shared.AppEntitlementInput{
                        ProvisionPolicy: &shared.ProvisionPolicy{
                            ConnectorProvision: &shared.ConnectorProvision{},
                            DelegatedProvision: &shared.DelegatedProvision{
                                AppID: conductoronesdkterraform.String("quia"),
                                EntitlementID: conductoronesdkterraform.String("commodi"),
                                Implicit: conductoronesdkterraform.Bool(false),
                            },
                            ExternalTicketProvision: &shared.ExternalTicketProvision{
                                AppID: conductoronesdkterraform.String("neque"),
                                ConnectorID: conductoronesdkterraform.String("quibusdam"),
                                ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("numquam"),
                                Instructions: conductoronesdkterraform.String("rem"),
                            },
                            ManualProvision: &shared.ManualProvision{
                                Instructions: conductoronesdkterraform.String("officiis"),
                                UserIds: []string{
                                    "neque",
                                    "corporis",
                                    "quod",
                                },
                            },
                            MultiStep: &shared.MultiStep{
                                ProvisionSteps: []shared.ProvisionPolicy{
                                    shared.ProvisionPolicy{},
                                },
                            },
                            WebhookProvision: &shared.WebhookProvision{
                                WebhookID: conductoronesdkterraform.String("placeat"),
                            },
                        },
                        AppID: conductoronesdkterraform.String("excepturi"),
                        AppResourceID: conductoronesdkterraform.String("recusandae"),
                        AppResourceTypeID: conductoronesdkterraform.String("quos"),
                        CertifyPolicyID: conductoronesdkterraform.String("dicta"),
                        ComplianceFrameworkValueIds: []string{
                            "ipsum",
                            "consequatur",
                            "soluta",
                            "necessitatibus",
                        },
                        DefaultValuesApplied: conductoronesdkterraform.Bool(false),
                        Description: conductoronesdkterraform.String("sequi"),
                        DisplayName: conductoronesdkterraform.String("recusandae"),
                        DurationGrant: conductoronesdkterraform.String("labore"),
                        DurationUnset: &shared.AppEntitlementDurationUnset{},
                        EmergencyGrantEnabled: conductoronesdkterraform.Bool(false),
                        EmergencyGrantPolicyID: conductoronesdkterraform.String("adipisci"),
                        GrantPolicyID: conductoronesdkterraform.String("magni"),
                        IsManuallyManaged: conductoronesdkterraform.Bool(false),
                        OverrideAccessRequestsDefaults: conductoronesdkterraform.Bool(false),
                        RevokePolicyID: conductoronesdkterraform.String("aperiam"),
                        RiskLevelValueID: conductoronesdkterraform.String("dolores"),
                        Slug: conductoronesdkterraform.String("illum"),
                        SourceConnectorIds: map[string]string{
                            "magni": "beatae",
                            "aliquid": "ad",
                        },
                        UserEditedMask: conductoronesdkterraform.String("voluptate"),
                    },
                },
                AppIds: []string{
                    "minima",
                    "sit",
                },
                CreatedByUserID: conductoronesdkterraform.String("vel"),
                Description: conductoronesdkterraform.String("laboriosam"),
                DisplayName: conductoronesdkterraform.String("quaerat"),
                ID: conductoronesdkterraform.String("1870d9d2-1f9a-4d03-8c4e-cc11a0836429"),
                Published: conductoronesdkterraform.Bool(false),
                RequestBundle: conductoronesdkterraform.Bool(false),
                VisibleToEveryone: conductoronesdkterraform.Bool(false),
            },
            RequestCatalogExpandMask: &shared.RequestCatalogExpandMask{
                Paths: []string{
                    "suscipit",
                },
            },
            UpdateMask: conductoronesdkterraform.String("laudantium"),
        },
        ID: "b8502a55-e7f7-43bc-845e-320a319f4bad",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogManagementServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `request`                                                                                                                                                                    | [operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateRequest](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdaterequest.md) | :heavy_check_mark:                                                                                                                                                           | The request object to use for the request.                                                                                                                                   |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogManagementServiceUpdateResponse](../../models/operations/c1apirequestcatalogv1requestcatalogmanagementserviceupdateresponse.md), error**

