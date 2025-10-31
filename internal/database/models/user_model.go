package models

import (
	"log"
	. "pasteBin/pkg/hash"

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


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	  if u.Password!=nil {
				hashedPassword, err := Hash(*u.Password)
					if err != nil {
						log.Print("Error hashing password before create:", err)
						tx.Rollback()
						return err
				}
				u.Password = hashedPassword
		}

			  return nil
}
