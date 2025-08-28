# SendEmail

The SendEmail message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Body`                                                                        | **string*                                                                     | :heavy_minus_sign:                                                            | The body field.                                                               |
| `Subject`                                                                     | **string*                                                                     | :heavy_minus_sign:                                                            | The subject field.                                                            |
| `Title`                                                                       | **string*                                                                     | :heavy_minus_sign:                                                            | The title field.                                                              |
| `UseSubjectUser`                                                              | **bool*                                                                       | :heavy_minus_sign:                                                            | If true, the step will use the subject user of the automation as the subject. |
| `UserIdsCel`                                                                  | **string*                                                                     | :heavy_minus_sign:                                                            | The userIdsCel field.                                                         |
| `UserRefs`                                                                    | [][components.UserRef](../../models/components/userref.md)                    | :heavy_minus_sign:                                                            | The userRefs field.                                                           |