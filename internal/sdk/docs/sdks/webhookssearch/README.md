# WebhooksSearch

### Available Operations

* [Search](#search) - Search

## Search

Invokes the c1.api.webhooks.v1.WebhooksSearch.Search method.

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
    res, err := s.WebhooksSearch.Search(ctx, shared.WebhooksSearchRequest{
        PageSize: conductoronesdkterraform.Int(401844),
        PageToken: conductoronesdkterraform.String("neque"),
        Query: conductoronesdkterraform.String("quod"),
        Refs: []shared.WebhookRef{
            shared.WebhookRef{
                ID: conductoronesdkterraform.String("05fda840-774a-468a-9a35-d086b6f66fef"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WebhooksSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.WebhooksSearchRequest](../../models/shared/webhookssearchrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.C1APIWebhooksV1WebhooksSearchSearchResponse](../../models/operations/c1apiwebhooksv1webhookssearchsearchresponse.md), error**

