# TaskRevokeSource

The TaskRevokeSource message indicates the source of the revoke task is one of expired, nonUsage, request, or review.

This message contains a oneof named origin. Only a single field of the following list may be set at a time:
  - review
  - request
  - expired
  - nonUsage



## Fields

| Field                                                                                                                 | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `TaskRevokeSourceExpired`                                                                                             | [*TaskRevokeSourceExpired](../../models/shared/taskrevokesourceexpired.md)                                            | :heavy_minus_sign:                                                                                                    | The TaskRevokeSourceExpired message indicates that the source of the revoke task is due to a grant expiring.          |
| `TaskRevokeSourceNonUsage`                                                                                            | [*TaskRevokeSourceNonUsage](../../models/shared/taskrevokesourcenonusage.md)                                          | :heavy_minus_sign:                                                                                                    | The TaskRevokeSourceNonUsage message indicates that the source of the revoke task is due to the grant not being used. |
| `TaskRevokeSourceRequest`                                                                                             | [*TaskRevokeSourceRequest](../../models/shared/taskrevokesourcerequest.md)                                            | :heavy_minus_sign:                                                                                                    | The TaskRevokeSourceRequest message indicates that the source of the revoke task was a request.                       |
| `TaskRevokeSourceReview`                                                                                              | [*TaskRevokeSourceReview](../../models/shared/taskrevokesourcereview.md)                                              | :heavy_minus_sign:                                                                                                    | The TaskRevokeSourceReview message tracks which access review was the source of the specificed revoke ticket.         |