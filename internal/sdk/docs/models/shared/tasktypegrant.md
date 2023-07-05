# TaskTypeGrant

The TaskTypeGrant message.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppEntitlementID`                                                   | **string*                                                            | :heavy_minus_sign:                                                   | The appEntitlementId field.                                          |
| `AppID`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The appId field.                                                     |
| `AppUserID`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | The appUserId field.                                                 |
| `GrantDuration`                                                      | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `IdentityUserID`                                                     | **string*                                                            | :heavy_minus_sign:                                                   | The identityUserId field.                                            |
| `Outcome`                                                            | [*TaskTypeGrantOutcome](../../models/shared/tasktypegrantoutcome.md) | :heavy_minus_sign:                                                   | The outcome field.                                                   |
| `OutcomeTime`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |