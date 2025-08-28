# AccountProvision

The AccountProvision message.

This message contains a oneof named storage_type. Only a single field of the following list may be set at a time:
  - saveToVault
  - doNotSave



## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Config`                                                          | map[string]*any*                                                  | :heavy_minus_sign:                                                | N/A                                                               |
| `ConnectorID`                                                     | **string*                                                         | :heavy_minus_sign:                                                | The connectorId field.                                            |
| `DoNotSave`                                                       | [*components.DoNotSave](../../models/components/donotsave.md)     | :heavy_minus_sign:                                                | The DoNotSave message.                                            |
| `SaveToVault`                                                     | [*components.SaveToVault](../../models/components/savetovault.md) | :heavy_minus_sign:                                                | The SaveToVault message.                                          |
| `SchemaID`                                                        | **string*                                                         | :heavy_minus_sign:                                                | The schemaId field.                                               |