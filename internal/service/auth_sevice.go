package service

import (
	"be-blog/internal/dto/request"
	"be-blog/internal/repository"
	"be-blog/pkg"
	"errors"
)

type AuthService interface {
	Login(req request.AuthRequest) (string, error)
	Logout(userID string) error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(req request.AuthRequest) (string, error) {
	user, err := s.repo.CheckIdentifier(req.Identifier)

	if err != nil || !pkg.CompareHash(user.Password, req.Password) {
		return "", errors.New("Email/Username atau password salah!")
	}

	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	var _ = s.repo.UpdateTokenVersion(user)
	token, err := pkg.GenerateJWT(user.ID, roleNames, user.TokenVersion)

	return token, err
}

func (s *authService) Logout(userID string) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	return s.repo.UpdateTokenVersion(user)
}
