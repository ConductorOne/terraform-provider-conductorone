# AppAccessRequestDefaults1

The AppAccessRequestDefaults message.

This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
  - durationUnset
  - durationGrant



## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `CatalogIds`                                                                             | []*string*                                                                               | :heavy_minus_sign:                                                                       | The request catalog ids for the app access request rule.                                 |
| `DefaultsEnabled`                                                                        | **bool*                                                                                  | :heavy_minus_sign:                                                                       | If true the app level request configuration will be applied to specified resource types. |
| `DurationGrant`                                                                          | **string*                                                                                | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `DurationUnset`                                                                          | [*components.DurationUnset](../../models/components/durationunset.md)                    | :heavy_minus_sign:                                                                       | N/A                                                                                      |
| `EmergencyGrantEnabled`                                                                  | **bool*                                                                                  | :heavy_minus_sign:                                                                       | If emergency grants are enabled for this app access request rule.                        |
| `EmergencyGrantPolicyID`                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | The policy id for the emergency grant policy.                                            |
| `RequestPolicyID`                                                                        | **string*                                                                                | :heavy_minus_sign:                                                                       | The requestPolicyId field.                                                               |
| `ResourceTypeIds`                                                                        | []*string*                                                                               | :heavy_minus_sign:                                                                       | The app resource type ids for which the app access request defaults are applied.         |
| `State`                                                                                  | [*components.State](../../models/components/state.md)                                    | :heavy_minus_sign:                                                                       | The last applied state of the app access request defaults.                               |