# ExternalTicketProvision

This provision step indicates that we should check an external ticket to provision this entitlement


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AppID`                                                                           | **string*                                                                         | :heavy_minus_sign:                                                                | The appId field.                                                                  |
| `ConnectorID`                                                                     | **string*                                                                         | :heavy_minus_sign:                                                                | The connectorId field.                                                            |
| `ExternalTicketProvisionerConfigID`                                               | **string*                                                                         | :heavy_minus_sign:                                                                | The externalTicketProvisionerConfigId field.                                      |
| `Instructions`                                                                    | **string*                                                                         | :heavy_minus_sign:                                                                | This field indicates a text body of instructions for the provisioner to indicate. |