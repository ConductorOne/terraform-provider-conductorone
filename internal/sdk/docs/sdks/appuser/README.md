# AppUser
(*AppUser*)

## Overview

### Available Operations

* [List](#list) - List
* [ListAppUserCredentials](#listappusercredentials) - List App User Credentials
* [ListAppUsersForUser](#listappusersforuser) - List App Users For User
* [Update](#update) - Update
* [Search](#search) - Search

## List

Invokes the c1.api.app.v1.AppUserService.List method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserService.List" method="get" path="/api/v1/apps/{app_id}/app_users" -->
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

    res, err := s.AppUser.List(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUserServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppUserServiceListResponse](../../models/operations/c1apiappv1appuserservicelistresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAppUserCredentials

Invokes the c1.api.app.v1.AppUserService.ListAppUserCredentials method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserService.ListAppUserCredentials" method="get" path="/api/v1/apps/{app_id}/app_users/{app_user_id}/credentials" -->
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

    res, err := s.AppUser.ListAppUserCredentials(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUserServiceListCredentialsResponse != nil {
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

**[*operations.C1APIAppV1AppUserServiceListAppUserCredentialsResponse](../../models/operations/c1apiappv1appuserservicelistappusercredentialsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAppUsersForUser

Invokes the c1.api.app.v1.AppUserService.ListAppUsersForUser method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserService.ListAppUsersForUser" method="get" path="/api/v1/apps/{app_id}/users/{user_id}/app_users" -->
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

    res, err := s.AppUser.ListAppUsersForUser(ctx, "<id>", "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUsersForUserServiceListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `appID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAppV1AppUserServiceListAppUsersForUserResponse](../../models/operations/c1apiappv1appuserservicelistappusersforuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an app user by ID. Only the fields specified in the update mask are updated.
 Currently, only the appUserType, and identityUserId fields can be updated.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserService.Update" method="post" path="/api/v1/apps/{app_user_app_id}/app_users/{app_user_id}" -->
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

    res, err := s.AppUser.Update(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUserServiceUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `appUserAppID`                                                                                    | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `appUserID`                                                                                       | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `appUserServiceUpdateRequest`                                                                     | [*components.AppUserServiceUpdateRequest](../../models/components/appuserserviceupdaterequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.C1APIAppV1AppUserServiceUpdateResponse](../../models/operations/c1apiappv1appuserserviceupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Search

Search app users based on filters specified in the request body.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.app.v1.AppUserService.Search" method="post" path="/api/v1/search/app_users" -->
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

    res, err := s.AppUser.Search(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AppUserServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.AppUserServiceSearchRequest](../../models/components/appuserservicesearchrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.C1APIAppV1AppUserServiceSearchResponse](../../models/operations/c1apiappv1appuserservicesearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |