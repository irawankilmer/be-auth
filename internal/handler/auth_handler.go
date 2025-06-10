package handler

import (
	"be-blog/internal/dto/request"
	"be-blog/internal/service"
	"be-blog/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	user, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !pkg.CompareHash(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Password salah!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"role": user.Roles,
	})
}
