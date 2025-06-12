package routes

import (
	"be-blog/internal"
	"be-blog/internal/handler"
	"be-blog/internal/middleware"
	validates2 "be-blog/pkg/validates"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func InitRouter(r *gin.Engine, app *internal.AppContainer) {
	v := validator.New()
	validates := validates2.NewValidates(v)
	authHandler := handler.NewAuthHandler(app.AuthService, validates)

	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	api.POST("/login", authHandler.Login)

	// Route setelah ini harus memiliki authentication
	api.Use(middleware.AuthMiddleware())
	api.POST("/logout", authHandler.Logout)

	api.GET("/me", func(c *gin.Context) { // Coba saja, nanti dihapus lagi
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	// Route setelah ini harus user dengan admin dan editor
	api.Use(middleware.RoleMiddleware("admin", "editor"))
	api.GET("/nulis", func(c *gin.Context) { // Coba saja, nanti dihapus lagi
		c.JSON(http.StatusOK, gin.H{"message": "Selamat datang admin dan editor"})
	})
}
