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
