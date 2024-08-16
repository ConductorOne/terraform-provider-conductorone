# RoleInput

Role is a role that can be assigned to a user in ConductorOne.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `DisplayName`                                | **string*                                    | :heavy_minus_sign:                           | The display name of the role.                |
| `Permissions`                                | []*string*                                   | :heavy_minus_sign:                           | The list of permissions this role has.       |
| `ServiceRoles`                               | []*string*                                   | :heavy_minus_sign:                           | The list of serviceRoles that this role has. |