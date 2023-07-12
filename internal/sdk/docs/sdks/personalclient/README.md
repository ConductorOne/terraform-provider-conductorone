# PersonalClient

### Available Operations

* [Create](#create) - Invokes the c1.api.iam.v1.PersonalClientService.Create method.

## Create

Invokes the c1.api.iam.v1.PersonalClientService.Create method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.PersonalClient.Create(ctx, shared.PersonalClientServiceCreateRequest{
        AllowSourceCidr: []string{
            "modi",
            "qui",
        },
        DisplayName: conductoroneapi.String("aliquid"),
        Expires: conductoroneapi.String("cupiditate"),
        ScopedRoles: []string{
            "perferendis",
            "magni",
            "assumenda",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PersonalClientServiceCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [shared.PersonalClientServiceCreateRequest](../../models/shared/personalclientservicecreaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.C1APIIamV1PersonalClientServiceCreateResponse](../../models/operations/c1apiiamv1personalclientservicecreateresponse.md), error**

