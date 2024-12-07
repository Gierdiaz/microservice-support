package integrations

import (
	"github.com/Gierdiaz/internal/handlers"
	"github.com/Gierdiaz/internal/services"
)

type CorreiosIntegration struct {
	FreteHandler *handlers.FreteHandler
}

func NewCorreiosIntegration() *CorreiosIntegration {
	correiosService := services.NewCorreiosService("https://api.correios.com.br")
	freteHandler := handlers.NewFreteHandler(*correiosService)

	return &CorreiosIntegration{
		FreteHandler: freteHandler,
	}
}
