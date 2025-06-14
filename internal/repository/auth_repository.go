package repository

import (
	"be-blog/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CheckIdentifier(identifier string) (*model.User, error)
	UpdateTokenVersion(user *model.User) error
	FindByID(id string) (*model.User, error)
	IsUsernameExists(username string) (bool, error)
	IsEmailExists(email string) (bool, error)
	GetRoleByNames(name string) (*model.Role, error)
	CreateGuestUser(user *model.User, role *model.Role) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) CheckIdentifier(identifier string) (*model.User, error) {
	var user model.User
	err := r.db.
		Preload("Roles").
		Where("username = ? OR email = ?", identifier, identifier).
		First(&user).Error

	return &user, err
}

func (r *authRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) UpdateTokenVersion(user *model.User) error {
	newTokenVersion := uuid.New().String()
	user.TokenVersion = newTokenVersion
	return r.db.Model(&model.User{}).
		Where("id = ?", user.ID).
		Update("token_version", newTokenVersion).Error
}

func (r *authRepository) IsUsernameExists(username string) (bool, error) {
	var user model.User
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
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == nil {
		return true, err
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, err
}

func (r *authRepository) GetRoleByNames(name string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err == nil {
		return nil, err
	}
	return &role, nil
}

func (r *authRepository) CreateGuestUser(user *model.User, role *model.Role) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	if err := r.db.Model(user).Association("Roles").Replace(role); err != nil {
		return err
	}

	return nil
}
