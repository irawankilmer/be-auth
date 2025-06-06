package bootstrap

import (
	"be-blog/config"
)

func InitAPP() {
	config.LoadENV()
	config.InitDB()
	db := config.DB
	MigrateAndSeed(db)
}
