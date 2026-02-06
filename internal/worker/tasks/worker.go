package worker

import (
	"context"
	"log"
	"time"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

type Worker struct {
	repo     taskV1.Repository
	registry *Registry
}

func (w *Worker) Run(ctx context.Context, tasks <-chan taskV1.Task) {
	for {
		time.Sleep(time.Second * 5)
		select {
		case <-ctx.Done():
			return
		case task, ok := <-tasks:
			if !ok {
				continue
			}

			executor, err := w.registry.Get(task.Type)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			err = executor.Execute(ctx, task)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}
	}
}

func NewWorker(repo taskV1.Repository, registry Registry) *Worker {
	return &Worker{repo: repo, registry: &registry}
}
