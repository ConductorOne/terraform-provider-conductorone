# CancelledAction

The outcome of a provision instance that is cancelled.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `CancelledAt`                                                       | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | N/A                                                                 |
| `CancelledByUserID`                                                 | **string*                                                           | :heavy_minus_sign:                                                  | The userID, usually the system, that cancells a provision instance. |