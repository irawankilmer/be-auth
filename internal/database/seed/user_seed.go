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

	var role model.Role
	if err := db.Where("name = ?", "admin").First(&role).Error; err != nil {
		panic(err)
	}

	var user = model.User{
		FullName: "Admin Pertama",
		Username: "admin1",
		Email:    "admin1@gmail.com",
		Password: password,
		Roles:    []*model.Role{&role},
	}

	db.Create(&user)
}
