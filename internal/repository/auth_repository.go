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
