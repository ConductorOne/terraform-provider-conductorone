# Approval

The Approval message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - users
  - manager
  - appOwners
  - group
  - self
  - entitlementOwners



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AllowReassignment`                                                          | **bool*                                                                      | :heavy_minus_sign:                                                           | The allowReassignment field.                                                 |
| `AppOwners`                                                                  | [*AppOwnerApproval](../../models/shared/appownerapproval.md)                 | :heavy_minus_sign:                                                           | The AppOwnerApproval message.                                                |
| `Assigned`                                                                   | **bool*                                                                      | :heavy_minus_sign:                                                           | The assigned field.                                                          |
| `EntitlementOwners`                                                          | [*EntitlementOwnerApproval](../../models/shared/entitlementownerapproval.md) | :heavy_minus_sign:                                                           | The EntitlementOwnerApproval message.                                        |
| `Group`                                                                      | [*AppGroupApproval](../../models/shared/appgroupapproval.md)                 | :heavy_minus_sign:                                                           | The AppGroupApproval message.                                                |
| `Manager`                                                                    | [*ManagerApproval](../../models/shared/managerapproval.md)                   | :heavy_minus_sign:                                                           | The ManagerApproval message.                                                 |
| `RequireApprovalReason`                                                      | **bool*                                                                      | :heavy_minus_sign:                                                           | The requireApprovalReason field.                                             |
| `RequireReassignmentReason`                                                  | **bool*                                                                      | :heavy_minus_sign:                                                           | The requireReassignmentReason field.                                         |
| `Self`                                                                       | [*SelfApproval](../../models/shared/selfapproval.md)                         | :heavy_minus_sign:                                                           | The SelfApproval message.                                                    |
| `Users`                                                                      | [*UserApproval](../../models/shared/userapproval.md)                         | :heavy_minus_sign:                                                           | The UserApproval message.                                                    |