package worker

import (
	"context"
	"time"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

func StartWorkerPool(ctx context.Context, repo taskV1.Repository, registry *Registry, taskSource TaskSource, numWorkers int) {
	tasks := make(chan taskV1.Task, 100)

	for i := 0; i < numWorkers; i++ {
		worker := NewWorker(repo, *registry)
		go worker.Run(ctx, tasks)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(tasks)
				return
			default:
				task, err := taskSource.Fetch(ctx)
				if err != nil {
					time.Sleep(time.Second * 5)
					continue
				}

				tasks <- *task
			}
		}
	}()
}
