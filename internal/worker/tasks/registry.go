package worker

import (
	"fmt"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

type Registry struct {
	Tasks map[taskV1.TaskType]Executor
}

func NewRegistry(executors ...Executor) *Registry {
	tasks := make(map[taskV1.TaskType]Executor)
	for _, e := range executors {
		tasks[e.Type()] = e
	}
	return &Registry{Tasks: tasks}
}

func (r *Registry) Get(taskType taskV1.TaskType) (Executor, error) {
	executor, ok := r.Tasks[taskType]
	if !ok {
		return nil, fmt.Errorf("executor '%v' not found", taskType)
	}
	return executor, nil
}
