package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func InitLogger() error {
	// Carrega o nível de log a partir da variável de ambiente LOG_LEVEL
	logLevel := zapcore.InfoLevel
	if os.Getenv("LOG_LEVEL") == "debug" {
		logLevel = zapcore.DebugLevel
	}

	// Configura o encoder e o log file (caso queira salvar os logs em arquivo)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	file, err := os.Create("app.log")
	if err != nil {
		return err
	}

	// Cria um escritor para o log
	writeSyncer := zapcore.AddSync(file)
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// Inicializa o logger
	Logger = zap.New(core)
	return nil
}

func CloseLogger() {
	_ = Logger.Sync() // Garante que os logs pendentes sejam gravados
}
