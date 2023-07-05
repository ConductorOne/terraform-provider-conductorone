# TaskSearch

### Available Operations

* [Search](#search) - Invokes the c1.api.task.v1.TaskSearchService.Search method.

## Search

Invokes the c1.api.task.v1.TaskSearchService.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/shared"
	"terraform/internal/sdk/pkg/types"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequest{
        AccessReviewIds: []string{
            "natus",
        },
        AccountOwnerIds: []string{
            "aut",
            "voluptatibus",
            "exercitationem",
            "nulla",
        },
        ActorID: conductoroneapi.String("fugit"),
        AppEntitlementIds: []string{
            "maiores",
            "doloribus",
            "iusto",
            "eligendi",
        },
        AppResourceIds: []string{
            "alias",
            "officia",
        },
        AppResourceTypeIds: []string{
            "ipsam",
            "ea",
        },
        AppUserSubjectIds: []string{
            "vel",
        },
        ApplicationIds: []string{
            "magnam",
            "ratione",
            "ex",
            "laudantium",
        },
        AssigneesInIds: []string{
            "dolor",
        },
        CreatedAfter: types.MustTimeFromString("2022-09-15T20:38:47.955Z"),
        CreatedBefore: types.MustTimeFromString("2022-02-20T07:12:08.273Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepApproval.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "nostrum",
            "sapiente",
            "quisquam",
            "saepe",
        },
        ExcludeIds: []string{
            "impedit",
            "corporis",
        },
        ExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "aliquid",
                "inventore",
            },
        },
        IncludeDeleted: conductoroneapi.Bool(false),
        MyWorkUserIds: []string{
            "ea",
            "quo",
        },
        OpenerIds: []string{
            "recusandae",
        },
        PageSize: conductoroneapi.Float64(1324.87),
        PageToken: conductoroneapi.String("minima"),
        PreviouslyActedOnIds: []string{
            "a",
        },
        Query: conductoroneapi.String("libero"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoroneapi.String("08c42e14-1aac-4366-88dd-6b1442907474"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByAccount.ToPointer(),
        SubjectIds: []string{
            "rem",
            "fuga",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
            shared.TaskSearchRequestTaskStatesTaskStateClosed,
        },
        TaskTypes: []shared.TaskType{
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("eum"),
                    AccessReviewSelection: conductoroneapi.String("suscipit"),
                    AppEntitlementID: conductoroneapi.String("assumenda"),
                    AppID: conductoroneapi.String("eos"),
                    AppUserID: conductoroneapi.String("praesentium"),
                    IdentityUserID: conductoroneapi.String("quisquam"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-05-05T02:03:02.582Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("quidem"),
                    AppID: conductoroneapi.String("neque"),
                    AppUserID: conductoroneapi.String("quo"),
                    GrantDuration: conductoroneapi.String("illum"),
                    IdentityUserID: conductoroneapi.String("quo"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-10-27T21:30:06.318Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("voluptas"),
                    AppID: conductoroneapi.String("ab"),
                    AppUserID: conductoroneapi.String("cupiditate"),
                    IdentityUserID: conductoroneapi.String("consequatur"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeRevoked.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-11-20T21:59:17.660Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-10-21T01:48:15.498Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-08-15T10:59:14.485Z"),
                            LastLogin: types.MustTimeFromString("2022-11-13T03:35:18.820Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("distinctio"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("quod"),
                            CertTicketID: conductoroneapi.String("dignissimos"),
                        },
                    },
                },
            },
            shared.TaskType{
                Certify: &shared.TaskTypeCertify{
                    AccessReviewID: conductoroneapi.String("inventore"),
                    AccessReviewSelection: conductoroneapi.String("nihil"),
                    AppEntitlementID: conductoroneapi.String("totam"),
                    AppID: conductoroneapi.String("accusamus"),
                    AppUserID: conductoroneapi.String("aliquam"),
                    IdentityUserID: conductoroneapi.String("odio"),
                    Outcome: shared.TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-01-15T19:21:50.024Z"),
                },
                Grant: &shared.TaskTypeGrant{
                    AppEntitlementID: conductoroneapi.String("dolores"),
                    AppID: conductoroneapi.String("deserunt"),
                    AppUserID: conductoroneapi.String("molestiae"),
                    GrantDuration: conductoroneapi.String("accusantium"),
                    IdentityUserID: conductoroneapi.String("porro"),
                    Outcome: shared.TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-12-24T16:30:16.544Z"),
                },
                Revoke: &shared.TaskTypeRevoke{
                    AppEntitlementID: conductoroneapi.String("consequuntur"),
                    AppID: conductoroneapi.String("deleniti"),
                    AppUserID: conductoroneapi.String("fugit"),
                    IdentityUserID: conductoroneapi.String("fuga"),
                    Outcome: shared.TaskTypeRevokeOutcomeRevokeOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-06-18T04:24:04.187Z"),
                    Source: &shared.TaskRevokeSource{
                        Expired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-09-04T03:00:04.529Z"),
                        },
                        NonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-11-08T08:09:29.073Z"),
                            LastLogin: types.MustTimeFromString("2022-07-09T22:36:51.892Z"),
                        },
                        Request: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoroneapi.String("ratione"),
                        },
                        Review: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoroneapi.String("explicabo"),
                            CertTicketID: conductoroneapi.String("saepe"),
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.TaskSearchRequest](../../models/shared/tasksearchrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.C1APITaskV1TaskSearchServiceSearchResponse](../../models/operations/c1apitaskv1tasksearchservicesearchresponse.md), error**

