package initializers

import (
	"pasteBin/internal/database"

	"gorm.io/gorm"
)
func InitDb()( *gorm.DB,error ){
	db,err:=database.InitDB()
	if  err!=nil{
		return db,err
	}
	return db,nil
}
