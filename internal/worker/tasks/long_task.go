package worker

import (
	"context"
	"crypto/sha256"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

// cpu bound
type LongTask struct{}

func (LongTask) Name() taskV1.TaskType {
	return "long_task"
}

func (t *LongTask) Execute(ctx context.Context, task *taskV1.Task) error {
	payload := task.Payload
	hash := []byte(payload)
	for i := 0; i < 100000; i++ {
		sum := sha256.Sum256(hash)
		hash = sum[:]
	}
	task.Result = hash
	task.Status = "DONE"
	// сохранить
	return nil
}
