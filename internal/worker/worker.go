package worker

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/repository"
	"sync"
	"time"
)

type Worker struct {
	DB       *sql.DB
	Tasks    repository.Tasks
	Interval time.Duration
	stop     chan struct{}
	once     sync.Once
}

func New(db *sql.DB, d time.Duration) *Worker {
	return &Worker{DB: db, Tasks: repository.Tasks{DB: db}, Interval: d, stop: make(chan struct{})}
}
func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-w.stop:
			return
		case now := <-ticker.C:
			_ = w.Tasks.RequeueExpired(ctx, now.UTC())
		}
	}
}
func (w *Worker) Stop() { w.once.Do(func() { close(w.stop) }) }
