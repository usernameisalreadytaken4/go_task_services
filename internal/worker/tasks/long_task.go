package worker

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"time"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

// cpu bound
type LongTask struct {
	repo taskV1.Repository
}

func (LongTask) Type() taskV1.TaskType {
	return "long_task"
}

func (t *LongTask) Execute(ctx context.Context, task taskV1.Task) error {
	now := time.Now()
	task.Started = &now

	payload := task.Payload
	hash := []byte(payload)
	for i := 0; i < 100000; i++ {
		sum := sha256.Sum256(hash)
		hash = sum[:]
	}
	task.Result, _ = json.Marshal(hash)
	task.Status = "done"

	now = time.Now()
	task.Finished = &now
	t.repo.Update(ctx, &task)
	return nil
}

func NewLongTask(repo taskV1.Repository) *LongTask {
	return &LongTask{repo: repo}
}
