# SetBundleAutomationRequest

The SetBundleAutomationRequest message.

This message contains a oneof named conditions. Only a single field of the following list may be set at a time:
  - entitlements



## Fields

| Field                                                                                                     | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `CreateTasks`                                                                                             | **bool*                                                                                                   | :heavy_minus_sign:                                                                                        | The createTasks field.                                                                                    |
| `Enabled`                                                                                                 | **bool*                                                                                                   | :heavy_minus_sign:                                                                                        | The enabled field.                                                                                        |
| `BundleAutomationRuleEntitlement`                                                                         | [*components.BundleAutomationRuleEntitlement](../../models/components/bundleautomationruleentitlement.md) | :heavy_minus_sign:                                                                                        | The BundleAutomationRuleEntitlement message.                                                              |