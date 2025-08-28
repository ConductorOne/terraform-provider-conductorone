# IntrospectResponse

IntrospectResponse contains information about the current user who is authenticated.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Features`                                                                      | []*string*                                                                      | :heavy_minus_sign:                                                              | The list of feature flags enabled for the tenant the logged in user belongs to. |
| `Permissions`                                                                   | []*string*                                                                      | :heavy_minus_sign:                                                              | The list of permissions that the current logged in user has.                    |
| `PrincipleID`                                                                   | **string*                                                                       | :heavy_minus_sign:                                                              | The principleID of the current logged in user.                                  |
| `Roles`                                                                         | []*string*                                                                      | :heavy_minus_sign:                                                              | The list of roles that the current logged in user has.                          |
| `UserID`                                                                        | **string*                                                                       | :heavy_minus_sign:                                                              | The userID of the current logged in user.                                       |