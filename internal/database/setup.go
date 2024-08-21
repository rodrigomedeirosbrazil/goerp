package models

import (
	models "goerp/internal/auth/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("internal/database/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	DB = database
}
