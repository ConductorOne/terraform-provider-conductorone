# TaskAuditConnectorActionResult

The TaskAuditConnectorActionResult message.

This message contains a oneof named result. Only a single field of the following list may be set at a time:
  - success
  - error
  - cancelled



## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `AppEntitlementID`                                                                          | **string*                                                                                   | :heavy_minus_sign:                                                                          | The appEntitlementId field.                                                                 |
| `AppID`                                                                                     | **string*                                                                                   | :heavy_minus_sign:                                                                          | The appId field.                                                                            |
| `TaskAuditCancelledResult`                                                                  | [*components.TaskAuditCancelledResult](../../models/components/taskauditcancelledresult.md) | :heavy_minus_sign:                                                                          | The TaskAuditCancelledResult message.                                                       |
| `ConnectorActionID`                                                                         | **string*                                                                                   | :heavy_minus_sign:                                                                          | The connectorActionId field.                                                                |
| `ConnectorID`                                                                               | **string*                                                                                   | :heavy_minus_sign:                                                                          | The connectorId field.                                                                      |
| `TaskAuditErrorResult`                                                                      | [*components.TaskAuditErrorResult](../../models/components/taskauditerrorresult.md)         | :heavy_minus_sign:                                                                          | The TaskAuditErrorResult message.                                                           |
| `TaskAuditSuccessResult`                                                                    | [*components.TaskAuditSuccessResult](../../models/components/taskauditsuccessresult.md)     | :heavy_minus_sign:                                                                          | The TaskAuditSuccessResult message.                                                         |