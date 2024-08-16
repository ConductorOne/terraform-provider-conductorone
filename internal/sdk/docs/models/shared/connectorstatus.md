# ConnectorStatus

The status field on the connector is used to track the status of the connectors sync, and when syncing last started, completed, or caused the connector to update.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CompletedAt`                                                          | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `LastError`                                                            | **string*                                                              | :heavy_minus_sign:                                                     | The last error encountered by the connector.                           |
| `StartedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Status`                                                               | [*ConnectorStatusStatus](../../models/shared/connectorstatusstatus.md) | :heavy_minus_sign:                                                     | The status of the connector sync.                                      |
| `UpdatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |