package persistence

import (
	"database/sql"
	"go.uber.org/zap"

	"fmt"
	"gapi-agp/config"
	"gapi-agp/internal/logger"
	"github.com/go-sql-driver/mysql"
)

func NewSqlDB() (*sql.DB, error) {
	mySqlConfig := mysql.NewConfig()
	address := fmt.Sprintf("%s:%d", config.C.Database.Host, config.C.Database.Port)
	mySqlConfig.Addr = address
	mySqlConfig.DBName = config.C.Database.Name
	mySqlConfig.User = config.C.Database.User
	mySqlConfig.Passwd = config.C.Database.Password
	mySqlConfig.Net = config.C.Database.Net
	mySqlConfig.ParseTime = true

	logger.ZapLogger.Info("Connecting to database")
	sqlDB, err := sql.Open(config.C.Database.Dialect, mySqlConfig.FormatDSN())
	if err != nil {
		logger.ZapLogger.Error("Error connecting to database", zap.Error(err))
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(config.C.Database.MaxConnLifeTime)
	sqlDB.SetMaxOpenConns(config.C.Database.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(config.C.Database.MaxIdleConnections)
	return sqlDB, nil
}
