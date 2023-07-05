# Policies

### Available Operations

* [Create](#create) - Invokes the c1.api.policy.v1.Policies.Create method.
* [Delete](#delete) - Invokes the c1.api.policy.v1.Policies.Delete method.
* [Get](#get) - Invokes the c1.api.policy.v1.Policies.Get method.
* [List](#list) - Invokes the c1.api.policy.v1.Policies.List method.
* [Update](#update) - Invokes the c1.api.policy.v1.Policies.Update method.

## Create

Invokes the c1.api.policy.v1.Policies.Create method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Create(ctx, shared.CreatePolicyRequest{
        Description: conductoroneapi.String("ipsam"),
        DisplayName: conductoroneapi.String("alias"),
        PolicySteps: map[string]shared.PolicySteps{
            "dolorum": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AllowReassignment: conductoroneapi.Bool(false),
                            AppOwners: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            Assigned: conductoroneapi.Bool(false),
                            EntitlementOwners: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "facilis",
                                    "tempore",
                                },
                            },
                            Group: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("labore"),
                                AppID: conductoroneapi.String("delectus"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "non",
                                    "eligendi",
                                },
                            },
                            Manager: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "aliquid",
                                    "provident",
                                    "necessitatibus",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "officia",
                                    "dolor",
                                    "debitis",
                                },
                            },
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                            Self: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "dolorum",
                                    "in",
                                    "in",
                                    "illum",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "rerum",
                                    "dicta",
                                    "magnam",
                                    "cumque",
                                },
                            },
                            Users: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "ea",
                                    "aliquid",
                                    "laborum",
                                    "accusamus",
                                },
                            },
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("non"),
                                    EntitlementID: conductoroneapi.String("occaecati"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("enim"),
                                    UserIds: []string{
                                        "delectus",
                                        "quidem",
                                        "provident",
                                        "nam",
                                    },
                                },
                            },
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AllowReassignment: conductoroneapi.Bool(false),
                            AppOwners: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            Assigned: conductoroneapi.Bool(false),
                            EntitlementOwners: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "blanditiis",
                                    "deleniti",
                                    "sapiente",
                                },
                            },
                            Group: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("amet"),
                                AppID: conductoroneapi.String("deserunt"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "vel",
                                    "natus",
                                },
                            },
                            Manager: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "molestiae",
                                    "perferendis",
                                    "nihil",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "distinctio",
                                    "id",
                                },
                            },
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                            Self: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "labore",
                                    "suscipit",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "nobis",
                                    "eum",
                                    "vero",
                                },
                            },
                            Users: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "architecto",
                                },
                            },
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("magnam"),
                                    EntitlementID: conductoroneapi.String("et"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("excepturi"),
                                    UserIds: []string{
                                        "provident",
                                        "quos",
                                    },
                                },
                            },
                        },
                    },
                    shared.PolicyStep{
                        Approval: &shared.Approval{
                            AllowReassignment: conductoroneapi.Bool(false),
                            AppOwners: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                            },
                            Assigned: conductoroneapi.Bool(false),
                            EntitlementOwners: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "accusantium",
                                    "mollitia",
                                    "reiciendis",
                                },
                            },
                            Group: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AppGroupID: conductoroneapi.String("mollitia"),
                                AppID: conductoroneapi.String("ad"),
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "dolor",
                                    "necessitatibus",
                                },
                            },
                            Manager: &shared.ManagerApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                AssignedUserIds: []string{
                                    "nemo",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "iure",
                                },
                            },
                            RequireApprovalReason: conductoroneapi.Bool(false),
                            RequireReassignmentReason: conductoroneapi.Bool(false),
                            Self: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "debitis",
                                    "eius",
                                    "maxime",
                                    "deleniti",
                                },
                                Fallback: conductoroneapi.Bool(false),
                                FallbackUserIds: []string{
                                    "in",
                                    "architecto",
                                    "architecto",
                                },
                            },
                            Users: &shared.UserApproval{
                                AllowSelfApproval: conductoroneapi.Bool(false),
                                UserIds: []string{
                                    "ullam",
                                    "expedita",
                                    "nihil",
                                    "repellat",
                                },
                            },
                        },
                        Provision: &shared.Provision{
                            Assigned: conductoroneapi.Bool(false),
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                Connector: &shared.ConnectorProvision{},
                                Delegated: &shared.DelegatedProvision{
                                    AppID: conductoroneapi.String("quibusdam"),
                                    EntitlementID: conductoroneapi.String("sed"),
                                },
                                Manual: &shared.ManualProvision{
                                    Instructions: conductoroneapi.String("saepe"),
                                    UserIds: []string{
                                        "accusantium",
                                        "consequuntur",
                                        "praesentium",
                                        "natus",
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeGrant.ToPointer(),
        PostActions: []shared.PolicyPostActions{
            shared.PolicyPostActions{
                CertifyRemediateImmediately: conductoroneapi.Bool(false),
            },
        },
        ReassignTasksToDelegates: conductoroneapi.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.CreatePolicyRequest](../../models/shared/createpolicyrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.C1APIPolicyV1PoliciesCreateResponse](../../models/operations/c1apipolicyv1policiescreateresponse.md), error**


## Delete

Invokes the c1.api.policy.v1.Policies.Delete method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/operations"
	"terraform/internal/sdk/pkg/models/shared"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Delete(ctx, operations.C1APIPolicyV1PoliciesDeleteRequest{
        DeletePolicyRequest: &shared.DeletePolicyRequest{},
        ID: "cddc6926-01fb-4576-b0d5-f0d30c5fbb25",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIPolicyV1PoliciesDeleteRequest](../../models/operations/c1apipolicyv1policiesdeleterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.C1APIPolicyV1PoliciesDeleteResponse](../../models/operations/c1apipolicyv1policiesdeleteresponse.md), error**


## Get

Invokes the c1.api.policy.v1.Policies.Get method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/operations"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Get(ctx, operations.C1APIPolicyV1PoliciesGetRequest{
        ID: "87053202-c73d-45fe-9b90-c28909b3fe49",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.C1APIPolicyV1PoliciesGetRequest](../../models/operations/c1apipolicyv1policiesgetrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.C1APIPolicyV1PoliciesGetResponse](../../models/operations/c1apipolicyv1policiesgetresponse.md), error**


## List

Invokes the c1.api.policy.v1.Policies.List method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListPolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.C1APIPolicyV1PoliciesListResponse](../../models/operations/c1apipolicyv1policieslistresponse.md), error**


## Update

Invokes the c1.api.policy.v1.Policies.Update method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"terraform/internal/sdk"
	"terraform/internal/sdk/pkg/models/operations"
	"terraform/internal/sdk/pkg/models/shared"
	"terraform/internal/sdk/pkg/types"
)

