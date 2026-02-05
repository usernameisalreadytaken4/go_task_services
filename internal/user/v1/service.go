package user

import (
	"context"
	"log"
)

type service struct {
	repo Repository
}

type Service interface {
	GetUser(context.Context, string, string) (*User, error) // для создания не факт, что лучше передавать string string
	CreateUser(context.Context, string, string) (*User, error)
	GetTokenByUser(context.Context, *User) (string, error)
}

func (s *service) CreateUser(ctx context.Context, email, password string) (*User, error) {
	_, err := s.repo.GetByEmail(ctx, email)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	return s.repo.Create(ctx, email, password)
}

func (s *service) GetUser(ctx context.Context, email, password string) (*User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		log.Println(err.Error())
		return nil, ErrInternalError
	}
	if err := user.CheckPassword(password); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetTokenByUser(ctx context.Context, user *User) (string, error) {
	token, err := s.repo.GetToken(ctx, *user)
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
