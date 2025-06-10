package service

import (
	"be-blog/internal/dto/request"
	"be-blog/internal/model"
	"be-blog/internal/repository"
	"errors"
)

type AuthService interface {
	Login(req request.AuthRequest) (*model.User, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(req request.AuthRequest) (*model.User, error) {
	user, err := s.repo.CheckIdentifier(req.Identifier)

	if err != nil {
		return user, errors.New("Email/Username tidak terdaftar!")
	}

	return user, err
}
