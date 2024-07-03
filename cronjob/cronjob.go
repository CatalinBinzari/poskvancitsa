package cronjob

import (
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"
)

// Function type definition
type Task func()

func StartCronjob(t Task, scheduled_time string) {
	s := gocron.NewScheduler(time.Local)
	slog.Info("cronjob started at", "scheduled_time", scheduled_time)
	_, err := s.Every(1).Day().At(scheduled_time).Do(t)
	if err != nil {
		slog.Error("cronjob", "err", err)
		return
	}

	s.StartAsync()
}
