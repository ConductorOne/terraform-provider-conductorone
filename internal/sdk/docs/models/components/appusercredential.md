# AppUserCredential

A credentials for the Application User that represents an account in the application.

This message contains a oneof named credential. Only a single field of the following list may be set at a time:
  - encryptedData



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AppID`                                                                | **string*                                                              | :heavy_minus_sign:                                                     | The ID of the application.                                             |
| `AppUserID`                                                            | **string*                                                              | :heavy_minus_sign:                                                     | A unique identifier of the application user.                           |
| `CreatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `DeletedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `EncryptedData`                                                        | [*components.EncryptedData](../../models/components/encrypteddata.md)  | :heavy_minus_sign:                                                     | EncryptedData is a message that contains encrypted bytes and metadata. |
| `ExpiresAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |
| `ID`                                                                   | **string*                                                              | :heavy_minus_sign:                                                     | A unique identifier of the credential.                                 |
| `UpdatedAt`                                                            | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | N/A                                                                    |