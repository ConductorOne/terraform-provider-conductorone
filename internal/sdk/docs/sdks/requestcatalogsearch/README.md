# RequestCatalogSearch

### Available Operations

* [SearchEntitlements](#searchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

## SearchEntitlements

Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

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
    res, err := s.RequestCatalogSearch.SearchEntitlements(ctx, shared.RequestCatalogSearchServiceSearchEntitlementsRequest{
        EntitlementAlias: conductoroneapi.String("laborum"),
        ExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "incidunt",
                "aspernatur",
                "dolores",
            },
        },
        PageSize: conductoroneapi.Float64(7168.6),
        PageToken: conductoroneapi.String("facilis"),
        Query: conductoroneapi.String("aliquid"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RequestCatalogSearchServiceSearchEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [shared.RequestCatalogSearchServiceSearchEntitlementsRequest](../../models/shared/requestcatalogsearchservicesearchentitlementsrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogsearchservicesearchentitlementsresponse.md), error**

