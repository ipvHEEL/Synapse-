package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	TeamId            int    `json:"TeamId"`
	TeamName          string `json:"TeamName"`
	TeamSpecification string `json:"TeamSpecification"`
}
