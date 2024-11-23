package crons

import (
	"github.com/robfig/cron/v3"
)

type AppCron struct {
}

func NewAppCron() *AppCron {
	return &AppCron{}
}

func (ap *AppCron) RunCron() {
	c := cron.New()

	c.AddFunc("@every 6h", func() {

	})

	c.AddFunc("@every 15m", func() {

	})

	// Start the cron scheduler
	c.Start()
}
