package persistence

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewGormDB(sqlDB *sql.DB) (*gorm.DB, error) {
	fmt.Printf("%+v\n", sqlDB)
	//gormDB, err := gorm.Open(mysql.New(mysql.Config{
	//	Conn: sqlDB,
	//}), &gorm.Config{})
	//var err error
	//if err != nil {
	fmt.Println("retrying connection different way")
	dsn := "crypto:crypto@tcp(db:10306)/cryptowatcher"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//}
	if err != nil {
		log.Fatalln(err)
	}
	return gormDB, nil
}
