# Escalation

The Escalation message.

This message contains a oneof named escalation_policy. Only a single field of the following list may be set at a time:
  - replacePolicy
  - reassignToApprovers



## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `EscalationComment`                                                               | **string*                                                                         | :heavy_minus_sign:                                                                | The escalationComment field.                                                      |
| `Expiration`                                                                      | **int64*                                                                          | :heavy_minus_sign:                                                                | The expiration field.                                                             |
| `ReassignToApprovers`                                                             | [*components.ReassignToApprovers](../../models/components/reassigntoapprovers.md) | :heavy_minus_sign:                                                                | The ReassignToApprovers message.                                                  |
| `ReplacePolicy`                                                                   | [*components.ReplacePolicy](../../models/components/replacepolicy.md)             | :heavy_minus_sign:                                                                | The ReplacePolicy message.                                                        |