package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateJWT(userID uint, roles []string, tokenVersion string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id":       userID,
		"roles":         roles,
		"token_version": tokenVersion,
		"exp":           now.Add(time.Hour * 24).Unix(),
		"iat":           now.Unix(),
		"iss":           os.Getenv("APP_NAME"),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
