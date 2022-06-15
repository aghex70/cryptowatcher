package persistence

import (
	"fmt"
	"cwapi/config"
	"cwapi/internal/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.C.Database.User, config.C.Database.Password, config.C.Database.Net, config.C.Database.Host, config.C.Database.Port, config.C.Database.Name)
	//dsn = "crypto:crypto@tcp(localhost:10306)/cryptowatcher"
	logger.ZapLogger.Info("Connecting ORM to database")
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ZapLogger.Error("Error connecting ORM to database", zap.Error(err))
		return nil, err
	}
	return gormDB, nil
}
