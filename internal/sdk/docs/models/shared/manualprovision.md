# ManualProvision

Manual provisioning indicates that a human must intervene for the provisioning of this step.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `Instructions`                                                                    | **string*                                                                         | :heavy_minus_sign:                                                                | This field indicates a text body of instructions for the provisioner to indicate. |
| `UserIds`                                                                         | []*string*                                                                        | :heavy_minus_sign:                                                                | An array of users that are required to provision during this step.                |