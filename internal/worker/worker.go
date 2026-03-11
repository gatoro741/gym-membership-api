package worker

import (
	"context"
	"log"
	"time"
)

type Storage interface {
	DeactivateExpiredMemberships(ctx context.Context) error
}
type Worker struct {
	storage Storage
}

func New(storage Storage) *Worker {
	return &Worker{storage: storage}
}

func (w *Worker) Start(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {

			case <-ticker.C:
				err := w.storage.DeactivateExpiredMemberships(ctx)
				if err != nil {
					log.Println("worker: error DeactivateExpiredMemberships")
				}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}

	}()
}
