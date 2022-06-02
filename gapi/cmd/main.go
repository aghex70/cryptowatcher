package main

import (
	"gapi-agp/config"
	"gapi-agp/infrastructure/persistence"
	"gapi-agp/internal/core/usecases"
	"gapi-agp/internal/repositories/gorm"
	"gapi-agp/internal/repositories/redis"
	"gapi-agp/pkg/providers"
	"gapi-agp/server"
)

func main() {
	// Initialize config
	config.LoadConfig(config.CONFIG_PATH)

	// Intialize database
	db, err := persistence.NewSqlDB()
	if err != nil {
		panic(err)
	}

	// Initialize ORM
	gormDB, err := persistence.NewGormDB(db)
	if err != nil {
		panic(err)
	}

	// Initialize redis cache
	cache, err := persistence.NewRedisCache()
	if err != nil {
		panic(err)
	}

	// Initialize repositories
	cacheRepo, err := redis.NewRedisRepo(cache)
	if err != nil {
		panic(err)
	}

	tradeRepo, err := gorm.NewTradeGormRepo(gormDB)
	if err != nil {
		panic(err)
	}

	userRepo, err := gorm.NewUserGormRepo(gormDB)
	if err != nil {
		panic(err)
	}

	providerManager := providers.NewProviderManager()
	logger := config.NewLogger()
	tradeService := usecases.NewTradeInteractor(tradeRepo, cacheRepo, providerManager, logger.ZapLogger)
	userService := usecases.NewUserInteractor(userRepo, cacheRepo, providerManager)

	// Start server
	s := server.NewServer(tradeService, userService)
	err = s.StartServer()
	if err != nil {
		panic(err)
	}

}
