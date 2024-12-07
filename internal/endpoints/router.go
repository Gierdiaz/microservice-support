package endpoints

import (
	"github.com/Gierdiaz/config"
	"github.com/Gierdiaz/internal/handlers"
	"github.com/Gierdiaz/internal/integrations"
	"github.com/Gierdiaz/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config *config.Config) *gin.Engine {
	r := gin.Default()

	correiosIntegration := integrations.NewCorreiosIntegration()

	SupportHandler := handlers.NewSupportHandler()
	r.GET("/support", SupportHandler.GetSupport)

	correios := r.Group("/correios")
	{
		correios.POST("/frete/calcular", correiosIntegration.FreteHandler.CalcularFrete)
	}

	logger.Logger.Info("Servidor iniciado na porta 8080")

	return r
}

