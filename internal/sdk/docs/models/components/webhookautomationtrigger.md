# WebhookAutomationTrigger

The WebhookAutomationTrigger message.

This message contains a oneof named auth_config. Only a single field of the following list may be set at a time:
  - jwt
  - hmac



## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `WebhookListenerAuthHMAC`                                                                 | [*components.WebhookListenerAuthHMAC](../../models/components/webhooklistenerauthhmac.md) | :heavy_minus_sign:                                                                        | The WebhookListenerAuthHMAC message.                                                      |
| `WebhookListenerAuthJWT`                                                                  | [*components.WebhookListenerAuthJWT](../../models/components/webhooklistenerauthjwt.md)   | :heavy_minus_sign:                                                                        | The WebhookListenerAuthJWT message.                                                       |
| `ListenerID`                                                                              | **string*                                                                                 | :heavy_minus_sign:                                                                        | Optional existing listener ID (hidden field from frontend)                                |