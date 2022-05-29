package gorm

import "gapi-agp/internal/core/domain"

type User struct {
	ID         int    `gorm:"coulmn:id;type:int;auto_increment;primary_key"`
	ExternalID int    `gorm:"column:external_id"`
	Source     string `gorm:"column:source"`
}

func (g GormRepo) GetUser(userID int) (domain.User, error) {
	var user User
	tx := g.DB.Where("id = ?", userID).Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
