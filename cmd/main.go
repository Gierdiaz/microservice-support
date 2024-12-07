package main

import (
	"fmt"

	"os"

	"github.com/Gierdiaz/config"
	"github.com/Gierdiaz/internal/endpoints"
	"github.com/Gierdiaz/pkg/logger"

	"go.uber.org/zap"
)

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

	r := endpoints.SetupRouter(config)
	logger.Logger.Info("Servidor iniciado na porta", zap.String("port", config.Server.APP_PORT))
	if err := r.Run(":" + config.Server.APP_PORT); err != nil {
		logger.Logger.Fatal("Erro ao iniciar o servidor", zap.Error(err))
	}
}
