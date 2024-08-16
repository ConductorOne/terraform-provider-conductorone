# ReassignedAction

The ReassignedAction object describes the outcome of a policy step that has been reassigned.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `NewPolicyStepID`                                                            | **string*                                                                    | :heavy_minus_sign:                                                           | The ID of the policy step that was created as a result of this reassignment. |
| `ReassignedAt`                                                               | [*time.Time](https://pkg.go.dev/time#Time)                                   | :heavy_minus_sign:                                                           | N/A                                                                          |
| `UserID`                                                                     | **string*                                                                    | :heavy_minus_sign:                                                           | The UserID of the person who reassigned this step.                           |