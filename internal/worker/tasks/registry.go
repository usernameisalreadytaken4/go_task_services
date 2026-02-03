package worker

import (
	"errors"

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

func (r *Registry) Get(taskName taskV1.TaskType) (Executor, error)																																																																																																																																																																																																					xecutor, error) {
	executor, ok := r.Tasks[taskName]
	if !ok {
		return nil, errors.New("Task not found")
	}
	return executor, nil
}
