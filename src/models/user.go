package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Phone    string `gorm:"unique"`
	Password string
	Role     string
	Birth    string
	Gender   string
}
