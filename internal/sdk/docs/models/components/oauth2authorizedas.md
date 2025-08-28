# OAuth2AuthorizedAs

OAuth2AuthorizedAs tracks the user that OAuthed with the connector.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AuthEmail`                                                                   | **string*                                                                     | :heavy_minus_sign:                                                            | authEmail is the email of the user that authorized the connector using OAuth. |
| `AuthorizedAt`                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |