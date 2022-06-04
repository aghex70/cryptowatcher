package main

import (
	"gapi-agp/config"
	"gapi-agp/infrastructure/persistence"
	"gapi-agp/internal/core/usecases"
	"gapi-agp/internal/logger"
	"gapi-agp/internal/repositories/gorm"
	"gapi-agp/internal/repositories/redis"
	"gapi-agp/pkg/providers"
	"gapi-agp/server"
	"go.uber.org/zap"
)

func main() {
	logger.ZapLogger.Info("Starting server...")

	// Initialize config
	logger.ZapLogger.Info("Loading configuration")
	config.LoadConfig(config.CONFIG_PATH)
	logger.ZapLogger.Info("Configuration loaded")

	// Intialize database
	logger.ZapLogger.Info("Starting application database")
	_, err := persistence.NewSqlDB()
	if err != nil {
		logger.ZapLogger.Fatal("error starting application database", zap.Error(err))
	}

	// Initialize ORM
	logger.ZapLogger.Info("Starting GORM ORM")
	gormDB, err := persistence.NewGormDB()
	if err != nil {
		logger.ZapLogger.Fatal("error starting GORM ORM", zap.Error(err))
	}
	logger.ZapLogger.Info("GORM ORM started")

	// Initialize redis cache
	logger.ZapLogger.Info("Starting redis cache")
	cache, err := persistence.NewRedisCache()
	if err != nil {
		logger.ZapLogger.Fatal("error starting redis cache", zap.Error(err))
	}
	logger.ZapLogger.Info("Redis cache started")

	// Initialize repositories
	logger.ZapLogger.Info("Initializing repositories")
	cacheRepo, err := redis.NewRedisRepo(cache)
	if err != nil {
		logger.ZapLogger.Fatal("error initializing redis repository", zap.Error(err))
	}

	logger.ZapLogger.Info("Initializing trade repository")
	tradeRepo, err := gorm.NewTradeGormRepo(gormDB)
	if err != nil {
		logger.ZapLogger.Fatal("error initializing trade repository", zap.Error(err))
	}

	logger.ZapLogger.Info("Initializing user repository")
	userRepo, err := gorm.NewUserGormRepo(gormDB)
	if err != nil {
		logger.ZapLogger.Fatal("error initializing user repository", zap.Error(err))
	}

	logger.ZapLogger.Info("Initializing provider manager")
	providerManager := providers.NewProviderManager()
	logger.ZapLogger.Info("Provider manager initialized")

	logger.ZapLogger.Info("Initializing trade service")
	tradeService := usecases.NewTradeInteractor(tradeRepo, cacheRepo, providerManager, logger.ZapLogger)
	logger.ZapLogger.Info("Trade service initialized")

	logger.ZapLogger.Info("Initializing user service")
	userService := usecases.NewUserInteractor(userRepo, cacheRepo, providerManager)
	logger.ZapLogger.Info("User service initialized")

	// Start server
	logger.ZapLogger.Info("Starting server")
	s := server.NewServer(tradeService, userService)
	err = s.StartServer()
	if err != nil {
		logger.ZapLogger.Fatal("error starting server", zap.Error(err))
	}

}
