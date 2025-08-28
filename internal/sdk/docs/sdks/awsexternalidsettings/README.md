# AWSExternalIDSettings
(*AWSExternalIDSettings*)

## Overview

### Available Operations

* [Get](#get) - Get

## Get

Invokes the c1.api.settings.v1.AWSExternalIDSettings.Get method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.settings.v1.AWSExternalIDSettings.Get" method="get" path="/api/v1/settings/aws-external-id" -->
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

    res, err := s.AWSExternalIDSettings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAWSExternalIDResponse != nil {
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

**[*operations.C1APISettingsV1AWSExternalIDSettingsGetResponse](../../models/operations/c1apisettingsv1awsexternalidsettingsgetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |