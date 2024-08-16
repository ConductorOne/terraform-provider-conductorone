# BundleAutomation

The BundleAutomation message.

This message contains a oneof named conditions. Only a single field of the following list may be set at a time:
  - entitlements



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `BundleAutomationLastRunState`                                                             | [*BundleAutomationLastRunState](../../models/shared/bundleautomationlastrunstate.md)       | :heavy_minus_sign:                                                                         | The BundleAutomationLastRunState message.                                                  |
| `BundleAutomationRuleEntitlement`                                                          | [*BundleAutomationRuleEntitlement](../../models/shared/bundleautomationruleentitlement.md) | :heavy_minus_sign:                                                                         | The BundleAutomationRuleEntitlement message.                                               |
| `CreatedAt`                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `DeletedAt`                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `RequestCatalogID`                                                                         | **string*                                                                                  | :heavy_minus_sign:                                                                         | The requestCatalogId field.                                                                |
| `TenantID`                                                                                 | **string*                                                                                  | :heavy_minus_sign:                                                                         | The tenantId field.                                                                        |
| `UpdatedAt`                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                 | :heavy_minus_sign:                                                                         | N/A                                                                                        |