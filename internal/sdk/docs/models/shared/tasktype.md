# TaskType

Task Type provides configuration for the type of task: certify, grant, or revoke

This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
  - grant
  - revoke
  - certify
  - offboarding



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `TaskTypeCertify`                                                                            | [*TaskTypeCertify1](../../models/shared/tasktypecertify1.md)                                 | :heavy_minus_sign:                                                                           | The TaskTypeCertify message indicates that a task is a certify task and all related details. |
| `TaskTypeGrant`                                                                              | [*TaskTypeGrant](../../models/shared/tasktypegrant.md)                                       | :heavy_minus_sign:                                                                           | The TaskTypeGrant message indicates that a task is a grant task and all related details.     |
| `TaskTypeOffboarding`                                                                        | [*TaskTypeOffboarding1](../../models/shared/tasktypeoffboarding1.md)                         | :heavy_minus_sign:                                                                           | The TaskTypeOffboarding message.                                                             |
| `TaskTypeRevoke`                                                                             | [*TaskTypeRevoke](../../models/shared/tasktyperevoke.md)                                     | :heavy_minus_sign:                                                                           | The TaskTypeRevoke message indicates that a task is a revoke task and all related details.   |