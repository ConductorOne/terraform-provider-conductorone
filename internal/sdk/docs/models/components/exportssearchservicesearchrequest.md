# ExportsSearchServiceSearchRequest

The ExportsSearchServiceSearchRequest message.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `DisplayName`                                                                      | **string*                                                                          | :heavy_minus_sign:                                                                 | Search for system log exporters with a case insensitive match on the display name. |
| `PageSize`                                                                         | **int*                                                                             | :heavy_minus_sign:                                                                 | The pageSize field.                                                                |
| `PageToken`                                                                        | **string*                                                                          | :heavy_minus_sign:                                                                 | The pageToken field.                                                               |
| `Query`                                                                            | **string*                                                                          | :heavy_minus_sign:                                                                 | The query field.                                                                   |
| `Refs`                                                                             | [][components.ExporterRef](../../models/components/exporterref.md)                 | :heavy_minus_sign:                                                                 | The refs field.                                                                    |