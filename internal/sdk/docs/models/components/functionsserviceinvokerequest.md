# FunctionsServiceInvokeRequest

The FunctionsServiceInvokeRequest message.

This message contains a oneof named arg. Only a single field of the following list may be set at a time:
  - json



## Fields

| Field                                                                                                                                                  | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `CommitID`                                                                                                                                             | **string*                                                                                                                                              | :heavy_minus_sign:                                                                                                                                     | The commitId field.                                                                                                                                    |
| `JSON`                                                                                                                                                 | **string*                                                                                                                                              | :heavy_minus_sign:                                                                                                                                     | The json field.<br/>This field is part of the `arg` oneof.<br/>See the documentation for `c1.api.functions.v1.FunctionsServiceInvokeRequest` for more details. |