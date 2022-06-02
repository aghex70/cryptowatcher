package config

import "log"

const (
	CONFIG_NAME string = ".env_gapi"
	CONFIG_TYPE string = "env"
	CONFIG_PATH string = "./config/"
)

type Config struct {
	Cache     CacheConfig
	Database  DatabaseConfig
	Logger    LoggerConfig
	Providers ProvidersConfig
	Server    ServerConfig
}

var C Config

func LoadConfig(path string) {
	Config := &C
	cacheConfig, err := LoadCacheConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	databaseConfig, err := LoadDatabaseConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	loggerConfig, err := LoadLoggerConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	providersConfig, err := LoadProvidersConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	serverConfig, err := LoadServerConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	Config.Logger = loggerConfig
	Config.Cache = cacheConfig
	Config.Database = databaseConfig
	Config.Providers = providersConfig
	Config.Server = serverConfig
}
