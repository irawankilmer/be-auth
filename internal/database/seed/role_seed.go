package seed

import (
	"be-blog/internal/model"
	"gorm.io/gorm"
)

func RoleSeed(db *gorm.DB) {
	role := []*model.Role{
		{Name: "super admin"},
		{Name: "admin"},
		{Name: "penulis"},
		{Name: "editor"},
	}

	db.Create(role)
}
