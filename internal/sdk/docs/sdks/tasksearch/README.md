# TaskSearch

### Available Operations

* [Search](#search) - Search

## Search

Search tasks based on filters specified in the request body.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/types"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.TaskSearch.Search(ctx, shared.TaskSearchRequestInput{
        TaskExpandMask: &shared.TaskExpandMask{
            Paths: []string{
                "quidem",
                "illum",
                "reiciendis",
            },
        },
        AccessReviewIds: []string{
            "dolores",
            "consequatur",
            "nesciunt",
            "quia",
        },
        AccountOwnerIds: []string{
            "voluptas",
            "quo",
            "laudantium",
        },
        AccountTypes: []shared.TaskSearchRequestAccountTypes{
            shared.TaskSearchRequestAccountTypesAppUserTypeServiceAccount,
            shared.TaskSearchRequestAccountTypesAppUserTypeServiceAccount,
        },
        ActorID: conductoronesdkterraform.String("fugit"),
        AppEntitlementIds: []string{
            "quidem",
        },
        AppResourceIds: []string{
            "debitis",
            "vitae",
        },
        AppResourceTypeIds: []string{
            "ad",
        },
        AppUserSubjectIds: []string{
            "ut",
            "asperiores",
            "reprehenderit",
        },
        ApplicationIds: []string{
            "itaque",
            "et",
            "eos",
        },
        AssigneesInIds: []string{
            "ex",
            "praesentium",
            "natus",
            "vitae",
        },
        CreatedAfter: types.MustTimeFromString("2021-06-15T21:46:49.215Z"),
        CreatedBefore: types.MustTimeFromString("2022-03-31T08:06:02.567Z"),
        CurrentStep: shared.TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        EmergencyStatus: shared.TaskSearchRequestEmergencyStatusUnspecified.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "ullam",
        },
        ExcludeIds: []string{
            "inventore",
            "voluptate",
        },
        IncludeDeleted: conductoronesdkterraform.Bool(false),
        MyWorkUserIds: []string{
            "dolorem",
        },
        OlderThanDuration: conductoronesdkterraform.String("eaque"),
        OpenerIds: []string{
            "amet",
            "voluptate",
        },
        PageSize: conductoronesdkterraform.Int(454232),
        PageToken: conductoronesdkterraform.String("pariatur"),
        PreviouslyActedOnIds: []string{
            "a",
            "fuga",
            "totam",
            "cupiditate",
        },
        Query: conductoronesdkterraform.String("at"),
        Refs: []shared.TaskRef{
            shared.TaskRef{
                ID: conductoronesdkterraform.String("975e3566-8609-42e9-83dd-c5f111dea102"),
            },
            shared.TaskRef{
                ID: conductoronesdkterraform.String("6d541a4d-190f-4eb2-9780-bccc0dbbddb4"),
            },
            shared.TaskRef{
                ID: conductoronesdkterraform.String("84708fb4-e391-4e6b-8158-c4c4e54599ea"),
            },
            shared.TaskRef{
                ID: conductoronesdkterraform.String("342260e9-b200-4ce7-8a1b-d8fb7a0a116c"),
            },
        },
        SortBy: shared.TaskSearchRequestSortByTaskSearchSortByReverseTicketID.ToPointer(),
        SubjectIds: []string{
            "fugit",
            "ratione",
        },
        TaskStates: []shared.TaskSearchRequestTaskStates{
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
            shared.TaskSearchRequestTaskStatesTaskStateOpen,
        },
        TaskTypes: []shared.TaskTypeInput{
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoronesdkterraform.String("deserunt"),
                        IntegrationID: conductoronesdkterraform.String("ratione"),
                        RequestID: conductoronesdkterraform.String("ipsa"),
                    },
                },
                TaskTypeOffboarding: &shared.TaskTypeOffboarding{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2021-02-26T20:02:05.121Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-02-10T14:02:23.370Z"),
                            LastLogin: types.MustTimeFromString("2022-11-03T12:07:22.406Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoronesdkterraform.String("ipsam"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoronesdkterraform.String("libero"),
                            CertTicketID: conductoronesdkterraform.String("quia"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoronesdkterraform.String("omnis"),
                        IntegrationID: conductoronesdkterraform.String("dicta"),
                        RequestID: conductoronesdkterraform.String("qui"),
                    },
                },
                TaskTypeOffboarding: &shared.TaskTypeOffboarding{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-12-29T17:39:27.998Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-12-28T14:40:04.720Z"),
                            LastLogin: types.MustTimeFromString("2021-05-28T08:36:40.882Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoronesdkterraform.String("velit"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoronesdkterraform.String("hic"),
                            CertTicketID: conductoronesdkterraform.String("ullam"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoronesdkterraform.String("deserunt"),
                        IntegrationID: conductoronesdkterraform.String("itaque"),
                        RequestID: conductoronesdkterraform.String("distinctio"),
                    },
                },
                TaskTypeOffboarding: &shared.TaskTypeOffboarding{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-07-04T17:44:17.469Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-11-02T05:18:20.126Z"),
                            LastLogin: types.MustTimeFromString("2022-08-15T15:03:40.347Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoronesdkterraform.String("odit"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoronesdkterraform.String("vero"),
                            CertTicketID: conductoronesdkterraform.String("deleniti"),
                        },
                    },
                },
            },
            shared.TaskTypeInput{
                TaskTypeCertify: &shared.TaskTypeCertify{},
                TaskTypeGrant: &shared.TaskTypeGrantInput{
                    TaskGrantSource: &shared.TaskGrantSource{
                        ExternalURL: conductoronesdkterraform.String("optio"),
                        IntegrationID: conductoronesdkterraform.String("quasi"),
                        RequestID: conductoronesdkterraform.String("repellat"),
                    },
                },
                TaskTypeOffboarding: &shared.TaskTypeOffboarding{},
                TaskTypeRevoke: &shared.TaskTypeRevokeInput{
                    TaskRevokeSource: &shared.TaskRevokeSource{
                        TaskRevokeSourceExpired: &shared.TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-05-28T15:53:37.829Z"),
                        },
                        TaskRevokeSourceNonUsage: &shared.TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-07-19T06:38:54.434Z"),
                            LastLogin: types.MustTimeFromString("2022-10-01T15:10:03.132Z"),
                        },
                        TaskRevokeSourceRequest: &shared.TaskRevokeSourceRequest{
                            RequestUserID: conductoronesdkterraform.String("nemo"),
                        },
                        TaskRevokeSourceReview: &shared.TaskRevokeSourceReview{
                            AccessReviewID: conductoronesdkterraform.String("delectus"),
                            CertTicketID: conductoronesdkterraform.String("illum"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.TaskSearchRequestInput](../../models/shared/tasksearchrequestinput.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.C1APITaskV1TaskSearchServiceSearchResponse](../../models/operations/c1apitaskv1tasksearchservicesearchresponse.md), error**

