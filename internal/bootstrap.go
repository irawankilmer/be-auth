package internal

import (
	"be-blog/internal/config"
	"be-blog/internal/migration"
	"os"
)

type AppContainer struct {
}

func InitApp() *AppContainer {
	config.InitDB()
	db := config.DB

	if os.Getenv("GIN_MODE") == "debug" {
		migration.AutoMigrate(db)
	}

	return &AppContainer{}
}
