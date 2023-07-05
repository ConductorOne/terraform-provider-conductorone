# AppSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.app.v1.AppSearch.Search method.

## Search

Invokes the c1.api.app.v1.AppSearch.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.AppSearch.Search(ctx, shared.SearchAppsRequest{
        AppIds: []string{
            "laborum",
            "quasi",
            "reiciendis",
            "voluptatibus",
        },
        ExcludeAppIds: []string{
            "nihil",
            "praesentium",
            "voluptatibus",
            "ipsa",
        },
        PageSize: conductoroneapi.Float64(6048.46),
        PageToken: conductoroneapi.String("voluptate"),
        Query: conductoroneapi.String("cum"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.SearchAppsRequest](../../models/shared/searchappsrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.C1APIAppV1AppSearchSearchResponse](../../models/operations/c1apiappv1appsearchsearchresponse.md), error**

