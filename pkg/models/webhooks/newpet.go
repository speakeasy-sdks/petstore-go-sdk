package webhooks

import (
	"github.com/speakeasy-sdks/petstore-go-sdk/pkg/models/shared"
)

type NewPetResponse struct {
	ContentType string
	StatusCode  int
}

type NewPetRequest struct {
	Request *shared.NewPet `request:"mediaType=application/json"`
}
