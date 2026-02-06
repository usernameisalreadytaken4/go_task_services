package worker

import (
	"context"
	"time"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

// i/o bound
type ShortTask struct {
	repo taskV1.Repository
}

func (ShortTask) Type() taskV1.TaskType {
	return "short_task"
}

func (t *ShortTask) Execute(ctx context.Context, task taskV1.Task) error {
	now := time.Now()
	task.Started = &now

	time.Sleep(time.Second * 3)

	task.Status = "done"
	now = time.Now()
	task.Finished = &now
	t.repo.Update(ctx, &task)
	return nil
}

func NewShortTask(repo taskV1.Repository) *ShortTask {
	return &ShortTask{repo: repo}
}
