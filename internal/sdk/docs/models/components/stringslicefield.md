# StringSliceField

The StringSliceField message.

This message contains a oneof named view. Only a single field of the following list may be set at a time:
  - chipsField


This message contains a oneof named _rules. Only a single field of the following list may be set at a time:
  - rules



## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ChipsField`                                                          | [*components.ChipsField](../../models/components/chipsfield.md)       | :heavy_minus_sign:                                                    | The ChipsField message.                                               |
| `DefaultValues`                                                       | []*string*                                                            | :heavy_minus_sign:                                                    | The defaultValues field.                                              |
| `Placeholder`                                                         | **string*                                                             | :heavy_minus_sign:                                                    | The placeholder field.                                                |
| `RepeatedRules`                                                       | [*components.RepeatedRules](../../models/components/repeatedrules.md) | :heavy_minus_sign:                                                    | RepeatedRules describe the constraints applied to `repeated` values   |