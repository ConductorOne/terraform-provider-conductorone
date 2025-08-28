# DirectoryServiceCreateRequest

Uplevel an app into a full directory.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AppID`                                                                           | **string*                                                                         | :heavy_minus_sign:                                                                | The AppID to make into a directory, providing identities and more for the C1 app. |
| `DirectoryExpandMask`                                                             | [*components.DirectoryExpandMask](../../models/components/directoryexpandmask.md) | :heavy_minus_sign:                                                                | The fields to be included in the directory response.                              |