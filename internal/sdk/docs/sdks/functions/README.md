# Functions
(*Functions*)

## Overview

### Available Operations

* [ListFunctions](#listfunctions) - List Functions
* [CreateFunction](#createfunction) - Create Function
* [ListCommits](#listcommits) - List Commits
* [Commit](#commit) - Commit
* [GetCommit](#getcommit) - Get Commit
* [Invoke](#invoke) - Invoke
* [ListTags](#listtags) - List Tags
* [CreateTag](#createtag) - Create Tag
* [DeleteFunction](#deletefunction) - Delete Function
* [GetFunction](#getfunction) - Get Function
* [UpdateFunction](#updatefunction) - Update Function

## ListFunctions

List retrieves all functions with pagination

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.ListFunctions" method="get" path="/api/v1/functions" -->
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

    res, err := s.Functions.ListFunctions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceListFunctionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceListFunctionsResponse](../../models/operations/c1apifunctionsv1functionsservicelistfunctionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateFunction

Invokes the c1.api.functions.v1.FunctionsService.CreateFunction method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.CreateFunction" method="post" path="/api/v1/functions" -->
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

    res, err := s.Functions.CreateFunction(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceCreateFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [components.FunctionsServiceCreateFunctionRequest](../../models/components/functionsservicecreatefunctionrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceCreateFunctionResponse](../../models/operations/c1apifunctionsv1functionsservicecreatefunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListCommits

ListCommits retrieves the commit history

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.ListCommits" method="get" path="/api/v1/functions/{function_id}/commits" -->
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

    res, err := s.Functions.ListCommits(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceListCommitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `functionID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceListCommitsResponse](../../models/operations/c1apifunctionsv1functionsservicelistcommitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Commit

Commit saves a new version of the function code

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.Commit" method="post" path="/api/v1/functions/{function_id}/commits" -->
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

    res, err := s.Functions.Commit(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceCommitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `functionID`                                                                                          | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `functionsServiceCommitRequest`                                                                       | [*components.FunctionsServiceCommitRequest](../../models/components/functionsservicecommitrequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceCommitResponse](../../models/operations/c1apifunctionsv1functionsservicecommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCommit

GetCommit retrieves the commit and its code content for a specific version

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.GetCommit" method="get" path="/api/v1/functions/{function_id}/commits/{id}" -->
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

    res, err := s.Functions.GetCommit(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceGetCommitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `functionID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceGetCommitResponse](../../models/operations/c1apifunctionsv1functionsservicegetcommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Invoke

Invokes the c1.api.functions.v1.FunctionsService.Invoke method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.Invoke" method="post" path="/api/v1/functions/{function_id}/invoke" -->
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

    res, err := s.Functions.Invoke(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceInvokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `functionID`                                                                                          | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `functionsServiceInvokeRequest`                                                                       | [*components.FunctionsServiceInvokeRequest](../../models/components/functionsserviceinvokerequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceInvokeResponse](../../models/operations/c1apifunctionsv1functionsserviceinvokeresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListTags

ListTags lists all tags for a function

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.ListTags" method="get" path="/api/v1/functions/{function_id}/tags" -->
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

    res, err := s.Functions.ListTags(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceListTagsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `functionID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceListTagsResponse](../../models/operations/c1apifunctionsv1functionsservicelisttagsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateTag

CreateTag creates a named reference to a specific commit

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.CreateTag" method="post" path="/api/v1/functions/{function_id}/tags" -->
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

    res, err := s.Functions.CreateTag(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceCreateTagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `functionID`                                                                                                | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `functionsServiceCreateTagRequest`                                                                          | [*components.FunctionsServiceCreateTagRequest](../../models/components/functionsservicecreatetagrequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceCreateTagResponse](../../models/operations/c1apifunctionsv1functionsservicecreatetagresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteFunction

Delete removes a function

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.DeleteFunction" method="delete" path="/api/v1/functions/{id}" -->
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

    res, err := s.Functions.DeleteFunction(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceDeleteFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                             | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                 | :heavy_check_mark:                                                                                                    | The context to use for the request.                                                                                   |
| `id`                                                                                                                  | *string*                                                                                                              | :heavy_check_mark:                                                                                                    | N/A                                                                                                                   |
| `functionsServiceDeleteFunctionRequest`                                                                               | [*components.FunctionsServiceDeleteFunctionRequest](../../models/components/functionsservicedeletefunctionrequest.md) | :heavy_minus_sign:                                                                                                    | N/A                                                                                                                   |
| `opts`                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                              | :heavy_minus_sign:                                                                                                    | The options for this request.                                                                                         |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceDeleteFunctionResponse](../../models/operations/c1apifunctionsv1functionsservicedeletefunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetFunction

Get retrieves a specific function by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.GetFunction" method="get" path="/api/v1/functions/{id}" -->
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

    res, err := s.Functions.GetFunction(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceGetFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceGetFunctionResponse](../../models/operations/c1apifunctionsv1functionsservicegetfunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateFunction

Update updates an existing function's metadata

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.functions.v1.FunctionsService.UpdateFunction" method="post" path="/api/v1/functions/update" -->
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

    res, err := s.Functions.UpdateFunction(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.FunctionsServiceUpdateFunctionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [components.FunctionsServiceUpdateFunctionRequest](../../models/components/functionsserviceupdatefunctionrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.C1APIFunctionsV1FunctionsServiceUpdateFunctionResponse](../../models/operations/c1apifunctionsv1functionsserviceupdatefunctionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |