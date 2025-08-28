# TaskActions
(*TaskActions*)

## Overview

### Available Operations

* [Approve](#approve) - Approve
* [ApproveWithStepUp](#approvewithstepup) - Approve With Step Up
* [Close](#close) - Close
* [Comment](#comment) - Comment
* [Deny](#deny) - Deny
* [EscalateToEmergencyAccess](#escalatetoemergencyaccess) - Escalate To Emergency Access
* [ProcessNow](#processnow) - Process Now
* [Reassign](#reassign) - Reassign
* [HardReset](#hardreset) - Hard Reset
* [Restart](#restart) - Restart
* [SkipStep](#skipstep) - Skip Step
* [UpdateGrantDuration](#updategrantduration) - Update Grant Duration
* [UpdateRequestData](#updaterequestdata) - Update Request Data

## Approve

Invokes the c1.api.task.v1.TaskActionsService.Approve method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Approve" method="post" path="/api/v1/tasks/{task_id}/action/approve" -->
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

    res, err := s.TaskActions.Approve(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceApproveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `taskID`                                                                                                    | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `taskActionsServiceApproveRequest`                                                                          | [*components.TaskActionsServiceApproveRequest](../../models/components/taskactionsserviceapproverequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.C1APITaskV1TaskActionsServiceApproveResponse](../../models/operations/c1apitaskv1taskactionsserviceapproveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ApproveWithStepUp

Invokes the c1.api.task.v1.TaskActionsService.ApproveWithStepUp method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.ApproveWithStepUp" method="post" path="/api/v1/tasks/{task_id}/action/approve-with-step-up" -->
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

    res, err := s.TaskActions.ApproveWithStepUp(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceApproveWithStepUpResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |
| `taskID`                                                                                                                        | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | N/A                                                                                                                             |
| `taskActionsServiceApproveWithStepUpRequest`                                                                                    | [*components.TaskActionsServiceApproveWithStepUpRequest](../../models/components/taskactionsserviceapprovewithstepuprequest.md) | :heavy_minus_sign:                                                                                                              | N/A                                                                                                                             |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |

### Response

**[*operations.C1APITaskV1TaskActionsServiceApproveWithStepUpResponse](../../models/operations/c1apitaskv1taskactionsserviceapprovewithstepupresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Close

Invokes the c1.api.task.v1.TaskActionsService.Close method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Close" method="post" path="/api/v1/tasks/{task_id}/action/close" -->
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

    res, err := s.TaskActions.Close(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceCloseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |
| `taskID`                                                                                                | *string*                                                                                                | :heavy_check_mark:                                                                                      | N/A                                                                                                     |
| `taskActionsServiceCloseRequest`                                                                        | [*components.TaskActionsServiceCloseRequest](../../models/components/taskactionsservicecloserequest.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |

### Response

**[*operations.C1APITaskV1TaskActionsServiceCloseResponse](../../models/operations/c1apitaskv1taskactionsservicecloseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Comment

Invokes the c1.api.task.v1.TaskActionsService.Comment method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Comment" method="post" path="/api/v1/tasks/{task_id}/action/comment" -->
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

    res, err := s.TaskActions.Comment(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceCommentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `taskID`                                                                                                    | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `taskActionsServiceCommentRequest`                                                                          | [*components.TaskActionsServiceCommentRequest](../../models/components/taskactionsservicecommentrequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.C1APITaskV1TaskActionsServiceCommentResponse](../../models/operations/c1apitaskv1taskactionsservicecommentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Deny

Invokes the c1.api.task.v1.TaskActionsService.Deny method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Deny" method="post" path="/api/v1/tasks/{task_id}/action/deny" -->
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

    res, err := s.TaskActions.Deny(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceDenyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |
| `taskID`                                                                                              | *string*                                                                                              | :heavy_check_mark:                                                                                    | N/A                                                                                                   |
| `taskActionsServiceDenyRequest`                                                                       | [*components.TaskActionsServiceDenyRequest](../../models/components/taskactionsservicedenyrequest.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |

### Response

**[*operations.C1APITaskV1TaskActionsServiceDenyResponse](../../models/operations/c1apitaskv1taskactionsservicedenyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## EscalateToEmergencyAccess

Invokes the c1.api.task.v1.TaskActionsService.EscalateToEmergencyAccess method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.EscalateToEmergencyAccess" method="post" path="/api/v1/tasks/{task_id}/action/escalate" -->
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

    res, err := s.TaskActions.EscalateToEmergencyAccess(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                       | Type                                                                                                                                            | Required                                                                                                                                        | Description                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                           | :heavy_check_mark:                                                                                                                              | The context to use for the request.                                                                                                             |
| `taskID`                                                                                                                                        | *string*                                                                                                                                        | :heavy_check_mark:                                                                                                                              | N/A                                                                                                                                             |
| `taskActionsServiceEscalateToEmergencyAccessRequest`                                                                                            | [*components.TaskActionsServiceEscalateToEmergencyAccessRequest](../../models/components/taskactionsserviceescalatetoemergencyaccessrequest.md) | :heavy_minus_sign:                                                                                                                              | N/A                                                                                                                                             |
| `opts`                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                        | :heavy_minus_sign:                                                                                                                              | The options for this request.                                                                                                                   |

### Response

**[*operations.C1APITaskV1TaskActionsServiceEscalateToEmergencyAccessResponse](../../models/operations/c1apitaskv1taskactionsserviceescalatetoemergencyaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ProcessNow

Invokes the c1.api.task.v1.TaskActionsService.ProcessNow method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.ProcessNow" method="post" path="/api/v1/tasks/{task_id}/action/process" -->
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

    res, err := s.TaskActions.ProcessNow(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceProcessNowResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |
| `taskID`                                                                                                          | *string*                                                                                                          | :heavy_check_mark:                                                                                                | N/A                                                                                                               |
| `taskActionsServiceProcessNowRequest`                                                                             | [*components.TaskActionsServiceProcessNowRequest](../../models/components/taskactionsserviceprocessnowrequest.md) | :heavy_minus_sign:                                                                                                | N/A                                                                                                               |
| `opts`                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                          | :heavy_minus_sign:                                                                                                | The options for this request.                                                                                     |

### Response

**[*operations.C1APITaskV1TaskActionsServiceProcessNowResponse](../../models/operations/c1apitaskv1taskactionsserviceprocessnowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Reassign

Invokes the c1.api.task.v1.TaskActionsService.Reassign method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Reassign" method="post" path="/api/v1/tasks/{task_id}/action/reassign" -->
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

    res, err := s.TaskActions.Reassign(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceReassignResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |
| `taskID`                                                                                                      | *string*                                                                                                      | :heavy_check_mark:                                                                                            | N/A                                                                                                           |
| `taskActionsServiceReassignRequest`                                                                           | [*components.TaskActionsServiceReassignRequest](../../models/components/taskactionsservicereassignrequest.md) | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `opts`                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                      | :heavy_minus_sign:                                                                                            | The options for this request.                                                                                 |

### Response

**[*operations.C1APITaskV1TaskActionsServiceReassignResponse](../../models/operations/c1apitaskv1taskactionsservicereassignresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## HardReset

Invokes the c1.api.task.v1.TaskActionsService.HardReset method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.HardReset" method="post" path="/api/v1/tasks/{task_id}/action/reset" -->
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

    res, err := s.TaskActions.HardReset(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceHardResetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |
| `taskID`                                                                                                        | *string*                                                                                                        | :heavy_check_mark:                                                                                              | N/A                                                                                                             |
| `taskActionsServiceHardResetRequest`                                                                            | [*components.TaskActionsServiceHardResetRequest](../../models/components/taskactionsservicehardresetrequest.md) | :heavy_minus_sign:                                                                                              | N/A                                                                                                             |
| `opts`                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                        | :heavy_minus_sign:                                                                                              | The options for this request.                                                                                   |

### Response

**[*operations.C1APITaskV1TaskActionsServiceHardResetResponse](../../models/operations/c1apitaskv1taskactionsservicehardresetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Restart

Invokes the c1.api.task.v1.TaskActionsService.Restart method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.Restart" method="post" path="/api/v1/tasks/{task_id}/action/restart" -->
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

    res, err := s.TaskActions.Restart(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskActionsServiceRestartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `taskID`                                                                                                    | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `taskActionsServiceRestartRequest`                                                                          | [*components.TaskActionsServiceRestartRequest](../../models/components/taskactionsservicerestartrequest.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.C1APITaskV1TaskActionsServiceRestartResponse](../../models/operations/c1apitaskv1taskactionsservicerestartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SkipStep

Invokes the c1.api.task.v1.TaskActionsService.SkipStep method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.SkipStep" method="post" path="/api/v1/tasks/{task_id}/action/skip-step" -->
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

    res, err := s.TaskActions.SkipStep(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                     | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                         | :heavy_check_mark:                                                                                            | The context to use for the request.                                                                           |
| `taskID`                                                                                                      | *string*                                                                                                      | :heavy_check_mark:                                                                                            | N/A                                                                                                           |
| `taskActionsServiceSkipStepRequest`                                                                           | [*components.TaskActionsServiceSkipStepRequest](../../models/components/taskactionsserviceskipsteprequest.md) | :heavy_minus_sign:                                                                                            | N/A                                                                                                           |
| `opts`                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                      | :heavy_minus_sign:                                                                                            | The options for this request.                                                                                 |

### Response

**[*operations.C1APITaskV1TaskActionsServiceSkipStepResponse](../../models/operations/c1apitaskv1taskactionsserviceskipstepresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateGrantDuration

Invokes the c1.api.task.v1.TaskActionsService.UpdateGrantDuration method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.UpdateGrantDuration" method="post" path="/api/v1/tasks/{task_id}/action/update-grant-duration" -->
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

    res, err := s.TaskActions.UpdateGrantDuration(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |
| `taskID`                                                                                                                            | *string*                                                                                                                            | :heavy_check_mark:                                                                                                                  | N/A                                                                                                                                 |
| `taskActionsServiceUpdateGrantDurationRequest`                                                                                      | [*components.TaskActionsServiceUpdateGrantDurationRequest](../../models/components/taskactionsserviceupdategrantdurationrequest.md) | :heavy_minus_sign:                                                                                                                  | N/A                                                                                                                                 |
| `opts`                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                            | :heavy_minus_sign:                                                                                                                  | The options for this request.                                                                                                       |

### Response

**[*operations.C1APITaskV1TaskActionsServiceUpdateGrantDurationResponse](../../models/operations/c1apitaskv1taskactionsserviceupdategrantdurationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateRequestData

Invokes the c1.api.task.v1.TaskActionsService.UpdateRequestData method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.task.v1.TaskActionsService.UpdateRequestData" method="post" path="/api/v1/tasks/{task_id}/action/update-request-data" -->
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

    res, err := s.TaskActions.UpdateRequestData(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskServiceActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |
| `taskID`                                                                                                                        | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | N/A                                                                                                                             |
| `taskActionsServiceUpdateRequestDataRequest`                                                                                    | [*components.TaskActionsServiceUpdateRequestDataRequest](../../models/components/taskactionsserviceupdaterequestdatarequest.md) | :heavy_minus_sign:                                                                                                              | N/A                                                                                                                             |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |

### Response

**[*operations.C1APITaskV1TaskActionsServiceUpdateRequestDataResponse](../../models/operations/c1apitaskv1taskactionsserviceupdaterequestdataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |