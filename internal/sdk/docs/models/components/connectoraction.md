# ConnectorAction

The ConnectorAction message.

This message contains a oneof named connector_identifier. Only a single field of the following list may be set at a time:
  - connectorRef



## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ActionName`                                                        | **string*                                                           | :heavy_minus_sign:                                                  | The actionName field.                                               |
| `ArgsTemplate`                                                      | map[string]*any*                                                    | :heavy_minus_sign:                                                  | N/A                                                                 |
| `ConnectorRef`                                                      | [*components.ConnectorRef](../../models/components/connectorref.md) | :heavy_minus_sign:                                                  | The ConnectorRef message.                                           |