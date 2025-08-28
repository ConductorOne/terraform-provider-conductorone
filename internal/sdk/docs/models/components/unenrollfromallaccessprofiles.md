# UnenrollFromAllAccessProfiles

The UnenrollFromAllAccessProfiles message.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `CatalogIds`                                                                          | []*string*                                                                            | :heavy_minus_sign:                                                                    | Optional list of catalog IDs to unenroll from. If empty, unenroll from all catalogs.  |
| `CatalogIdsCel`                                                                       | **string*                                                                             | :heavy_minus_sign:                                                                    | CEL expression to dynamically select catalog IDs. If provided, overrides catalog_ids. |
| `UseSubjectUser`                                                                      | **bool*                                                                               | :heavy_minus_sign:                                                                    | If true, the step will use the subject user of the automation as the subject.         |
| `UserIdsCel`                                                                          | **string*                                                                             | :heavy_minus_sign:                                                                    | The userIdsCel field.                                                                 |
| `UserRefs`                                                                            | [][components.UserRef](../../models/components/userref.md)                            | :heavy_minus_sign:                                                                    | The userRefs field.                                                                   |