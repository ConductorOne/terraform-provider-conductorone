# AppEntitlementSearch

### Available Operations

* [Search](#search) - Search
* [SearchAppEntitlementsWithExpired](#searchappentitlementswithexpired) - Search App Entitlements With Expired

## Search

Search app entitlements based on filters specified in the request body.

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
    res, err := s.AppEntitlementSearch.Search(ctx, shared.AppEntitlementSearchServiceSearchRequest{
        AppEntitlementExpandMask: &shared.AppEntitlementExpandMask{
            Paths: []string{
                "deserunt",
            },
        },
        AccessReviewID: conductoronesdkterraform.String("perferendis"),
        Alias: conductoronesdkterraform.String("ipsam"),
        AppIds: []string{
            "sapiente",
            "quo",
            "odit",
            "at",
        },
        AppUserIds: []string{
            "maiores",
            "molestiae",
            "quod",
            "quod",
        },
        ComplianceFrameworkIds: []string{
            "totam",
            "porro",
        },
        ExcludeAppIds: []string{
            "dicta",
            "nam",
            "officia",
        },
        ExcludeAppUserIds: []string{
            "fugit",
            "deleniti",
            "hic",
        },
        ExcludeResourceTypeIds: []string{
            "totam",
            "beatae",
            "commodi",
            "molestiae",
        },
        IncludeDeleted: conductoronesdkterraform.Bool(false),
        IsAutomated: conductoronesdkterraform.Bool(false),
        MembershipType: []shared.AppEntitlementSearchServiceSearchRequestMembershipType{
            shared.AppEntitlementSearchServiceSearchRequestMembershipTypeAppEntitlementMembershipTypeUnspecified,
            shared.AppEntitlementSearchServiceSearchRequestMembershipTypeAppEntitlementMembershipTypeExclusion,
        },
        OnlyGetExpiring: conductoronesdkterraform.Bool(false),
        PageSize: conductoronesdkterraform.Int(736918),
        PageToken: conductoronesdkterraform.String("esse"),
        Query: conductoronesdkterraform.String("ipsum"),
        Refs: []shared.AppEntitlementRef{
            shared.AppEntitlementRef{
                AppID: conductoronesdkterraform.String("aspernatur"),
                ID: conductoronesdkterraform.String("05929396-fea7-4596-ab10-faaa2352c595"),
            },
            shared.AppEntitlementRef{
                AppID: conductoronesdkterraform.String("minima"),
                ID: conductoronesdkterraform.String("907aff1a-3a2f-4a94-a773-9251aa52c3f5"),
            },
            shared.AppEntitlementRef{
                AppID: conductoronesdkterraform.String("id"),
                ID: conductoronesdkterraform.String("d019da1f-fe78-4f09-bb00-74f15471b5e6"),
            },
        },
        ResourceIds: []string{
            "quae",
            "ipsum",
            "quidem",
            "molestias",
        },
        ResourceTraitIds: []string{
            "pariatur",
            "modi",
            "praesentium",
        },
        ResourceTypeIds: []string{
            "voluptates",
            "quasi",
            "repudiandae",
        },
        RiskLevelIds: []string{
            "veritatis",
            "itaque",
            "incidunt",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppEntitlementSearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.AppEntitlementSearchServiceSearchRequest](../../models/shared/appentitlementsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchresponse.md), error**


## SearchAppEntitlementsWithExpired

Search app entitlements, include app users, users, expires, discovered.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppEntitlementSearch.SearchAppEntitlementsWithExpired(ctx, operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest{
        AppEntitlementID: "enim",
        AppID: "consequatur",
        PageSize: conductoronesdkterraform.Int(667411),
        PageToken: conductoronesdkterraform.String("quibusdam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchAppEntitlementsWithExpiredResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                          | Type                                                                                                                                                                                               | Required                                                                                                                                                                                           | Description                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                |
| `request`                                                                                                                                                                                          | [operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredRequest](../../models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementswithexpiredrequest.md) | :heavy_check_mark:                                                                                                                                                                                 | The request object to use for the request.                                                                                                                                                         |


### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementswithexpiredresponse.md), error**

