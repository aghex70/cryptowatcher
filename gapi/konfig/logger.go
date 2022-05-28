package konfig

import (
	"fmt"
	"github.com/spf13/viper"
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
	fmt.Println("CONFIG_NAME:", CONFIG_NAME)
	fmt.Println("CONFIG_TYPE:", CONFIG_TYPE)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
