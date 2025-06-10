package repository

import (
	"be-blog/internal/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CheckIdentifier(identifier string) (*model.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) CheckIdentifier(identifier string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ? OR email = ?", identifier, identifier).First(&user).Error

	return &user, err
}
