# Exporter

The Exporter message.

This message contains a oneof named export_to. Only a single field of the following list may be set at a time:
  - datasource



## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `CreatedAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | N/A                                                                             |
| `ExportToDatasource`                                                            | [*components.ExportToDatasource](../../models/components/exporttodatasource.md) | :heavy_minus_sign:                                                              | The ExportToDatasource message.                                                 |
| `DeletedAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | N/A                                                                             |
| `DisplayName`                                                                   | **string*                                                                       | :heavy_minus_sign:                                                              | The displayName field.                                                          |
| `ExportID`                                                                      | **string*                                                                       | :heavy_minus_sign:                                                              | The exportId field.                                                             |
| `State`                                                                         | [*components.ExporterState](../../models/components/exporterstate.md)           | :heavy_minus_sign:                                                              | The state field.                                                                |
| `UpdatedAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | N/A                                                                             |
| `WatermarkEventID`                                                              | **string*                                                                       | :heavy_minus_sign:                                                              | we've synchorized this far                                                      |