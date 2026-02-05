package task

import (
	"context"

	userV1 "github.com/usernameisalreadytaken4/go_task_services/internal/user/v1"
)

type service struct {
	repo Repository
}

type Service interface {
	Create(context.Context, *userV1.User, *Task) (*Task, error)
	Get(context.Context, int, int) (*Task, error)
	GetByUserID(context.Context, int) ([]Task, error)
}

func (s *service) Create(ctx context.Context, user *userV1.User, task *Task) (*Task, error) {
	task, err := s.repo.Save(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (s *service) Get(ctx context.Context, ID, userID int) (*Task, error) {
	task, err := s.repo.GetByID(ctx, ID, userID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *service) GetByUserID(ctx context.Context, userID int) ([]Task, error) {
	tasks, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
