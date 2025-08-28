# UserApproval

The user approval object describes the approval configuration of a policy step that needs to be approved by a specific list of users.


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `AllowSelfApproval`                                                                                 | **bool*                                                                                             | :heavy_minus_sign:                                                                                  | Configuration to allow self approval of if the user is specified and also the target of the ticket. |
| `UserIds`                                                                                           | []*string*                                                                                          | :heavy_minus_sign:                                                                                  | Array of users configured for approval.                                                             |