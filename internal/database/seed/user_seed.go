package seed

import (
	"be-blog/internal/model"
	"be-blog/pkg"
	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {
	password, err := pkg.GenerateHash("admin1")
	if err != nil {
		panic(err)
	}

	var roles []*model.Role
	roleNames := []string{"admin", "editor"}

	var tempRoles []model.Role
	if err := db.Where("name IN ?", roleNames).Find(&tempRoles).Error; err != nil {
		panic(err)
	}

	for _, role := range tempRoles {
		roles = append(roles, &role)
	}

	var user = model.User{
		FullName: "Admin Pertama",
		Username: "admin1",
		Email:    "admin1@gmail.com",
		Password: password,
		Roles:    roles,
	}

	db.Create(&user)
}
