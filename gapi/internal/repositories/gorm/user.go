package gorm

import (
	"database/sql"
	"gapi-agp/internal/core/domain"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type UserGormRepo struct {
	*gorm.DB
	SqlDB  *sql.DB
	logger *zap.Logger
}

type User struct {
	ID         uint   `gorm:"column:id;type:int;auto_increment;primary_key"`
	ExternalID int    `gorm:"column:external_id"`
	Source     string `gorm:"column:source"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUserGormRepo(db *gorm.DB) (*UserGormRepo, error) {
	return &UserGormRepo{
		DB: db,
	}, nil
}

func (g UserGormRepo) GetUser(userID int) (user domain.User, err error) {
	var u User
	tx := g.DB.Where("id = ?", userID).Find(&u)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}
	return u.toDto(), nil
}
