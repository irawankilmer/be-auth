package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/be-auth/internal/config"
	"github.com/irawankilmer/be-auth/pkg/auth/container"
	"github.com/irawankilmer/be-auth/pkg/auth/routes"
	"os"
)

func main() {
	config.LoadENV()
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	app := container.InitApp()
	routes.InitRouter(r, app)

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
