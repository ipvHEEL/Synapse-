package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"size:255;uniqueIndex"`
	PasswordHash string `json:"-" gorm:"size:255"`
}
