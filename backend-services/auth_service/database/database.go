package database

import (
	users "auth_service/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("auth_database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных авторизации")
	}

	DB.AutoMigrate(&users.User{})
}
