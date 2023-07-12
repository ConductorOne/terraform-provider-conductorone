# UserSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.user.v1.UserSearch.Search method.

## Search

Invokes the c1.api.user.v1.UserSearch.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"conductorone/internal/sdk/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.UserSearch.Search(ctx, shared.SearchUsersRequest{
        ExcludeIds: []string{
            "consectetur",
            "esse",
            "blanditiis",
            "provident",
        },
        ExpandMask: &shared.UserExpandMask{
            Paths: []string{
                "nulla",
                "quas",
                "esse",
                "quasi",
            },
        },
        Ids: []string{
            "error",
            "sint",
            "pariatur",
            "possimus",
        },
        PageSize: conductoroneapi.Float64(1576.32),
        PageToken: conductoroneapi.String("eveniet"),
        Query: conductoroneapi.String("asperiores"),
        Refs: []shared.UserRef{
            shared.UserRef{
                ID: conductoroneapi.String("121aa6f1-e674-4bdb-84f1-5756082d68ea"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("19f1d170-5133-49d0-8086-a1840394c260"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("71f93f5f-0642-4dac-baf5-15cc413aa63a"),
            },
            shared.UserRef{
                ID: conductoroneapi.String("ae8d6786-4dbb-4675-bd5e-60b375ed4f6f"),
            },
        },
        RoleIds: []string{
            "saepe",
            "necessitatibus",
            "dolore",
        },
        UserStatuses: []shared.SearchUsersRequestUserStatuses{
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

