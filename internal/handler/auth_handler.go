package handler

import (
	"be-blog/internal/dto/request"
	"be-blog/internal/service"
	"be-blog/pkg/response"
	"github.com/gin-gonic/gin"
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
		response.BadRequest(c, nil, "Gagal validasi input")
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.OK(c, token, "Anda berhasil login", nil)
}
