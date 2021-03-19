package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rommel96/mock-auth-server/config"
)

var db *gorm.DB

func Connect() {
	dialect, args := config.GetConfigDB()
	conn, err := gorm.Open(dialect, args)
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}
	conn.LogMode(true)
	db = conn
	log.Println("DB Connected...")
}

func GetConnectionDB() *gorm.DB {
	return db
}
