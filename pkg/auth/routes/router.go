package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irawankilmer/be-auth/internal/container"
	"github.com/irawankilmer/be-auth/pkg/auth/handler"
	middleware2 "github.com/irawankilmer/be-auth/pkg/auth/middleware"
	validates2 "github.com/irawankilmer/be-auth/pkg/validates"
	"net/http"
)

func InitRouter(r *gin.Engine, app *container.AppContainer) {
	v := validator.New()
	validates := validates2.NewValidates(v)
	authHandler := handler.NewAuthHandler(app.AuthService, validates)

	r.Use(middleware2.CORSMiddleware())
	api := r.Group("/api")
	api.POST("/login", authHandler.Login)
	api.POST("/guest-register", authHandler.RegisterGuest)

	// Route setelah ini harus memiliki authentication
	api.Use(middleware2.AuthMiddleware())
	api.POST("/logout", authHandler.Logout)

	api.GET("/me", func(c *gin.Context) { // Coba saja, nanti dihapus lagi
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	// Route setelah ini harus user dengan admin dan editor
	api.Use(middleware2.RoleMiddleware("admin", "editor"))
	api.GET("/nulis", func(c *gin.Context) { // Coba saja, nanti dihapus lagi
		c.JSON(http.StatusOK, gin.H{"message": "Selamat datang admin dan editor"})
	})
}
