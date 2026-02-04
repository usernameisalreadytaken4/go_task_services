package task

import "errors"

type service struct {
	repo Repository
}

func (s *service) Create(task *Task) error {
	return nil
}
func (s *service) Get(ID int) (*Task, error) {
	return nil, errors.New("fail")
}

func (s *service) GetByUserID(userID int) ([]*Task, error) {
	return nil, errors.New("fail")
}

type Service interface {
	Create(*Task) error
	Get(int) (*Task, error)
	GetByUserID(int) ([]*Task, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
