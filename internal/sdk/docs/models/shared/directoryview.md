# DirectoryView

The directory view contains a directory and an app_path which is a JSONPATH set to the location in the expand mask that the expanded app will live if requested by the expander.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Directory`                                                                  | [*Directory](../../models/shared/directory.md)                               | :heavy_minus_sign:                                                           | This object indicates that an app is also a directory.                       |
| `AppPath`                                                                    | **string*                                                                    | :heavy_minus_sign:                                                           | JSONPATH expression indicating the location of the App object in the  array. |