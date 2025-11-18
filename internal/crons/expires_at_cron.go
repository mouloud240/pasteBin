package crons

import (
	"log"
	"pasteBin/internal/database/repository"

	"github.com/robfig/cron"
)
type ExpiresAtCron struct{
	pasteRepo *repository.PastesRepository
}

func NewExpiresAtCron(pasteRepo *repository.PastesRepository) *ExpiresAtCron{
	return &ExpiresAtCron{pasteRepo: pasteRepo}
}
func (r *ExpiresAtCron) SetupExpiresAtCron(c *cron.Cron) error {
    return c.AddFunc("0 0 1 * * *", func() {
        log.Print("Running ExpiresAtCron Job")
        if err := r.pasteRepo.DeleteExpiredPastes(); err != nil {
            log.Printf("ERROR: Failed to delete expired pastes: %v", err)
        } else {
            log.Print("Successfully deleted expired pastes")
        }
    })
}
