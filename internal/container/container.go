package container

import (
	"github.com/irawankilmer/be-auth/internal/config"
	"github.com/irawankilmer/be-auth/internal/database/migration"
	"github.com/irawankilmer/be-auth/internal/database/seed"
	"github.com/irawankilmer/be-auth/pkg/auth/repository"
	"github.com/irawankilmer/be-auth/pkg/auth/service"
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
		seed.MainSeed(db)
	}

	authRepo := repository.NewAuthRepository(db)

	authService := service.NewAuthService(authRepo)

	return &AppContainer{
		AuthService: authService,
	}
}
