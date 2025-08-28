# UsageBasedRevocationTrigger

The UsageBasedRevocationTrigger message.

This message contains a oneof named cold_start_schedule. Only a single field of the following list may be set at a time:
  - runImmediately
  - runDelayed



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `AppID`                                                                              | **string*                                                                            | :heavy_minus_sign:                                                                   | The appId field.                                                                     |
| `EnabledAt`                                                                          | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `ExcludedGroupRefs`                                                                  | [][components.AppEntitlementRef](../../models/components/appentitlementref.md)       | :heavy_minus_sign:                                                                   | The excludedGroupRefs field.                                                         |
| `ExcludedUserRefs`                                                                   | [][components.UserRef](../../models/components/userref.md)                           | :heavy_minus_sign:                                                                   | The excludedUserRefs field.                                                          |
| `IncludeUsersWithNoActivity`                                                         | **bool*                                                                              | :heavy_minus_sign:                                                                   | The includeUsersWithNoActivity field.                                                |
| `RunDelayed`                                                                         | [*components.RunDelayed](../../models/components/rundelayed.md)                      | :heavy_minus_sign:                                                                   | The RunDelayed message.                                                              |
| `RunImmediately`                                                                     | [*components.RunImmediately](../../models/components/runimmediately.md)              | :heavy_minus_sign:                                                                   | No fields needed; this just indicates the trigger should run immediately             |
| `TargetedAppUserTypes`                                                               | [][components.TargetedAppUserTypes](../../models/components/targetedappusertypes.md) | :heavy_minus_sign:                                                                   | The targetedAppUserTypes field.                                                      |
| `TargetedEntitlementRefs`                                                            | [][components.AppEntitlementRef](../../models/components/appentitlementref.md)       | :heavy_minus_sign:                                                                   | The targetedEntitlementRefs field.                                                   |
| `UnusedForDays`                                                                      | **int64*                                                                             | :heavy_minus_sign:                                                                   | The unusedForDays field.                                                             |