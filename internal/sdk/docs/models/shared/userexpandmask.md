# UserExpandMask

The user expand mask is used to indicate which related objects should be expanded in the response.
 The supported paths are 'role_ids', 'manager_ids', 'delegated_user_id', 'directory_ids', and '*'.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Paths`                                           | []*string*                                        | :heavy_minus_sign:                                | An array of paths to be expanded in the response. |