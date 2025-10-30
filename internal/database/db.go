package database

import (
	"pasteBin/internal/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{},&models.Paste{})
	println("Migrated Succefully")
	
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	
	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(0)

	return db, nil
}
