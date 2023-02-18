package operations

import (
	"github.com/speakeasy-sdks/petstore-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/petstore-go-sdk/pkg/models/utils"
)

type FindPetByIDPathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindPetByIDRequest struct {
	PathParams FindPetByIDPathParams
	Retries    *utils.RetryConfig
}

type FindPetByIDResponse struct {
	ContentType string
	Error       *shared.Error
	Pet         *shared.Pet
	StatusCode  int
}
