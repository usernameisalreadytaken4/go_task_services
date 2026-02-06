package worker

import (
	"context"

	taskV1 "github.com/usernameisalreadytaken4/go_task_services/internal/task/v1"
)

type TaskSource interface {
	Fetch(ctx context.Context) (*taskV1.Task, error)
	Complete(ctx context.Context, task *taskV1.Task) error
}

type PostgresSource struct {
	repo taskV1.Repository
}

func (p *PostgresSource) Fetch(ctx context.Context) (*taskV1.Task, error) {
	return p.repo.GetNextTask(ctx)
}

func (p *PostgresSource) Complete(ctx context.Context, task *taskV1.Task) error {
	return p.repo.Update(ctx, task)
}

func NewPostgresSource(repo taskV1.Repository) *PostgresSource {
	return &PostgresSource{repo: repo}
}
