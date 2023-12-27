package jwt

import (
	"authService/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	AppId  int `json:"app_id"`
}

func GenerateAccessToken(user *entity.User, app *entity.App, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		app.Id,
	})

	return token.SignedString([]byte(app.Secret))
}
