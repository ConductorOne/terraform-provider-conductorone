# WaitInstance

Used by the policy engine to describe an instantiated wait step.

This message contains a oneof named until. Only a single field of the following list may be set at a time:
  - condition


This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
  - succeeded
  - timedOut
  - skipped



## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `CommentOnFirstWait`                                                                   | **string*                                                                              | :heavy_minus_sign:                                                                     | The comment to post on first failed check.                                             |
| `CommentOnTimeout`                                                                     | **string*                                                                              | :heavy_minus_sign:                                                                     | The comment to post if we timeout.                                                     |
| `WaitConditionInstance`                                                                | [*components.WaitConditionInstance](../../models/components/waitconditioninstance.md)  | :heavy_minus_sign:                                                                     | Used by the policy engine to describe an instantiated condition to wait on.            |
| `Name`                                                                                 | **string*                                                                              | :heavy_minus_sign:                                                                     | The name field.                                                                        |
| `SkippedAction`                                                                        | [*components.SkippedAction](../../models/components/skippedaction.md)                  | :heavy_minus_sign:                                                                     | The SkippedAction object describes the outcome of a policy step that has been skipped. |
| `StartedWaitingAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `State`                                                                                | [*components.WaitInstanceState](../../models/components/waitinstancestate.md)          | :heavy_minus_sign:                                                                     | The state field.                                                                       |
| `ConditionSucceeded`                                                                   | [*components.ConditionSucceeded](../../models/components/conditionsucceeded.md)        | :heavy_minus_sign:                                                                     | The ConditionSucceeded message.                                                        |
| `ConditionTimedOut`                                                                    | [*components.ConditionTimedOut](../../models/components/conditiontimedout.md)          | :heavy_minus_sign:                                                                     | The ConditionTimedOut message.                                                         |
| `Timeout`                                                                              | [*time.Time](https://pkg.go.dev/time#Time)                                             | :heavy_minus_sign:                                                                     | N/A                                                                                    |
| `TimeoutDuration`                                                                      | **string*                                                                              | :heavy_minus_sign:                                                                     | N/A                                                                                    |