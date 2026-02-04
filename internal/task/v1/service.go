package task

import (
	"errors"

	userV1 "github.com/usernameisalreadytaken4/go_task_services/internal/user/v1"
)

type service struct {
	repo Repository
}

type Service interface {
	Create(*userV1.User, *Task) (*Task, error)
	Get(int) (*Task, error)
	GetByUserID(int) ([]*Task, error)
}

func (s *service) Create(user *userV1.User, task *Task) (*Task, error) {
	task, err := s.repo.Save(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (s *service) Get(ID int) (*Task, error) {
	return nil, errors.New("fail")
}

func (s *service) GetByUserID(userID int) ([]*Task, error) {
	return nil, errors.New("fail")
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
