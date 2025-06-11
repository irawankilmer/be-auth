package routes

import (
	"be-blog/internal"
	"be-blog/internal/handler"
	"be-blog/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine, app *internal.AppContainer) {
	authHandler := handler.NewAuthHandler(app.AuthService)

	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	api.POST("/login", authHandler.Login)

	// Route setelah ini harus memiliki authentication
	api.Use(middleware.AuthMiddleware())
	api.GET("/me", func(c *gin.Context) { // Coba saja
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
}
