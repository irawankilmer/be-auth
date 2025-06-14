package seed

import (
	"github.com/irawankilmer/be-auth/pkg/auth/model"
	"gorm.io/gorm"
)

func RoleSeed(db *gorm.DB) {
	role := []*model.Role{
		{Name: "super admin"},
		{Name: "admin"},
		{Name: "penulis"},
		{Name: "editor"},
		{Name: "tamu"},
	}

	db.Create(role)
}
