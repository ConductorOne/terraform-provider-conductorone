# MessageRules

MessageRules describe the constraints applied to embedded message values.
 For message-type fields, validation is performed recursively.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Required`                                                                      | **bool*                                                                         | :heavy_minus_sign:                                                              | Required specifies that this field must be set                                  |
| `Skip`                                                                          | **bool*                                                                         | :heavy_minus_sign:                                                              | Skip specifies that the validation rules of this field should not be<br/> evaluated |