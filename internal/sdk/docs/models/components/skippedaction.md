# SkippedAction

The SkippedAction object describes the outcome of a policy step that has been skipped.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `NewPolicyStepID`                                                        | **string*                                                                | :heavy_minus_sign:                                                       | The ID of the policy step that was created as a result of this skipping. |
| `SkippedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `UserID`                                                                 | **string*                                                                | :heavy_minus_sign:                                                       | The UserID of the user who skipped this step.                            |