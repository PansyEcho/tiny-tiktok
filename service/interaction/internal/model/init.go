package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:shi@tcp(43.139.184.22:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	err = db.AutoMigrate(&User{})
	err = db.AutoMigrate(&Favorite{})
	err = db.AutoMigrate(&Video{})
	err = db.AutoMigrate(&Comment{})
	err = db.AutoMigrate(&Follow{})
	if err != nil {
		log.Fatalln("[DB ERROR] : ", err)
	}
	DB = db
}
