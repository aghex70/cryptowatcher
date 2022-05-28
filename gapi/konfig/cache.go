package konfig

import (
	"github.com/spf13/viper"
)

type CacheConfig struct {
	Host     string `mapstructure:"CACHE_HOST"`
	Port     int    `mapstructure:"CACHE_PORT"`
	Name     string `mapstructure:"CACHE_NAME"`
	User     string `mapstructure:"CACHE_USER"`
	Password string `mapstructure:"CACHE_PASSWORD"`
}

func LoadCacheConfig(path string) (config CacheConfig, err error) {
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
