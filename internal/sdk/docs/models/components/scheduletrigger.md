# ScheduleTrigger

The ScheduleTrigger message.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Advanced`                                 | **bool*                                    | :heavy_minus_sign:                         | The advanced field.                        |
| `CronSpec`                                 | **string*                                  | :heavy_minus_sign:                         | The cronSpec field.                        |
| `SkipIfTrueCel`                            | **string*                                  | :heavy_minus_sign:                         | The skipIfTrueCel field.                   |
| `Start`                                    | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |