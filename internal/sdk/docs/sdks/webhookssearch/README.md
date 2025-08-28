# WebhooksSearch
(*WebhooksSearch*)

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Invokes the c1.api.webhooks.v1.WebhooksSearch.Search method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.webhooks.v1.WebhooksSearch.Search" method="post" path="/api/v1/search/webhooks" -->
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

    res, err := s.WebhooksSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhooksSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.WebhooksSearchRequest](../../models/components/webhookssearchrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.C1APIWebhooksV1WebhooksSearchSearchResponse](../../models/operations/c1apiwebhooksv1webhookssearchsearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |