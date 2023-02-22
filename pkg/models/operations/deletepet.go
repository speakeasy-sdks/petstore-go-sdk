package operations

import (
	"github.com/speakeasy-sdks/petstore-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/petstore-go-sdk/pkg/utils"
)

type DeletePetPathParams struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePetRequest struct {
	PathParams DeletePetPathParams
	Retries    *utils.RetryConfig
}

type DeletePetResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int
}
