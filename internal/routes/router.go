package routes

import (
	"be-blog/internal"
	"be-blog/internal/handler"
	"be-blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, app *internal.AppContainer) {
	authHandler := handler.NewAuthHandler(app.AuthService)

	r.Use(middleware.CORSMiddleware())
	api := r.Group("/api")
	api.POST("/login", authHandler.Login)
}
