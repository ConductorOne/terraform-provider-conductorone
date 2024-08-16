# Policies

### Available Operations

* [Create](#create) - Create
* [Delete](#delete) - Delete
* [Get](#get) - Get
* [List](#list) - List
* [Update](#update) - Update

## Create

Create a policy.

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
    res, err := s.Policies.Create(ctx, shared.CreatePolicyRequest{
        Description: conductoronesdkterraform.String("ipsam"),
        DisplayName: conductoronesdkterraform.String("debitis"),
        PolicySteps: map[string]shared.PolicySteps{
            "sit": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("error"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("veniam"),
                                AppID: conductoronesdkterraform.String("minima"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "reiciendis",
                                    "nulla",
                                    "magni",
                                    "aperiam",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "numquam",
                                    "veniam",
                                    "in",
                                    "officiis",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "laudantium",
                                },
                                Expressions: []string{
                                    "praesentium",
                                    "cum",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "dolorum",
                                    "voluptatum",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "hic",
                                    "expedita",
                                    "debitis",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "dolorum",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "officia",
                                    "dolorum",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "accusamus",
                                    "tempora",
                                    "atque",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "ut",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("fugiat"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("voluptatem"),
                                    EntitlementID: conductoronesdkterraform.String("culpa"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("expedita"),
                                    ConnectorID: conductoronesdkterraform.String("magnam"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("consequatur"),
                                    Instructions: conductoronesdkterraform.String("esse"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("ipsam"),
                                    UserIds: []string{
                                        "voluptatum",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("repudiandae"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("corporis"),
                                AppID: conductoronesdkterraform.String("et"),
                                AppUserID: conductoronesdkterraform.String("blanditiis"),
                                GrantDuration: conductoronesdkterraform.String("ex"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("sed"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("sit"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("vel"),
                            CommentOnTimeout: conductoronesdkterraform.String("nostrum"),
                            Name: conductoronesdkterraform.String("Evan Altenwerth"),
                            TimeoutDuration: conductoronesdkterraform.String("dolorem"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("harum"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("dicta"),
                                AppID: conductoronesdkterraform.String("architecto"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "labore",
                                    "quidem",
                                    "atque",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "nam",
                                    "tenetur",
                                    "laboriosam",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "amet",
                                },
                                Expressions: []string{
                                    "voluptate",
                                    "unde",
                                    "reiciendis",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "repellendus",
                                    "delectus",
                                    "voluptates",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "est",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "reprehenderit",
                                    "facere",
                                    "fuga",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "mollitia",
                                    "veniam",
                                    "voluptatem",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "repudiandae",
                                    "quasi",
                                    "atque",
                                    "reprehenderit",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "totam",
                                    "suscipit",
                                    "quidem",
                                    "maxime",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("et"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("esse"),
                                    EntitlementID: conductoronesdkterraform.String("amet"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("assumenda"),
                                    ConnectorID: conductoronesdkterraform.String("ea"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("atque"),
                                    Instructions: conductoronesdkterraform.String("error"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("officiis"),
                                    UserIds: []string{
                                        "accusamus",
                                        "natus",
                                        "minima",
                                        "aspernatur",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("maiores"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("corrupti"),
                                AppID: conductoronesdkterraform.String("at"),
                                AppUserID: conductoronesdkterraform.String("error"),
                                GrantDuration: conductoronesdkterraform.String("blanditiis"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("suscipit"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("repudiandae"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("atque"),
                            CommentOnTimeout: conductoronesdkterraform.String("atque"),
                            Name: conductoronesdkterraform.String("Delia Parisian"),
                            TimeoutDuration: conductoronesdkterraform.String("reiciendis"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("doloremque"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("repudiandae"),
                                AppID: conductoronesdkterraform.String("dicta"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "beatae",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "enim",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "velit",
                                    "a",
                                },
                                Expressions: []string{
                                    "magnam",
                                    "saepe",
                                    "consequuntur",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "officiis",
                                    "perspiciatis",
                                    "in",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "eveniet",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "consequuntur",
                                    "fugit",
                                    "id",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "reprehenderit",
                                    "error",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "corporis",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "eveniet",
                                    "non",
                                    "vero",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("doloremque"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("iure"),
                                    EntitlementID: conductoronesdkterraform.String("ipsa"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("totam"),
                                    ConnectorID: conductoronesdkterraform.String("quae"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("molestiae"),
                                    Instructions: conductoronesdkterraform.String("eveniet"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("qui"),
                                    UserIds: []string{
                                        "iure",
                                        "necessitatibus",
                                        "ratione",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("distinctio"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("voluptatum"),
                                AppID: conductoronesdkterraform.String("rem"),
                                AppUserID: conductoronesdkterraform.String("aliquam"),
                                GrantDuration: conductoronesdkterraform.String("ad"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("repellat"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("alias"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("corporis"),
                            CommentOnTimeout: conductoronesdkterraform.String("perspiciatis"),
                            Name: conductoronesdkterraform.String("Dr. Iris Hodkiewicz"),
                            TimeoutDuration: conductoronesdkterraform.String("dolores"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("id"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("minima"),
                                AppID: conductoronesdkterraform.String("dolore"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "nesciunt",
                                    "quae",
                                    "recusandae",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "quaerat",
                                    "molestiae",
                                    "ex",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "culpa",
                                    "adipisci",
                                },
                                Expressions: []string{
                                    "laudantium",
                                    "eum",
                                    "nemo",
                                    "recusandae",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "provident",
                                    "quis",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "reiciendis",
                                    "provident",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "ullam",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "animi",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "mollitia",
                                    "provident",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "animi",
                                    "ex",
                                    "aliquid",
                                    "accusantium",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("repellat"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("doloribus"),
                                    EntitlementID: conductoronesdkterraform.String("ullam"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("in"),
                                    ConnectorID: conductoronesdkterraform.String("nam"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("earum"),
                                    Instructions: conductoronesdkterraform.String("officia"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("laborum"),
                                    UserIds: []string{
                                        "modi",
                                        "voluptatibus",
                                        "molestias",
                                        "officiis",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("cumque"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("vitae"),
                                AppID: conductoronesdkterraform.String("rerum"),
                                AppUserID: conductoronesdkterraform.String("tempora"),
                                GrantDuration: conductoronesdkterraform.String("quis"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("inventore"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("fugit"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("cumque"),
                            CommentOnTimeout: conductoronesdkterraform.String("quae"),
                            Name: conductoronesdkterraform.String("Josephine Collins"),
                            TimeoutDuration: conductoronesdkterraform.String("rem"),
                        },
                    },
                },
            },
            "at": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("eos"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("sapiente"),
                                AppID: conductoronesdkterraform.String("eum"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "minima",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "cupiditate",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "earum",
                                    "soluta",
                                    "hic",
                                },
                                Expressions: []string{
                                    "eaque",
                                    "earum",
                                    "perspiciatis",
                                    "maiores",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "aliquid",
                                    "porro",
                                    "suscipit",
                                    "dolorem",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "cumque",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "ratione",
                                    "animi",
                                    "necessitatibus",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "consequatur",
                                    "quasi",
                                    "et",
                                    "ducimus",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "occaecati",
                                    "suscipit",
                                    "adipisci",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "magni",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("doloribus"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("nulla"),
                                    EntitlementID: conductoronesdkterraform.String("necessitatibus"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("ipsa"),
                                    ConnectorID: conductoronesdkterraform.String("tempora"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("nihil"),
                                    Instructions: conductoronesdkterraform.String("molestiae"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("dicta"),
                                    UserIds: []string{
                                        "esse",
                                        "praesentium",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("reiciendis"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("vel"),
                                AppID: conductoronesdkterraform.String("architecto"),
                                AppUserID: conductoronesdkterraform.String("fugiat"),
                                GrantDuration: conductoronesdkterraform.String("doloremque"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("dicta"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("odio"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("tempora"),
                            CommentOnTimeout: conductoronesdkterraform.String("esse"),
                            Name: conductoronesdkterraform.String("Miss Kim Jacobs II"),
                            TimeoutDuration: conductoronesdkterraform.String("fugiat"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("expedita"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("aliquid"),
                                AppID: conductoronesdkterraform.String("officia"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "aliquid",
                                    "perferendis",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "voluptas",
                                    "iste",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "ab",
                                    "error",
                                    "possimus",
                                },
                                Expressions: []string{
                                    "mollitia",
                                    "laborum",
                                    "libero",
                                    "ad",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "enim",
                                    "vitae",
                                    "repellendus",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "quo",
                                    "ex",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "ad",
                                    "expedita",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "molestias",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "aliquid",
                                    "beatae",
                                    "voluptatum",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "veritatis",
                                    "rerum",
                                    "est",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("culpa"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("voluptatem"),
                                    EntitlementID: conductoronesdkterraform.String("sapiente"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("officiis"),
                                    ConnectorID: conductoronesdkterraform.String("architecto"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("fuga"),
                                    Instructions: conductoronesdkterraform.String("pariatur"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("debitis"),
                                    UserIds: []string{
                                        "alias",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("earum"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("ex"),
                                AppID: conductoronesdkterraform.String("sapiente"),
                                AppUserID: conductoronesdkterraform.String("rem"),
                                GrantDuration: conductoronesdkterraform.String("minus"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("nemo"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("asperiores"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("ratione"),
                            CommentOnTimeout: conductoronesdkterraform.String("ullam"),
                            Name: conductoronesdkterraform.String("Desiree Leannon"),
                            TimeoutDuration: conductoronesdkterraform.String("nam"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("ipsam"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("culpa"),
                                AppID: conductoronesdkterraform.String("dolor"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "inventore",
                                    "deleniti",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "tempora",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "consequatur",
                                },
                                Expressions: []string{
                                    "sit",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "fugit",
                                    "ab",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "quae",
                                    "dolor",
                                    "fugiat",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "consequuntur",
                                    "ipsa",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "eveniet",
                                    "impedit",
                                    "officiis",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "necessitatibus",
                                    "sed",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "nesciunt",
                                    "expedita",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("eum"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("vel"),
                                    EntitlementID: conductoronesdkterraform.String("voluptatum"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("magnam"),
                                    ConnectorID: conductoronesdkterraform.String("exercitationem"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("ab"),
                                    Instructions: conductoronesdkterraform.String("porro"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("autem"),
                                    UserIds: []string{
                                        "laboriosam",
                                        "recusandae",
                                        "consequuntur",
                                        "voluptatem",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("necessitatibus"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("quasi"),
                                AppID: conductoronesdkterraform.String("nisi"),
                                AppUserID: conductoronesdkterraform.String("at"),
                                GrantDuration: conductoronesdkterraform.String("vero"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("est"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("harum"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("sequi"),
                            CommentOnTimeout: conductoronesdkterraform.String("doloribus"),
                            Name: conductoronesdkterraform.String("Earnest McClure"),
                            TimeoutDuration: conductoronesdkterraform.String("blanditiis"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("officia"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("voluptas"),
                                AppID: conductoronesdkterraform.String("numquam"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "quos",
                                    "eius",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "ducimus",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "fuga",
                                },
                                Expressions: []string{
                                    "incidunt",
                                    "quasi",
                                    "rem",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "dicta",
                                    "nisi",
                                    "consequuntur",
                                    "consectetur",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "cupiditate",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "soluta",
                                    "alias",
                                    "omnis",
                                    "eos",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "iste",
                                    "magni",
                                    "inventore",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "accusamus",
                                    "voluptatibus",
                                    "distinctio",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "delectus",
                                    "minima",
                                    "praesentium",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("maxime"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("magnam"),
                                    EntitlementID: conductoronesdkterraform.String("temporibus"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("quos"),
                                    ConnectorID: conductoronesdkterraform.String("commodi"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("itaque"),
                                    Instructions: conductoronesdkterraform.String("commodi"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("totam"),
                                    UserIds: []string{
                                        "modi",
                                        "nam",
                                        "vero",
                                        "voluptatem",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("vel"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("alias"),
                                AppID: conductoronesdkterraform.String("quasi"),
                                AppUserID: conductoronesdkterraform.String("non"),
                                GrantDuration: conductoronesdkterraform.String("maiores"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("enim"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("sint"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("nulla"),
                            CommentOnTimeout: conductoronesdkterraform.String("deserunt"),
                            Name: conductoronesdkterraform.String("Ida Kilback"),
                            TimeoutDuration: conductoronesdkterraform.String("sint"),
                        },
                    },
                },
            },
            "accusamus": shared.PolicySteps{
                Steps: []shared.PolicyStep{
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("hic"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("necessitatibus"),
                                AppID: conductoronesdkterraform.String("asperiores"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "voluptas",
                                    "debitis",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "quae",
                                    "minus",
                                    "fuga",
                                    "laborum",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "velit",
                                },
                                Expressions: []string{
                                    "ipsum",
                                    "impedit",
                                    "magni",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "repudiandae",
                                    "nam",
                                    "dolore",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "voluptate",
                                    "sequi",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "neque",
                                    "quo",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "quibusdam",
                                    "iure",
                                    "odit",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "vel",
                                    "magnam",
                                    "quibusdam",
                                    "inventore",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "libero",
                                    "architecto",
                                    "voluptatibus",
                                    "quia",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("porro"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("aliquam"),
                                    EntitlementID: conductoronesdkterraform.String("velit"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("illo"),
                                    ConnectorID: conductoronesdkterraform.String("accusantium"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("vel"),
                                    Instructions: conductoronesdkterraform.String("ea"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("beatae"),
                                    UserIds: []string{
                                        "excepturi",
                                        "eum",
                                        "velit",
                                        "ut",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("earum"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("dicta"),
                                AppID: conductoronesdkterraform.String("impedit"),
                                AppUserID: conductoronesdkterraform.String("voluptatibus"),
                                GrantDuration: conductoronesdkterraform.String("iste"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("itaque"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("alias"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("nisi"),
                            CommentOnTimeout: conductoronesdkterraform.String("itaque"),
                            Name: conductoronesdkterraform.String("Maggie Friesen"),
                            TimeoutDuration: conductoronesdkterraform.String("sit"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("doloremque"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("consequatur"),
                                AppID: conductoronesdkterraform.String("officia"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "ea",
                                    "quidem",
                                    "voluptas",
                                    "facilis",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "perspiciatis",
                                    "expedita",
                                    "deleniti",
                                    "a",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "ullam",
                                    "unde",
                                },
                                Expressions: []string{
                                    "animi",
                                    "impedit",
                                    "ipsam",
                                    "corporis",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "error",
                                    "esse",
                                    "labore",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "vero",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "vitae",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "dolorem",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "qui",
                                    "iste",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "nemo",
                                    "soluta",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("libero"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("rem"),
                                    EntitlementID: conductoronesdkterraform.String("dolorum"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("odio"),
                                    ConnectorID: conductoronesdkterraform.String("fugit"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("alias"),
                                    Instructions: conductoronesdkterraform.String("magni"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("vel"),
                                    UserIds: []string{
                                        "quae",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("neque"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("exercitationem"),
                                AppID: conductoronesdkterraform.String("itaque"),
                                AppUserID: conductoronesdkterraform.String("et"),
                                GrantDuration: conductoronesdkterraform.String("ipsum"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("unde"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("nulla"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("distinctio"),
                            CommentOnTimeout: conductoronesdkterraform.String("maxime"),
                            Name: conductoronesdkterraform.String("Marilyn Heaney"),
                            TimeoutDuration: conductoronesdkterraform.String("dicta"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("id"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("libero"),
                                AppID: conductoronesdkterraform.String("fugiat"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "quos",
                                    "placeat",
                                    "sit",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "ipsa",
                                    "voluptates",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "aperiam",
                                },
                                Expressions: []string{
                                    "dolore",
                                    "eligendi",
                                    "distinctio",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "autem",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "dolores",
                                    "assumenda",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "est",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "corrupti",
                                    "molestiae",
                                    "provident",
                                    "accusamus",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "tempore",
                                    "sint",
                                    "ea",
                                    "autem",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "rerum",
                                    "laudantium",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("corporis"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("officiis"),
                                    EntitlementID: conductoronesdkterraform.String("voluptatibus"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("cum"),
                                    ConnectorID: conductoronesdkterraform.String("at"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("alias"),
                                    Instructions: conductoronesdkterraform.String("quia"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("quidem"),
                                    UserIds: []string{
                                        "repudiandae",
                                        "accusantium",
                                        "expedita",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("eos"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("quibusdam"),
                                AppID: conductoronesdkterraform.String("odio"),
                                AppUserID: conductoronesdkterraform.String("praesentium"),
                                GrantDuration: conductoronesdkterraform.String("odit"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("explicabo"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("corporis"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("error"),
                            CommentOnTimeout: conductoronesdkterraform.String("earum"),
                            Name: conductoronesdkterraform.String("Delia Murray"),
                            TimeoutDuration: conductoronesdkterraform.String("quis"),
                        },
                    },
                    shared.PolicyStep{
                        Accept: &shared.Accept{
                            AcceptMessage: conductoronesdkterraform.String("beatae"),
                        },
                        Approval: &shared.Approval{
                            AppGroupApproval: &shared.AppGroupApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AppGroupID: conductoronesdkterraform.String("unde"),
                                AppID: conductoronesdkterraform.String("molestiae"),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "cupiditate",
                                    "fugit",
                                    "numquam",
                                    "numquam",
                                },
                            },
                            AppOwnerApproval: &shared.AppOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                            },
                            EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "at",
                                },
                            },
                            ExpressionApproval: &shared.ExpressionApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "dignissimos",
                                    "optio",
                                    "necessitatibus",
                                },
                                Expressions: []string{
                                    "qui",
                                    "expedita",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "cupiditate",
                                    "minima",
                                    "placeat",
                                },
                            },
                            ManagerApproval: &shared.ManagerApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                AssignedUserIds: []string{
                                    "neque",
                                    "in",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "eum",
                                    "modi",
                                    "corporis",
                                    "magnam",
                                },
                            },
                            SelfApproval: &shared.SelfApproval{
                                AssignedUserIds: []string{
                                    "maiores",
                                    "tempore",
                                    "aperiam",
                                    "libero",
                                },
                                Fallback: conductoronesdkterraform.Bool(false),
                                FallbackUserIds: []string{
                                    "labore",
                                },
                            },
                            UserApproval: &shared.UserApproval{
                                AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                UserIds: []string{
                                    "occaecati",
                                    "voluptas",
                                    "quo",
                                },
                            },
                            WebhookApproval: &shared.WebhookApproval{
                                WebhookID: conductoronesdkterraform.String("velit"),
                            },
                            AllowReassignment: conductoronesdkterraform.Bool(false),
                            Assigned: conductoronesdkterraform.Bool(false),
                            RequireApprovalReason: conductoronesdkterraform.Bool(false),
                            RequireDenialReason: conductoronesdkterraform.Bool(false),
                            RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                        },
                        Provision: &shared.Provision{
                            ProvisionPolicy: &shared.ProvisionPolicy{
                                ConnectorProvision: &shared.ConnectorProvision{},
                                DelegatedProvision: &shared.DelegatedProvision{
                                    AppID: conductoronesdkterraform.String("minus"),
                                    EntitlementID: conductoronesdkterraform.String("fuga"),
                                    Implicit: conductoronesdkterraform.Bool(false),
                                },
                                ExternalTicketProvision: &shared.ExternalTicketProvision{
                                    AppID: conductoronesdkterraform.String("nostrum"),
                                    ConnectorID: conductoronesdkterraform.String("est"),
                                    ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("impedit"),
                                    Instructions: conductoronesdkterraform.String("delectus"),
                                },
                                ManualProvision: &shared.ManualProvision{
                                    Instructions: conductoronesdkterraform.String("tempore"),
                                    UserIds: []string{
                                        "odit",
                                        "repellat",
                                        "pariatur",
                                        "nemo",
                                    },
                                },
                                MultiStep: &shared.MultiStep{
                                    ProvisionSteps: []shared.ProvisionPolicy{
                                        shared.ProvisionPolicy{},
                                        shared.ProvisionPolicy{},
                                    },
                                },
                                WebhookProvision: &shared.WebhookProvision{
                                    WebhookID: conductoronesdkterraform.String("aperiam"),
                                },
                            },
                            ProvisionTarget: &shared.ProvisionTarget{
                                AppEntitlementID: conductoronesdkterraform.String("odio"),
                                AppID: conductoronesdkterraform.String("minima"),
                                AppUserID: conductoronesdkterraform.String("in"),
                                GrantDuration: conductoronesdkterraform.String("ducimus"),
                            },
                            Assigned: conductoronesdkterraform.Bool(false),
                        },
                        Reject: &shared.Reject{
                            RejectMessage: conductoronesdkterraform.String("excepturi"),
                        },
                        Wait: &shared.Wait{
                            WaitCondition: &shared.WaitCondition{
                                Condition: conductoronesdkterraform.String("dolores"),
                            },
                            CommentOnFirstWait: conductoronesdkterraform.String("error"),
                            CommentOnTimeout: conductoronesdkterraform.String("veritatis"),
                            Name: conductoronesdkterraform.String("Colleen Streich"),
                            TimeoutDuration: conductoronesdkterraform.String("optio"),
                        },
                    },
                },
            },
        },
        PolicyType: shared.CreatePolicyRequestPolicyTypePolicyTypeRevoke.ToPointer(),
        PostActions: []shared.PolicyPostActions{
            shared.PolicyPostActions{
                CertifyRemediateImmediately: conductoronesdkterraform.Bool(false),
            },
            shared.PolicyPostActions{
                CertifyRemediateImmediately: conductoronesdkterraform.Bool(false),
            },
        },
        ReassignTasksToDelegates: conductoronesdkterraform.Bool(false),
        Rules: []shared.Rule{
            shared.Rule{
                Condition: conductoronesdkterraform.String("officiis"),
                PolicyKey: conductoronesdkterraform.String("placeat"),
            },
            shared.Rule{
                Condition: conductoronesdkterraform.String("quidem"),
                PolicyKey: conductoronesdkterraform.String("exercitationem"),
            },
        },
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

Delete a policy by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Delete(ctx, operations.C1APIPolicyV1PoliciesDeleteRequest{
        DeletePolicyRequest: &shared.DeletePolicyRequest{},
        ID: "73409e3e-b1e5-4a2b-92eb-07f116db9954",
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

Get a policy by ID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Get(ctx, operations.C1APIPolicyV1PoliciesGetRequest{
        ID: "5fc95fa8-8970-4e18-9dbb-30fcb33ea055",
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

List policies.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.List(ctx, operations.C1APIPolicyV1PoliciesListRequest{
        PageSize: conductoronesdkterraform.Int(727481),
        PageToken: conductoronesdkterraform.String("architecto"),
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.C1APIPolicyV1PoliciesListRequest](../../models/operations/c1apipolicyv1policieslistrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.C1APIPolicyV1PoliciesListResponse](../../models/operations/c1apipolicyv1policieslistresponse.md), error**


## Update

Update a policy by providing a policy object and an update mask.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Policies.Update(ctx, operations.C1APIPolicyV1PoliciesUpdateRequest{
        UpdatePolicyRequestInput: &shared.UpdatePolicyRequestInput{
            Policy: &shared.PolicyInput{
                Description: conductoronesdkterraform.String("cupiditate"),
                DisplayName: conductoronesdkterraform.String("molestiae"),
                PolicySteps: map[string]shared.PolicySteps{
                    "possimus": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("magnam"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("itaque"),
                                        AppID: conductoronesdkterraform.String("sed"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "veniam",
                                            "consequuntur",
                                            "facere",
                                            "laudantium",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "pariatur",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "exercitationem",
                                        },
                                        Expressions: []string{
                                            "velit",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "tempore",
                                            "nisi",
                                            "voluptatibus",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "blanditiis",
                                            "distinctio",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "quis",
                                            "nisi",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "minus",
                                            "facere",
                                            "facilis",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "ad",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "voluptatibus",
                                            "consequuntur",
                                            "debitis",
                                            "labore",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("rerum"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("eos"),
                                            EntitlementID: conductoronesdkterraform.String("reprehenderit"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("nostrum"),
                                            ConnectorID: conductoronesdkterraform.String("neque"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("iusto"),
                                            Instructions: conductoronesdkterraform.String("est"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("rem"),
                                            UserIds: []string{
                                                "fugiat",
                                                "unde",
                                                "officiis",
                                                "ducimus",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("dicta"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("error"),
                                        AppID: conductoronesdkterraform.String("porro"),
                                        AppUserID: conductoronesdkterraform.String("vitae"),
                                        GrantDuration: conductoronesdkterraform.String("dignissimos"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("esse"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("fugiat"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("ad"),
                                    CommentOnTimeout: conductoronesdkterraform.String("aspernatur"),
                                    Name: conductoronesdkterraform.String("Marta Kub"),
                                    TimeoutDuration: conductoronesdkterraform.String("illo"),
                                },
                            },
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("ab"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("incidunt"),
                                        AppID: conductoronesdkterraform.String("accusamus"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "tempore",
                                            "veniam",
                                            "eos",
                                            "reiciendis",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "reprehenderit",
                                            "praesentium",
                                            "nemo",
                                            "repellat",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "sequi",
                                            "nihil",
                                            "deleniti",
                                            "illo",
                                        },
                                        Expressions: []string{
                                            "assumenda",
                                            "aliquam",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "provident",
                                            "laudantium",
                                            "repudiandae",
                                            "consequatur",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "aspernatur",
                                            "nam",
                                            "expedita",
                                            "quas",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "repudiandae",
                                            "rerum",
                                            "dignissimos",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "vero",
                                            "similique",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "iure",
                                            "dolorem",
                                            "commodi",
                                            "impedit",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "aut",
                                            "voluptatem",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("ad"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("quae"),
                                            EntitlementID: conductoronesdkterraform.String("amet"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("illum"),
                                            ConnectorID: conductoronesdkterraform.String("praesentium"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("quidem"),
                                            Instructions: conductoronesdkterraform.String("cum"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("amet"),
                                            UserIds: []string{
                                                "dicta",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("doloremque"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("earum"),
                                        AppID: conductoronesdkterraform.String("iusto"),
                                        AppUserID: conductoronesdkterraform.String("amet"),
                                        GrantDuration: conductoronesdkterraform.String("provident"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("dolorum"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("necessitatibus"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("provident"),
                                    CommentOnTimeout: conductoronesdkterraform.String("repudiandae"),
                                    Name: conductoronesdkterraform.String("Vivian Kreiger"),
                                    TimeoutDuration: conductoronesdkterraform.String("corrupti"),
                                },
                            },
                        },
                    },
                    "aperiam": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("accusamus"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("eos"),
                                        AppID: conductoronesdkterraform.String("totam"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "voluptatem",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "dolor",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "a",
                                        },
                                        Expressions: []string{
                                            "occaecati",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "beatae",
                                            "at",
                                            "labore",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "esse",
                                            "voluptatem",
                                            "perferendis",
                                            "rerum",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "aperiam",
                                            "dignissimos",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "velit",
                                            "porro",
                                            "provident",
                                            "consectetur",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "dignissimos",
                                            "consectetur",
                                            "soluta",
                                            "natus",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "officia",
                                            "amet",
                                            "tenetur",
                                            "aspernatur",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("quo"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("itaque"),
                                            EntitlementID: conductoronesdkterraform.String("illum"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("laborum"),
                                            ConnectorID: conductoronesdkterraform.String("dignissimos"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("vero"),
                                            Instructions: conductoronesdkterraform.String("qui"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("consectetur"),
                                            UserIds: []string{
                                                "explicabo",
                                                "explicabo",
                                                "exercitationem",
                                                "nihil",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("ab"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("illo"),
                                        AppID: conductoronesdkterraform.String("hic"),
                                        AppUserID: conductoronesdkterraform.String("deserunt"),
                                        GrantDuration: conductoronesdkterraform.String("delectus"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("non"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("distinctio"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("in"),
                                    CommentOnTimeout: conductoronesdkterraform.String("exercitationem"),
                                    Name: conductoronesdkterraform.String("Marjorie Waelchi"),
                                    TimeoutDuration: conductoronesdkterraform.String("explicabo"),
                                },
                            },
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("accusamus"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("rem"),
                                        AppID: conductoronesdkterraform.String("aperiam"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "deleniti",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "voluptate",
                                            "similique",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "libero",
                                            "magnam",
                                        },
                                        Expressions: []string{
                                            "modi",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "nesciunt",
                                            "mollitia",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "fugiat",
                                            "nostrum",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "veniam",
                                            "reiciendis",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "modi",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "aut",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "odio",
                                            "commodi",
                                            "numquam",
                                            "dolorum",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("possimus"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("voluptate"),
                                            EntitlementID: conductoronesdkterraform.String("consectetur"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("nesciunt"),
                                            ConnectorID: conductoronesdkterraform.String("quaerat"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("itaque"),
                                            Instructions: conductoronesdkterraform.String("minus"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("sunt"),
                                            UserIds: []string{
                                                "iusto",
                                                "quas",
                                                "et",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("amet"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("autem"),
                                        AppID: conductoronesdkterraform.String("fuga"),
                                        AppUserID: conductoronesdkterraform.String("alias"),
                                        GrantDuration: conductoronesdkterraform.String("rem"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("aut"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("quos"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("laudantium"),
                                    CommentOnTimeout: conductoronesdkterraform.String("repellendus"),
                                    Name: conductoronesdkterraform.String("Brenda Bechtelar"),
                                    TimeoutDuration: conductoronesdkterraform.String("mollitia"),
                                },
                            },
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("nulla"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("officia"),
                                        AppID: conductoronesdkterraform.String("sed"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "alias",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "hic",
                                            "voluptatem",
                                            "incidunt",
                                            "qui",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "necessitatibus",
                                        },
                                        Expressions: []string{
                                            "explicabo",
                                            "beatae",
                                            "aliquid",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "optio",
                                            "voluptatibus",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "officia",
                                            "libero",
                                            "totam",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "aliquid",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "impedit",
                                            "ducimus",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "velit",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "repellat",
                                            "nulla",
                                            "laborum",
                                            "natus",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("accusamus"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("doloremque"),
                                            EntitlementID: conductoronesdkterraform.String("nisi"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("rerum"),
                                            ConnectorID: conductoronesdkterraform.String("recusandae"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("voluptates"),
                                            Instructions: conductoronesdkterraform.String("non"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("rem"),
                                            UserIds: []string{
                                                "ullam",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("dicta"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("voluptatibus"),
                                        AppID: conductoronesdkterraform.String("eligendi"),
                                        AppUserID: conductoronesdkterraform.String("quae"),
                                        GrantDuration: conductoronesdkterraform.String("officiis"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("architecto"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("architecto"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("enim"),
                                    CommentOnTimeout: conductoronesdkterraform.String("optio"),
                                    Name: conductoronesdkterraform.String("Joseph Purdy"),
                                    TimeoutDuration: conductoronesdkterraform.String("iste"),
                                },
                            },
                        },
                    },
                    "dicta": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("ullam"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("dolore"),
                                        AppID: conductoronesdkterraform.String("modi"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "maxime",
                                            "modi",
                                            "consequuntur",
                                            "assumenda",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "doloribus",
                                            "impedit",
                                            "porro",
                                            "accusamus",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "reiciendis",
                                            "ab",
                                            "sint",
                                        },
                                        Expressions: []string{
                                            "esse",
                                            "iure",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "nesciunt",
                                            "debitis",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "neque",
                                            "corporis",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "consequuntur",
                                            "officia",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "distinctio",
                                            "eius",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "rem",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "accusantium",
                                            "veniam",
                                            "saepe",
                                            "neque",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("facere"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("aliquam"),
                                            EntitlementID: conductoronesdkterraform.String("quos"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("doloribus"),
                                            ConnectorID: conductoronesdkterraform.String("fugiat"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("est"),
                                            Instructions: conductoronesdkterraform.String("delectus"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("velit"),
                                            UserIds: []string{
                                                "nesciunt",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("illo"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("repellat"),
                                        AppID: conductoronesdkterraform.String("nemo"),
                                        AppUserID: conductoronesdkterraform.String("doloribus"),
                                        GrantDuration: conductoronesdkterraform.String("possimus"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("unde"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("incidunt"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("explicabo"),
                                    CommentOnTimeout: conductoronesdkterraform.String("ipsam"),
                                    Name: conductoronesdkterraform.String("Lucas Abbott"),
                                    TimeoutDuration: conductoronesdkterraform.String("commodi"),
                                },
                            },
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("sapiente"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("consequuntur"),
                                        AppID: conductoronesdkterraform.String("veniam"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "officia",
                                            "sint",
                                            "ut",
                                            "numquam",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "adipisci",
                                            "libero",
                                            "in",
                                            "minima",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "minus",
                                            "ab",
                                        },
                                        Expressions: []string{
                                            "hic",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "quisquam",
                                            "dolor",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "fuga",
                                            "minima",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "qui",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "magni",
                                            "incidunt",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "praesentium",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "exercitationem",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("expedita"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("facilis"),
                                            EntitlementID: conductoronesdkterraform.String("impedit"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("sit"),
                                            ConnectorID: conductoronesdkterraform.String("nemo"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("culpa"),
                                            Instructions: conductoronesdkterraform.String("consequuntur"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("amet"),
                                            UserIds: []string{
                                                "modi",
                                                "veniam",
                                                "quod",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("a"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("quisquam"),
                                        AppID: conductoronesdkterraform.String("enim"),
                                        AppUserID: conductoronesdkterraform.String("doloribus"),
                                        GrantDuration: conductoronesdkterraform.String("assumenda"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("officiis"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("architecto"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("alias"),
                                    CommentOnTimeout: conductoronesdkterraform.String("culpa"),
                                    Name: conductoronesdkterraform.String("Della Treutel III"),
                                    TimeoutDuration: conductoronesdkterraform.String("perspiciatis"),
                                },
                            },
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("debitis"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("ullam"),
                                        AppID: conductoronesdkterraform.String("architecto"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "perferendis",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "provident",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "iure",
                                            "quibusdam",
                                            "quod",
                                            "nemo",
                                        },
                                        Expressions: []string{
                                            "velit",
                                            "magnam",
                                            "dignissimos",
                                            "laboriosam",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "odio",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "provident",
                                            "cum",
                                            "doloribus",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "quidem",
                                            "itaque",
                                            "laboriosam",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "modi",
                                            "perspiciatis",
                                            "hic",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "aspernatur",
                                            "libero",
                                            "nam",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "recusandae",
                                            "quod",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("id"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("saepe"),
                                            EntitlementID: conductoronesdkterraform.String("autem"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("quo"),
                                            ConnectorID: conductoronesdkterraform.String("nesciunt"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("illum"),
                                            Instructions: conductoronesdkterraform.String("nemo"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("illum"),
                                            UserIds: []string{
                                                "non",
                                                "mollitia",
                                                "assumenda",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("distinctio"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("pariatur"),
                                        AppID: conductoronesdkterraform.String("ad"),
                                        AppUserID: conductoronesdkterraform.String("facere"),
                                        GrantDuration: conductoronesdkterraform.String("laborum"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("eveniet"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("laborum"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("incidunt"),
                                    CommentOnTimeout: conductoronesdkterraform.String("maxime"),
                                    Name: conductoronesdkterraform.String("Mary Hoeger"),
                                    TimeoutDuration: conductoronesdkterraform.String("laborum"),
                                },
                            },
                        },
                    },
                    "est": shared.PolicySteps{
                        Steps: []shared.PolicyStep{
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("labore"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("quo"),
                                        AppID: conductoronesdkterraform.String("perferendis"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "aliquid",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "quaerat",
                                            "eligendi",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "nostrum",
                                            "officiis",
                                            "unde",
                                            "nulla",
                                        },
                                        Expressions: []string{
                                            "mollitia",
                                            "magnam",
                                            "nostrum",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "corrupti",
                                            "fuga",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "impedit",
                                            "quasi",
                                            "deserunt",
                                            "quod",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "doloremque",
                                            "voluptatem",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "necessitatibus",
                                            "maxime",
                                            "consequatur",
                                            "eaque",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "similique",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "blanditiis",
                                            "quae",
                                            "magni",
                                            "officiis",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("sed"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("necessitatibus"),
                                            EntitlementID: conductoronesdkterraform.String("impedit"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("ipsa"),
                                            ConnectorID: conductoronesdkterraform.String("excepturi"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("a"),
                                            Instructions: conductoronesdkterraform.String("maiores"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("laudantium"),
                                            UserIds: []string{
                                                "alias",
                                                "asperiores",
                                                "rem",
                                                "dicta",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("earum"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("doloribus"),
                                        AppID: conductoronesdkterraform.String("velit"),
                                        AppUserID: conductoronesdkterraform.String("eius"),
                                        GrantDuration: conductoronesdkterraform.String("esse"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("in"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("eligendi"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("quasi"),
                                    CommentOnTimeout: conductoronesdkterraform.String("neque"),
                                    Name: conductoronesdkterraform.String("Andy Bashirian"),
                                    TimeoutDuration: conductoronesdkterraform.String("beatae"),
                                },
                            },
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("incidunt"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("dicta"),
                                        AppID: conductoronesdkterraform.String("odit"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "rerum",
                                            "alias",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "vel",
                                            "accusantium",
                                            "id",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "ex",
                                            "quas",
                                        },
                                        Expressions: []string{
                                            "ullam",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "similique",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "quam",
                                            "magni",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "delectus",
                                            "omnis",
                                            "sed",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "maxime",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "cupiditate",
                                            "aliquam",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "maiores",
                                            "laudantium",
                                            "velit",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("reiciendis"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("amet"),
                                            EntitlementID: conductoronesdkterraform.String("nemo"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("ipsa"),
                                            ConnectorID: conductoronesdkterraform.String("quisquam"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("tenetur"),
                                            Instructions: conductoronesdkterraform.String("quas"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("molestiae"),
                                            UserIds: []string{
                                                "asperiores",
                                                "a",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("perspiciatis"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("accusantium"),
                                        AppID: conductoronesdkterraform.String("dicta"),
                                        AppUserID: conductoronesdkterraform.String("minus"),
                                        GrantDuration: conductoronesdkterraform.String("commodi"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("eveniet"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("porro"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("tempore"),
                                    CommentOnTimeout: conductoronesdkterraform.String("quidem"),
                                    Name: conductoronesdkterraform.String("Silvia Considine"),
                                    TimeoutDuration: conductoronesdkterraform.String("eligendi"),
                                },
                            },
                            shared.PolicyStep{
                                Accept: &shared.Accept{
                                    AcceptMessage: conductoronesdkterraform.String("asperiores"),
                                },
                                Approval: &shared.Approval{
                                    AppGroupApproval: &shared.AppGroupApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AppGroupID: conductoronesdkterraform.String("esse"),
                                        AppID: conductoronesdkterraform.String("blanditiis"),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "repellat",
                                            "a",
                                            "animi",
                                        },
                                    },
                                    AppOwnerApproval: &shared.AppOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                    },
                                    EntitlementOwnerApproval: &shared.EntitlementOwnerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "itaque",
                                            "nulla",
                                            "deserunt",
                                            "corporis",
                                        },
                                    },
                                    ExpressionApproval: &shared.ExpressionApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "officiis",
                                        },
                                        Expressions: []string{
                                            "officia",
                                            "saepe",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "repudiandae",
                                            "accusantium",
                                        },
                                    },
                                    ManagerApproval: &shared.ManagerApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        AssignedUserIds: []string{
                                            "impedit",
                                            "quasi",
                                            "blanditiis",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "quisquam",
                                            "eos",
                                        },
                                    },
                                    SelfApproval: &shared.SelfApproval{
                                        AssignedUserIds: []string{
                                            "natus",
                                            "minus",
                                            "quia",
                                        },
                                        Fallback: conductoronesdkterraform.Bool(false),
                                        FallbackUserIds: []string{
                                            "reprehenderit",
                                            "quod",
                                        },
                                    },
                                    UserApproval: &shared.UserApproval{
                                        AllowSelfApproval: conductoronesdkterraform.Bool(false),
                                        UserIds: []string{
                                            "corrupti",
                                            "amet",
                                            "molestiae",
                                        },
                                    },
                                    WebhookApproval: &shared.WebhookApproval{
                                        WebhookID: conductoronesdkterraform.String("amet"),
                                    },
                                    AllowReassignment: conductoronesdkterraform.Bool(false),
                                    Assigned: conductoronesdkterraform.Bool(false),
                                    RequireApprovalReason: conductoronesdkterraform.Bool(false),
                                    RequireDenialReason: conductoronesdkterraform.Bool(false),
                                    RequireReassignmentReason: conductoronesdkterraform.Bool(false),
                                },
                                Provision: &shared.Provision{
                                    ProvisionPolicy: &shared.ProvisionPolicy{
                                        ConnectorProvision: &shared.ConnectorProvision{},
                                        DelegatedProvision: &shared.DelegatedProvision{
                                            AppID: conductoronesdkterraform.String("laborum"),
                                            EntitlementID: conductoronesdkterraform.String("modi"),
                                            Implicit: conductoronesdkterraform.Bool(false),
                                        },
                                        ExternalTicketProvision: &shared.ExternalTicketProvision{
                                            AppID: conductoronesdkterraform.String("perferendis"),
                                            ConnectorID: conductoronesdkterraform.String("necessitatibus"),
                                            ExternalTicketProvisionerConfigID: conductoronesdkterraform.String("architecto"),
                                            Instructions: conductoronesdkterraform.String("molestias"),
                                        },
                                        ManualProvision: &shared.ManualProvision{
                                            Instructions: conductoronesdkterraform.String("dolore"),
                                            UserIds: []string{
                                                "maiores",
                                            },
                                        },
                                        MultiStep: &shared.MultiStep{
                                            ProvisionSteps: []shared.ProvisionPolicy{
                                                shared.ProvisionPolicy{},
                                            },
                                        },
                                        WebhookProvision: &shared.WebhookProvision{
                                            WebhookID: conductoronesdkterraform.String("odit"),
                                        },
                                    },
                                    ProvisionTarget: &shared.ProvisionTarget{
                                        AppEntitlementID: conductoronesdkterraform.String("earum"),
                                        AppID: conductoronesdkterraform.String("veniam"),
                                        AppUserID: conductoronesdkterraform.String("ipsam"),
                                        GrantDuration: conductoronesdkterraform.String("eaque"),
                                    },
                                    Assigned: conductoronesdkterraform.Bool(false),
                                },
                                Reject: &shared.Reject{
                                    RejectMessage: conductoronesdkterraform.String("exercitationem"),
                                },
                                Wait: &shared.Wait{
                                    WaitCondition: &shared.WaitCondition{
                                        Condition: conductoronesdkterraform.String("veniam"),
                                    },
                                    CommentOnFirstWait: conductoronesdkterraform.String("nihil"),
                                    CommentOnTimeout: conductoronesdkterraform.String("ad"),
                                    Name: conductoronesdkterraform.String("Essie Hayes"),
                                    TimeoutDuration: conductoronesdkterraform.String("suscipit"),
                                },
                            },
                        },
                    },
                },
                PolicyType: shared.PolicyPolicyTypePolicyTypeProvision.ToPointer(),
                PostActions: []shared.PolicyPostActions{
                    shared.PolicyPostActions{
                        CertifyRemediateImmediately: conductoronesdkterraform.Bool(false),
                    },
                },
                ReassignTasksToDelegates: conductoronesdkterraform.Bool(false),
                Rules: []shared.Rule{
                    shared.Rule{
                        Condition: conductoronesdkterraform.String("repellendus"),
                        PolicyKey: conductoronesdkterraform.String("perferendis"),
                    },
                    shared.Rule{
                        Condition: conductoronesdkterraform.String("id"),
                        PolicyKey: conductoronesdkterraform.String("sapiente"),
                    },
                    shared.Rule{
                        Condition: conductoronesdkterraform.String("sed"),
                        PolicyKey: conductoronesdkterraform.String("possimus"),
                    },
                },
            },
            UpdateMask: conductoronesdkterraform.String("repellat"),
        },
        ID: "e13db4f6-2cba-43f8-941a-ebc0b80a6924",
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

