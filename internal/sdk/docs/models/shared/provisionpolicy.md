# ProvisionPolicy

The ProvisionPolicy message.

This message contains a oneof named typ. Only a single field of the following list may be set at a time:
  - connector
  - manual
  - delegated



## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Connector`                                                      | [*ConnectorProvision](../../models/shared/connectorprovision.md) | :heavy_minus_sign:                                               | The ConnectorProvision message.                                  |
| `Delegated`                                                      | [*DelegatedProvision](../../models/shared/delegatedprovision.md) | :heavy_minus_sign:                                               | The DelegatedProvision message.                                  |
| `Manual`                                                         | [*ManualProvision](../../models/shared/manualprovision.md)       | :heavy_minus_sign:                                               | The ManualProvision message.                                     |