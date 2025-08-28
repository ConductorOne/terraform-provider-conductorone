# EncryptedData

EncryptedData is a message that contains encrypted bytes and metadata.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Description`                                         | **string*                                             | :heavy_minus_sign:                                    | The human-readable description of the encrypted data. |
| `EncryptedBytes`                                      | **string*                                             | :heavy_minus_sign:                                    | The encrypted bytes.                                  |
| `KeyID`                                               | **string*                                             | :heavy_minus_sign:                                    | The key ID used to encrypt the data.                  |
| `Name`                                                | **string*                                             | :heavy_minus_sign:                                    | The human-readable name of the encrypted data.        |
| `Provider`                                            | **string*                                             | :heavy_minus_sign:                                    | The encryption provider used to encrypt the data.     |
| `Schema`                                              | **string*                                             | :heavy_minus_sign:                                    | The (optional) JSON schema of the encrypted data.     |