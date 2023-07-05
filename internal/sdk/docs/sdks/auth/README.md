# Auth

### Available Operations

* [Introspect](#introspect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

## Introspect

Invokes the c1.api.auth.v1.Auth.Introspect method.

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
    res, err := s.Auth.Introspect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.IntrospectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIAuthV1AuthIntrospectResponse](../../models/operations/c1apiauthv1authintrospectresponse.md), error**

