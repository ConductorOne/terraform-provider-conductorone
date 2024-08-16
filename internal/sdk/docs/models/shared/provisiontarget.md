# ProvisionTarget

ProvisionTarget indicates the specific app, app entitlement, and if known, the app user and grant duration of this provision step


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AppEntitlementID`                                                               | **string*                                                                        | :heavy_minus_sign:                                                               | The app entitlement that should be provisioned.                                  |
| `AppID`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | The app in which the entitlement should be provisioned                           |
| `AppUserID`                                                                      | **string*                                                                        | :heavy_minus_sign:                                                               | The app user that should be provisioned. May be unset if the app user is unknown |
| `GrantDuration`                                                                  | **string*                                                                        | :heavy_minus_sign:                                                               | N/A                                                                              |