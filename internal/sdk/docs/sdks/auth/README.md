# Auth
(*Auth*)

## Overview

### Available Operations

* [Introspect](#introspect) - Introspect

## Introspect

Introspect returns the current user's principle_id, user_id and a list of roles, permissions, and enabled features.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.auth.v1.Auth.Introspect" method="get" path="/api/v1/auth/introspect" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAuthV1AuthIntrospectResponse](../../models/operations/c1apiauthv1authintrospectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |