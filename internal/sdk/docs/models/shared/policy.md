# Policy

The Policy message.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `CreatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |
| `DeletedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |
| `Description`                                                   | **string*                                                       | :heavy_minus_sign:                                              | The description field.                                          |
| `DisplayName`                                                   | **string*                                                       | :heavy_minus_sign:                                              | The displayName field.                                          |
| `ID`                                                            | **string*                                                       | :heavy_minus_sign:                                              |  Properties<br/>                                                |
| `PolicySteps`                                                   | map[string][PolicySteps](../../models/shared/policysteps.md)    | :heavy_minus_sign:                                              | The policySteps field.                                          |
| `PolicyType`                                                    | [*PolicyPolicyType](../../models/shared/policypolicytype.md)    | :heavy_minus_sign:                                              | The policyType field.                                           |
| `PostActions`                                                   | [][PolicyPostActions](../../models/shared/policypostactions.md) | :heavy_minus_sign:                                              | The postActions field.                                          |
| `ReassignTasksToDelegates`                                      | **bool*                                                         | :heavy_minus_sign:                                              | The reassignTasksToDelegates field.                             |
| `SystemBuiltin`                                                 | **bool*                                                         | :heavy_minus_sign:                                              | The systemBuiltin field.                                        |
| `UpdatedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                      | :heavy_minus_sign:                                              | N/A                                                             |