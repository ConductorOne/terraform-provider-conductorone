# ProvisionInstance

The ProvisionInstance message.

This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - completed
  - cancelled
  - errored
  - reassignedByError



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Cancelled`                                                                | [*CancelledAction](../../models/shared/cancelledaction.md)                 | :heavy_minus_sign:                                                         | The CancelledAction message.                                               |
| `Completed`                                                                | [*CompletedAction](../../models/shared/completedaction.md)                 | :heavy_minus_sign:                                                         | The CompletedAction message.                                               |
| `Errored`                                                                  | [*ErroredAction](../../models/shared/erroredaction.md)                     | :heavy_minus_sign:                                                         | The ErroredAction message.                                                 |
| `NotificationID`                                                           | **string*                                                                  | :heavy_minus_sign:                                                         | The notificationId field.                                                  |
| `Provision`                                                                | [*Provision](../../models/shared/provision.md)                             | :heavy_minus_sign:                                                         | The Provision message.                                                     |
| `ReassignedByError`                                                        | [*ReassignedByErrorAction](../../models/shared/reassignedbyerroraction.md) | :heavy_minus_sign:                                                         | The ReassignedByErrorAction message.                                       |
| `State`                                                                    | [*ProvisionInstanceState](../../models/shared/provisioninstancestate.md)   | :heavy_minus_sign:                                                         | The state field.                                                           |