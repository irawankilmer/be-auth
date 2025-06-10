package internal

import (
	"be-blog/internal/config"
	"be-blog/internal/migration"
	"be-blog/internal/repository"
	"be-blog/internal/service"
	"os"
)

type AppContainer struct {
	AuthService service.AuthService
}

func InitApp() *AppContainer {
	config.InitDB()
	db := config.DB

	if os.Getenv("GIN_MODE") == "debug" {
		migration.AutoMigrate(db)
	}

	authRepo := repository.NewAuthRepository(db)

	authService := service.NewAuthService(authRepo)

	return &AppContainer{
		AuthService: authService,
	}
}
