# ConnectorServiceUpdateRequestInput

The ConnectorServiceUpdateRequest message contains the fields required to update a connector.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Connector`                                                               | [*ConnectorInput](../../models/shared/connectorinput.md)                  | :heavy_minus_sign:                                                        | A Connector is used to sync objects into Apps                             |
| `ConnectorExpandMask`                                                     | [*ConnectorExpandMask](../../models/shared/connectorexpandmask.md)        | :heavy_minus_sign:                                                        | The ConnectorExpandMask is used to expand related objects on a connector. |
| `UpdateMask`                                                              | **string*                                                                 | :heavy_minus_sign:                                                        | N/A                                                                       |