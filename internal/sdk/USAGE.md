<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
)

func main() {
    s := conductoronesdkterraform.New(
        conductoronesdkterraform.WithSecurity(shared.Security{
            BearerAuth: "",
            Oauth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AWSExternalIDSettings.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAWSExternalIDResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->