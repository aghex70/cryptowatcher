package gorm

import (
	"database/sql"
	"gorm.io/gorm"
)

type GormRepo struct {
	*gorm.DB
	SqlDb *sql.DB
}
