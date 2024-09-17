package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userId int) interface{} {
	expires_at := time.Now().Add(time.Hour * 72).Unix()
	claims := jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    expires_at,
	}
	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte("secret"))

	data := map[string]interface{}{
		"token":      accessToken,
		"expires_at": time.Unix(expires_at, 0),
		"error":        err,
	}
	return data
}
