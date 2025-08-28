# ApprovedAction

The approved action indicates that the approvalinstance had an outcome of approved.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ApprovedAt`                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                  | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `Entitlements`                                                                              | [][components.AppEntitlementReference](../../models/components/appentitlementreference.md)  | :heavy_minus_sign:                                                                          | The entitlements that were approved. This will only ever be a list of one entitlement.      |
| `StepUpTransactionID`                                                                       | **string*                                                                                   | :heavy_minus_sign:                                                                          | The ID of the step-up transaction that was used for this approval, if step-up was required. |
| `UserID`                                                                                    | **string*                                                                                   | :heavy_minus_sign:                                                                          | The UserID that approved this step.                                                         |