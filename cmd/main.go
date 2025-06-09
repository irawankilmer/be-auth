package main

import (
	"be-blog/internal"
	"be-blog/internal/config"
	"be-blog/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	config.LoadENV()
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	app := internal.InitApp()
	routes.InitRouter(r, app)

	port := os.Getenv("APP_PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
