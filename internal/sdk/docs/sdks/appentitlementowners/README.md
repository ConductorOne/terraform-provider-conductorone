# AppEntitlementOwners
(*AppEntitlementOwners*)

## Overview

### Available Operations

* [List](#list) - List
* [Add](#add) - Add
* [Set](#set) - Set
* [Remove](#remove) - Remove

## List

List owners for a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.List" method="get" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwners.List(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `entitlementID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersListResponse](../../models/operations/c1apiappv1appentitlementownerslistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Add

Add an owner to a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.Add" method="post" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwners.Add(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AddAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `appID`                                                                                               | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `entitlementID`                                                                                       | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `addAppEntitlementOwnerRequest`                                                                       | [*components.AddAppEntitlementOwnerRequest](../../models/components/addappentitlementownerrequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersAddResponse](../../models/operations/c1apiappv1appentitlementownersaddresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Set

Sets the owners for a given app entitlement to the specified list of users.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.Set" method="put" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners" -->
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

    res, err := s.AppEntitlementOwners.Set(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SetAppEntitlementOwnersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |
| `appID`                                                                                                 | *string*                                                                                                | :heavy_check_mark:                                                                                      | N/A                                                                                                     |
| `entitlementID`                                                                                         | *string*                                                                                                | :heavy_check_mark:                                                                                      | N/A                                                                                                     |
| `setAppEntitlementOwnersRequest`                                                                        | [*components.SetAppEntitlementOwnersRequest](../../models/components/setappentitlementownersrequest.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersSetResponse](../../models/operations/c1apiappv1appentitlementownerssetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Remove

Remove an owner from a given app entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppEntitlementOwners.Remove" method="delete" path="/api/v1/apps/{app_id}/entitlements/{entitlement_id}/owners/{user_id}" -->
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

    res, err := s.AppEntitlementOwners.Remove(ctx, "<id>", "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.RemoveAppEntitlementOwnerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `appID`                                                                                                     | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `entitlementID`                                                                                             | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `userID`                                                                                                    | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `removeAppEntitlementOwnerRequest`                                                                          | [*components.RemoveAppEntitlementOwnerRequest](../../models/components/removeappentitlementownerrequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.C1APIAppV1AppEntitlementOwnersRemoveResponse](../../models/operations/c1apiappv1appentitlementownersremoveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |