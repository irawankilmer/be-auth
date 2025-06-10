package migration

import (
	"be-blog/internal/model"
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
	)

	if err != nil {
		log.Fatalf("Migrasi database gagal %v", err.Error())
	}
}
