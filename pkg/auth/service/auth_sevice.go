package service

import (
	"errors"
	"github.com/irawankilmer/be-auth/pkg"
	"github.com/irawankilmer/be-auth/pkg/auth/dto/request"
	"github.com/irawankilmer/be-auth/pkg/auth/model"
	"github.com/irawankilmer/be-auth/pkg/auth/repository"
)

type AuthService interface {
	Login(req request.AuthRequest) (string, error)
	Logout(userID string) error
	RegisterGuest(req request.GuestRegisterRequest) error
	IsUsernameTaken(username string) bool
	IsEmailTaken(username string) bool
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

func (s *authService) RegisterGuest(req request.GuestRegisterRequest) error {
	exists, err := s.repo.IsUsernameExists(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username sudah digunakan")
	}

	exists, err = s.repo.IsEmailExists(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email sudah digunakan")
	}

	hashedPassword, err := pkg.GenerateHash(req.Password)
	if err != nil {
		return errors.New("gagal enkripsi password")
	}

	// Ambil role tamu
	role, err := s.repo.GetRoleByNames("tamu")
	if err != nil {
		return errors.New("role 'tamu' tidak ditemukan")
	}

	user := &model.User{
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Buat user
	if err := s.repo.CreateGuestUser(user, role); err != nil {
		return err
	}

	return nil
}

func (s *authService) IsUsernameTaken(username string) bool {
	exists, _ := s.repo.IsUsernameExists(username)
	return exists
}

func (s *authService) IsEmailTaken(email string) bool {
	exists, _ := s.repo.IsEmailExists(email)
	return exists
}
