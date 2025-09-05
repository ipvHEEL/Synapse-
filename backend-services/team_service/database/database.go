package database

import (
	"log"
	"team_service/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("TeamDB"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных команд", err)
	}

	DB.AutoMigrate(&models.Team{})
}
