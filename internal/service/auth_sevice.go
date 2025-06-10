package service

import (
	"be-blog/internal/dto/request"
	"be-blog/internal/repository"
	"errors"
)

type AuthService interface {
	Login(req request.AuthRequest) (string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(req request.AuthRequest) (string, error) {
	_, err := s.repo.CheckIdentifier(req.Identifier)

	if err != nil {
		return "", errors.New("Email/Username tidak terdaftar!")
	}

	return "", err
}
