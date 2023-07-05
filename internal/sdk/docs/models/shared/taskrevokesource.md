# TaskRevokeSource

The TaskRevokeSource message.

This message contains a oneof named origin. Only a single field of the following list may be set at a time:
  - review
  - request
  - expired
  - nonUsage



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Expired`                                                                    | [*TaskRevokeSourceExpired](../../models/shared/taskrevokesourceexpired.md)   | :heavy_minus_sign:                                                           | The TaskRevokeSourceExpired message.                                         |
| `NonUsage`                                                                   | [*TaskRevokeSourceNonUsage](../../models/shared/taskrevokesourcenonusage.md) | :heavy_minus_sign:                                                           | The TaskRevokeSourceNonUsage message.                                        |
| `Request`                                                                    | [*TaskRevokeSourceRequest](../../models/shared/taskrevokesourcerequest.md)   | :heavy_minus_sign:                                                           | The TaskRevokeSourceRequest message.                                         |
| `Review`                                                                     | [*TaskRevokeSourceReview](../../models/shared/taskrevokesourcereview.md)     | :heavy_minus_sign:                                                           | The TaskRevokeSourceReview message.                                          |