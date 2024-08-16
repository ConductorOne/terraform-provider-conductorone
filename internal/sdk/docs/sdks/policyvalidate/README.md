# PolicyValidate

### Available Operations

* [ValidateCEL](#validatecel) - Validate Cel

## ValidateCEL

Validate policies

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
    res, err := s.PolicyValidate.ValidateCEL(ctx, shared.ValidatePolicyCELRequest{
        Cel: conductoronesdkterraform.String("magni"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ValidatePolicyCELResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.ValidatePolicyCELRequest](../../models/shared/validatepolicycelrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.C1APIPolicyV1PolicyValidateValidateCELResponse](../../models/operations/c1apipolicyv1policyvalidatevalidatecelresponse.md), error**

