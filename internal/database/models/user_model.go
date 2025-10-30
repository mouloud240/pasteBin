package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique",`
	Email        string  `gorm:"unique",`
	Pastes 		 []Paste `gorm:"foreignKey:UserID"`
	Password *string 
	IsAdmin 		bool `gorm:"default:false"`
}
