<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/petstore-go-sdk"
    "github.com/speakeasy-sdks/petstore-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/petstore-go-sdk/pkg/models/operations"
)

func main() {
    s := petstore.New()
    
    req := operations.AddPetRequest{
        Request: shared.NewPet{
            Name: "unde",
            Tag: "deserunt",
        },
    }
    
    res, err := s.AddPet(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->