# AppResourceSearch

### Available Operations

* [SearchAppResourceTypes](#searchappresourcetypes) - Search App Resource Types

## SearchAppResourceTypes

Search app resources based on filters specified in the request body.

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
    res, err := s.AppResourceSearch.SearchAppResourceTypes(ctx, shared.SearchAppResourceTypesRequest{
        AppIds: []string{
            "rerum",
            "adipisci",
            "asperiores",
        },
        AppUserIds: []string{
            "modi",
            "iste",
            "dolorum",
            "deleniti",
        },
        ExcludeResourceTypeIds: []string{
            "provident",
            "nobis",
            "libero",
            "delectus",
        },
        ExcludeResourceTypeTraitIds: []string{
            "quos",
            "aliquid",
        },
        PageSize: conductoronesdkterraform.Int(212390),
        PageToken: conductoronesdkterraform.String("dolorem"),
        Query: conductoronesdkterraform.String("dolor"),
        ResourceTypeIds: []string{
            "ipsum",
        },
        ResourceTypeTraitIds: []string{
            "excepturi",
            "cum",
            "voluptate",
            "dignissimos",
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

