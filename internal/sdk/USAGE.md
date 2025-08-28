<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"undefined"
	"undefined/models/components"
)

func main() {
	ctx := context.Background()

	s := undefined.New(
		undefined.WithSecurity(components.Security{
			BearerAuth: "<YOUR_BEARER_TOKEN_HERE>",
			Oauth:      "<YOUR_OAUTH_HERE>",
		}),
	)

	res, err := s.AccessConflict.CreateMonitor(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ConflictMonitor != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->