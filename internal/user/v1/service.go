package user

import (
	"log"
)

type service struct {
	repo Repository
}

type Service interface {
	GetUser(string, string) (*User, error) // для создания не факт, что лучше передавать string string
	CreateUser(string, string) (*User, error)
	GetTokenByUser(*User) (string, error)
}

func (s *service) CreateUser(email, password string) (*User, error) {
	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	return s.repo.Create(email, password)
}

func (s *service) GetUser(email, password string) (*User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		log.Println(err.Error())
		return nil, ErrInternalError
	}
	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetTokenByUser(user *User) (string, error) {
	token, err := s.repo.GetToken(*user)
	if err != nil {
		log.Println(err.Error())
		return "", ErrInternalError
	}
	return token, nil
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
