# AppEntitlementSearch
(*AppEntitlementSearch*)

## Overview

### Available Operations

* [SearchAppEntitlementsWithExpired](#searchappentitlementswithexpired) - Search App Entitlements With Expired
* [SearchAppEntitlementsForAppUser](#searchappentitlementsforappuser) - Search App Entitlements For App User
* [Search](#search) - Search
* [SearchGrants](#searchgrants) - Search Grants

## SearchAppEntitlementsWithExpired

Search app entitlements, include app users, users, expires, discovered.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsWithExpired" method="get" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/grants" -->
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

    res, err := s.AppEntitlementSearch.SearchAppEntitlementsWithExpired(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAppEntitlementsWithExpiredResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `appEntitlementID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsWithExpiredResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementswithexpiredresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SearchAppEntitlementsForAppUser

Invokes the c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsForAppUser method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchAppEntitlementsForAppUser" method="get" path="/api/v1/search/apps/{app_id}/entitlements/users/{app_user_id}" -->
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

    res, err := s.AppEntitlementSearch.SearchAppEntitlementsForAppUser(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `appUserID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchAppEntitlementsForAppUserResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchappentitlementsforappuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Search

Search app entitlements based on filters specified in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.Search" method="post" path="/api/v1/search/entitlements" -->
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

    res, err := s.AppEntitlementSearch.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceSearchResponse != nil {
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [components.AppEntitlementSearchServiceSearchRequest](../../models/components/appentitlementsearchservicesearchrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SearchGrants

Invokes the c1.api.app.v1.AppEntitlementSearchService.SearchGrants method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementSearchService.SearchGrants" method="post" path="/api/v1/search/grants" -->
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

    res, err := s.AppEntitlementSearch.SearchGrants(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppEntitlementSearchServiceSearchGrantsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [components.AppEntitlementSearchServiceSearchGrantsRequest](../../models/components/appentitlementsearchservicesearchgrantsrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIAppV1AppEntitlementSearchServiceSearchGrantsResponse](../../models/operations/c1apiappv1appentitlementsearchservicesearchgrantsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |