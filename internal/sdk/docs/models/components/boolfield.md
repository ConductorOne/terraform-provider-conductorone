# BoolField

The BoolField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - checkboxField


This message contains a oneof named _rules. Only a single field of the following list may be set at a time:
  - rules



## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `CheckboxField`                                                       | [*components.CheckboxField](../../models/components/checkboxfield.md) | :heavy_minus_sign:                                                    | The CheckboxField message.                                            |
| `DefaultValue`                                                        | **bool*                                                               | :heavy_minus_sign:                                                    | The defaultValue field.                                               |
| `BoolRules`                                                           | [*components.BoolRules](../../models/components/boolrules.md)         | :heavy_minus_sign:                                                    | BoolRules describes the constraints applied to `bool` values          |