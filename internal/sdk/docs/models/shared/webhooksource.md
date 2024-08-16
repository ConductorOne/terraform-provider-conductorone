# WebhookSource

The WebhookSource message.

This message contains a oneof named source. Only a single field of the following list may be set at a time:
  - test
  - policyPostAction
  - approvalStep
  - provisionStep



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `WebhookSourceApprovalStep`                                                            | [*WebhookSourceApprovalStep](../../models/shared/webhooksourceapprovalstep.md)         | :heavy_minus_sign:                                                                     | The WebhookSourceApprovalStep message.                                                 |
| `WebhookSourcePolicyPostAction`                                                        | [*WebhookSourcePolicyPostAction](../../models/shared/webhooksourcepolicypostaction.md) | :heavy_minus_sign:                                                                     | The WebhookSourcePolicyPostAction message.                                             |
| `WebhookSourceProvisionStep`                                                           | [*WebhookSourceProvisionStep](../../models/shared/webhooksourceprovisionstep.md)       | :heavy_minus_sign:                                                                     | The WebhookSourceProvisionStep message.                                                |
| `WebhookSourceTest`                                                                    | [*WebhookSourceTest](../../models/shared/webhooksourcetest.md)                         | :heavy_minus_sign:                                                                     | The WebhookSourceTest message.                                                         |