package repository

import (
	"github.com/google/uuid"
	model2 "github.com/irawankilmer/be-auth/pkg/auth/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CheckIdentifier(identifier string) (*model2.User, error)
	UpdateTokenVersion(user *model2.User) error
	FindByID(id string) (*model2.User, error)
	IsUsernameExists(username string) (bool, error)
	IsEmailExists(email string) (bool, error)
	GetRoleByNames(name string) (*model2.Role, error)
	CreateGuestUser(user *model2.User, role *model2.Role) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) CheckIdentifier(identifier string) (*model2.User, error) {
	var user model2.User
	err := r.db.
		Preload("Roles").
		Where("username = ? OR email = ?", identifier, identifier).
		First(&user).Error

	return &user, err
}

func (r *authRepository) FindByID(id string) (*model2.User, error) {
	var user model2.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) UpdateTokenVersion(user *model2.User) error {
	newTokenVersion := uuid.New().String()
	user.TokenVersion = newTokenVersion
	return r.db.Model(&model2.User{}).
		Where("id = ?", user.ID).
		Update("token_version", newTokenVersion).Error
}

func (r *authRepository) IsUsernameExists(username string) (bool, error) {
	var user model2.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err == nil {
		return true, nil
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, nil
}

func (r *authRepository) IsEmailExists(email string) (bool, error) {
	var user model2.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == nil {
		return true, err
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, err
}

func (r *authRepository) GetRoleByNames(name string) (*model2.Role, error) {
	var role model2.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err == nil {
		return nil, err
	}
	return &role, nil
}

func (r *authRepository) CreateGuestUser(user *model2.User, role *model2.Role) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	if err := r.db.Model(user).Association("Roles").Replace(role); err != nil {
		return err
	}

	return nil
}