func main() {
    s := conductoroneapi.New()

    ctx := context.Background()
    res, err := s.Policies.Update(ctx, operations.C1APIPolicyV1PoliciesUpdateRequest{
        UpdatePolicyRequest: &shared.UpdatePolicyRequest{
            Policy: &shared.Policy{
                CreatedAt: types.MustTimeFromString("2021-12-05T23:42:47.745Z"),
                DeletedAt: types.MustTimeFromString("2021-03-26T01:09:56.415Z"),
                Description: conductoroneapi.String("nobis"),
                DisplayName: conductoroneapi.String("libero"),
                ID: conductoroneapi.String("f4863332-3f9b-477f-ba41-00674ebf6928"),
                PolicySteps: map[string]shared.PolicySteps{
                    "fugiat": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Approval: &shared.Approval{
                                    AllowReassignment: conductoroneapi.Bool(false),
                                    AppOwners: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                    },
                                    Assigned: conductoroneapi.Bool(false),
                                    EntitlementOwners: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "dolorum",
                                            "iusto",
                                            "voluptate",
                                        },
                                    },
                                    Group: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AppGroupID: conductoroneapi.String("dolorum"),
                                        AppID: conductoroneapi.String("deleniti"),
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "necessitatibus",
                                            "distinctio",
                                            "asperiores",
                                        },
                                    },
                                    Manager: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        AssignedUserIds: []string{
                                            "ipsum",
                                            "voluptate",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "saepe",
                                            "eius",
                                            "aspernatur",
                                        },
                                    },
                                    RequireApprovalReason: conductoroneapi.Bool(false),
                                    RequireReassignmentReason: conductoroneapi.Bool(false),
                                    Self: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "amet",
                                        },
                                        Fallback: conductoroneapi.Bool(false),
                                        FallbackUserIds: []string{
                                            "accusamus",
                                            "ad",
                                            "saepe",
                                            "suscipit",
                                        },
                                    },
                                    Users: &shared.UserApproval{
                                        AllowSelfApproval: conductoroneapi.Bool(false),
                                        UserIds: []string{
                                            "provident",
                                            "minima",
                                            "repellendus",
                                        },
                                    },
                                },
                                Provision: &shared.Provision{
                                    Assigned: conductoroneapi.Bool(false),
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        Connector: &shared.ConnectorProvision{},
                                        Delegated: &shared.DelegatedProvision{
                                            AppID: conductoroneapi.String("totam"),
                                            EntitlementID: conductoroneapi.String("similique"),
                                        },
                                        Manual: &shared.ManualProvision{
                                            Instructions: conductoroneapi.String("alias"),
                                            UserIds: []string{
                                                "quaerat",
                                                "tempora",
                                                "vel",
                                                "quod",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeProvision.ToPointer(),
                PostActions: []shared.PolicyPostActions{
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoroneapi.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoroneapi.Bool(false),
                SystemBuiltin: conductoroneapi.Bool(false),
                UpdatedAt: types.MustTimeFromString("2021-02-04T11:05:24.484Z"),
            },
            UpdateMask: conductoroneapi.String("esse"),
        },
        ID: "a73cf3be-453f-4870-b326-b5a73429cdb1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePolicyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.C1APIPolicyV1PoliciesUpdateRequest](../../models/operations/c1apipolicyv1policiesupdaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.C1APIPolicyV1PoliciesUpdateResponse](../../models/operations/c1apipolicyv1policiesupdateresponse.md), error**

