# SessionSettings

### Available Operations

* [Get](#get) - Get
* [Update](#update) - Update

## Get

Invokes the c1.api.settings.v1.SessionSettingsService.Get method.

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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APISettingsV1SessionSettingsServiceGetResponse](../../models/operations/c1apisettingsv1sessionsettingsservicegetresponse.md), error**


## Update

Invokes the c1.api.settings.v1.SessionSettingsService.Update method.

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
    res, err := s.SessionSettings.Update(ctx, shared.UpdateSessionSettingsRequest{
        SessionSettings: &shared.SessionSettings{
            MaxSessionLength: conductoronesdkterraform.String("suscipit"),
        },
        UpdateMask: conductoronesdkterraform.String("quibusdam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSessionSettingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.UpdateSessionSettingsRequest](../../models/shared/updatesessionsettingsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APISettingsV1SessionSettingsServiceUpdateResponse](../../models/operations/c1apisettingsv1sessionsettingsserviceupdateresponse.md), error**

