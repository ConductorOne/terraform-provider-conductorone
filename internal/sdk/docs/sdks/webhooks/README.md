# Webhooks

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [Test](#test) - Test
* [Update](#update) - Update

## Create

Invokes the c1.api.webhooks.v1.WebhooksService.Create method.

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
    res, err := s.Webhooks.Create(ctx, shared.WebhooksServiceCreateRequest{
        Description: conductoronesdkterraform.String("sapiente"),
        DisplayName: conductoronesdkterraform.String("alias"),
        URL: conductoronesdkterraform.String("impedit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.WebhooksServiceCreateRequest](../../models/shared/webhooksservicecreaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APIWebhooksV1WebhooksServiceCreateResponse](../../models/operations/c1apiwebhooksv1webhooksservicecreateresponse.md), error**


## Delete

Invokes the c1.api.webhooks.v1.WebhooksService.Delete method.

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
    res, err := s.Webhooks.Delete(ctx, operations.C1APIWebhooksV1WebhooksServiceDeleteRequest{
        WebhooksServiceDeleteRequest: &shared.WebhooksServiceDeleteRequest{},
        ID: "42b78f15-6263-498a-8dc7-66324ccb06c8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksServiceDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIWebhooksV1WebhooksServiceDeleteRequest](../../models/operations/c1apiwebhooksv1webhooksservicedeleterequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APIWebhooksV1WebhooksServiceDeleteResponse](../../models/operations/c1apiwebhooksv1webhooksservicedeleteresponse.md), error**


## Get

Invokes the c1.api.webhooks.v1.WebhooksService.Get method.

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
    res, err := s.Webhooks.Get(ctx, operations.C1APIWebhooksV1WebhooksServiceGetRequest{
        ID: "ca12d025-2927-40b8-9572-2dd895b8bcf2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.C1APIWebhooksV1WebhooksServiceGetRequest](../../models/operations/c1apiwebhooksv1webhooksservicegetrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.C1APIWebhooksV1WebhooksServiceGetResponse](../../models/operations/c1apiwebhooksv1webhooksservicegetresponse.md), error**


## List

Invokes the c1.api.webhooks.v1.WebhooksService.List method.

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
    res, err := s.Webhooks.List(ctx, operations.C1APIWebhooksV1WebhooksServiceListRequest{
        PageSize: conductoronesdkterraform.Int(295307),
        PageToken: conductoronesdkterraform.String("pariatur"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIWebhooksV1WebhooksServiceListRequest](../../models/operations/c1apiwebhooksv1webhooksservicelistrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.C1APIWebhooksV1WebhooksServiceListResponse](../../models/operations/c1apiwebhooksv1webhooksservicelistresponse.md), error**


## Test

Invokes the c1.api.webhooks.v1.WebhooksService.Test method.

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
    res, err := s.Webhooks.Test(ctx, operations.C1APIWebhooksV1WebhooksServiceTestRequest{
        WebhooksServiceTestRequest: &shared.WebhooksServiceTestRequest{},
        ID: "b9596933-52f7-4453-b994-d78de3b6e938",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksServiceTestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.C1APIWebhooksV1WebhooksServiceTestRequest](../../models/operations/c1apiwebhooksv1webhooksservicetestrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |


### Response

**[*operations.C1APIWebhooksV1WebhooksServiceTestResponse](../../models/operations/c1apiwebhooksv1webhooksservicetestresponse.md), error**


## Update

Invokes the c1.api.webhooks.v1.WebhooksService.Update method.

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
    res, err := s.Webhooks.Update(ctx, operations.C1APIWebhooksV1WebhooksServiceUpdateRequest{
        WebhooksServiceUpdateRequestInput: &shared.WebhooksServiceUpdateRequestInput{
            Webhook: &shared.WebhookInput{
                Description: conductoronesdkterraform.String("cupiditate"),
                DisplayName: conductoronesdkterraform.String("voluptatibus"),
                ID: conductoronesdkterraform.String("5abb7f66-2550-4a28-b82a-c483afd2315b"),
                URL: conductoronesdkterraform.String("facilis"),
            },
            UpdateMask: conductoronesdkterraform.String("deserunt"),
        },
        ID: "650164e0-6f5b-4f6a-a591-bc8bdef3612b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.C1APIWebhooksV1WebhooksServiceUpdateRequest](../../models/operations/c1apiwebhooksv1webhooksserviceupdaterequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.C1APIWebhooksV1WebhooksServiceUpdateResponse](../../models/operations/c1apiwebhooksv1webhooksserviceupdateresponse.md), error**

