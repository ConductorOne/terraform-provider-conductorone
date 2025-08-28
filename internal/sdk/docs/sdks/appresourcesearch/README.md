# AppResourceSearch
(*AppResourceSearch*)

## Overview

### Available Operations

* [SearchAppResourceTypes](#searchappresourcetypes) - Search App Resource Types
* [SearchAppResources](#searchappresources) - Search App Resources

## SearchAppResourceTypes

Search app resources based on filters specified in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes" method="post" path="/api/v1/search/app_resource_types" -->
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

    res, err := s.AppResourceSearch.SearchAppResourceTypes(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppResourceTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.SearchAppResourceTypesRequest](../../models/components/searchappresourcetypesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.C1APIAppV1AppResourceSearchSearchAppResourceTypesResponse](../../models/operations/c1apiappv1appresourcesearchsearchappresourcetypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SearchAppResources

Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResources method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppResourceSearch.SearchAppResources" method="post" path="/api/v1/search/app_resources" -->
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

    res, err := s.AppResourceSearch.SearchAppResources(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppResourcesResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.SearchAppResourcesRequest](../../models/components/searchappresourcesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.C1APIAppV1AppResourceSearchSearchAppResourcesResponse](../../models/operations/c1apiappv1appresourcesearchsearchappresourcesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |