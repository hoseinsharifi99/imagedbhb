package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect(host string, port int, username, password, dbname string) (err error) {
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		dbname,
	))
	return
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate(model interface{}) {
	db.AutoMigrate(model)
}
