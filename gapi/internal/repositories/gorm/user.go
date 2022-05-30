package gorm

import (
	"gapi-agp/internal/core/domain"
	"time"
)

type User struct {
	ID         uint   `gorm:"column:id;type:int;auto_increment;primary_key"`
	ExternalID int    `gorm:"column:external_id"`
	Source     string `gorm:"column:source"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (g GormRepo) GetUser(userID int) (user domain.User, err error) {
	var u User
	tx := g.DB.Where("id = ?", userID).Find(&u)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}
	return u.toDto(), nil
}
