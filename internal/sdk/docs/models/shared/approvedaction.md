# ApprovedAction

The approved action indicates that the approvalinstance had an outcome of approved.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ApprovedAt`                                                                           | [*time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `Entitlements`                                                                         | [][AppEntitlementReference](../../models/shared/appentitlementreference.md)            | :heavy_minus_sign:                                                                     | The entitlements that were approved. This will only ever be a list of one entitlement. |
| `UserID`                                                                               | **string*                                                                              | :heavy_minus_sign:                                                                     | The UserID that approved this step.                                                    |