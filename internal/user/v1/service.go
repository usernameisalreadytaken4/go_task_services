package user

import (
	"log"
)

type UserService struct {
	repo *UserRepository
}

func (s *UserService) CreateUser(email, password string) (*User, error) {
	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	return s.repo.CreateUser(email, password)
}

func (s *UserService) GetUser(email, password string) (*User, error) {
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

func (s *UserService) GetTokenByUser(user *User) (string, error) {
	token, err := s.repo.GetTokenByUser(*user)
	if err != nil {
		log.Println(err.Error())
		return "", ErrInternalError
	}
	return token, nil
}

func NewService(repo *UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
