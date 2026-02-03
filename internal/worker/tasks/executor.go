package worker

import (
	"context"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

type Executor interface {
	Type() taskV1.TaskType
	Execute(ctx context.Context, task taskV1.Task) error
}
