package config

import (
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Name               string        `mapstructure:"DB_NAME"`
	Host               string        `mapstructure:"DB_HOST"`
	Port               int           `mapstructure:"DB_PORT"`
	User               string        `mapstructure:"DB_USER"`
	Password           string        `mapstructure:"DB_PASSWORD"`
	MaxOpenConnections int           `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int           `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	MaxConnLifeTime    time.Duration `mapstructure:"DB_MAX_CONN_LIFE_TIME"`
	Dialect            string        `mapstructure:"DB_DIALECT"`
	MigrationDir       string        `mapstructure:"DB_MIGRATION_DIR"`
	LogQuery           bool
}

func LoadDatabaseConfig(path string) (config DatabaseConfig, err error) {
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
