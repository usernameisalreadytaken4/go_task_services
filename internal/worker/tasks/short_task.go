package worker

import (
	"context"
	"time"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

// i/o bound
type ShortTask struct{}

func (ShortTask) Name() taskV1.TaskType {
	return "short_task"
}

func (t *ShortTask) Execute(ctx context.Context, task taskV1.Task) error {
	time.Sleep(time.Second * 3)
	task.Status = "DONE"
	return nil
}
