package routes

import (
	"be-blog/internal"
	"be-blog/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, app *internal.AppContainer) {
	authHandler := handler.NewAuthHandler(app.AuthService)

	api := r.Group("/api")
	api.POST("/login", authHandler.Login)
}
