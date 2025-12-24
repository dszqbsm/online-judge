package jwtc

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtPayload struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type JwtConfig struct {
	AccessSecret string
	AccessExpire int64
}

func GenerateToken(cfg JwtConfig, userId int64) (string, error) {
	claims := JwtPayload{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(cfg.AccessExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.AccessSecret))
}
