package models

import (
	"database/sql"
	"pasteBin/pkg/hash"

	. "pasteBin/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Paste struct {
	
	ID string `gorm:"primaryKey;type:uuid`
	Content 	 string
	Password *string `json:"-"`
	MaxViews  *int
	UserID		*uint 
	author User  `gorm:"foreignKey:UserID" json:"Author"`
	ExpirationDate *sql.NullTime	
	TimeModel
}

func (p *Paste) BeforeCreate(tx *gorm.DB)(err error){
	if p.ID==""{
		p.ID=uuid.NewString()
	}
	if p.Password!=nil{
		p.Password,err=hash.Hash(*p.Password)

		if err!=nil{
			tx.Rollback()
		  return err
		}
	}
	return;
	
}

func (p *Paste) AfterFind(tx *gorm.DB)(err error){
	if p.Password!=nil {
		p.Content="Password protected!"
	}
	return;
}

