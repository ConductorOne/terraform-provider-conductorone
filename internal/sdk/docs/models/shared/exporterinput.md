# ExporterInput

The Exporter message.

This message contains a oneof named export_to. Only a single field of the following list may be set at a time:
  - datasource



## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ExportToDatasource`                                             | [*ExportToDatasource](../../models/shared/exporttodatasource.md) | :heavy_minus_sign:                                               | The ExportToDatasource message.                                  |
| `DisplayName`                                                    | **string*                                                        | :heavy_minus_sign:                                               | The displayName field.                                           |