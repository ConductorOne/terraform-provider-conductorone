# AppUserServiceUpdateRequestInput

The AppUserServiceUpdateRequest message contains the app user and the fields to be updated.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AppUser`                                                                         | [*AppUserInput](../../models/shared/appuserinput.md)                              | :heavy_minus_sign:                                                                | Application User that represents an account in the application.                   |
| `AppUserExpandMask`                                                               | [*AppUserExpandMask](../../models/shared/appuserexpandmask.md)                    | :heavy_minus_sign:                                                                | The AppUserExpandMask message contains a list of paths to expand in the response. |
| `UpdateMask`                                                                      | **string*                                                                         | :heavy_minus_sign:                                                                | N/A                                                                               |