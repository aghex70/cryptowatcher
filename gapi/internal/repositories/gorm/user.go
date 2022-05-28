package gorm

type User struct {
	ID         int    `gorm:"coulmn:id;type:int;auto_increment;primary_key"`
	ExternalID int    `gorm:"column:external_id"`
	Source     string `gorm:"column:source"`
}
