package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewJWT(mySigningKey, userName string) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"user_name": userName,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte(mySigningKey))
}
