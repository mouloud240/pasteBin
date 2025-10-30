package initializers

import (
	"pasteBin/internal/database"

	"gorm.io/gorm"
)
var DB *gorm.DB
func InitDb()error{

	db,err:=database.InitDB()

	
	if  err!=nil{
		return err
	}
	DB=db
	return nil
}
