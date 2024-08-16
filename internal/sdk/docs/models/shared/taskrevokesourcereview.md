# TaskRevokeSourceReview

The TaskRevokeSourceReview message tracks which access review was the source of the specificed revoke ticket.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AccessReviewID`                                                           | **string*                                                                  | :heavy_minus_sign:                                                         | The ID of the access review associated with the revoke task.               |
| `CertTicketID`                                                             | **string*                                                                  | :heavy_minus_sign:                                                         | The ID of the certify ticket that was denied and created this revoke task. |