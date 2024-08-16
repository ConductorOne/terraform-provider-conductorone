# AppResourceType

The AppResourceType is referenced by an app entitlement defining its resource types. Commonly things like Group or Role.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AppID`                                                         | **string*                                                       | :heavy_minus_sign:                                              | The ID of the app that is associated with the app resource type |
| `CreatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |
| `DeletedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |
| `DisplayName`                                                   | **string*                                                       | :heavy_minus_sign:                                              | The display name of the app resource type.                      |
| `ID`                                                            | **string*                                                       | :heavy_minus_sign:                                              | The unique ID for the app resource type.                        |
| `UpdatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |