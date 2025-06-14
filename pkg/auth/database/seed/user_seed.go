package seed

import (
	"github.com/irawankilmer/be-auth/pkg"
	model2 "github.com/irawankilmer/be-auth/pkg/auth/model"
	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {
	password, err := pkg.GenerateHash("superadmin1")
	if err != nil {
		panic(err)
	}

	var roles []*model2.Role
	roleNames := []string{"super admin", "admin", "penulis", "editor"}

	var tempRoles []model2.Role
	if err := db.Where("name IN ?", roleNames).Find(&tempRoles).Error; err != nil {
		panic(err)
	}

	for i := range tempRoles {
		roles = append(roles, &tempRoles[i])
	}

	var user = model2.User{
		FullName: "Super Admin 1",
		Username: "superadmin1",
		Email:    "superadmin1@gmail.com",
		Password: password,
		Roles:    roles,
	}

	db.Omit("Roles.*").Create(&user)
	_ = db.Model(&user).Association("Roles").Replace(roles)
}
