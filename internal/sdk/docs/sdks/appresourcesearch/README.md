# AppResourceSearch

### Available Operations

* [SearchAppResourceTypes](#searchappresourcetypes) - Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes method.

## SearchAppResourceTypes

Invokes the c1.api.app.v1.AppResourceSearch.SearchAppResourceTypes method.

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
    res, err := s.AppResourceSearch.SearchAppResourceTypes(ctx, shared.SearchAppResourceTypesRequest{
        AppIds: []string{
            "quidem",
            "architecto",
            "ipsa",
            "reiciendis",
        },
        ExcludeResourceTypeIds: []string{
            "mollitia",
            "laborum",
            "dolores",
        },
        ExcludeResourceTypeTraitIds: []string{
            "corporis",
        },
        PageSize: conductoroneapi.Float64(1289.26),
        PageToken: conductoroneapi.String("nobis"),
        Query: conductoroneapi.String("enim"),
        ResourceTypeIds: []string{
            "nemo",
            "minima",
            "excepturi",
        },
        ResourceTypeTraitIds: []string{
            "iure",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchAppResourceTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.SearchAppResourceTypesRequest](../../models/shared/searchappresourcetypesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.C1APIAppV1AppResourceSearchSearchAppResourceTypesResponse](../../models/operations/c1apiappv1appresourcesearchsearchappresourcetypesresponse.md), error**

