# FieldRelationship

FieldRelationships can be used during form validation, or they can represent
 information that is necessary to when it comes to visually rendering the form

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - requiredTogether
  - atLeastOne
  - mutuallyExclusive



## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AtLeastOne`                                                                  | [*components.AtLeastOne](../../models/components/atleastone.md)               | :heavy_minus_sign:                                                            | The AtLeastOne message.                                                       |
| `FieldNames`                                                                  | []*string*                                                                    | :heavy_minus_sign:                                                            | The names of the fields that share this relationship                          |
| `MutuallyExclusive`                                                           | [*components.MutuallyExclusive](../../models/components/mutuallyexclusive.md) | :heavy_minus_sign:                                                            | The MutuallyExclusive message.                                                |
| `RequiredTogether`                                                            | [*components.RequiredTogether](../../models/components/requiredtogether.md)   | :heavy_minus_sign:                                                            | The RequiredTogether message.                                                 |