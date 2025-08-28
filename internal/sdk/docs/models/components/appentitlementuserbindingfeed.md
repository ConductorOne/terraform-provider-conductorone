# AppEntitlementUserBindingFeed

The AppEntitlementUserBindingFeed message.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `AppEntitlementID`                                            | **string*                                                     | :heavy_minus_sign:                                            | The ID of the app entitlement that the app user has access to |
| `AppID`                                                       | **string*                                                     | :heavy_minus_sign:                                            | The ID of the app associated with the app entitlement         |
| `AppUserID`                                                   | **string*                                                     | :heavy_minus_sign:                                            | The ID of the app user that has access to the app entitlement |
| `Date`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | N/A                                                           |
| `EventType`                                                   | [*components.EventType](../../models/components/eventtype.md) | :heavy_minus_sign:                                            | The eventType field.                                          |
| `TicketID`                                                    | **string*                                                     | :heavy_minus_sign:                                            | The ticketId field.                                           |