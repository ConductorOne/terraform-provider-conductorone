# AppEntitlementUserBinding
(*AppEntitlementUserBinding*)

## Overview

### Available Operations

* [RemoveGrantDuration](#removegrantduration) - Remove Grant Duration
* [UpdateGrantDuration](#updategrantduration) - Update Grant Duration
* [ListAppUsersForIdentityWithGrant](#listappusersforidentitywithgrant) - List App Users For Identity With Grant
* [SearchGrantFeed](#searchgrantfeed) - Search Grant Feed
* [SearchPastGrants](#searchpastgrants) - Search Past Grants

## RemoveGrantDuration

Invokes the c1.api.app.v1.AppEntitlementUserBindingService.RemoveGrantDuration method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementUserBindingService.RemoveGrantDuration" method="post" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/users/{app_user_id}/remove-grant-duration" -->
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

    res, err := s.AppEntitlementUserBinding.RemoveGrantDuration(ctx, "<id>", "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveGrantDurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |
| `appID`                                                                                         | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `appEntitlementID`                                                                              | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `appUserID`                                                                                     | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `removeGrantDurationRequest`                                                                    | [*components.RemoveGrantDurationRequest](../../models/components/removegrantdurationrequest.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |

### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceRemoveGrantDurationResponse](../../models/operations/c1apiappv1appentitlementuserbindingserviceremovegrantdurationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateGrantDuration

Invokes the c1.api.app.v1.AppEntitlementUserBindingService.UpdateGrantDuration method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementUserBindingService.UpdateGrantDuration" method="post" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/users/{app_user_id}/update-grant-duration" -->
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

    res, err := s.AppEntitlementUserBinding.UpdateGrantDuration(ctx, "<id>", "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateGrantDurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |
| `appID`                                                                                         | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `appEntitlementID`                                                                              | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `appUserID`                                                                                     | *string*                                                                                        | :heavy_check_mark:                                                                              | N/A                                                                                             |
| `updateGrantDurationRequest`                                                                    | [*components.UpdateGrantDurationRequest](../../models/components/updategrantdurationrequest.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |

### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceUpdateGrantDurationResponse](../../models/operations/c1apiappv1appentitlementuserbindingserviceupdategrantdurationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAppUsersForIdentityWithGrant

Returns a list of app users for the identity in the app. If that app user also has a grant to the entitlement from the request, data about the grant is also returned. It will always return ALL app users for this identity, but only SOME may have grant data.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant" method="get" path="/api/v1/apps/{app_id}/entitlements/{app_entitlement_id}/users/{identity_user_id}/grants" -->
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

    res, err := s.AppEntitlementUserBinding.ListAppUsersForIdentityWithGrant(ctx, "<id>", "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppUsersForIdentityWithGrantResponse != nil {
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
| `identityUserID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantResponse](../../models/operations/c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrantresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SearchGrantFeed

Invokes the c1.api.app.v1.AppEntitlementUserBindingService.SearchGrantFeed method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementUserBindingService.SearchGrantFeed" method="post" path="/api/v1/grants/feed" -->
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

    res, err := s.AppEntitlementUserBinding.SearchGrantFeed(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchGrantFeedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.SearchGrantFeedRequest](../../models/components/searchgrantfeedrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceSearchGrantFeedResponse](../../models/operations/c1apiappv1appentitlementuserbindingservicesearchgrantfeedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SearchPastGrants

Invokes the c1.api.app.v1.AppEntitlementUserBindingService.SearchPastGrants method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementUserBindingService.SearchPastGrants" method="post" path="/api/v1/search/past-grants" -->
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

    res, err := s.AppEntitlementUserBinding.SearchPastGrants(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchPastGrantsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.SearchPastGrantsRequest](../../models/components/searchpastgrantsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceSearchPastGrantsResponse](../../models/operations/c1apiappv1appentitlementuserbindingservicesearchpastgrantsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |