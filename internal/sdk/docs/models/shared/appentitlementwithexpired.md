# AppEntitlementWithExpired

The AppEntitlementWithExpired message.


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `AppUser`                                                                               | [*AppUser](../../models/shared/appuser.md)                                              | :heavy_minus_sign:                                                                      | Application User that represents an account in the application.                         |
| `User`                                                                                  | [*User](../../models/shared/user.md)                                                    | :heavy_minus_sign:                                                                      | The User object provides all of the details for an user, as well as some configuration. |
| `Discovered`                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |
| `Expired`                                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                              | :heavy_minus_sign:                                                                      | N/A                                                                                     |