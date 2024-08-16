# CreateAppRequest

The CreateAppRequest message is used to create a new app.


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `CertifyPolicyID`                                | **string*                                        | :heavy_minus_sign:                               | Creates the app with this certify policy.        |
| `Description`                                    | **string*                                        | :heavy_minus_sign:                               | Creates the app with this description.           |
| `DisplayName`                                    | **string*                                        | :heavy_minus_sign:                               | Creates the app with this display name.          |
| `GrantPolicyID`                                  | **string*                                        | :heavy_minus_sign:                               | Creates the app with this grant policy.          |
| `MonthlyCostUsd`                                 | **int*                                           | :heavy_minus_sign:                               | Creates the app with this monthly cost per seat. |
| `Owners`                                         | []*string*                                       | :heavy_minus_sign:                               | Creates the app with this array of owners.       |
| `RevokePolicyID`                                 | **string*                                        | :heavy_minus_sign:                               | Creates the app with this revoke policy.         |