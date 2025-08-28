# AccountProvisionPolicyTest
(*AccountProvisionPolicyTest*)

## Overview

### Available Operations

* [Test](#test) - Test

## Test

Invokes the c1.api.policy.v1.AccountProvisionPolicyTest.Test method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.policy.v1.AccountProvisionPolicyTest.Test" method="post" path="/api/v1/policies/test-account-provision-policy" -->
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

    res, err := s.AccountProvisionPolicyTest.Test(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TestAccountProvisionPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [components.TestAccountProvisionPolicyRequest](../../models/components/testaccountprovisionpolicyrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.C1APIPolicyV1AccountProvisionPolicyTestTestResponse](../../models/operations/c1apipolicyv1accountprovisionpolicytesttestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |