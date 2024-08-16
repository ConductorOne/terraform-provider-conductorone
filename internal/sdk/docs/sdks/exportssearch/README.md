# ExportsSearch

### Available Operations

* [Search](#search) - Search

## Search

Invokes the c1.api.systemlog.v1.ExportsSearchService.Search method.

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
    res, err := s.ExportsSearch.Search(ctx, shared.ExportsSearchServiceSearchRequest{
        DisplayName: conductoronesdkterraform.String("necessitatibus"),
        PageSize: conductoronesdkterraform.Int(215529),
        PageToken: conductoronesdkterraform.String("ea"),
        Query: conductoronesdkterraform.String("occaecati"),
        Refs: []shared.ExporterRef{
            shared.ExporterRef{
                ExportID: conductoronesdkterraform.String("voluptatibus"),
            },
            shared.ExporterRef{
                ExportID: conductoronesdkterraform.String("tempora"),
            },
            shared.ExporterRef{
                ExportID: conductoronesdkterraform.String("tempora"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ExportsSearchServiceSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [shared.ExportsSearchServiceSearchRequest](../../models/shared/exportssearchservicesearchrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.C1APISystemlogV1ExportsSearchServiceSearchResponse](../../models/operations/c1apisystemlogv1exportssearchservicesearchresponse.md), error**

