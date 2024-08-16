# AppEntitlementUserBinding

The AppEntitlementUserBinding represents the relationship that gives an app user access to an app entitlement


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AppEntitlementID`                                            | **string*                                                     | :heavy_minus_sign:                                            | The ID of the app entitlement that the app user has access to |
| `AppID`                                                       | **string*                                                     | :heavy_minus_sign:                                            | The ID of the app associated with the app entitlement         |
| `AppUserID`                                                   | **string*                                                     | :heavy_minus_sign:                                            | The ID of the app user that has access to the app entitlement |
| `CreatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |
| `DeletedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |
| `DeprovisionAt`                                               | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |