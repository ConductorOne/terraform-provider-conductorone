# PolicySearch

### Available Operations

* [Search](#search) - Search

## Search

Search policies based on filters specified in the request body.

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
    res, err := s.PolicySearch.Search(ctx, shared.SearchPoliciesRequest{
        DisplayName: conductoronesdkterraform.String("possimus"),
        IncludeDeleted: conductoronesdkterraform.Bool(false),
        PageSize: conductoronesdkterraform.Int(220824),
        PageToken: conductoronesdkterraform.String("rerum"),
        PolicyTypes: []shared.SearchPoliciesRequestPolicyTypes{
            shared.SearchPoliciesRequestPolicyTypesPolicyTypeProvision,
        },
        Query: conductoronesdkterraform.String("optio"),
        Refs: []shared.PolicyRef{
            shared.PolicyRef{
                ID: conductoronesdkterraform.String("cc8f8950-10f5-4dd3-96fa-1804e54c82f1"),
            },
            shared.PolicyRef{
                ID: conductoronesdkterraform.String("68a363c8-873e-4484-b80b-1f6b8ca275a6"),
            },
            shared.PolicyRef{
                ID: conductoronesdkterraform.String("0a04c495-cc69-4917-9b51-c1bdb1cf4b88"),
            },
            shared.PolicyRef{
                ID: conductoronesdkterraform.String("8ebdfc4c-cca9-49bc-bfc0-b2dce10873e4"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.SearchPoliciesRequest](../../models/shared/searchpoliciesrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.C1APIPolicyV1PolicySearchSearchResponse](../../models/operations/c1apipolicyv1policysearchsearchresponse.md), error**

