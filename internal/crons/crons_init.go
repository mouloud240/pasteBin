package crons

import (
	"fmt"
	"log"
	"pasteBin/internal/database/repository"

	"github.com/robfig/cron"
	"gorm.io/gorm"
)
type CronsManager struct {
	db *gorm.DB
	cron *cron.Cron
}
func NewCronsManager(db *gorm.DB,cron *cron.Cron) *CronsManager{
	return &CronsManager{db: db,cron: cron}
}
func (r *CronsManager) InitCrons() error {
    r.cron = cron.New()
    pasteRepo := repository.NewPastesRepository(r.db)
    
    if err := NewExpiresAtCron(pasteRepo).SetupExpiresAtCron(r.cron); err != nil {
        return fmt.Errorf("failed to setup expires cron: %w", err)
    }
    
    r.cron.Start()
    log.Print("Cron jobs initialized and started")
    return nil
}
