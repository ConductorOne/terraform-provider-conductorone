# TaskTypeInput

Task Type provides configuration for the type of task: certify, grant, or revoke

This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
  - grant
  - revoke
  - certify
  - offboarding



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `TaskTypeCertify`                                                                            | [*components.TaskTypeCertifyInput](../../models/components/tasktypecertifyinput.md)          | :heavy_minus_sign:                                                                           | The TaskTypeCertify message indicates that a task is a certify task and all related details. |
| `TaskTypeGrant`                                                                              | [*components.TaskTypeGrantInput](../../models/components/tasktypegrantinput.md)              | :heavy_minus_sign:                                                                           | The TaskTypeGrant message indicates that a task is a grant task and all related details.     |
| `TaskTypeOffboarding`                                                                        | [*components.TaskTypeOffboardingInput](../../models/components/tasktypeoffboardinginput.md)  | :heavy_minus_sign:                                                                           | The TaskTypeOffboarding message.                                                             |
| `TaskTypeRevoke`                                                                             | [*components.TaskTypeRevokeInput](../../models/components/tasktyperevokeinput.md)            | :heavy_minus_sign:                                                                           | The TaskTypeRevoke message indicates that a task is a revoke task and all related details.   |