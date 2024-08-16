# PersonalClientServiceCreateRequest

The PersonalClientServiceCreateRequest message contains the fields for creating a new personal client.


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `AllowSourceCidr`                                | []*string*                                       | :heavy_minus_sign:                               | A list of CIDRs to restrict this credential to.  |
| `DisplayName`                                    | **string*                                        | :heavy_minus_sign:                               | The display name for the new personal client.    |
| `Expires`                                        | **string*                                        | :heavy_minus_sign:                               | N/A                                              |
| `ScopedRoles`                                    | []*string*                                       | :heavy_minus_sign:                               | The list of roles to restrict the credential to. |