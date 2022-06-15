package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port int    `mapstructure:"SERVER_PORT"`
}

func LoadServerConfig(path string) (config ServerConfig, err error) {
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
