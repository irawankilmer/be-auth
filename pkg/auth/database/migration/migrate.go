package migration

import (
	model2 "github.com/irawankilmer/be-auth/pkg/auth/model"
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model2.User{},
		&model2.Role{},
	)

	if err != nil {
		log.Fatalf("Migrasi database gagal %v", err.Error())
	}
}
