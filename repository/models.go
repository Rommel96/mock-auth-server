package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rommel96/mock-auth-server/database"
)

var db *gorm.DB

func Init() {
	db = database.GetConnectionDB()
	err := db.DropTableIfExists(&User{}).Error
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(&User{}).Error
	if err != nil {
		log.Println(err)
	}
}

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
