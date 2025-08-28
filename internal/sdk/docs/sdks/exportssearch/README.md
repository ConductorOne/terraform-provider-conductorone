# ExportsSearch
(*ExportsSearch*)

## Overview

### Available Operations

* [Search](#search) - Search

## Search

Invokes the c1.api.systemlog.v1.ExportsSearchService.Search method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.systemlog.v1.ExportsSearchService.Search" method="post" path="/api/v1/search/systemlog/exports" -->
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

    res, err := s.ExportsSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ExportsSearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [components.ExportsSearchServiceSearchRequest](../../models/components/exportssearchservicesearchrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.C1APISystemlogV1ExportsSearchServiceSearchResponse](../../models/operations/c1apisystemlogv1exportssearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |