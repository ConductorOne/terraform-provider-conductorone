# UserSearch

### Available Operations

* [Search](#search) - Search

## Search

Search users based on filters specified in the request body.

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
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        UserExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "eos",
                "dolorem",
                "hic",
                "repellendus",
            },
        },
        Email: conductoronesdkterraform.String("Annie_Gulgowski41@hotmail.com"),
        ExcludeIds: []string{
            "recusandae",
            "nostrum",
            "officia",
        },
        Ids: []string{
            "laudantium",
            "corporis",
        },
        PageSize: conductoronesdkterraform.Int(567241),
        PageToken: conductoronesdkterraform.String("natus"),
        Query: conductoronesdkterraform.String("deleniti"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoronesdkterraform.String("22ae20da-16fc-42b2-b1a2-89c57e854e90"),
            },
            shared.UserRef{
                ID: conductoronesdkterraform.String("439d2224-6569-4462-8070-84f7ab37cef0"),
            },
            shared.UserRef{
                ID: conductoronesdkterraform.String("2225194d-b554-410a-9c66-9af90a26c7cd"),
            },
            shared.UserRef{
                ID: conductoronesdkterraform.String("c981f068-981d-46bb-b3cf-aa348c31bf40"),
            },
        },
        RoleIds: []string{
            "itaque",
            "necessitatibus",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
            shared.SearchUsersRequestUserStatusesDeleted,
            shared.SearchUsersRequestUserStatusesDeleted,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchUsersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.SearchUsersRequest](../../models/shared/searchusersrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.C1APIUserV1UserSearchSearchResponse](../../models/operations/c1apiuserv1usersearchsearchresponse.md), error**

