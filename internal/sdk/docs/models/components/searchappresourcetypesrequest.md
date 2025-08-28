# SearchAppResourceTypesRequest

Search for app resources based on some filters.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AppIds`                                                      | []*string*                                                    | :heavy_minus_sign:                                            | A list of app IDs to restrict the search by.                  |
| `AppUserIds`                                                  | []*string*                                                    | :heavy_minus_sign:                                            | A list of app user IDs to restrict the search by.             |
| `DisplayName`                                                 | **string*                                                     | :heavy_minus_sign:                                            | Exact match on display name                                   |
| `ExcludeResourceTypeIds`                                      | []*string*                                                    | :heavy_minus_sign:                                            | A list of resource type IDs to exclude from the search.       |
| `ExcludeResourceTypeTraitIds`                                 | []*string*                                                    | :heavy_minus_sign:                                            | A list of resource type trait IDs to exclude from the search. |
| `PageSize`                                                    | **int*                                                        | :heavy_minus_sign:                                            | The pageSize where 10 <= pageSize <= 100, default 25.         |
| `PageToken`                                                   | **string*                                                     | :heavy_minus_sign:                                            | The pageToken field.                                          |
| `Query`                                                       | **string*                                                     | :heavy_minus_sign:                                            | Fuzzy search the display name of resource types.              |
| `ResourceTypeIds`                                             | []*string*                                                    | :heavy_minus_sign:                                            | A list of resource type IDs to restrict the search by.        |
| `ResourceTypeTraitIds`                                        | []*string*                                                    | :heavy_minus_sign:                                            | A list of resource type trait IDs to restrict the search by.  |