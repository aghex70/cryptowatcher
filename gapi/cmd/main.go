package main

import (
	"gapi-agp/infrastructure/config"
	"gapi-agp/infrastructure/persistence"
	"gapi-agp/internal/repositories/gorm"
)

func main() {
	config.LoadConfig(config.CONFIG_PATH)

	db, err := persistence.NewSqlDB()
	if err != nil {
		panic(err)
	}

	gormDB, err := persistence.NewGormDB(db)
	if err != nil {
		panic(err)
	}

	cache, err := persistence.NewRedisCache()
	if err != nil {
		panic(err)
	}

	tradeRepo, err := gorm.NewTradeGormRepo(gormDB)
	if err != nil {
		panic(err)
	}

}
