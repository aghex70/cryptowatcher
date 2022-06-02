package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
)

type LoggerConfig struct {
	FluentdHost    string `mapstructure:"FLUENTD_HOST"`
	FluentdPort    int    `mapstructure:"FLUENTD_PORT"`
	FluentdEnabled bool   `mapstructure:"FLUENTD_ENABLED"`
}

func LoadLoggerConfig(path string) (config LoggerConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

type Logger struct {
	ZapLogger *zap.Logger
}

func NewLogger() *Logger {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Status: Failed to initialize zap logger: %v", err)
	}
	l.Info("Zap log is now initialized")
	return &Logger{ZapLogger: l}
}
