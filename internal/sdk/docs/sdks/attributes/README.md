# Attributes
(*Attributes*)

## Overview

### Available Operations

* [DeleteAttributeValue](#deleteattributevalue) - Delete Attribute Value
* [CreateAttributeValue](#createattributevalue) - Create Attribute Value
* [GetAttributeValue](#getattributevalue) - Get Attribute Value
* [CreateComplianceFrameworkAttributeValue](#createcomplianceframeworkattributevalue) - Create Compliance Framework Attribute Value
* [DeleteComplianceFrameworkAttributeValue](#deletecomplianceframeworkattributevalue) - Delete Compliance Framework Attribute Value
* [GetComplianceFrameworkAttributeValue](#getcomplianceframeworkattributevalue) - Get Compliance Framework Attribute Value
* [CreateRiskLevelAttributeValue](#createrisklevelattributevalue) - Create Risk Level Attribute Value
* [DeleteRiskLevelAttributeValue](#deleterisklevelattributevalue) - Delete Risk Level Attribute Value
* [GetRiskLevelAttributeValue](#getrisklevelattributevalue) - Get Risk Level Attribute Value
* [ListAttributeTypes](#listattributetypes) - List Attribute Types
* [ListAttributeValues](#listattributevalues) - List Attribute Values

## DeleteAttributeValue

Delete an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.DeleteAttributeValue" method="delete" path="/api/v1/attribute/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.DeleteAttributeValue(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `id`                                                                                              | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `deleteAttributeValueRequest`                                                                     | [*components.DeleteAttributeValueRequest](../../models/components/deleteattributevaluerequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.C1APIAttributeV1AttributesDeleteAttributeValueResponse](../../models/operations/c1apiattributev1attributesdeleteattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateAttributeValue

Create a new attribute value.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.CreateAttributeValue" method="post" path="/api/v1/attributes" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.CreateAttributeValue(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [components.CreateAttributeValueRequest](../../models/components/createattributevaluerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.C1APIAttributeV1AttributesCreateAttributeValueResponse](../../models/operations/c1apiattributev1attributescreateattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAttributeValue

Get an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.GetAttributeValue" method="get" path="/api/v1/attributes/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.GetAttributeValue(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAttributeV1AttributesGetAttributeValueResponse](../../models/operations/c1apiattributev1attributesgetattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateComplianceFrameworkAttributeValue

Create a compliance framework value.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.CreateComplianceFrameworkAttributeValue" method="post" path="/api/v1/attributes/compliance_frameworks" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.CreateComplianceFrameworkAttributeValue(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateComplianceFrameworkAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [components.CreateComplianceFrameworkAttributeValueRequest](../../models/components/createcomplianceframeworkattributevaluerequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.C1APIAttributeV1AttributesCreateComplianceFrameworkAttributeValueResponse](../../models/operations/c1apiattributev1attributescreatecomplianceframeworkattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteComplianceFrameworkAttributeValue

Delete an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.DeleteComplianceFrameworkAttributeValue" method="delete" path="/api/v1/attributes/compliance_frameworks/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.DeleteComplianceFrameworkAttributeValue(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteComplianceFrameworkAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                               | Type                                                                                                                                    | Required                                                                                                                                | Description                                                                                                                             |
| --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                   | :heavy_check_mark:                                                                                                                      | The context to use for the request.                                                                                                     |
| `id`                                                                                                                                    | *string*                                                                                                                                | :heavy_check_mark:                                                                                                                      | N/A                                                                                                                                     |
| `deleteComplianceFrameworkAttributeValueRequest`                                                                                        | [*components.DeleteComplianceFrameworkAttributeValueRequest](../../models/components/deletecomplianceframeworkattributevaluerequest.md) | :heavy_minus_sign:                                                                                                                      | N/A                                                                                                                                     |
| `opts`                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                | :heavy_minus_sign:                                                                                                                      | The options for this request.                                                                                                           |

### Response

**[*operations.C1APIAttributeV1AttributesDeleteComplianceFrameworkAttributeValueResponse](../../models/operations/c1apiattributev1attributesdeletecomplianceframeworkattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetComplianceFrameworkAttributeValue

Get an attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.GetComplianceFrameworkAttributeValue" method="get" path="/api/v1/attributes/compliance_frameworks/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.GetComplianceFrameworkAttributeValue(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetComplianceFrameworkAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAttributeV1AttributesGetComplianceFrameworkAttributeValueResponse](../../models/operations/c1apiattributev1attributesgetcomplianceframeworkattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateRiskLevelAttributeValue

Create a risk level attribute.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.CreateRiskLevelAttributeValue" method="post" path="/api/v1/attributes/risk_levels" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.CreateRiskLevelAttributeValue(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateRiskLevelAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [components.CreateRiskLevelAttributeValueRequest](../../models/components/createrisklevelattributevaluerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.C1APIAttributeV1AttributesCreateRiskLevelAttributeValueResponse](../../models/operations/c1apiattributev1attributescreaterisklevelattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteRiskLevelAttributeValue

Delete a risk level attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.DeleteRiskLevelAttributeValue" method="delete" path="/api/v1/attributes/risk_levels/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.DeleteRiskLevelAttributeValue(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteRiskLevelAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                           | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                               | :heavy_check_mark:                                                                                                  | The context to use for the request.                                                                                 |
| `id`                                                                                                                | *string*                                                                                                            | :heavy_check_mark:                                                                                                  | N/A                                                                                                                 |
| `deleteRiskLevelAttributeValueRequest`                                                                              | [*components.DeleteRiskLevelAttributeValueRequest](../../models/components/deleterisklevelattributevaluerequest.md) | :heavy_minus_sign:                                                                                                  | N/A                                                                                                                 |
| `opts`                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                            | :heavy_minus_sign:                                                                                                  | The options for this request.                                                                                       |

### Response

**[*operations.C1APIAttributeV1AttributesDeleteRiskLevelAttributeValueResponse](../../models/operations/c1apiattributev1attributesdeleterisklevelattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetRiskLevelAttributeValue

Get a risk level attribute value by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.GetRiskLevelAttributeValue" method="get" path="/api/v1/attributes/risk_levels/{id}" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.GetRiskLevelAttributeValue(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRiskLevelAttributeValueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAttributeV1AttributesGetRiskLevelAttributeValueResponse](../../models/operations/c1apiattributev1attributesgetrisklevelattributevalueresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAttributeTypes

List all attribute types.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.ListAttributeTypes" method="get" path="/api/v1/attributes/types" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.ListAttributeTypes(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAttributeTypesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAttributeV1AttributesListAttributeTypesResponse](../../models/operations/c1apiattributev1attributeslistattributetypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAttributeValues

List all attribute values for a given attribute type.

### Example Usage

<!-- UsageSnippet language="go" operationID="c1.api.attribute.v1.Attributes.ListAttributeValues" method="get" path="/api/v1/attributes/types/{attribute_type_id}/values" -->
```go
package main

import(
	"context"
	"undefined"
	"undefined/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := undefined.New(
        undefined.WithSecurity(components.Security{
            BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
            Oauth: "<YOUR_OAUTH_HERE>",
        }),
    )

    res, err := s.Attributes.ListAttributeValues(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAttributeValuesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `attributeTypeID`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `pageSize`                                               | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `pageToken`                                              | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.C1APIAttributeV1AttributesListAttributeValuesResponse](../../models/operations/c1apiattributev1attributeslistattributevaluesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |