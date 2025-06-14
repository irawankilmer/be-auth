package request

type AuthRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func (a *AuthRequest) Sanitize() map[string]any {
	return map[string]any{
		"identifier": a.Identifier, // field yang masuk ke old data. Password jangan karena sensitive
	}
}

type GuestRegister struct {
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (a *GuestRegister) Sanitize() map[string]any {
	return map[string]any{
		"full_name": a.FullName,
		"username":  a.Username,
		"email":     a.Email,
	}
}
