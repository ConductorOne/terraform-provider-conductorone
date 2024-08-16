# AppUserServiceUpdateResponse

The AppUserServiceUpdateResponse message.


## Fields

| Field                                                                                                              | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `AppUserView`                                                                                                      | [*AppUserView](../../models/shared/appuserview.md)                                                                 | :heavy_minus_sign:                                                                                                 | The AppUserView contains an app user as well as paths for apps, identity users, and last usage in expanded arrays. |
| `Expanded`                                                                                                         | []map[string]*interface{}*                                                                                         | :heavy_minus_sign:                                                                                                 | The expanded field.                                                                                                |