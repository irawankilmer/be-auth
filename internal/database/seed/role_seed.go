package seed

import (
	"be-blog/internal/model"
	"gorm.io/gorm"
)

func RoleSeed(db *gorm.DB) {
	role := []*model.Role{
		{Name: "admin"},
		{Name: "penulis"},
		{Name: "editor"},
	}

	db.Create(role)
}
