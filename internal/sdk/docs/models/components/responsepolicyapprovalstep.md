# ResponsePolicyApprovalStep

The ResponsePolicyApprovalStep message.

This message contains a oneof named action. Only a single field of the following list may be set at a time:
  - approve
  - deny
  - reassign
  - replacePolicy



## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ResponsePolicyApprovalStepApprove`                                                                               | [*components.ResponsePolicyApprovalStepApprove](../../models/components/responsepolicyapprovalstepapprove.md)     | :heavy_minus_sign:                                                                                                | The ResponsePolicyApprovalStepApprove message.                                                                    |
| `ResponsePolicyApprovalStepDeny`                                                                                  | [*components.ResponsePolicyApprovalStepDeny](../../models/components/responsepolicyapprovalstepdeny.md)           | :heavy_minus_sign:                                                                                                | The ResponsePolicyApprovalStepDeny message.                                                                       |
| `ResponsePolicyApprovalStepReassign`                                                                              | [*components.ResponsePolicyApprovalStepReassign](../../models/components/responsepolicyapprovalstepreassign.md)   | :heavy_minus_sign:                                                                                                | The ResponsePolicyApprovalStepReassign message.                                                                   |
| `ResponsePolicyApprovalReplacePolicy`                                                                             | [*components.ResponsePolicyApprovalReplacePolicy](../../models/components/responsepolicyapprovalreplacepolicy.md) | :heavy_minus_sign:                                                                                                | The ResponsePolicyApprovalReplacePolicy message.                                                                  |
| `Version`                                                                                                         | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | version contains the constant value "v1". Future versions of the Webhook Response<br/> will use a different string. |