package seed

import (
	"gorm.io/gorm"
)

func MainSeed(db *gorm.DB) {
	RoleSeed(db)
	UserSeed(db)
}
