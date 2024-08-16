# Wait

Define a Wait step for a policy to wait on a condition to be met.

This message contains a oneof named until. Only a single field of the following list may be set at a time:
  - condition



## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `WaitCondition`                                            | [*WaitCondition](../../models/shared/waitcondition.md)     | :heavy_minus_sign:                                         | The WaitCondition message.                                 |
| `CommentOnFirstWait`                                       | **string*                                                  | :heavy_minus_sign:                                         | The comment to post on first failed check.                 |
| `CommentOnTimeout`                                         | **string*                                                  | :heavy_minus_sign:                                         | The comment to post if we timeout.                         |
| `Name`                                                     | **string*                                                  | :heavy_minus_sign:                                         | The name of our condition to show on the task details page |
| `TimeoutDuration`                                          | **string*                                                  | :heavy_minus_sign:                                         | N/A                                                        |