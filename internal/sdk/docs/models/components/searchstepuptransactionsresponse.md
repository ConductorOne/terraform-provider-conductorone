# SearchStepUpTransactionsResponse

Response message for searching step-up transactions


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `List`                                                                         | [][components.StepUpTransaction](../../models/components/stepuptransaction.md) | :heavy_minus_sign:                                                             | List of transactions matching the search criteria                              |
| `NextPageToken`                                                                | **string*                                                                      | :heavy_minus_sign:                                                             | Token for retrieving the next page of results                                  |