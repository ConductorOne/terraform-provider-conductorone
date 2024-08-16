# RejectInstance

This policy step indicates that a ticket should have a denied outcome. This is a terminal approval state and is used to explicitly define the end of approval steps.
 The instance is just a marker for it being copied into an active policy.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `RejectMessage`                                                                       | **string*                                                                             | :heavy_minus_sign:                                                                    | An optional message to include in the comments when a task is automatically rejected. |