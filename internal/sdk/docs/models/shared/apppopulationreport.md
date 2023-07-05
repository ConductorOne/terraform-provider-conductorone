# AppPopulationReport

The AppPopulationReport message.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AppID`                                                                      | **string*                                                                    | :heavy_minus_sign:                                                           | The appId field.                                                             |
| `CreatedAt`                                                                  | [*time.Time](https://pkg.go.dev/time#Time)                                   | :heavy_minus_sign:                                                           | N/A                                                                          |
| `DownloadURL`                                                                | **string*                                                                    | :heavy_minus_sign:                                                           | The downloadUrl field.                                                       |
| `Hashes`                                                                     | map[string]*string*                                                          | :heavy_minus_sign:                                                           | The hashes field.                                                            |
| `ID`                                                                         | **string*                                                                    | :heavy_minus_sign:                                                           | The id field.                                                                |
| `State`                                                                      | [*AppPopulationReportState](../../models/shared/apppopulationreportstate.md) | :heavy_minus_sign:                                                           | The state field.                                                             |