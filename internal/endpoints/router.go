package endpoints

import (
	"github.com/Gierdiaz/config"
	"github.com/Gierdiaz/internal/handlers"
	"github.com/Gierdiaz/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config *config.Config) *gin.Engine {
	r := gin.Default()

	SupportHandler := handlers.NewSupportHandler()
	r.GET("/support", SupportHandler.GetSupport)

	logger.Logger.Info("Servidor iniciado na porta 8080")

	return r
}

