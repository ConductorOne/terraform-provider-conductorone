# WebhookSource

The WebhookSource message.

This message contains a oneof named source. Only a single field of the following list may be set at a time:
  - test
  - policyPostAction
  - approvalStep
  - provisionStep
  - workflowStep



## Fields

| Field                                                                                                 | Type                                                                                                  | Required                                                                                              | Description                                                                                           |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `WebhookSourceApprovalStep`                                                                           | [*components.WebhookSourceApprovalStep](../../models/components/webhooksourceapprovalstep.md)         | :heavy_minus_sign:                                                                                    | The WebhookSourceApprovalStep message.                                                                |
| `WebhookSourcePolicyPostAction`                                                                       | [*components.WebhookSourcePolicyPostAction](../../models/components/webhooksourcepolicypostaction.md) | :heavy_minus_sign:                                                                                    | The WebhookSourcePolicyPostAction message.                                                            |
| `WebhookSourceProvisionStep`                                                                          | [*components.WebhookSourceProvisionStep](../../models/components/webhooksourceprovisionstep.md)       | :heavy_minus_sign:                                                                                    | The WebhookSourceProvisionStep message.                                                               |
| `WebhookSourceTest`                                                                                   | [*components.WebhookSourceTest](../../models/components/webhooksourcetest.md)                         | :heavy_minus_sign:                                                                                    | The WebhookSourceTest message.                                                                        |
| `WebhookSourceWorkflowStep`                                                                           | [*components.WebhookSourceWorkflowStep](../../models/components/webhooksourceworkflowstep.md)         | :heavy_minus_sign:                                                                                    | The WebhookSourceWorkflowStep message.                                                                |