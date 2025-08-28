# AutomationSearch
(*AutomationSearch*)

## Overview

### Available Operations

* [SearchAutomationTemplateVersions](#searchautomationtemplateversions) - Search Automation Template Versions
* [SearchAutomations](#searchautomations) - Search Automations

## SearchAutomationTemplateVersions

Invokes the c1.api.automations.v1.AutomationSearchService.SearchAutomationTemplateVersions method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationSearchService.SearchAutomationTemplateVersions" method="post" path="/api/v1/automation_versions/search" -->
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

    res, err := s.AutomationSearch.SearchAutomationTemplateVersions(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAutomationTemplateVersionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [components.SearchAutomationTemplateVersionsRequest](../../models/components/searchautomationtemplateversionsrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.C1APIAutomationsV1AutomationSearchServiceSearchAutomationTemplateVersionsResponse](../../models/operations/c1apiautomationsv1automationsearchservicesearchautomationtemplateversionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SearchAutomations

Invokes the c1.api.automations.v1.AutomationSearchService.SearchAutomations method.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.automations.v1.AutomationSearchService.SearchAutomations" method="post" path="/api/v1/automations/search" -->
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

    res, err := s.AutomationSearch.SearchAutomations(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAutomationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.SearchAutomationsRequest](../../models/components/searchautomationsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.C1APIAutomationsV1AutomationSearchServiceSearchAutomationsResponse](../../models/operations/c1apiautomationsv1automationsearchservicesearchautomationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |