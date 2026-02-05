package task

import (
	"context"
	"errors"

	userV1 "github.com/usernameisalreadytaken4/go_task_services/internal/user/v1"
)

type service struct {
	repo Repository
}

type Service interface {
	Create(context.Context, *userV1.User, *Task) (*Task, error)
	Get(context.Context, int) (*Task, error)
	GetByUserID(context.Context, int) ([]*Task, error)
}

func (s *service) Create(ctx context.Context, user *userV1.User, task *Task) (*Task, error) {
	task, err := s.repo.Save(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (s *service) Get(ctx context.Context, ID int) (*Task, error) {
	return nil, errors.New("fail")
}

func (s *service) GetByUserID(ctx context.Context, userID int) ([]*Task, error) {
	return nil, errors.New("fail")
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
