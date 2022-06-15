package config

import (
	"github.com/spf13/viper"
)

type ProvidersConfig struct {
	Fetcher FetcherConfig
}

type FetcherConfig struct {
	Host string `mapstructure:"FETCHER_HOST"`
	Port int    `mapstructure:"FETCHER_PORT"`
}

func LoadFetcherConfig(path string) (config FetcherConfig, err error) {
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

func LoadProvidersConfig(path string) (config ProvidersConfig, err error) {
	config.Fetcher, err = LoadFetcherConfig(path)
	if err != nil {
		return config, err
	}
	return config, err
}
