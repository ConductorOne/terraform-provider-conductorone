# WebhooksServiceUpdateRequest

The WebhooksServiceUpdateRequest message contains the webhook object to update and a field mask to indicate which fields to update. It uses URL value for input.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `UpdateMask`                                                        | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Webhook`                                                           | [*components.WebhookInput](../../models/components/webhookinput.md) | :heavy_minus_sign:                                                  | The Webhook message.                                                |