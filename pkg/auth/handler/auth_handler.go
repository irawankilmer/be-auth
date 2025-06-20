package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/irawankilmer/be-auth/pkg/auth/dto/request"
	"github.com/irawankilmer/be-auth/pkg/auth/service"
	"github.com/irawankilmer/be-auth/pkg/response"
	"github.com/irawankilmer/be-auth/pkg/validates"
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

func (h *AuthHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User tidak ditemukan!")
		return
	}

	idStr, ok := userID.(string)
	if !ok {
		response.BadRequest(c, nil, "ID user tidak valid!")
		return
	}

	if err := h.service.Logout(idStr); err != nil {
		response.ServerError(c, "Gagal Logout")
		return
	}

	response.OK(c, nil, "Logout berhasil", nil)
}

func (h *AuthHandler) RegisterGuest(c *gin.Context) {
	var req request.GuestRegisterRequest

	// Validasi input
	if !h.Validator.ValidateJSON(c, &req) {
		return
	}

	// Validasi bussines login
	fieldErrors := make(map[string]string)

	if h.service.IsUsernameTaken(req.Username) {
		fieldErrors["username"] = "Username sudah terdaftar"
	}

	if h.service.IsEmailTaken(req.Email) {
		fieldErrors["email"] = "Email sudah terdaftar"
	}

	if !h.Validator.ValidateBussiness(c, &req, fieldErrors) {
		return
	}

	// Buat user tamu baru
	if err := h.service.RegisterGuest(req); err != nil {
		response.BadRequest(c, err.Error(), "user tamu gagal dibuat")
		return
	}

	response.Created(c, nil, "user tamu berhasil dibuat")
}
