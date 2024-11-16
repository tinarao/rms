package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var Client *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to establish db connection: %s", err)
	}

	if err = db.AutoMigrate(&User{}, &Admin{}, &Order{}, &Product{}); err != nil {
		log.Fatalf("failed to run automigrations: %s", err)
	}

	Client = db
}
