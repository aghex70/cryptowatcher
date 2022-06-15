package logger

import (
	"go.uber.org/zap"
	"log"
)

var ZapLogger *zap.Logger

func NewLogger() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Status: Failed to initialize zap logger: %v", err)
	}
	l.Info("Zap logger is now initialized")
	return l
}

func init() {
	ZapLogger = NewLogger()
}
