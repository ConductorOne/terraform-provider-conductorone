# ConnectorServiceRotateCredentialResponse

ConnectorServiceRotateCredentialResponse is the response returned by the rotate method.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ClientSecret`                                                                    | **string*                                                                         | :heavy_minus_sign:                                                                | The new clientSecret returned after rotating the connector credential.            |
| `ConnectorCredential`                                                             | [*components.ConnectorCredential](../../models/components/connectorcredential.md) | :heavy_minus_sign:                                                                | ConnectorCredential is used by a connector to authenticate with conductor one.    |