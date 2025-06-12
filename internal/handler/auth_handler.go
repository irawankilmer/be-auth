package handler

import (
	"be-blog/internal/dto/request"
	"be-blog/internal/service"
	"be-blog/pkg/response"
	"be-blog/pkg/validates"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service   service.AuthService
	Validator *validates.Validates
}

func NewAuthHandler(s service.AuthService, v *validates.Validates) *AuthHandler {
	return &AuthHandler{
		service:   s,
		Validator: v,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.AuthRequest
	if !h.Validator.ValidateJSON(c, &req) {
		return
	}

	token, err := h.service.Login(req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.OK(c, token, "Anda berhasil login", nil)
}
