package api

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

const key = "jamma-box"

func jwtGenerate(id string) (string, error) {
	var claims Claims
	claims.Id = id
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().AddDate(0, 1, 0))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func jwtVerify(str string) (*Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(str, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}
