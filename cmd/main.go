package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/Gierdiaz/config"
	"github.com/Gierdiaz/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	err := logger.InitLogger()
	if err != nil {
		fmt.Printf("Erro ao inicializar o logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.CloseLogger()

	config, err := config.LoadConfig()
	if err != nil {
		logger.Logger.Error("Erro ao carregar configuração", zap.Error(err))
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/support", handleTest)

	logger.Logger.Info("Servidor iniciado na porta", zap.String("port", config.Server.APP_PORT))
	r.Run(":" + config.Server.APP_PORT)
}

func handleTest(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "microservice support",
		Status:  "success",
	})
}